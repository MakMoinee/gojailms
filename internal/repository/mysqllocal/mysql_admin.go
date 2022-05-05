package mysqllocal

import (
	"fmt"
	"log"

	"github.com/MakMoinee/go-mith/pkg/encrypt"
	"github.com/MakMoinee/gojailms/internal/common"
	"github.com/MakMoinee/gojailms/internal/gojailms/models"
)

func (svc *mySqlService) CreateAdmin(user models.Users) (bool, error) {
	log.Println("Inside mysql:CreateAdmin()")
	isCreated := true
	var err error
	hashPas, err := encrypt.HashPassword(user.UserPassword)
	if err != nil {
		return isCreated, err
	}
	query := fmt.Sprintf(common.InsertAdminUserQuery, user.UserName, hashPas)
	svc.Db = svc.openDBConnection()

	_, err = svc.Db.Query(query)
	defer svc.Db.Close()
	if err != nil {
		isCreated = false
		return isCreated, err
	}

	return isCreated, err
}
