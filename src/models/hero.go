package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hero struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	TrueName  string             `json:"trueName"`
	LastName  string             `json:"lastName"`
	HeroName  string             `json:"heroName"`
	HeroRank  int64              `json:"heroRank,omitempty"`
	Age       int64              `json:"age"`
	HeroImage []HeroImage        `json:"images,omitempty"`
}
