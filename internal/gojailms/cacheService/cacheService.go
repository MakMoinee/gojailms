package cacheservice

import (
	"sync"

	"github.com/MakMoinee/gojailms/internal/gojailms/models"
	"github.com/MakMoinee/gojailms/internal/repository/mysqllocal"
)

type cacheService struct {
	Mysql   mysqllocal.MysqlIntf
	UserMap map[int]models.Users
}

type CacheIntf interface {
	LoadCache() error
	GetUserMap() map[int]models.Users
}

func NewCacheService() CacheIntf {
	svc := cacheService{}
	svc.Set()
	return &svc
}

func (svc *cacheService) Set() {
	svc.Mysql = mysqllocal.NewUserMySqlService()
	svc.LoadCache()
}

func (svc *cacheService) LoadCache() error {
	usersList := []models.Users{}
	var err error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		usersList, err = svc.Mysql.GetUsers()
	}()
	wg.Wait()
	svc.UserMap = make(map[int]models.Users)
	for _, user := range usersList {
		svc.UserMap[user.UserID] = user
	}

	return err
}

func (svc *cacheService) GetUserMap() map[int]models.Users {
	return svc.UserMap
}
