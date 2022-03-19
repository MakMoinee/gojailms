package mysqllocal

import (
	"fmt"
	"log"

	"github.com/MakMoinee/gojailms/internal/common"
	"github.com/MakMoinee/gojailms/internal/gojailms/models"
)

func (svc *mySqlService) CreateInmate(inmate models.Inmates) (bool, error) {
	log.Println("Inside mysql:CreateInmate()")
	isUpdated := true
	svc.Db = svc.openDBConnection()
	query := fmt.Sprintf(common.CreateInmateQuery, inmate.CrimeID, inmate.FirstName, inmate.LastName, inmate.MiddleName, inmate.Address, inmate.BirthPlace, inmate.BirthDate)
	_, err := svc.Db.Query(query)
	if err != nil {
		isUpdated = false
	}
	defer svc.Db.Close()
	return isUpdated, err
}

func (svc *mySqlService) GetInmates() ([]models.Inmates, error) {
	log.Println("Inside mysql:GetInmates()")
	inmatesList := []models.Inmates{}
	svc.Db = svc.openDBConnection()
	query := common.GetInmatesQuery
	result, err := svc.Db.Query(query)
	if err != nil {
		return inmatesList, err
	}
	defer svc.Db.Close()
	for result.Next() {
		inmate := models.Inmates{}
		err = result.Scan(
			&inmate.InmateID,
			&inmate.CrimeID,
			&inmate.FirstName,
			&inmate.LastName,
			&inmate.MiddleName,
			&inmate.Address,
			&inmate.BirthDate,
			&inmate.BirthPlace,
			&inmate.LastModifiedDate,
			&inmate.CreatedDate,
		)
		if err != nil {
			break
		}
		inmatesList = append(inmatesList, inmate)
	}

	return inmatesList, err
}
