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
