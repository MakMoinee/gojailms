package service

import (
	"errors"
	"log"
	"sync"

	"github.com/MakMoinee/gojailms/internal/gojailms/models"
	"github.com/MakMoinee/gojailms/internal/repository/mysqllocal"
)

func SendCreateInmate(inmate models.Inmates, mysql mysqllocal.MysqlIntf) (bool, error) {
	log.Println("Inside mysql:SendCreateInmate()")
	isUpdated := false
	var err error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		isUpdated, err = mysql.CreateInmate(inmate)
	}()
	wg.Wait()

	return isUpdated, err
}

func SendGetInmates(mysql mysqllocal.MysqlIntf) ([]models.Inmates, error) {
	log.Println("Inside service:SendGetInmates()")
	inmatesList := []models.Inmates{}
	var err error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		inmatesList, err = mysql.GetInmates()
	}()
	wg.Wait()

	return inmatesList, err
}

func ValidateInmateRequest(inmate models.Inmates) error {
	var err error
	if len(inmate.FirstName) == 0 ||
		len(inmate.LastName) == 0 ||
		inmate.CrimeID == 0 ||
		len(inmate.Address) == 0 ||
		len(inmate.BirthDate) == 0 ||
		len(inmate.BirthPlace) == 0 {
		err = errors.New("invalid or empty required paramaters")
	}
	return err
}
