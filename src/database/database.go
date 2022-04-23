package database

import (
	"MyHeroAcademiaApi/src/config"
	"context"
	"log"
	"os"

	"github.com/juju/mgo/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//usando driver custom
func ConnectCustom() (*mgo.Session, error) {

	session, erro := mgo.Dial(config.UrlConexao)
	if erro != nil {
		log.Fatalf("Error to connect to server %s", erro)
		return nil, erro
	}
	if erro = session.Ping(); erro != nil {
		session.Close()
		return nil, erro
	}

	return session.Clone(), nil
}

//function to connect on mongo db server
func Connect() (*mongo.Client, error) {
	//doing connection with mongo
	clientOptions := options.Client().SetDirect(true).ApplyURI(os.Getenv("URL"))
	client, erro := mongo.Connect(context.TODO(), clientOptions)
	if erro != nil {
		return nil, erro
	}
	if erro = client.Ping(context.TODO(), nil); erro != nil {
		client.Disconnect(context.TODO())
		log.Fatal(erro)
		return nil, erro
	}

	return client, nil

}
