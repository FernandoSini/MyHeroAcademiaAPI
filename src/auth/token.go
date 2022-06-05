package auth

import (
	"MyHeroAcademiaApi/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

//criando o token com as permissions do usuario
func CriarToken(userId string) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId
	//gerando o secret pra garantir authenticidade dele
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey)) //retornando token assinado, com o secret
}

//verify if token is valid
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	//fazendo o parse do token
	token, erro := jwt.Parse(tokenString, returnVerificationKey)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New(" Inválid token")

}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	//verificando se é um bearer token
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf(" Inválid signature method %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

func ExtractUserID(r *http.Request) (string, error) {
	tokenString := extractToken(r)
	//fazendo o parse do token
	token, erro := jwt.Parse(tokenString, returnVerificationKey)
	if erro != nil {
		return "", erro
	}

	//ver o id que está dentro do token
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := fmt.Sprintf("%s", permissions["userId"])
		if len(userID) <= 0 || userID == "" {
			return "", errors.New(" Inválid token")
		}
		return userID, nil
	}

	return "", errors.New(" Inválid token")

}
