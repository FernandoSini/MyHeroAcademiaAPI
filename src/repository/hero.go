package repository

import (
	"MyHeroAcademiaApi/src/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Hero struct {
	db *mongo.Client
}

func NewHeroRepository(db *mongo.Client) *Hero {
	return &Hero{db}
}

//old method trying to send primitive id to controller after created hero in db
/*func (repos Hero) CreateHero(hero models.Hero) (string, error) {
	result, err := repos.db.Database("MyHeroDataBase").Collection("Hero").InsertOne(context.Background(),
		hero,
	)
	if err != nil {
		return "", err
	}
	fmt.Println(result.InsertedID)

	return fmt.Sprintf("%v", result.InsertedID), nil
}*/
func (repos Hero) CreateHero(hero models.Hero) error {
	result, err := repos.db.Database("MyHeroDataBase").Collection("Hero").InsertOne(context.Background(),
		hero,
	)
	if err != nil {
		return err
	}
	fmt.Println(result.InsertedID)

	if err != nil {
		fmt.Println("caiu aqui")
		return err
	}
	return nil
}

func (repos Hero) FindHeroByID(id string) (models.Hero, error) {
	hero := models.Hero{}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		return hero, err
	}
	err = repos.db.Database("MyHeroDataBase").
		Collection("Hero").
		FindOne(context.Background(),
			bson.D{{"_id", objectId}},
		).
		Decode(&hero)

	if err != nil {
		return hero, err
	}
	return hero, nil
}

func (repos Hero) UpdateHero(id string, hero models.Hero) error {
	//result, erro := repos.db.Database("MyHeroDatabase").Collection("Hero").FindOne(context.Background())
	return nil
}
