package mysqllocal

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/MakMoinee/go-mith/pkg/encrypt"
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
	DeleteUser(id string) (bool, error)
	UpdateUser(user models.Users) (bool, error)
	LogUser(user models.Users) (bool, models.Users, error)
	GetUserById(userId string) (models.Users, error)
	GetUserVisitor(userVisitor models.UserVisitor) (models.UserVisitor, error)

	CreateVisitor(visitor models.Visitor) (bool, error)
	GetVisitors() ([]models.Visitor, error)
	GetVisitorById(id string) (models.Visitor, error)
	DeleteVisitor(id string) (bool, error)

	GetInmates() ([]models.Inmates, error)
	CreateInmate(inmate models.Inmates) (bool, error)

	CreateAdmin(user models.Users) (bool, error)
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

func (svc *mySqlService) LogUser(user models.Users) (bool, models.Users, error) {
	log.Println("Inside mysql:LogUser()")
	userLogin := false
	users := models.Users{}
	usersList := []models.Users{}
	svc.Db = svc.openDBConnection()
	query := fmt.Sprintf(common.LogUserQuery, user.UserName)
	result, err := svc.Db.Query(query)
	if err != nil {
		log.Println(err.Error())
		return userLogin, user, err
	}
	defer svc.Db.Close()

	for result.Next() {
		err := result.Scan(&users.UserID, &users.UserName, &users.UserPassword, &users.UserType)
		if err != nil {
			log.Println(err.Error())
			return userLogin, user, err
		}
		usersList = append(usersList, users)
	}
	defer result.Close()

	userLogin, list := svc.checkLoggingUser(usersList, user)

	return userLogin, list, nil
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
		err := result.Scan(&users.UserID, &users.UserName, &users.UserPassword, &users.UserType)
		if err != nil {
			log.Println(err.Error())
			return usersList, err
		}
		usersList = append(usersList, users)
	}
	defer result.Close()

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
	defer svc.Db.Close()
	return userCreated, nil
}

func (svc *mySqlService) DeleteUser(id string) (bool, error) {
	log.Println("Inside mysql: DeleteUser()")
	isDeleted := true
	var err error
	svc.Db = svc.openDBConnection()
	query := fmt.Sprintf(common.DeleteUserQuery, id)
	_, err = svc.Db.Query(query)
	if err != nil {
		log.Println("mysql:DeleteUser() -> " + err.Error())
		isDeleted = false
	}
	defer svc.Db.Close()
	return isDeleted, err
}

func (svc *mySqlService) UpdateUser(user models.Users) (bool, error) {
	log.Println("Inside mysql: UpdateUser()")
	isUpdated := true
	svc.Db = svc.openDBConnection()
	query := fmt.Sprintf(common.UpdateUserQuery, user.UserName, user.UserPassword, user.UserID)
	_, err := svc.Db.Query(query)
	if err != nil {
		log.Println("mysql:UpdateUser() -> " + err.Error())
		isUpdated = false
	}
	defer svc.Db.Close()
	return isUpdated, err
}

func (svc *mySqlService) openDBConnection() *sql.DB {
	db, err := sql.Open(svc.DbDriver, svc.ConnectionString)
	if err != nil {
		log.Println(err.Error())
	}
	return db
}

func (svc *mySqlService) checkLoggingUser(dbUser []models.Users, currentUser models.Users) (bool, models.Users) {
	log.Println("Inside CheckLoggingUser()")
	isPresent := false
	user := models.Users{}
	if len(dbUser) > 0 {
		for _, data := range dbUser {
			isMatch := encrypt.CheckPasswordHash(currentUser.UserPassword, data.UserPassword)
			if isMatch && strings.EqualFold(currentUser.UserName, data.UserName) {
				isPresent = true
				user = data
				break
			}
		}
	}
	return isPresent, user
}

func (svc *mySqlService) GetUserVisitor(requestUserVisitor models.UserVisitor) (models.UserVisitor, error) {
	log.Println("Inside GetUserVisitor ...")
	userVisitor := models.UserVisitor{}
	var err error
	svc.Db = svc.openDBConnection()
	query := fmt.Sprintf(common.GetUserVisitorQuery, requestUserVisitor.FirstName, requestUserVisitor.LastName, requestUserVisitor.MiddleName, requestUserVisitor.UserName)
	result, err := svc.Db.Query(query)
	defer svc.Db.Close()
	defer result.Close()
	for result.Next() {
		err := result.Scan(
			&userVisitor.UserID,
			&userVisitor.UserName,
			&userVisitor.FirstName,
			&userVisitor.LastName,
			&userVisitor.MiddleName,
			&userVisitor.BirthPlace,
		)
		if err != nil {
			log.Println(err.Error())
			return userVisitor, err
		}
	}

	return userVisitor, err
}

func (svc *mySqlService) GetUserById(userId string) (models.Users, error) {
	users := models.Users{}
	var err error

	svc.Db = svc.openDBConnection()
	query := fmt.Sprintf(common.SelectUserByIDQuery, userId)
	result, err := svc.Db.Query(query)
	if err != nil {
		log.Println(err.Error())
		return users, err
	}
	defer svc.Db.Close()

	for result.Next() {
		err = result.Scan(&users.UserID, &users.UserName, &users.UserPassword, &users.UserType)
		if err != nil {
			log.Println(err.Error())
			return users, err
		}
	}
	defer result.Close()
	return users, err
}
