package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*type HeroFile struct {
 	Id        primitive.ObjectID `bson:"_id, omitempty" json:"id,omitempty"`
	FileName  string             `bson:"filename, omitempty" json:"filename,omitempty"`
	Content  []byte             `bson:"content,omitempty" json:"content,omitempty"`
	IdHeroRef primitive.ObjectID `bson:"idHeroRef, omitempty" json:"idHeroRef,omitempty"`
} */

type HeroFile struct {
	Id           primitive.ObjectID `bson:"_id, omitempty" json:"id,omitempty"`
	FileName     string             `bson:"filename,omitempty" json:"filename,omitempty"`
	Content      []byte             `bson:"content,omitempty" json:"content,omitempty"`
	IdVillainRef primitive.ObjectID `bson:"idVillainRef,omitempty" json:"idVillainRef,omitempty"`
}
