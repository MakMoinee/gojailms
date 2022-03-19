package common

import "time"

const (
	GetUsersPath   = "/get/jms/users"
	CreateUserPath = "/create/jms/user"
	DeleteUserPath = "/delete/jms/user"

	CreateVisitorPath = "/create/jms/visitor"
	GetVisitorsPath   = "/get/jms/visitors"
	ContentTypeKey    = "Content-Type"
	ContentTypeValue  = "application/json; charset=UTF-8"
	TimeFormat        = "2006-01-02 15:04:05"

	GetUsersQuery   = "SELECT * FROM users;"
	InsertUserQuery = "INSERT INTO users (userName,password,userType) VALUES ('%v','%v',2);"
	DeleteUserQuery = "DELETE FROM users where userID=%v;"

	CreateVisitorQuery = "INSERT INTO visitors (userID,firstName,lastName,middleName,address,birthPlace,birthDate,lastModifiedDate,createdDate) VALUES(%v,'%v','%v','%v','%v','%v','%v',NOW(),NOW());"
	GetVisitorsQuery   = "SELECT * FROM visitors;"
)

var (
	SERVER_PORT             string
	SERVER_ENABLE_PROFILING bool
	DB_NAME                 string
	DB_DRIVER               string
	MYSQL_USERNAME          string
	MYSQL_PASSWORD          string
	CONNECTION_STRING       string
	RETRY_SLEEP             time.Duration
)
