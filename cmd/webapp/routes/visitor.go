package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MakMoinee/go-mith/pkg/response"
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

func (svc *routesHandler) DeleteVisitor(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside routes:DeleteVisitor()")
	query := r.URL.Query()
	visitorID := query.Get("id")
	errorBuilder := response.ErrorResponse{}

	if len(visitorID) == 0 {
		errorBuilder.ErrorMessage = "Missing Required Parameters"
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		response.Error(w, errorBuilder)
		return
	}

	isDeleted, err := svc.JailMs.DeleteVisitor(visitorID)
	if err != nil {
		errorBuilder.ErrorMessage = err.Error()
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		response.Error(w, errorBuilder)
		return
	}

	if !isDeleted {
		errorBuilder.ErrorMessage = "Failed to Delete Visitor"
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		response.Error(w, errorBuilder)
		return
	}
	response.Success(w, "Successfully Deleted Visitor")
}

func (svc *routesHandler) GetVisitorByUserID(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside routes:GetVisitorByUserID()")
	errorBuilder := response.ErrorResponse{}
	query := r.URL.Query()
	userID := query.Get("uid")

	if len(userID) == 0 {
		log.Println("Error in routes:GetVisitorByUserID() -> Empty Userid passed")
		errorBuilder.ErrorMessage = "Missing Required Parameters"
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		response.Error(w, errorBuilder)
		return
	}

	visitor, err := svc.JailMs.GetVisitorByUserID(userID)
	if err != nil {
		log.Println("Error in routes:GetVisitorByUserID() -> Erorr in getting visitor by userid")
		errorBuilder.ErrorMessage = err.Error()
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		response.Error(w, errorBuilder)
		return
	}
	response.Success(w, visitor)
}
