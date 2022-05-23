package service

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
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

func SendCreateAdminUser(user models.Users, mysql mysqllocal.MysqlIntf) (bool, error) {
	log.Println("Inside service:SendCreateAdminUser()")

	isValid, err := mysql.CreateAdmin(user)
	return isValid, err
}

func SendUpdateUserVisitor(userVisitor models.UserVisitor, mysql mysqllocal.MysqlIntf) (bool, error) {
	log.Println("Inside service:SendUpdateUserVisitor() ...")
	isUpdated := false
	var err error
	var wg sync.WaitGroup

	newUserVisitor := models.UserVisitor{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		newUserVisitor, err = mysql.GetUserVisitor(userVisitor)
	}()
	wg.Wait()

	if err != nil {
		return isUpdated, err
	}

	if newUserVisitor.UserID == 1 && len(newUserVisitor.Token) == 0 && !strings.EqualFold(common.AUTH_TOKEN, newUserVisitor.Token) {
		log.Println("Not Authorized")
		return isUpdated, errors.New("not authorized")
	}

	if reflect.DeepEqual(newUserVisitor, models.UserVisitor{}) {
		log.Println("Inside service:SendUpdateUserVisitor() -> Error: Empty Result")
		return isUpdated, errors.New("empty result")
	}

	userVisitor.UserID = newUserVisitor.UserID
	isValid := ValidateUserVisitor(userVisitor)
	if !isValid {
		log.Println("Inside service:SendUpdateUserVisitor() -> Error: Not Valid Parameters")
		return isUpdated, errors.New("not valid parameters")
	}

	user := models.Users{}
	user.UserID = userVisitor.UserID
	user.UserName = userVisitor.UserName
	user.UserType = 2
	hashPass, _ := encrypt.HashPassword(userVisitor.UserPassword)
	user.UserPassword = hashPass
	isUpdated, err = mysql.UpdateUser(user)

	return isUpdated, err
}

func ValidateUserVisitor(userVisitor models.UserVisitor) bool {
	log.Println("Inside service:ValidateUserVisitor() ...")
	isValid := false
	isUserValid := false
	isVisitorValid := false

	//check user
	if len(userVisitor.UserName) != 0 && userVisitor.UserID > 0 && len(userVisitor.UserPassword) > 0 {
		isUserValid = true
	}

	//check visitor
	if len(userVisitor.FirstName) != 0 && len(userVisitor.LastName) > 0 && len(userVisitor.MiddleName) > 0 && len(userVisitor.BirthPlace) > 0 {
		isVisitorValid = true
	}

	isValid = (isUserValid == isVisitorValid)

	return isValid
}
