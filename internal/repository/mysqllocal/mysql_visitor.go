package mysqllocal

import (
	"fmt"
	"log"

	"github.com/MakMoinee/gojailms/internal/common"
	"github.com/MakMoinee/gojailms/internal/gojailms/models"
)

func (svc *mySqlService) CreateVisitor(visitor models.Visitor) (bool, error) {
	log.Println("Inside mysql:CreateVisitor()")
	isCreated := true
	svc.Db = svc.openDBConnection()
	query := fmt.Sprintf(common.CreateVisitorQuery, visitor.UserID, visitor.FirstName, visitor.LastName, visitor.MiddleName, visitor.Address, visitor.BirthPlace, visitor.BirthDate)
	_, err := svc.Db.Query(query)
	if err != nil {
		log.Println("mysql:CreateVisitor()-> Query Error")
		isCreated = false
	}
	defer svc.Db.Close()
	return isCreated, err
}

func (svc *mySqlService) GetVisitors() ([]models.Visitor, error) {
	log.Println("Inside mysql:GetVisitors()")
	visitorsList := []models.Visitor{}

	svc.Db = svc.openDBConnection()
	query := common.GetVisitorsQuery
	result, err := svc.Db.Query(query)
	if err != nil {
		log.Println("mysql:GetVisitors() -> Query Error")
		return visitorsList, err
	}

	defer svc.Db.Close()

	for result.Next() {
		visitor := models.Visitor{}
		err := result.Scan(
			&visitor.VisitorID,
			&visitor.UserID,
			&visitor.FirstName,
			&visitor.LastName,
			&visitor.MiddleName,
			&visitor.Address,
			&visitor.BirthPlace,
			&visitor.BirthDate,
			&visitor.LastModifiedDate,
			&visitor.CreatedDate,
		)
		if err != nil {
			log.Println("mysql:GetVisitors() -> Error in Scanning the Result")
			break
		}
		visitorsList = append(visitorsList, visitor)
	}

	defer result.Close()

	return visitorsList, err
}
