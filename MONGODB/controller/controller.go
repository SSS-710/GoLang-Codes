package controller

import (
	"context"
	"fmt"
	model "hello/MONGODB/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://shashanksss710_db_user:<Sss236710>@cluster0.fwnmvdg.mongodb.net/?appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

// most important

var collection *mongo.Collection

// connect with mongoDB

func init() {
	// client option
	clientOption := options.Client().ApplyURL(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connecgtion success")

	collection = client.Databse(dbName).Collection(colName)

	// collection instance

	fmt.Println("Collection instance is ready")
}

func init() {

}

// mongodb helpers - file
// insert 1 record

func inserOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count:", result.ModifiedCount)

}
