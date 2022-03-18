package routes

import (
	"log"
	"net/http"

	"github.com/MakMoinee/go-mith/pkg/goserve"
	"github.com/MakMoinee/gojailms/cmd/webapp/response"
	"github.com/MakMoinee/gojailms/internal/common"
	"github.com/MakMoinee/gojailms/internal/gojailms"
	"github.com/go-chi/cors"
)

type routesHandler struct {
	JailMs gojailms.JailIntf
}

type RoutesIntf interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
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
}

func (svc *routesHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Innside Routes: GetUsers()")
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
