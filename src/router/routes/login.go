package routes

import (
	"MyHeroAcademiaApi/src/controllers"
	"net/http"
)

var routeLogin = []Route{
	{
		URI:      "/login",
		Method:   http.MethodPost,
		Function: controllers.Login,
		NeedAuth: false,
	},
}
