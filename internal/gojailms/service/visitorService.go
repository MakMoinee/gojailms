package service

import (
	"log"
	"sync"

	"github.com/MakMoinee/gojailms/internal/gojailms/models"
	"github.com/MakMoinee/gojailms/internal/repository/mysqllocal"
)

func SendCreateVisitor(visitor models.Visitor, mysql mysqllocal.MysqlIntf) (bool, error) {
	log.Println("Inside visitorService:SendCreateVisitor()")
	isCreated := false
	var err error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		isCreated, err = mysql.CreateVisitor(visitor)
	}()
	wg.Wait()

	return isCreated, err
}

func SendGetVisitors(mysql mysqllocal.MysqlIntf) ([]models.Visitor, error) {
	log.Println("Inside visitorService:SendGetVisitors()")
	visitorsList := []models.Visitor{}
	var err error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		visitorsList, err = mysql.GetVisitors()
	}()
	wg.Wait()

	return visitorsList, err
}

func SendGetVisitorByUserID(userID string, mysql mysqllocal.MysqlIntf) (models.Visitor, error) {
	log.Println("Inside visitorService:SendGetVisitorByUserID()")
	visitor := models.Visitor{}
	var err error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		visitor, err = mysql.GetVisitorById(userID)
	}()
	wg.Wait()

	return visitor, err
}

func SendDeleteVisitor(visitorID string, mysql mysqllocal.MysqlIntf) (bool, error) {
	log.Println("Inside visitorService:SendDeleteVisitor()")
	isDeleted := false
	var err error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		isDeleted, err = mysql.DeleteVisitor(visitorID)
	}()
	wg.Wait()

	return isDeleted, err
}

func SendUpdateVisitor(visitor models.Visitor, mysql mysqllocal.MysqlIntf) (bool, error) {
	log.Println("Inside visitorService:SendUpdateVisitor()")
	return mysql.UpdateVisitor(visitor)
}
