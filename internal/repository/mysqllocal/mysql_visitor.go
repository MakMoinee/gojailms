package mysqllocal

import (
	"errors"
	"fmt"
	"log"

	"github.com/MakMoinee/gojailms/internal/common"
	"github.com/MakMoinee/gojailms/internal/gojailms/models"
)

func (svc *mySqlService) CreateVisitor(visitor models.Visitor) (bool, error) {
	log.Println("Inside mysql:CreateVisitor()")
	isCreated := true
	svc.Db = svc.openDBConnection()
	query := fmt.Sprintf(common.CreateVisitorQuery, visitor.UserID, visitor.FirstName, visitor.LastName, visitor.MiddleName, visitor.Address, visitor.BirthPlace, visitor.BirthDate, visitor.ContactNumber)
	_, err := svc.Db.Query(query)
	if err != nil {
		log.Println("mysql:CreateVisitor()-> Query Error >> ", err)
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

func (svc *mySqlService) GetVisitorById(userID string) (models.Visitor, error) {
	visitor := models.Visitor{}
	var err error

	if len(userID) == 0 {
		return visitor, errors.New("empty id")
	}

	svc.Db = svc.openDBConnection()
	query := fmt.Sprintf(common.GetVisitorByUserID, userID)
	result, err := svc.Db.Query(query)
	if err != nil {
		return visitor, err
	}
	defer svc.Db.Close()
	for result.Next() {
		errs := result.Scan(
			&visitor.VisitorID,
			&visitor.UserID,
			&visitor.FirstName,
			&visitor.MiddleName,
			&visitor.LastName,
			&visitor.Address,
			&visitor.BirthPlace,
			&visitor.BirthDate,
			&visitor.ContactNumber,
			&visitor.LastModifiedDate,
			&visitor.CreatedDate,
		)

		if errs != nil {
			log.Println("mysql_visitor:GetVisitorById() -> Error in Scanning")
			err = errs
			break
		}
	}
	defer result.Close()

	return visitor, err
}

func (svc *mySqlService) DeleteVisitor(id string) (bool, error) {
	isDeleted := true
	var err error

	if len(id) == 0 {
		return isDeleted, errors.New("empty id")
	}

	svc.Db = svc.openDBConnection()
	query := fmt.Sprintf(common.DeleteVisitor, id)
	_, err = svc.Db.Query(query)
	if err != nil {
		isDeleted = false
		return isDeleted, err
	}

	return isDeleted, err
}
