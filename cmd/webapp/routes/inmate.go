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

func (svc *routesHandler) CreateInmate(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside routes:CreateInmate()")
	inmate := models.Inmates{}
	byteBody, err := ioutil.ReadAll(r.Body)
	errorBuilder := response.ErrorResponse{}
	if err != nil {
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	err = json.Unmarshal(byteBody, &inmate)
	if err != nil {
		log.Println("routes:CreateInmate() -> Unmarshal Error")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	//validate request
	err = service.ValidateInmateRequest(inmate)
	if err != nil {
		log.Println("routes:CreateInmate() -> Invalid Paramaters")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	isCreated, err := svc.JailMs.CreateInmate(inmate)
	if err != nil {
		log.Println("routes:CreateInmate() -> Error Creating Inmate Record")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	if !isCreated {
		log.Println("routes:CreateInmate() -> Failed to Create Inmate Record")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	response.Success(w, "Successfully Created Record")
}

func (svc *routesHandler) GetInmates(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside routes:GetInmates()")
	errorBuilder := response.ErrorResponse{}
	inmatesList, err := svc.JailMs.GetInmates()
	if err != nil {
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	if len(inmatesList) == 0 {
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = "No Inmates Recorded"
		response.Error(w, errorBuilder)
		return
	}

	response.Success(w, inmatesList)

}
