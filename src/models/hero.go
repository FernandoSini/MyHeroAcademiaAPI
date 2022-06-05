package models

import (
	"errors"
	"regexp"
	"strings"

	"github.com/juju/mgo/v2/bson"
)

type Hero struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	TrueName    string        `bson:"trueName,omitempty" json:"trueName,omitempty"`
	LastName    string        `bson:"lastName,omitempty" json:"lastName,omitempty"`
	HeroName    string        `bson:"heroName,omitempty" json:"heroName,omitempty"`
	HeroRank    int64         `bson:"heroRank,omitempty" json:"heroRank,omitempty"`
	Age         int64         `bson:"age,omitempty" json:"age,omitempty"`
	Description string        `bson:"description,omitempty" json:"description,omitempty"`
	HeroFiles   []HeroFile    `bson:"files,omitempty" json:"files,omitempty"`
	Thumbnail   HeroThumbnail `bson:"heroThumbnail,omitempty" json:"heroThumbnail,omitempty"` //`bson:"thumbnail,omitempty" json:"thumbnail,omitempty"`
}

func (hero *Hero) Preparar() error {
	//ai ele vai formatar os dados
	if erro := hero.formatar(); erro != nil {
		return erro
	}
	//validando se os campos não estão em branco
	/* if erro := hero.validar(); erro != nil {
		return erro
	} */

	return nil
}

//controlando os erros de campo no json
func (hero *Hero) validar() error {

	if hero.TrueName == "" {
		return errors.New(" Name needed and can't be empty")
	}

	if hero.LastName == "" {
		return errors.New(" LastName needed and can't be empty")
	}

	if hero.HeroName == "" {
		return errors.New(" HeroName needed and can't be empty")
	}

	return nil
}

func (hero *Hero) formatar() error {
	hero.TrueName = strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(hero.TrueName, " "))
	hero.LastName = strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(hero.LastName, " "))
	hero.HeroName = strings.ToLower(strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(hero.HeroName, " ")))
	hero.Description = strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(hero.Description, " "))

	/* if step == "registerUser" {
		passwordHashed, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}

		user.Password = string(passwordHashed)
	} */
	return nil
}
