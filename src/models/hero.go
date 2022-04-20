package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hero struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TrueName  string             `bson:"trueName,omitempty" json:"trueName,omitempty"`
	LastName  string             `bson:"lastName,omitempty" json:"lastName,omitempty"`
	HeroName  string             `bson:"heroName,omitempty" json:"heroName,omitempty"`
	HeroRank  int64              `bson:"heroRank,omitempty" json:"heroRank,omitempty"`
	Age       int64              `bson:"age,omitempty" json:"age,omitempty"`
	HeroImage []HeroImage        `bson:"images,omitempty" json:"images,omitempty"`
}
