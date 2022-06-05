package repository

import (
	"MyHeroAcademiaApi/src/models"
	"context"
	"fmt"
	"os"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Hero struct {
	db *mongo.Client
}
type HeroFile struct {
	db *mongo.Client
}

func NewHeroRepository(db *mongo.Client) *Hero {
	return &Hero{db}
}
func NewHeroFileRepository(db *mongo.Client) *HeroFile {
	return &HeroFile{db}
}

//old method trying to send primitive id to controller after created hero in db
/*func (repos Hero) CreateHero(hero models.Hero) (string, error) {
	result, err := repos.db.Database(os.Getenv("DBNAME")).Collection("Hero").InsertOne(context.Background(),
		hero,
	)
	if err != nil {
		return "", err
	}
	fmt.Println(result.InsertedID)

	return fmt.Sprintf("%v", result.InsertedID), nil
}*/
func (repos Hero) CreateHero(hero models.Hero) error {
	result, err := repos.db.Database(os.Getenv("DBNAME")).
		Collection("Hero").
		InsertOne(context.Background(),
			hero,
		)
	if err != nil {
		return err
	}
	if reflect.ValueOf(hero.Thumbnail).IsValid() {
		
		_, err = repos.db.Database(os.Getenv("DBNAME")).
			Collection("HeroThumbnail").
			InsertOne(context.Background(),
				&hero.Thumbnail,
			)
		if err != nil {
			return err
		}
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
	err = repos.db.Database(os.Getenv("DBNAME")).
		Collection("Hero").
		FindOne(context.Background(),
			bson.D{{Key: "_id", Value: objectId}},
		).
		Decode(&hero)

	if err != nil {
		return hero, err
	}
	return hero, nil
}

func (repos Hero) UpdateHero(id string, hero models.Hero) error {
	//result, erro := repos.db.Database(os.Getenv("DBNAME")).Collection("Hero").FindOne(context.Background())
	/* bson.D{{Key: "$trueName", Value: &hero.TrueName},
	{Key: "lastName", Value: &hero.LastName},
	{Key: "heroName", Value: &hero.HeroName},
	{Key: "heroRank", Value: &hero.HeroRank},
	{Key: "age", Value: &hero.Age}} */

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = repos.db.Database(os.Getenv("DBNAME")).
		Collection("Hero").UpdateOne(context.Background(), bson.D{{Key: "_id", Value: objectId}},
		bson.D{{Key: "$set", Value: &hero}})
	/*  bson.M{"trueName": &hero.TrueName,
	"lastName": &hero.LastName,
	"heroName": &hero.HeroName,
	"heroRank": &hero.HeroRank,
	"age":      &hero.Age}). */
	// Decode(&hero)
	if err != nil {
		return err
	}

	return nil
}

func (repos Hero) FindHeroes() ([]models.Hero, error) {
	hero := models.Hero{}
	heroes := []models.Hero{}

	result, err := repos.db.Database(os.Getenv("DBNAME")).
		Collection("Hero").
		Find(context.Background(), bson.D{})

	if err != nil {

		return nil, err
	}
	defer result.Close(context.Background())

	for result.Next(context.Background()) {
		if err = result.Decode(&hero); err != nil {
			return nil, err
		}

		heroes = append(heroes, hero)
	}

	return heroes, nil
}

func (repos Hero) DeleteHero(id string) error {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		return err
	}
	_, err = repos.db.Database(os.Getenv("DBNAME")).
		Collection("Hero").
		DeleteOne(context.Background(),
			bson.D{{Key: "_id", Value: objectId}},
		)

	if err != nil {
		return err
	}
	return nil
}

func (repo HeroFile) AddHeroFile(heroFile models.HeroFile) error {
	result, err := repo.db.Database(os.Getenv("DBNAME")).Collection("HeroFile").InsertOne(context.Background(), heroFile)
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

func (repo Hero) FindHeroByHeroName(heroName string) (models.Hero, error) {
	hero := models.Hero{}

	err := repo.db.Database(os.Getenv("DBNAME")).
		Collection("Hero").
		FindOne(context.Background(),
			bson.D{{Key: "heroName", Value: heroName}},
		).
		Decode(&hero)

	if err != nil {
		return hero, err
	}
	return hero, nil

}
