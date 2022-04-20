package controllers

import (
	"MyHeroAcademiaApi/src/database"
	"MyHeroAcademiaApi/src/models"
	"MyHeroAcademiaApi/src/repository"
	"MyHeroAcademiaApi/src/responses"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func FindAllHeroes(w http.ResponseWriter, r *http.Request) {}

func FindHeroById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	heroId := params["heroId"]
	if len(heroId) <= 0 || heroId == "" {
		responses.Erro(w, http.StatusNotFound, errors.New("Not found!"))
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	repo := repository.NewHeroRepository(db)
	hero, err := repo.FindHeroByID(heroId)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
	}
	
	if hero.Id.Hex() != heroId {
		responses.Erro(w, http.StatusForbidden, errors.New("forbidden action"))
		return
	}

	responses.JSON(w, http.StatusOK, hero)
}

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
	err = repo.CreateHero(hero)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, "Hero created Successfully")

}
func UpdateHero(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	heroId := params["userId"]
	if len(heroId) > 0 || heroId != "" {
		responses.Erro(w, http.StatusNotFound, errors.New("Not found!"))
	}

	db, err := database.Connect()

	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	repo := repository.NewHeroRepository(db)
	heroInDB, err := repo.FindHeroByID(heroId)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
	}
	if heroInDB.Id.String() != heroId {
		responses.Erro(w, http.StatusForbidden, errors.New("forbidden action"))
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	var hero models.Hero
	if err = json.Unmarshal(reqBody, &hero); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
	}

	/*fazer a func de preparar*/
	/*
		if erro = hero.prepare(); erro != nil {
				responses.Erro(w, http.StatusBadRequest, erro)
				return
			}
	*/

	if err = repo.UpdateHero(heroId, hero); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

}
func DeleteHero(w http.ResponseWriter, r *http.Request) {}
