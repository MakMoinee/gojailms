package service

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/MakMoinee/go-mith/pkg/encrypt"
	"github.com/MakMoinee/gojailms/internal/common"
	"github.com/MakMoinee/gojailms/internal/gojailms/models"
	"github.com/MakMoinee/gojailms/internal/repository/mysqllocal"
)

func SendCreateUser(user models.Users, visitor models.Visitor, mysql mysqllocal.MysqlIntf) (bool, error) {
	retries := 3
	isUserCreated := false
	userTemp := user
	var err error
	var wg sync.WaitGroup
	hashPas, _ := encrypt.HashPassword(userTemp.UserPassword)
	userTemp.UserPassword = hashPas
	for retries > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			isUserCreated, err = mysql.CreateUser(userTemp)
		}()
		wg.Wait()
		if err == nil {
			break
		}
		retries--
		log.Println("Retry attempt " + fmt.Sprintf("%v", retries))
		time.Sleep(common.RETRY_SLEEP)
	}

	if isUserCreated {
		isValidUser := false
		isVisitorCreated := false
		validUser := models.Users{}

		wg.Add(1)
		go func() {
			defer wg.Done()
			isValidUser, validUser, err = SendLogUser(user, mysql)
		}()
		wg.Wait()

		if isValidUser {
			wg.Add(1)
			go func() {
				defer wg.Done()
				visitor.UserID = validUser.UserID
				isVisitorCreated, err = SendCreateVisitor(visitor, mysql)
			}()
			wg.Wait()
		}

		isUserCreated = isVisitorCreated

	}

	return isUserCreated, err
}

func SendDeleteUser(id string, mysql mysqllocal.MysqlIntf) (bool, error) {
	log.Println("Inside userService:SendDeleteUser()")
	isDeleted := false
	var err error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		isDeleted, err = mysql.DeleteUser(id)
	}()
	wg.Wait()
	return isDeleted, err
}

func SendUpdateUser(user models.Users, mysql mysqllocal.MysqlIntf) (bool, error) {
	log.Println("Inside userService:SendUpdateUser()")
	isUpdated := false
	var err error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		isUpdated, err = mysql.UpdateUser(user)
	}()
	wg.Wait()
	return isUpdated, err
}

func SendLogUser(user models.Users, mysql mysqllocal.MysqlIntf) (bool, models.Users, error) {
	log.Println("Inside userService:SendLogUser()")
	isValidUser := false
	validUser := models.Users{}
	var err error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		isValidUser, validUser, err = mysql.LogUser(user)
	}()
	wg.Wait()

	return isValidUser, validUser, err
}

func ValidateUserRequest(user models.Users) error {
	var err error
	if len(user.UserName) == 0 || len(user.UserPassword) == 0 {
		err = errors.New("empty username or password")
	}
	return err
}

func ValidateUpdateUserRequest(user models.Users) error {
	var err error
	if len(user.UserName) == 0 || len(user.UserPassword) == 0 || user.UserID == 0 {
		err = errors.New("empty username or password or userid")
	}
	return err
}

func ValidateVisitorRequest(visitor models.Visitor) error {
	var err error

	if len(visitor.FirstName) == 0 || len(visitor.LastName) == 0 || len(visitor.Address) == 0 || len(visitor.BirthPlace) == 0 || len(visitor.BirthDate) == 0 {
		err = errors.New("invalid or empty required parameters")
	}

	return err
}
