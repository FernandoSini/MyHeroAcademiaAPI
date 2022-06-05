package routes

import (
	"MyHeroAcademiaApi/src/controllers"
	"net/http"
)

var villainsRoute = []Route{

	{
		URI:      "/villains",
		Method:   http.MethodGet,
		Function: controllers.FindAllVillains,
		NeedAuth: false,
	},
	{
		URI:      "/villains/details/{villainId}",
		Method:   http.MethodGet,
		Function: controllers.FindVillainById,
		NeedAuth: false,
	},
	{
		URI:      "/villains/details/*",
		Method:   http.MethodGet,
		Function: controllers.NotFound,
		NeedAuth: false,
	},
	{
		URI:      "/villains/create",
		Method:   http.MethodPost,
		Function: controllers.CreateVillain,
		NeedAuth: false,
	},
	{
		URI:      "/villains/update/{villainId}",
		Method:   http.MethodPut,
		Function: controllers.UpdateVillain,
		NeedAuth: false,
	},
	{
		URI:      "/villains/delete/{villainId}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteVillain,
		NeedAuth: false,
	},
	// {
	// 	URI:      "/villains/{villainId}/addImage",
	// 	Method:   http.MethodPost,
	// 	Function: controllers.AddVillainImage,
	// 	NeedAuth: false,
	// },
	{
		URI:      "/villains/{villainName}",
		Method:   http.MethodGet,
		Function: controllers.FindVillainByVillainName,
		NeedAuth: false,
	},
}
