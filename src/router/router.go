package router

import (
	"MyHeroAcademiaApi/src/router/routes"
	"github.com/gorilla/mux"
)

// GenerateRoutes control the routes of api
//will return router with routes configured
func GenerateRoutes() *mux.Router {
	router := mux.NewRouter()
	return routes.ConfigureRoutes(router)

}
