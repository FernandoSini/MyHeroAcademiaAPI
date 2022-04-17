package routes

import (
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

	return r
}
