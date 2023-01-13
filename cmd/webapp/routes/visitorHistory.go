package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MakMoinee/go-mith/pkg/response"
	"github.com/MakMoinee/gojailms/internal/gojailms/models"
)

func (svc *routesHandler) InsertVisitorHistory(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside routes()->visitorHistory:InsertVisitorHistory()")
	visitorHistory := models.VisitorHistoryRequest{}
	byteBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error in routes:InsertVisitorHistory() -> Error reading the body")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}
	err = json.Unmarshal(byteBody, &visitorHistory)
	if err != nil {
		log.Println("Error in routes:InsertVisitorHistory() -> Error unmarshalling the body")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	isSaved, err := svc.JailMs.InsertVisitorHistory(visitorHistory)
	if err != nil {
		log.Println("Error in routes:InsertVisitorHistory() -> Error inserting visitor history")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	if !isSaved {
		errResponse := response.NewErrorBuilder("Failed to insert visitor history", http.StatusInternalServerError)
		response.Response(w, errResponse)
		return
	}

	successResponse := response.SuccessResponse{Message: "Successfully inserted visitor history"}
	response.Success(w, successResponse)
}

func (svc *routesHandler) GetAllVisitorHistory(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside routes->visitor:GetAllVisitorHistory()")
	errorBuilder := response.ErrorResponse{}
	searchKey := r.URL.Query().Get("search")
	if len(searchKey) == 0 {

	}
	list, err := svc.JailMs.GetAllVisitorHistory()
	if err != nil {
		log.Println("routes:GetAllVisitorHistory() -> Error in Getting the Visitor History")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}
	if len(list) == 0 {
		log.Println("routes:GetAllVisitorHistory() -> Empty Visitor History")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = "No Visitor History Recorded"
		response.Error(w, errorBuilder)
		return
	}

	response.Success(w, list)

}
