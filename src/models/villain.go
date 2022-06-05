package models

import (
	"errors"
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Villain struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TrueName     string             `bson:"trueName,omitempty" json:"trueName,omitempty"`
	LastName     string             `bson:"lastName,omitempty" json:"lastName,omitempty"`
	VillainName  string             `bson:"villainName,omitempty" json:"villainName,omitempty"`
	VillainRank  int64              `bson:"villainRank,omitempty" json:"villainRank,omitempty"`
	Age          int64              `bson:"age,omitempty" json:"age,omitempty"`
	Description  string             `bson:"description,omitempty" json:"description,omitempty"`
	VillainFiles []VillainFile      `bson:"files,omitempty" json:"files,omitempty"`
	Thumbnail    VillainThumbnail   `bson:"villainThumbnail,omitempty" json:"villainThumbnail,omitempty"` //`bson:"-" json:"-"`
}

func (villain *Villain) Preparar() error {
	//ai ele vai formatar os dados
	if erro := villain.formatar(); erro != nil {
		return erro
	}
	//validando se os campos não estão em branco
	/* if erro := hero.validar(); erro != nil {
		return erro
	} */

	return nil
}

//controlando os erros de campo no json
func (villain *Villain) validar() error {

	if villain.TrueName == "" {
		return errors.New(" Name needed and can't be empty")
	}

	if villain.LastName == "" {
		return errors.New(" LastName needed and can't be empty")
	}

	if villain.VillainName == "" {
		return errors.New(" HeroName needed and can't be empty")
	}

	return nil
}

func (villain *Villain) formatar() error {
	villain.TrueName = strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(villain.TrueName, " "))
	villain.LastName = strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(villain.LastName, " "))
	villain.VillainName = strings.ToLower(strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(villain.VillainName, " ")))
	villain.Description = strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(villain.Description, " "))

	/* if step == "registerUser" {
		passwordHashed, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}

		user.Password = string(passwordHashed)
	} */
	return nil
}
