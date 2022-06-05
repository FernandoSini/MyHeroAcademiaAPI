package routes

import (
	"MyHeroAcademiaApi/src/controllers"
	"net/http"
)

var heroesRoute = []Route{

	{
		URI:      "/heroes",
		Method:   http.MethodGet,
		Function: controllers.FindAllHeroes,
		NeedAuth: false,
	},
	{
		URI:      "/heroes/details/{heroId}",
		Method:   http.MethodGet,
		Function: controllers.FindHeroById,
		NeedAuth: false,
	},
	{
		URI:      "/heroes/details/*",
		Method:   http.MethodGet,
		Function: controllers.NotFound,
		NeedAuth: false,
	},
	{
		URI:      "/heroes/create",
		Method:   http.MethodPost,
		Function: controllers.CreateHero,
		NeedAuth: false,
	},
	{
		URI:      "/heroes/update/{heroId}",
		Method:   http.MethodPut,
		Function: controllers.UpdateHero,
		NeedAuth: false,
	},
	{
		URI:      "/heroes/delete/{heroId}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteHero,
		NeedAuth: false,
	},
	/* {
		URI:      "/heroes/{heroId}/addImage",
		Method:   http.MethodPost,
		Function: controllers.AddHeroImage,
		NeedAuth: false,
	}, */
	{
		URI:      "/heroes/{heroName}",
		Method:   http.MethodGet,
		Function: controllers.FindHeroByHeroName,
		NeedAuth: false,
	},
}
