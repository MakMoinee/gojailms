package mysqllocal

import (
	"fmt"
	"log"

	"github.com/MakMoinee/gojailms/internal/common"
	"github.com/MakMoinee/gojailms/internal/gojailms/models"
)

func (svc *mySqlService) InsertVisitorHistory(remarks string, visitorId string) (bool, error) {
	log.Println("Inside mysqllocal:InsertVisitorHistory()")
	isSave := true
	svc.Db = svc.openDBConnection()
	query := fmt.Sprintf(common.InsertVisitorHistoryQuery, remarks, visitorId)
	_, err := svc.Db.Query(query)
	if err != nil {
		log.Println(err.Error())
		isSave = false
	}
	defer svc.Db.Close()

	return isSave, err
}

func (svc *mySqlService) GetAllVisitorHistory() ([]models.VisitorHistory, error) {
	list := []models.VisitorHistory{}
	svc.Db = svc.openDBConnection()
	query := common.GetAllVisitorsHistoryQuery
	result, err := svc.Db.Query(query)
	if err != nil {
		log.Println("mysql:GetAllVisitorHistory() -> Query Error")
		return list, err
	}

	defer svc.Db.Close()

	for result.Next() {
		visitor := models.VisitorHistory{}
		err := result.Scan(
			&visitor.VisitorHistID,
			&visitor.UserID,
			&visitor.FirstName,
			&visitor.LastName,
			&visitor.MiddleName,
			&visitor.Remarks,
			&visitor.VisitDate,
		)
		if err != nil {
			log.Println("mysql:GetVisitors() -> Error in Scanning the Result")
			break
		}
		list = append(list, visitor)
	}

	defer result.Close()

	return list, nil
}
