package gojailms

import (
	"log"
	"sync"

	"github.com/MakMoinee/gojailms/internal/gojailms/models"
	"github.com/MakMoinee/gojailms/internal/repository/mysqllocal"
)

type JailMs struct {
	MySqlService mysqllocal.MysqlIntf
}

type JailIntf interface {
	GetUsers() ([]models.Users, error)
}

func NewJailMs() JailIntf {
	svc := JailMs{}
	svc.set()
	return &svc
}

func (svc *JailMs) set() {
	svc.MySqlService = mysqllocal.NewUserMySqlService()
}

func (svc *JailMs) GetUsers() ([]models.Users, error) {
	log.Println("Inside gojailms:GetUsers()")
	usersList := []models.Users{}
	var err error
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		usersList, err = svc.MySqlService.GetUsers()
	}()
	wg.Wait()

	if err != nil {
		log.Println("Error in retrieving users")
		return []models.Users{}, err
	}

	return usersList, nil
}
