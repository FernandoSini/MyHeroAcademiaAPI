package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/* type VillainThumbnail struct {
	Id           primitive.ObjectID `json:"id,omitempty"`
	ImagePath    string             `json:"imagePath,omitempty"`
	FileName     string             `json:"filename,omitempty"`
	ImgData      string             `json:"imgData,omitempty"`
	IdVillainRef primitive.ObjectID `json:"idVillainRef,omitempty"`
} */
type VillainThumbnail struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	//FileName     string             `bson:"filename,omitempty" json:"filename,omitempty"`
	//Content      []byte             `bson:"content,omitempty" json:"content,omitempty"`
	Content      string             `bson:"content,omitempty" json:"content,omitempty"`
	IdVillainRef primitive.ObjectID `bson:"idVillainRef,omitempty" json:"idVillainRef,omitempty"`
}
