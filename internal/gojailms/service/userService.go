package service

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/MakMoinee/go-mith/pkg/strings"
	"github.com/MakMoinee/gojailms/internal/common"
	"github.com/MakMoinee/gojailms/internal/gojailms/models"
	"github.com/MakMoinee/gojailms/internal/repository/mysqllocal"
)

func SendCreateUser(user models.Users, mysql mysqllocal.MysqlIntf) (bool, error) {
	retries := 3
	isUserCreated := false
	var err error
	var wg sync.WaitGroup
	for retries > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			isUserCreated, err = mysql.CreateUser(user)
		}()
		wg.Wait()
		if err == nil {
			break
		}
		retries--
		log.Println("Retry attempt " + fmt.Sprintf("%v", retries))
		time.Sleep(common.RETRY_SLEEP)
	}
	return isUserCreated, err
}

func ValidateUserRequest(user models.Users) error {
	var err error
	if !strings.IsStringEmpty(user.UserName) || !strings.IsStringEmpty(user.UserPassword) {
		err = errors.New("empty username or password")
	}
	return err
}
