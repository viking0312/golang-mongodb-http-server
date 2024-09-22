package db

import (
	"context"
	"errors"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DB_NAME = "sample_mflix"
	DB_URI  = "db.uri"
)

var dbClient *mongo.Client

func CreateMongoClient(uri string) *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	var res bson.M
	if err := client.Database(DB_NAME).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&res); err != nil {
		panic(err)
	}

	log.Println("Successfully connected to MongoDB")

	return client
}

func GetDbClient() (*mongo.Client, error) {
	if dbClient == nil {
		//TODO check how to have DB connection pool
		dbUriStr, ok := viper.Get(DB_URI).(string)
		log.Println("db uri:", dbUriStr, "ok:", ok)
		if !ok {
			return dbClient, errors.New("db uri not found")
		}
		dbClient = CreateMongoClient(dbUriStr)
	}
	return dbClient, nil
}
