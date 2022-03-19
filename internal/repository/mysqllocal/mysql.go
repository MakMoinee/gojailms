package mysqllocal

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/MakMoinee/gojailms/internal/common"
	"github.com/MakMoinee/gojailms/internal/gojailms/models"
)

type mySqlService struct {
	DatabaseName     string
	DatabaseServer   string
	DatabaseUser     string
	DatabasePassword string
	ConnectionString string
	Db               *sql.DB
	DbDriver         string
}

type MysqlIntf interface {
	GetUsers() ([]models.Users, error)
	CreateUser(user models.Users) (bool, error)
}

func NewUserMySqlService() MysqlIntf {
	svc := mySqlService{}
	svc.Set()
	return &svc
}

func (svc *mySqlService) Set() {
	svc.DatabaseName = common.DB_NAME
	svc.DatabaseUser = common.MYSQL_USERNAME
	svc.DatabasePassword = common.MYSQL_PASSWORD
	svc.DbDriver = common.DB_DRIVER
	svc.ConnectionString = svc.DatabaseUser + ":" + svc.DatabasePassword + "@" + common.CONNECTION_STRING + svc.DatabaseName
	svc.Db = svc.openDBConnection()
	defer svc.Db.Close()
}

// GetUsers - retrieve users
func (svc *mySqlService) GetUsers() ([]models.Users, error) {
	usersList := []models.Users{}
	var err error

	return usersList, err
}

func (svc *mySqlService) CreateUser(user models.Users) (bool, error) {
	log.Println("Inside mysql:CreateUser()")
	userCreated := true

	svc.Db = svc.openDBConnection()
	query := fmt.Sprintf(common.InsertUserQuery, user.UserName, user.UserPassword)

	_, err := svc.Db.Query(query)
	if err != nil {
		log.Println("Error in inserting the record to db")
		userCreated = false
		return userCreated, err
	}
	return userCreated, nil
}

func (svc *mySqlService) openDBConnection() *sql.DB {
	db, err := sql.Open(svc.DbDriver, svc.ConnectionString)
	if err != nil {
		log.Println(err.Error())
	}
	return db
}
