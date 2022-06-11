package routes

import (
	"MyHeroAcademiaApi/src/controllers"
	"net/http"
)

var routeAuth = []Route{
	{
		URI:      "/login",
		Method:   http.MethodPost,
		Function: controllers.Login,
		NeedAuth: false,
	},
	{
		URI:      "/register",
		Method:   http.MethodPost,
		Function: controllers.Register,
		NeedAuth: false,
	},
}
