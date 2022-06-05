package models

import (
	"MyHeroAcademiaApi/src/security"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` //quando vc for passar o usuario pra um json e o id estiver vazio, ele não sera passado pelo json
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	UserName  string             `bson:"username,omitempty" json:"username,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	Password  string             `bson:"password,omitempty" json:"password,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}

func (user *User) Preparar(step string) error {
	//ai ele vai formatar os dados
	if erro := user.formatar(step); erro != nil {
		return erro
	}
	//validando se os campos não estão em branco
	if erro := user.validar(step); erro != nil {
		return erro
	}

	return nil
}

//controlando os erros de campo no json
func (user *User) validar(step string) error {

	if user.Name == "" {
		return errors.New(" Name needed and can't be empty")
	}

	if user.UserName == "" {
		return errors.New(" Username needed and can't be empty")
	}

	if user.Email == "" {
		return errors.New(" Email needed and can't be empty")
	}
	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New(" Email invalid")
	}
	if step == "registerUser" && user.Password == "" {
		return errors.New(" Password needed and can't be empty")
	}
	return nil
}

func (user *User) formatar(step string) error {
	user.Name = strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(user.Name, " "))
	user.UserName = strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(user.UserName, " "))
	user.Email = strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(user.Email, " "))

	if step == "registerUser" {
		passwordHashed, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}

		user.Password = string(passwordHashed)
	}
	return nil
}
