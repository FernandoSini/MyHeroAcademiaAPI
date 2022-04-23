package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type HeroImage struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	ImagePath string             `json:"imagePath,omitempty"`
	FileName  string             `json:"filename,omitempty"`
	ImgData   string             `json:"imgByte,omitempty"`
	IdHeroRef primitive.ObjectID `json:"idHeroRef,omitempty"`
}
