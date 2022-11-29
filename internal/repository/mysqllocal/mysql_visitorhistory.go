package mysqllocal

import (
	"fmt"
	"log"

	"github.com/MakMoinee/gojailms/internal/common"
)

func (svc *mySqlService) InsertVisitorHistory(remarks string, visitorId string) (bool, error) {
	log.Println("Inside mysqllocal:InsertVisitorHistory()")
	isSave := true
	svc.Db = svc.openDBConnection()
	query := fmt.Sprintf(common.InsertVisitorHistoryQuery, remarks, visitorId)
	_, err := svc.Db.Query(query)
	if err != nil {
		log.Println(err.Error())
		isSave = false
	}
	defer svc.Db.Close()

	return isSave, err
}
