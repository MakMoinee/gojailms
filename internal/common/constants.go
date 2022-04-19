package common

import "time"

const (
	GetUsersPath   = "/get/jms/users"
	CreateUserPath = "/create/jms/user"
	DeleteUserPath = "/delete/jms/user"
	UpdateUserPath = "/update/jms/user"
	LogUserPath    = "/log/jms/user"

	CreateVisitorPath = "/create/jms/visitor"
	GetVisitorsPath   = "/get/jms/visitors"

	CreateInmatePath = "/create/jms/inmate"
	GetInmatePath    = "/get/jms/inmate"

	ContentTypeKey   = "Content-Type"
	ContentTypeValue = "application/json; charset=UTF-8"
	TimeFormat       = "2006-01-02 15:04:05"

	GetUsersQuery   = "SELECT * FROM users;"
	LogUserQuery    = "SELECT * FROM users where userName='%v';"
	InsertUserQuery = "INSERT INTO users (userName,password,userType) VALUES ('%v','%v',2);"
	DeleteUserQuery = "DELETE FROM users where userID=%v;"
	UpdateUserQuery = "UPDATE users SET userName='%v',password='%v' where userID=%v;"

	CreateVisitorQuery = "INSERT INTO visitors (userID,firstName,lastName,middleName,address,birthPlace,birthDate,lastModifiedDate,createdDate) VALUES(%v,'%v','%v','%v','%v','%v','%v',NOW(),NOW());"
	GetVisitorsQuery   = "SELECT * FROM visitors;"

	GetInmatesQuery   = "SELECT * FROM inmates;"
	CreateInmateQuery = "INSERT INTO inmates (crimeID,firstName,lastName,middleName,address,birthPlace,birthDate,lastModifiedDate,createdDate) VALUES (%v,'%v','%v','%v','%v','%v','%v',NOW(),NOW());"
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
	SERVICE_VERSION         string
)
