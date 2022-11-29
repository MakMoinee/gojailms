package service

import (
	"log"
	"strings"

	"github.com/MakMoinee/gojailms/internal/gojailms/models"
	"github.com/MakMoinee/gojailms/internal/repository/mysqllocal"
)

func SendInsertVisitorHistory(visitorHistory models.VisitorHistoryRequest, mysql mysqllocal.MysqlIntf) (bool, error) {
	log.Println("Inside service:SendInsertVisitorHistory()")
	visitorID := visitorHistory.VisitorID
	remarks := visitorHistory.Remarks
	remarks = strings.ReplaceAll(remarks, "'", "\\'")
	return mysql.InsertVisitorHistory(visitorID, remarks)
}
