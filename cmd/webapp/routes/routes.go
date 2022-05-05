package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/MakMoinee/go-mith/pkg/goserve"
	"github.com/MakMoinee/gojailms/cmd/webapp/response"
	"github.com/MakMoinee/gojailms/internal/common"
	"github.com/MakMoinee/gojailms/internal/gojailms"
	"github.com/MakMoinee/gojailms/internal/gojailms/models"
	"github.com/MakMoinee/gojailms/internal/gojailms/service"
	"github.com/go-chi/cors"
)

type routesStruct struct {
	RouteName  string
	RouteValue string
}
type routesHandler struct {
	JailMs gojailms.JailIntf
}

type RoutesIntf interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	LogUser(w http.ResponseWriter, r *http.Request)
	CreateAdminUser(w http.ResponseWriter, r *http.Request)

	CreateVisitor(w http.ResponseWriter, r *http.Request)
	GetVisitors(w http.ResponseWriter, r *http.Request)
	GetVisitorByUserID(w http.ResponseWriter, r *http.Request)
	DeleteVisitor(w http.ResponseWriter, r *http.Request)

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
	set := make(map[string]interface{})
	infos := []routesStruct{}
	httpService.Router.Get(common.GetUsersPath, handler.GetUsers)
	infos = append(infos, routesStruct{RouteName: "RetrieveUsers", RouteValue: common.GetUsersPath})

	httpService.Router.Post(common.CreateUserPath, handler.CreateUser)
	infos = append(infos, routesStruct{RouteName: "CreateUser", RouteValue: common.CreateUserPath})

	httpService.Router.Delete(common.DeleteUserPath, handler.DeleteUser)
	infos = append(infos, routesStruct{RouteName: "DeleteUser", RouteValue: common.DeleteUserPath})

	httpService.Router.Put(common.UpdateUserPath, handler.UpdateUser)
	infos = append(infos, routesStruct{RouteName: "UpdateUser", RouteValue: common.UpdateUserPath})

	httpService.Router.Post(common.LogUserPath, handler.LogUser)
	infos = append(infos, routesStruct{RouteName: "LogUser", RouteValue: common.LogUserPath})

	httpService.Router.Post(common.CreateAdminPath, handler.CreateAdminUser)
	infos = append(infos, routesStruct{RouteName: "CreateAdminUser", RouteValue: common.CreateAdminPath})

	//visitor
	httpService.Router.Post(common.CreateVisitorPath, handler.CreateVisitor)
	infos = append(infos, routesStruct{RouteName: "CreateVisitor", RouteValue: common.CreateVisitorPath})

	httpService.Router.Get(common.GetVisitorsPath, handler.GetVisitors)
	infos = append(infos, routesStruct{RouteName: "GetVisitors", RouteValue: common.GetVisitorsPath})

	httpService.Router.Get(common.GetVisitorsByUserIDPath, handler.GetVisitorByUserID)
	infos = append(infos, routesStruct{RouteName: "GetVisitorByUserID", RouteValue: common.GetVisitorsByUserIDPath})

	httpService.Router.Delete(common.DeleteVisitorPath, handler.DeleteVisitor)
	infos = append(infos, routesStruct{RouteName: "DeleteVisitor", RouteValue: common.DeleteVisitorPath})

	//inmate
	httpService.Router.Post(common.CreateInmatePath, handler.CreateInmate)
	infos = append(infos, routesStruct{RouteName: "CreateInmate", RouteValue: common.CreateInmatePath})

	httpService.Router.Get(common.GetInmatePath, handler.GetInmates)
	infos = append(infos, routesStruct{RouteName: "GetInmates", RouteValue: common.GetInmatePath})

	set["routes"] = infos
	set["version"] = common.SERVICE_VERSION

	httpService.SetInfo(set)

}

func (svc *routesHandler) LogUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside Routes: LogUser()")
	errorBuilder := response.ErrorResponse{}
	user := models.Users{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Inside Routes:LogUser() -> Error Reading Body")
		errorBuilder.ErrorMessage = err.Error()
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		response.Error(w, errorBuilder)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Inside Routes:LogUser() -> Json Unmarshal Error")
		errorBuilder.ErrorMessage = err.Error()
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		response.Error(w, errorBuilder)
		return
	}

	if errs := service.ValidateUserRequest(user); errs != nil {
		log.Println("Inside Routes:LogUser() -> Invalid User Request")
		errorBuilder.ErrorMessage = errs.Error()
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		response.Error(w, errorBuilder)
		return
	}
	log.Println(user)
	isValidUser, validUser, err := svc.JailMs.LogUser(user)
	if err != nil {
		log.Println("Inside routes:LogUser() -> Error in Logging In")
		errorBuilder.ErrorMessage = err.Error()
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		response.Error(w, errorBuilder)
		return
	}

	if !isValidUser {
		log.Println("Inside routes:LogUser() -> Wrong credentials")
		errorBuilder.ErrorMessage = "Wrong Username or Password"
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		response.Error(w, errorBuilder)
		return
	}
	validUser.UserPassword = user.UserPassword
	response.Success(w, validUser)

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
	request := models.CreateRequest{}
	errorBuilder := response.ErrorResponse{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Routes:CreateUser() -> Reading the body error")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Routes:CreateUser() -> Unmarshal Error")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	// validate request
	err = service.ValidateUserRequest(request.LocalUser)
	if err != nil {
		log.Println("Routes:CreateUser() -> Invalid Parameters")
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		errorBuilder.ErrorMessage = err.Error()
		response.Error(w, errorBuilder)
		return
	}

	isUserCreated, err := svc.JailMs.CreateUser(request.LocalUser, request.LocalVisitor)
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

func (svc *routesHandler) CreateAdminUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside routes:CreateAdminUser()")
	user := models.Users{}
	errorBuilder := response.ErrorResponse{}
	byteBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error in routes:CreateAdminUser() -> Error reading the body")
		errorBuilder.ErrorMessage = err.Error()
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		response.Error(w, errorBuilder)
		return
	}

	err = json.Unmarshal(byteBody, &user)
	if err != nil {
		log.Println("Error in routes:CreateAdminUser() -> Error unmarshalling the body")
		errorBuilder.ErrorMessage = err.Error()
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		response.Error(w, errorBuilder)
		return
	}

	token := r.Header.Get("auth-token")

	if !strings.EqualFold(common.AUTH_TOKEN, token) {
		log.Println("Error in routes:CreateAdminUser() -> Error token is not authorized")
		errorBuilder.ErrorMessage = "not authorized"
		errorBuilder.ErrorStatus = http.StatusForbidden
		response.Error(w, errorBuilder)
		return
	}

	isCreated, err := svc.JailMs.CreateAdminUser(user)
	if err != nil {
		log.Println("Error in routes:CreateAdminUser() -> Error creating the admin")
		errorBuilder.ErrorMessage = err.Error()
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		response.Error(w, errorBuilder)
		return
	}

	if !isCreated {
		log.Println("Error in routes:CreateAdminUser() -> Failed to create admin")
		errorBuilder.ErrorMessage = "Failed to create admin"
		errorBuilder.ErrorStatus = http.StatusInternalServerError
		response.Error(w, errorBuilder)
		return
	}

	response.Success(w, "Successfully created admin")

}
