package routes

import (
	"MyHeroAcademiaApi/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI      string
	Method   string
	Function func(w http.ResponseWriter, r *http.Request)
	NeedAuth bool
}

func ConfigureRoutes(r *mux.Router) *mux.Router {
	routes := heroesRoute
	routes = append(routes, villainsRoute...)
	routes = append(routes, routesUser...)
	routes = append(routes, routeAuth...)

	for _, route := range routes {
		//doing same middleware like node js
		if route.NeedAuth {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)

		}
	}
	return r
}
