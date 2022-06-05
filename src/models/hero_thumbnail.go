package models

import (
	"github.com/juju/mgo/v2/bson"
)

/* type HeroThumbnail struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ImagePath string             `bson:"imagePath,omitempty" json:"imagePath,omitempty"`
	FileName  string             `bson:"filename, omitempty" json:"filename,omitempty"`
	ImgData   string             `bson:"imgData,omitempty" json:"imgData,omitempty"`
	IdHeroRef primitive.ObjectID `bson:"idHeroRef, omitempty" json:"idHeroRef,omitempty"`
} */

type HeroThumbnail struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	FileName string        `bson:"filename,omitempty" json:"filename,omitempty"`
	//Content   []byte             `bson:"content,omitempty" json:"content,omitempty"`
	Content   string        `bson:"content,omitempty" json:"content,omitempty"`
	IdHeroRef bson.ObjectId `bson:"idHeroRef,omitempty" json:"idHeroRef,omitempty"`
}
