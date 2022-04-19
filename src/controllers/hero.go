package controllers

import (
	"MyHeroAcademiaApi/src/database"
	"MyHeroAcademiaApi/src/models"
	"MyHeroAcademiaApi/src/repository"
	"MyHeroAcademiaApi/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func FindAllHeroes(w http.ResponseWriter, r *http.Request) {}

func FindHeroById(w http.ResponseWriter, r *http.Request) {}

func CreateHero(w http.ResponseWriter, r *http.Request) {

	//getting the body of requisiton and converting to json
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var hero models.Hero
	//getting the json sent by the reqBody and converting to hero model
	//marshal --> converts byte/data to json
	//unmarshal -> converts json to model data
	if err = json.Unmarshal(reqBody, &hero); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	repo := repository.NewHeroRepository(db)
}
func UpdateHero(w http.ResponseWriter, r *http.Request) {}
func DeleteHero(w http.ResponseWriter, r *http.Request) {}
