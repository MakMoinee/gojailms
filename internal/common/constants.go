package common

import "time"

const (
	GetUsersPath          = "/get/jms/users"
	CreateUserPath        = "/create/jms/user"
	DeleteUserPath        = "/delete/jms/user"
	UpdateUserPath        = "/update/jms/user"
	LogUserPath           = "/log/jms/user"
	CreateAdminPath       = "/create/admin/user"
	UpdateUserVisitorPath = "/forgot/user/pass"

	CreateVisitorPath       = "/create/jms/visitor"
	GetVisitorsPath         = "/get/jms/visitors"
	GetVisitorsByUserIDPath = "/get/jms/visitor"
	DeleteVisitorPath       = "/delete/jms/visitor"

	CreateInmatePath         = "/create/jms/inmate"
	GetInmatePath            = "/get/jms/inmate"
	InsertVisitorHistoryPath = "/inser/visitor/history"

	ContentTypeKey   = "Content-Type"
	ContentTypeValue = "application/json; charset=UTF-8"
	TimeFormat       = "2006-01-02 15:04:05"

	GetUsersQuery             = "SELECT * FROM users;"
	LogUserQuery              = "SELECT * FROM users where userName='%v';"
	InsertUserQuery           = "INSERT INTO users (userName,password,userType) VALUES ('%v','%v',2);"
	InsertAdminUserQuery      = "INSERT INTO users (userName,password,userType) VALUES ('%v','%v',1);"
	DeleteUserQuery           = "DELETE FROM users where userID=%v;"
	UpdateUserQuery           = "UPDATE users SET userName='%v',password='%v' where userID=%v;"
	SelectUserByIDQuery       = "SELECT * FROM users where userID=%v;"
	InsertVisitorHistoryQuery = "INSERT INTO visitorhistory (visitorID,remarks,visitedDateTime) VALUES(%v,'%v',NOW())"
	GetUserVisitorQuery       = "SELECT * FROM vwuservisitor where firstName='%v' and lastName='%v' and middleName='%v' and userName='%v' LIMIT 1;"

	CreateVisitorQuery = "INSERT INTO visitors (userID,firstName,lastName,middleName,address,birthPlace,birthDate,contactNumber,lastModifiedDate,createdDate) VALUES(%v,'%v','%v','%v','%v','%v','%v','%v',NOW(),NOW());"
	GetVisitorsQuery   = "SELECT * FROM visitors;"
	GetVisitorByUserID = "SELECT * FROM visitors where userID=%v;"
	DeleteVisitor      = "DELETE FROM visitors where visitorID=%v;"

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
	AUTH_TOKEN              string
)
