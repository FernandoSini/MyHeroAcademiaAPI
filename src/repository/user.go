package repository

import (
	"MyHeroAcademiaApi/src/models"
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//representa o repositorio de usuarios
type User struct {
	db *mongo.Client
}

//faz a comunicacao direta com o banco de dados
func NovoRepositorioDeUsuarios(db *mongo.Client) *User {
	return &User{db}
}

//comando para caso queira dropar a collection se der bug no metodo criar
// teste:=repo.db.Database(os.Getenv("DBNAME")).Collection("User").Drop(context.TODO())
// print(teste)

//insere um usuario no banco de dados
func (repo User) Create(user models.User) (string, error) {
	statement, erro := repo.db.Database(os.Getenv("DBNAME")).
		Collection("User").
		InsertOne(context.Background(), user)

	if erro != nil {
		return "", erro
	}
	fmt.Sprintf(" noosso %s", statement.InsertedID)

	if erro != nil {
		return "", erro
	}

	//lastIDInserted :=statement.InsertedID.(primitive.ObjectID).Hex()
	//lastIDInserted := fmt.Sprintf("%s", statement.InsertedID.(primitive.ObjectID).Hex())
	lastIDInserted := statement.InsertedID.(primitive.ObjectID).Hex()
	if erro != nil || lastIDInserted == "" || len(lastIDInserted) <= 0 {
		return "", erro
	}
	return lastIDInserted, nil

}
func (repos Hero) FindUsers() ([]models.User, error) {
	user := models.User{}
	users := []models.User{}

	result, err := repos.db.Database(os.Getenv("DBNAME")).
		Collection("User").
		Find(context.Background(), bson.D{})

	if err != nil {

		return nil, err
	}
	defer result.Close(context.Background())

	for result.Next(context.Background()) {
		if err = result.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

//find user by id
func (repo User) FindUserById(ID string) (models.User, error) {
	user := models.User{}
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return user, err
	}
	err = repo.db.Database(os.Getenv("DBNAME")).
		Collection("User").
		FindOne(context.Background(),
			bson.D{{Key: "_id", Value: objectId}},
			options.FindOne().SetProjection(bson.M{"password": 0}),
		).
		Decode(&user)

	if err != nil {
		return user, err
	}
	return user, nil
}
func (repo User) FindUserByEmail(email string) (models.User, error) {
	user := models.User{}

	err := repo.db.Database(os.Getenv("DBNAME")).
		Collection("User").
		FindOne(context.Background(),
			bson.D{{Key: "email", Value: email}},
		).
		Decode(&user)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo User) Update(id string, user models.User) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = repo.db.Database(os.Getenv("DBNAME")).
		Collection("User").UpdateOne(context.Background(), bson.D{{Key: "_id", Value: objectId}},
		bson.D{{Key: "$set", Value: &user}})
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

func (repos User) DeleteUser(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		return err
	}
	_, err = repos.db.Database(os.Getenv("DBNAME")).
		Collection("User").
		DeleteOne(context.Background(),
			bson.D{{Key: "_id", Value: objectId}},
		)

	if err != nil {
		return err
	}
	return nil
}

func (repository User) FindPassword(userId string) (string, error) {
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {

		return "", err
	}
	var user models.User
	//linha, erro := repository.db.Query("select password from usuarios where id =? ", userId)
	erro := repository.db.Database(os.Getenv("DBNAME")).
		Collection("User").
		FindOne(
			context.Background(),
			bson.D{{Key: "_id", Value: objectId}},
		).
		Decode(&user)
	if erro != nil {
		return "", erro
	}
	return user.Password, nil
}

//altera a senha de um usuario
func (repository User) UpdatePassword(userId string, password string) error {

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	//statement, erro := repository.db.Prepare("UPDATE usuarios set password = ? where id = ?")
	_, erro := repository.db.Database(os.Getenv("DBNAME")).
		Collection("User").
		UpdateByID(context.Background(), bson.D{{Key: "_id", Value: objectId}}, bson.D{{Key: "password", Value: password}})
	if erro != nil {
		return erro
	}

	return nil
}
