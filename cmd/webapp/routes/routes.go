package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/MakMoinee/go-mith/pkg/goserve"
	"github.com/MakMoinee/go-mith/pkg/response"
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
	UpdateUserVisitor(w http.ResponseWriter, r *http.Request)

	CreateVisitor(w http.ResponseWriter, r *http.Request)
	GetVisitors(w http.ResponseWriter, r *http.Request)
	GetVisitorByUserID(w http.ResponseWriter, r *http.Request)
	DeleteVisitor(w http.ResponseWriter, r *http.Request)

	CreateInmate(w http.ResponseWriter, r *http.Request)
	GetInmates(w http.ResponseWriter, r *http.Request)
	InsertVisitorHistory(w http.ResponseWriter, r *http.Request)
	GetAllVisitorHistory(w http.ResponseWriter, r *http.Request)
	UpdateVisitor(w http.ResponseWriter, r *http.Request)
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

	httpService.Router.Post(common.UpdateUserVisitorPath, handler.UpdateUserVisitor)
	infos = append(infos, routesStruct{RouteName: "ForgotPassword", RouteValue: common.UpdateUserVisitorPath})

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

	httpService.Router.Post(common.InsertVisitorHistoryPath, handler.InsertVisitorHistory)
	infos = append(infos, routesStruct{RouteName: "InsertVisitorHistory", RouteValue: common.InsertVisitorHistoryPath})

	httpService.Router.Get(common.GetAllVisitorHistoryPath, handler.GetAllVisitorHistory)
	infos = append(infos, routesStruct{RouteName: "GetAllVisitorHistory", RouteValue: common.GetAllVisitorHistoryPath})

	httpService.Router.Post(common.UpdateVisitorPath, handler.UpdateVisitor)
	infos = append(infos, routesStruct{RouteName: "UpdateVisitor", RouteValue: common.UpdateVisitorPath})

	set["routes"] = infos
	set["version"] = common.SERVICE_VERSION

	httpService.SetInfo(set)

}

func (svc *routesHandler) LogUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside Routes: LogUser()")
	user := models.Users{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Inside Routes:LogUser() -> Error Reading Body")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Inside Routes:LogUser() -> Json Unmarshal Error")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	if errs := service.ValidateUserRequest(user); errs != nil {
		log.Println("Inside Routes:LogUser() -> Invalid User Request")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}
	log.Println(user)
	isValidUser, validUser, err := svc.JailMs.LogUser(user)
	if err != nil {
		log.Println("Inside routes:LogUser() -> Error in Logging In")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	if !isValidUser {
		log.Println("Inside routes:LogUser() -> Wrong credentials")
		errorBuilder := response.NewErrorBuilder("Wrong Username or Password", http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}
	validUser.UserPassword = user.UserPassword
	response.Success(w, validUser)

}

func (svc *routesHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside Routes: GetUsers()")
	usersList, err := svc.JailMs.GetUsers()
	if err != nil {

		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Routes:CreateUser() -> Reading the body error")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}
	fmt.Println("Body: ", string(body))
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Routes:CreateUser() -> Unmarshal Error")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	// validate request
	err = service.ValidateUserRequest(request.LocalUser)
	if err != nil {
		log.Println("Routes:CreateUser() -> Invalid Parameters")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	isUserCreated, err := svc.JailMs.CreateUser(request.LocalUser, request.LocalVisitor)
	if err != nil {
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	if !isUserCreated {
		errorBuilder := response.NewErrorBuilder("User Not Created", http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	response.Response(w, response.NewSuccessBuilder("Successfully Created User"))
}

func (svc *routesHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside Routes: DeleteUser()")
	query := r.URL.Query()
	userId := query.Get("id")
	if len(userId) == 0 {
		log.Println("routes:DeleteUser() -> Missing Required Parameters")
		errorBuilder := response.NewErrorBuilder("Missing Required Parameters", http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	isDeleted, err := svc.JailMs.DeleteUser(userId)
	if err != nil {
		log.Println("routes:DeleteUser() -> Error Deleting the User")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	if !isDeleted {
		log.Println("routes:DeleteUser() -> Error Deleting the User")
		errorBuilder := response.NewErrorBuilder("Can't Delete The User", http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	response.Response(w, response.NewSuccessBuilder("Successfully Deleted User"))
}

func (svc *routesHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside Routes: UpdateUser()")
	user := models.Users{}
	byteBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("routes:UpdateUser() -> Error Reading the Request Body")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}
	err = json.Unmarshal(byteBody, &user)
	if err != nil {
		log.Println("routes:UpdateUser() -> Unmarshal Error")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	// validate request
	err = service.ValidateUpdateUserRequest(user)
	if err != nil {
		log.Println("routes:UpdateUser() -> Invalid Parameters")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}
	isUpdated, err := svc.JailMs.UpdateUser(user)
	if err != nil {
		log.Println("routes:UpdateUser() -> Error Updating the User")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	if !isUpdated {
		log.Println("routes:UpdateUser() -> Failed to update user")
		errorBuilder := response.NewErrorBuilder("Failed to update user", http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}
	response.Response(w, response.NewSuccessBuilder("Successfully Updated User"))
}

func (svc *routesHandler) CreateAdminUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside routes:CreateAdminUser()")
	user := models.Users{}
	byteBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error in routes:CreateAdminUser() -> Error reading the body")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	err = json.Unmarshal(byteBody, &user)
	if err != nil {
		log.Println("Error in routes:CreateAdminUser() -> Error unmarshalling the body")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	token := r.Header.Get("auth-token")

	if !strings.EqualFold(common.AUTH_TOKEN, token) {
		log.Println("Error in routes:CreateAdminUser() -> Error token is not authorized")
		errorBuilder := response.NewErrorBuilder("not authorized", http.StatusForbidden)
		response.Response(w, errorBuilder)
		return
	}

	isCreated, err := svc.JailMs.CreateAdminUser(user)
	if err != nil {
		log.Println("Error in routes:CreateAdminUser() -> Error creating the admin")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	if !isCreated {
		log.Println("Error in routes:CreateAdminUser() -> Failed to create admin")
		errorBuilder := response.NewErrorBuilder("Failed to create admin", http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	response.Response(w, response.NewSuccessBuilder("Successfully created admin"))

}

func (svc *routesHandler) UpdateUserVisitor(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside routes:UpdateUserVisitor() ...")
	userVisitor := models.UserVisitor{}
	byteBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error in routes:UpdateUserVisitor() -> Error reading the body")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}
	err = json.Unmarshal(byteBody, &userVisitor)
	if err != nil {
		log.Println("Error in routes:UpdateUserVisitor() -> Error unmarshalling the body")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}
	token := r.Header.Get("auth-token")
	userVisitor.Token = token
	isUpdated, err := svc.JailMs.UpdateUserVisitor(userVisitor)
	if err != nil && strings.EqualFold(err.Error(), "not authorized") {
		log.Println("Error in routes:UpdateUserVisitor() -> not authorized")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusForbidden)
		response.Response(w, errorBuilder)
		return
	} else if err != nil {
		log.Println("Error in routes:UpdateUserVisitor() -> Error updating password")
		errorBuilder := response.NewErrorBuilder(err.Error(), http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	if !isUpdated {
		log.Println("Error in routes:UpdateUserVisitor() -> Error updating password")
		errorBuilder := response.NewErrorBuilder("failed to update password", http.StatusInternalServerError)
		response.Response(w, errorBuilder)
		return
	}

	response.Response(w, response.NewSuccessBuilder("Successfully updated password"))
}
