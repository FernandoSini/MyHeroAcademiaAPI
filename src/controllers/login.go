package controllers

import (
	"MyHeroAcademiaApi/src/auth"
	"MyHeroAcademiaApi/src/database"
	"MyHeroAcademiaApi/src/models"
	"MyHeroAcademiaApi/src/repository"
	"MyHeroAcademiaApi/src/responses"
	"MyHeroAcademiaApi/src/security"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//é responsável por autenticar um usuário na api
func Login(w http.ResponseWriter, r *http.Request) {
	reqBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(reqBody, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Disconnect(context.Background())
	repository := repository.NovoRepositorioDeUsuarios(db)
	savedUserInDB, erro := repository.FindUserByEmail(user.Email)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	//verificando se o hash bate com a senha vindo na requisicao
	if erro = security.VerifyPassword(savedUserInDB.Password, user.Password); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := auth.CriarToken(savedUserInDB.ID.Hex())

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	w.Write([]byte(token))
	//responses.JSON(w, http.StatusAccepted, token)

}
