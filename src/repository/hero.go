package repository

import (
	"MyHeroAcademiaApi/src/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type Hero struct {
	db *mongo.Client
}

func NewHeroRepository(db *mongo.Client) *Hero {
	return &Hero{db}
}

func (repos Hero) CreateHero(hero models.Hero) (*mongo.InsertOneResult, error) {
	result, err := repos.db.Database("MyHeroDataBase").Collection("Hero").InsertOne(context.Background(),
		hero,
	)
	if err != nil {
		return nil, err
	}
	fmt.Println(result.InsertedID)

	return result, nil
}
