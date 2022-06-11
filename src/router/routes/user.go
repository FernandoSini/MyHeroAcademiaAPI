package routes

import (
	"MyHeroAcademiaApi/src/controllers"
	"net/http"
)

var routesUser = []Route{

	{
		URI:      "/user/{userId}",
		Method:   http.MethodGet,
		Function: controllers.FindUser,
		NeedAuth: false,
	}, {
		URI:      "/users",
		Method:   http.MethodGet,
		Function: controllers.FindAllUsers,
		NeedAuth: true,
	},
	{
		URI:      "/user/{userId}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		NeedAuth: true,
	},
	{
		URI:      "/user/{userId}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		NeedAuth: true,
	},

	{
		URI:      "/user/{userId}/password/update",
		Method:   http.MethodPost,
		Function: controllers.UpdatePassword,
		NeedAuth: true,
	},
}
