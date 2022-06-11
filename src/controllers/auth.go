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
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserLoginData struct {
	ID        primitive.ObjectID `json:"id,omitempty"` //quando vc for passar o usuario pra um json e o id estiver vazio, ele não sera passado pelo json
	Name      string             ` json:"name,omitempty"`
	UserName  string             ` json:"username,omitempty"`
	Email     string             ` json:"email,omitempty"`
	Token     string             `json:"token,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty"`
}

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
	//w.Write([]byte(token))

	responses.JSON(w, http.StatusAccepted, UserLoginData{ID: savedUserInDB.ID, Name: savedUserInDB.Name, Email: savedUserInDB.Email, UserName: savedUserInDB.UserName, Token: token, CreatedAt: savedUserInDB.CreatedAt})

}

func Register(w http.ResponseWriter, r *http.Request) {
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

	id, erro := repository.Create(usuario)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuario.ID, erro = primitive.ObjectIDFromHex(id)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, " User created successfuly! ")

}
