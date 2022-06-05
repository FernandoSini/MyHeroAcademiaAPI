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

type Villain struct {
	db *mongo.Client
}
type VillainFile struct {
	db *mongo.Client
}
type VillainThumbnail struct {
	db *mongo.Client
}

func NewVillainRepository(db *mongo.Client) *Villain {
	return &Villain{db}
}
func NewVillainFileRepository(db *mongo.Client) *VillainFile {
	return &VillainFile{db}
}
func NewVillainThumbnailRepository(db *mongo.Client) *VillainFile {
	return &VillainFile{db}
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

func (repos Villain) CreateVillain(villain models.Villain) error {
	result, err := repos.db.Database(os.Getenv("DBNAME")).
		Collection("Villain").
		InsertOne(context.Background(),
			&villain,
		)
	if err != nil {
		return err
	}
	if reflect.ValueOf(villain.Thumbnail).IsValid() {
		_, err = repos.db.Database(os.Getenv("DBNAME")).
			Collection("VillainThumbnail").
			InsertOne(context.Background(),
				&villain.Thumbnail,
			)
		if err != nil {
			return err
		}
	}
	fmt.Println(result.InsertedID)

	/* 	err = VillainThumbnail.addVillainThumb(
	VillainThumbnail{},
	models.VillainThumbnail{
		Id:           villain.Thumbnail.Id,
		Content:      villain.Thumbnail.Content,
		IdVillainRef: villain.Thumbnail.IdVillainRef,
	}) */

	if err != nil {
		fmt.Println("caiu aqui")
		return err
	}
	return nil
}

func (repos Villain) FindVillainByID(id string) (models.Villain, error) {
	villain := models.Villain{}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		return villain, err
	}
	err = repos.db.Database(os.Getenv("DBNAME")).
		Collection("Villain").
		FindOne(context.Background(),
			bson.D{{Key: "_id", Value: objectId}},
		).
		Decode(&villain)

	if err != nil {
		return villain, err
	}
	return villain, nil
}

func (repos Villain) UpdateVillain(id string, villain models.Villain) error {
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
		Collection("Villain").UpdateOne(context.Background(), bson.D{{Key: "_id", Value: objectId}},
		bson.D{{Key: "$set", Value: &villain}})

	if err != nil {
		return err
	}

	return nil
}

//drop villains collection
/* teste := repos.db.Database(os.Getenv("DBNAME")).Collection("Villain").Drop(context.Background())
print(teste) */
func (repos Villain) FindVillains() ([]models.Villain, error) {

	villain := models.Villain{}
	villains := []models.Villain{}

	result, err := repos.db.Database(os.Getenv("DBNAME")).
		Collection("Villain").
		Find(context.Background(), bson.D{})

	if err != nil {

		return nil, err
	}
	defer result.Close(context.Background())

	for result.Next(context.Background()) {
		if err = result.Decode(&villain); err != nil {
			return nil, err
		}

		villains = append(villains, villain)
	}

	return villains, nil
}

func (repos Villain) DeleteVillain(id string) error {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		return err
	}
	_, err = repos.db.Database(os.Getenv("DBNAME")).
		Collection("Villain").
		DeleteOne(context.Background(),
			bson.D{{Key: "_id", Value: objectId}},
		)

	if err != nil {
		return err
	}
	return nil
}

func (repo VillainFile) AddVillainFile(villainFile models.VillainFile) error {
	result, err := repo.db.Database(os.Getenv("DBNAME")).Collection("VillainFile").InsertOne(context.Background(), villainFile)
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
func (repo VillainThumbnail) addVillainThumb(villainThumbnail models.VillainThumbnail) error {
	result, err := repo.db.Database(os.Getenv("DBNAME")).
		Collection("VillainThumbnail").
		InsertOne(context.Background(),
			&villainThumbnail,
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

func (repo Villain) FindVillainByVillainName(villainName string) (models.Villain, error) {
	villain := models.Villain{}

	err := repo.db.Database(os.Getenv("DBNAME")).
		Collection("Villain").
		FindOne(context.Background(),
			bson.D{{Key: "villainName", Value: villainName}},
		).
		Decode(&villain)

	if err != nil {
		return villain, err
	}
	return villain, nil

}
