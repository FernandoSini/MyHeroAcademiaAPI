package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type HeroImage struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	ImageUrl  string             `json:"imageUrl,omitempty"`
	IdHeroRef primitive.ObjectID `json:"idHeroRef,omitempty"`
}
