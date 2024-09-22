package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	log.Println("inside moviesdb init")
}

func CreateMovie(dbClient *mongo.Client, movie Movies) (string, error) {
	coll := dbClient.Database(DB_NAME).Collection(Collection.Movies)
	res, err := coll.InsertOne(context.TODO(), movie)
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func GetMovie(dbClient *mongo.Client, id primitive.ObjectID) (Movies, error) {
	coll := dbClient.Database(DB_NAME).Collection(Collection.Movies)
	filter := bson.D{{Key: "_id", Value: id}}

	var res Movies
	err := coll.FindOne(context.TODO(), filter).Decode(&res)
	return res, err
}

func UpdateMovie(dbClient *mongo.Client, id primitive.ObjectID, movie Movies) (int64, error) {
	coll := dbClient.Database(DB_NAME).Collection(Collection.Movies)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": movie}
	res, err := coll.UpdateOne(context.TODO(), filter, update)
	return res.ModifiedCount, err
}

func DeleteMovie(dbClient *mongo.Client, id primitive.ObjectID) (int64, error) {
	coll := dbClient.Database(DB_NAME).Collection(Collection.Movies)
	filter := bson.M{"_id": id}
	res, err := coll.DeleteOne(context.TODO(), filter)
	return res.DeletedCount, err
}
