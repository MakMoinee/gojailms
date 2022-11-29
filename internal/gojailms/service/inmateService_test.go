package service

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/MakMoinee/gojailms/internal/gojailms/models"
	"github.com/MakMoinee/gojailms/internal/repository/mysqllocal"
)

var mockInmate = models.Inmates{
	CrimeID:    1,
	FirstName:  "Sample",
	LastName:   "Sample",
	MiddleName: "X",
	Address:    "Purok4A Poblacion Valencia City, Bukidnon",
	BirthPlace: "Zamboanga City",
	BirthDate:  "1998-02-01",
}

type mockMySqlService struct {
}

func (svc *mockMySqlService) CreateUser(user models.Users) (bool, error) {
	return false, nil
}

func (svc *mockMySqlService) GetUsers() ([]models.Users, error) {
	return []models.Users{}, nil
}

func (svc *mockMySqlService) LogUser(user models.Users) (bool, models.Users, error) {
	return true, models.Users{}, nil
}

func (svc *mockMySqlService) DeleteUser(id string) (bool, error) {
	return false, nil
}

func (svc *mockMySqlService) UpdateUser(user models.Users) (bool, error) {
	return false, nil
}

func (svc *mockMySqlService) CreateVisitor(visitor models.Visitor) (bool, error) {
	return false, nil
}

func (svc *mockMySqlService) GetVisitors() ([]models.Visitor, error) {
	return []models.Visitor{}, nil
}
func (svc *mockMySqlService) DeleteVisitor(id string) (bool, error) {
	return false, nil
}
func (svc *mockMySqlService) GetVisitorById(id string) (models.Visitor, error) {
	return models.Visitor{}, nil
}

func (svc *mockMySqlService) GetInmates() ([]models.Inmates, error) {
	return []models.Inmates{mockInmate}, nil
}

func (svc *mockMySqlService) CreateAdmin(user models.Users) (bool, error) {
	return false, nil
}

func (svc *mockMySqlService) UpdateUserVisitor(userVisitor models.UserVisitor) (bool, error) {
	return false, nil
}

func (svc *mockMySqlService) GetUserVisitor(requestUserVisitor models.UserVisitor) (models.UserVisitor, error) {
	return models.UserVisitor{}, nil
}

func (svc *mockMySqlService) CreateInmate(inmate models.Inmates) (bool, error) {
	return true, nil
}

func (svc *mockMySqlService) GetUserById(userId string) (models.Users, error) {
	return models.Users{}, nil
}

func (svc *mockMySqlService) InsertVisitorHistory(remarks string, visitorId string) (bool, error) {
	return false, nil
}

func newMockMySqlService() mysqllocal.MysqlIntf {
	svc := mockMySqlService{}

	return &svc
}

func TestSendCreateInmate(t *testing.T) {
	type args struct {
		inmate models.Inmates
		mysql  mysqllocal.MysqlIntf
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Scenario1", args{models.Inmates{
			CrimeID:    1,
			FirstName:  "Sample",
			LastName:   "Sample",
			MiddleName: "X",
			Address:    "Purok4A Poblacion Valencia City, Bukidnon",
			BirthPlace: "Zamboanga City",
			BirthDate:  "1998-02-01",
		}, newMockMySqlService()}, true, false},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
				}
			}()
			got, err := SendCreateInmate(tt.args.inmate, tt.args.mysql)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendCreateInmate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SendCreateInmate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendGetInmates(t *testing.T) {
	type args struct {
		mysql mysqllocal.MysqlIntf
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Inmates
		wantErr bool
	}{
		{"Scenario1", args{newMockMySqlService()}, []models.Inmates{mockInmate}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SendGetInmates(tt.args.mysql)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendGetInmates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendGetInmates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateInmateRequest(t *testing.T) {
	type args struct {
		inmate models.Inmates
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Scenario1", args{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateInmateRequest(tt.args.inmate); (err != nil) != tt.wantErr {
				t.Errorf("ValidateInmateRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
