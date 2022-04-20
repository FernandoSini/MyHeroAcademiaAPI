package routes

import (
	"MyHeroAcademiaApi/src/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	URI      string
	Method   string
	Function func(w http.ResponseWriter, r *http.Request)
	NeedAuth bool
}

func ConfigureRoutes(r *mux.Router) *mux.Router {
	routes := heroesRoute

	for _, route := range routes {
		//doing same middleware like node js
		if route.NeedAuth {

		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}
	return r
}
