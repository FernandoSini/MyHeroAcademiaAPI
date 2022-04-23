package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	UrlConexao = ""
	Porta      = 0
	//chave pra assinar o token
	SecretKey []byte
)

//inicializando as variáveis de ambiente
func Carregar() {

	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatalln(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 5000
	}
	UrlConexao = fmt.Sprintln(os.Getenv("URL"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
