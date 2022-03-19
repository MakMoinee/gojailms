package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MakMoinee/gojailms/cmd/webapp/response"
	"github.com/MakMoinee/gojailms/internal/gojailms/models"
	"github.com/MakMoinee/gojailms/internal/gojailms/service"
)

func (svc *routesHandler) CreateVisitor(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside routes->visitor:CreateVisitor()")
	visitor := models.Visitor{}
	errorBuilder := response.ErrorResponse{}
	byteBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	err = json.Unmarshal(byteBody, &visitor)
	if err != nil {
		log.Println("routes:CreateVisitor() -> Unmarshal Error")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	err = service.ValidateVisitorRequest(visitor)
	if err != nil {
		log.Println("routes:CreateVisitor() -> Invalid Parameters")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	isCreated, err := svc.JailMs.CreateVisitor(visitor)
	if err != nil {
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	if !isCreated {
		log.Println("routes:CreateVisitor() -> Fail to Create User")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = "Failed to Create User"
		response.Error(w, errorBuilder)
		return
	}

	response.Success(w, "Successfully Created User")
}

func (svc *routesHandler) GetVisitors(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside routes->visitor:GetVisitors()")
	errorBuilder := response.ErrorResponse{}
	visitorsList, err := svc.JailMs.GetVisitors()
	if err != nil {
		log.Println("routes:GetVisitors() -> Error in Getting the Visitors")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}
	if len(visitorsList) == 0 {
		log.Println("routes:GetVisitors() -> Empty Visitors")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = "No Visitors Recorded"
		response.Error(w, errorBuilder)
		return
	}

	response.Success(w, visitorsList)

}
