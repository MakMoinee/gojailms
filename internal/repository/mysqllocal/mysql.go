package mysqllocal

import (
	"database/sql"
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
	users := models.Users{}
	var err error
	svc.Db = svc.openDBConnection()
	result, err := svc.Db.Query(common.GetUsersQuery)
	if err != nil {
		log.Println(err.Error())
		return usersList, err
	}

	defer svc.Db.Close()

	for result.Next() {
		err = result.Scan(&users.UserID, &users.UserName, &users.UserPassword, &users.UserType)
		if err != nil {
			log.Println(err.Error())
			break
		}
		usersList = append(usersList, users)
	}
	defer result.Close()

	return usersList, err
}

func (svc *mySqlService) openDBConnection() *sql.DB {
	db, err := sql.Open(svc.DbDriver, svc.ConnectionString)
	if err != nil {
		log.Println(err.Error())
	}
	return db
}
