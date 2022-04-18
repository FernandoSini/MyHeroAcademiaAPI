package controllers

import (
	"MyHeroAcademiaApi/src/responses"
	"errors"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {

	responses.Erro(w, http.StatusNotFound, errors.New("Not found!"))
}
