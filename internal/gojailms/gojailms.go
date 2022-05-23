package gojailms

import (
	"log"
	"sync"

	"github.com/MakMoinee/gojailms/internal/gojailms/models"
	"github.com/MakMoinee/gojailms/internal/gojailms/service"
	"github.com/MakMoinee/gojailms/internal/repository/mysqllocal"
)

type JailMs struct {
	MySqlService mysqllocal.MysqlIntf
}

type JailIntf interface {
	GetUsers() ([]models.Users, error)
	CreateUser(user models.Users, visitor models.Visitor) (bool, error)
	DeleteUser(id string) (bool, error)
	UpdateUser(user models.Users) (bool, error)
	LogUser(user models.Users) (bool, models.Users, error)
	CreateAdminUser(user models.Users) (bool, error)
	UpdateUserVisitor(userVisitor models.UserVisitor) (bool, error)

	CreateVisitor(visitor models.Visitor) (bool, error)
	GetVisitors() ([]models.Visitor, error)
	GetVisitorByUserID(userid string) (models.Visitor, error)
	DeleteVisitor(id string) (bool, error)

	CreateInmate(inmate models.Inmates) (bool, error)
	GetInmates() ([]models.Inmates, error)
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

func (svc *JailMs) LogUser(user models.Users) (bool, models.Users, error) {
	log.Println("Inside gojailms:LogUser()")
	isValidUser, validUser, err := service.SendLogUser(user, svc.MySqlService)
	return isValidUser, validUser, err
}

func (svc *JailMs) CreateUser(user models.Users, visitor models.Visitor) (bool, error) {
	log.Println("Inside gojailms:CreateUser()")
	userCreated, err := service.SendCreateUser(user, visitor, svc.MySqlService)
	return userCreated, err
}

func (svc *JailMs) DeleteUser(id string) (bool, error) {
	log.Println("Inside gojailms:DeleteUser()")
	isDeleted, err := service.SendDeleteUser(id, svc.MySqlService)
	return isDeleted, err
}

func (svc *JailMs) UpdateUser(user models.Users) (bool, error) {
	log.Println("Inside gojailms:UpdateUser()")
	isUpdated, err := service.SendUpdateUser(user, svc.MySqlService)
	return isUpdated, err
}

func (svc *JailMs) CreateVisitor(visitor models.Visitor) (bool, error) {
	log.Println("Inside gojailms:CreateVisitor()")
	isCreated, err := service.SendCreateVisitor(visitor, svc.MySqlService)
	return isCreated, err
}

func (svc *JailMs) GetVisitors() ([]models.Visitor, error) {
	log.Println("Inside gojailms:GetVisitors()")
	visitorsList, err := service.SendGetVisitors(svc.MySqlService)
	return visitorsList, err
}

func (svc *JailMs) CreateInmate(inmate models.Inmates) (bool, error) {
	log.Println("Inside gojailms:CreateInmate()")
	isUpdated, err := service.SendCreateInmate(inmate, svc.MySqlService)
	return isUpdated, err
}

func (svc *JailMs) GetInmates() ([]models.Inmates, error) {
	log.Println("Inside gojailms:GetInmates()")
	list, err := service.SendGetInmates(svc.MySqlService)
	return list, err
}

func (svc *JailMs) DeleteVisitor(id string) (bool, error) {
	log.Println("Inside gojailms:DeleteVisitor()")
	isDeleted, err := service.SendDeleteVisitor(id, svc.MySqlService)
	return isDeleted, err
}

func (svc *JailMs) GetVisitorByUserID(userid string) (models.Visitor, error) {
	log.Println("Inside gojailms:GetVisitorByUserID()")
	visitor, err := service.SendGetVisitorByUserID(userid, svc.MySqlService)
	return visitor, err
}

func (svc *JailMs) CreateAdminUser(user models.Users) (bool, error) {
	log.Println("Inside gojailms:CreateAdminUser()")
	isCreated, err := service.SendCreateAdminUser(user, svc.MySqlService)
	return isCreated, err
}

func (svc *JailMs) UpdateUserVisitor(userVisitor models.UserVisitor) (bool, error) {
	log.Println("Inside gojailms:UpdateUserVisitor()")
	isUpdated, err := service.SendUpdateUserVisitor(userVisitor, svc.MySqlService)
	return isUpdated, err
}
