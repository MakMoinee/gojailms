package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MakMoinee/go-mith/pkg/goserve"
	"github.com/MakMoinee/gojailms/cmd/webapp/response"
	"github.com/MakMoinee/gojailms/internal/common"
	"github.com/MakMoinee/gojailms/internal/gojailms"
	"github.com/MakMoinee/gojailms/internal/gojailms/models"
	"github.com/MakMoinee/gojailms/internal/gojailms/service"
	"github.com/go-chi/cors"
)

type routesHandler struct {
	JailMs gojailms.JailIntf
}

type RoutesIntf interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)

	CreateVisitor(w http.ResponseWriter, r *http.Request)
	GetVisitors(w http.ResponseWriter, r *http.Request)

	CreateInmate(w http.ResponseWriter, r *http.Request)
	GetInmates(w http.ResponseWriter, r *http.Request)
}

func newRoutes() RoutesIntf {
	svc := routesHandler{}
	svc.JailMs = gojailms.NewJailMs()
	return &svc
}

func Set(httpService *goserve.Service) {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "DELETE", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-TOKEN"},
		ExposedHeaders:   []string{"Link", "Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	newRoutesHandler := newRoutes()
	httpService.Router.Use(cors.Handler)
	initiateRoutes(httpService, newRoutesHandler)
}

// initiateRoutes initialize routes
func initiateRoutes(httpService *goserve.Service, handler RoutesIntf) {
	httpService.Router.Get(common.GetUsersPath, handler.GetUsers)
	httpService.Router.Post(common.CreateUserPath, handler.CreateUser)
	httpService.Router.Delete(common.DeleteUserPath, handler.DeleteUser)
	httpService.Router.Put(common.UpdateUserPath, handler.UpdateUser)

	//visitor
	httpService.Router.Post(common.CreateVisitorPath, handler.CreateVisitor)
	httpService.Router.Get(common.GetVisitorsPath, handler.GetVisitors)

	//inmate
	httpService.Router.Post(common.CreateInmatePath, handler.CreateInmate)
	httpService.Router.Get(common.GetInmatePath, handler.GetInmates)
}

func (svc *routesHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside Routes: GetUsers()")
	errorBuilder := response.ErrorResponse{}
	usersList, err := svc.JailMs.GetUsers()
	if err != nil {

		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	if len(usersList) == 0 {
		response.Success(w, "No Users Registered")
		return
	}

	response.Success(w, usersList)
}

func (svc *routesHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside Routes:CreateUser()")
	user := models.Users{}
	errorBuilder := response.ErrorResponse{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Routes:CreateUser() -> Reading the body error")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Routes:CreateUser() -> Unmarshal Error")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	// validate request
	err = service.ValidateUserRequest(user)
	if err != nil {
		log.Println("Routes:CreateUser() -> Invalid Parameters")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	isUserCreated, err := svc.JailMs.CreateUser(user)
	if err != nil {
		errorBuilder := response.ErrorResponse{}
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	if !isUserCreated {
		errorBuilder := response.ErrorResponse{}
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = "User Not Created"
		response.Error(w, errorBuilder)
		return
	}

	response.Success(w, "Successfully Created User")
}

func (svc *routesHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside Routes: DeleteUser()")
	query := r.URL.Query()
	userId := query.Get("id")
	errorBuilder := response.ErrorResponse{}
	if len(userId) == 0 {
		log.Println("routes:DeleteUser() -> Missing Required Parameters")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = "Missing Required Parameters"
		response.Error(w, errorBuilder)
		return
	}

	isDeleted, err := svc.JailMs.DeleteUser(userId)
	if err != nil {
		log.Println("routes:DeleteUser() -> Error Deleting the User")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	if !isDeleted {
		log.Println("routes:DeleteUser() -> Error Deleting the User")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = "Can't Delete The User"
		response.Error(w, errorBuilder)
		return
	}

	response.Success(w, "Successfully Deleted User")
}

func (svc *routesHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside Routes: UpdateUser()")
	user := models.Users{}
	errorBuilder := response.ErrorResponse{}
	byteBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("routes:UpdateUser() -> Error Reading the Request Body")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}
	err = json.Unmarshal(byteBody, &user)
	if err != nil {
		log.Println("routes:UpdateUser() -> Unmarshal Error")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	// validate request
	err = service.ValidateUpdateUserRequest(user)
	if err != nil {
		log.Println("routes:UpdateUser() -> Invalid Parameters")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}
	isUpdated, err := svc.JailMs.UpdateUser(user)
	if err != nil {
		log.Println("routes:UpdateUser() -> Error Updating the User")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	if !isUpdated {
		log.Println("routes:UpdateUser() -> Failed to update user")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = "Failed to update user"
		response.Error(w, errorBuilder)
		return
	}
	response.Success(w, "Successfully Updated User")
}
