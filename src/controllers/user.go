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
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

//controllers
//responsavel por controlar as operacoes, não necessariamente vao executar elas
/* func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.User
	usuario.CreatedAt = time.Now()
	if erro = json.Unmarshal(reqBody, &usuario); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = usuario.Preparar("registerUser"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Disconnect(context.Background())

	//passando a conexão pro repositório
	repository := repository.NovoRepositorioDeUsuarios(db)

	id, erro := repository.Criar(usuario)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	usuario.ID, erro = primitive.ObjectIDFromHex(id)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, usuario)

} */

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Disconnect(context.Background())

	repo := repository.NewHeroRepository(db)
	users, err := repo.FindUsers()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if len(users) <= 0 && err == nil {
		responses.Erro(w, http.StatusNotFound, errors.New("not found"))
		return
	}

	responses.JSON(w, http.StatusOK, users)

}

func FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId := params["userId"]
	if len(userId) <= 0 || userId == "" {
		responses.Erro(w, http.StatusInternalServerError, errors.New("not found"))
		return
	}
	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Disconnect(context.Background())
	repository := repository.NovoRepositorioDeUsuarios(db)
	user, erro := repository.FindUserById(userId)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	paramteros := mux.Vars(r)

	userId := paramteros["userId"]
	if len(userId) <= 0 || userId == "" {
		responses.Erro(w, http.StatusBadRequest, errors.New(" Id can't be null "))
		return
	}

	userIdInToken, erro := auth.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	if userId != userIdInToken {
		responses.Erro(w, http.StatusForbidden, errors.New(" Forbidden Action"))
		return
	}

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
	if erro = user.Preparar("updateUser"); erro != nil {
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
	if erro = repository.Update(userId, user); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userId := parametros["userId"]
	if len(userId) <= 0 || userId == "" {
		responses.Erro(w, http.StatusBadRequest, errors.New("Id can't be null"))
		return
	}
	userIdInToken, erro := auth.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userId != userIdInToken {
		responses.Erro(w, http.StatusForbidden, errors.New(" Forbidden Action"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Disconnect(context.Background())

	repository := repository.NovoRepositorioDeUsuarios(db)
	if erro = repository.DeleteUser(userId); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIdInToken, erro := auth.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)

	userId := params["userId"]
	if len(userId) <= 0 || userId == "" {
		responses.Erro(w, http.StatusBadRequest, errors.New("Id can't be null"))
		return
	}
	if userIdInToken != userId {
		responses.Erro(w, http.StatusForbidden, errors.New(" Forbidden action"))
		return
	}

	reqBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var password models.Password
	if erro = json.Unmarshal(reqBody, &password); erro != nil {
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
	//buscando a senha no banco para comparar sua hash com o valor vindo do json
	savedPassword, erro := repository.FindPassword(userId)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerifyPassword(savedPassword, password.CurrentPassword); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, errors.New(" Not correct password"))
		return
	}

	hashedPassword, erro := security.Hash(password.NewPassword)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	//inserido a senha atualizada no banco
	if erro = repository.UpdatePassword(userId, string(hashedPassword)); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
	}
	responses.JSON(w, http.StatusNoContent, nil)

}
