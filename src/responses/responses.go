package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

//retorna json como resposta
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatalln(erro)
		}

	}

}

func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
