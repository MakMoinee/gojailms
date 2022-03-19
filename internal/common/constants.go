package common

import "time"

const (
	GetUsersPath     = "/get/jms/users"
	CreateUserPath   = "/create/jms/user"
	ContentTypeKey   = "Content-Type"
	ContentTypeValue = "application/json; charset=UTF-8"
	TimeFormat       = "2006-01-02 15:04:05"

	GetUsersQuery   = "SELECT * FROM users;"
	InsertUserQuery = "INSERT INTO users (userName,password,userType) VALUES ('%v','%v',2);"
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
