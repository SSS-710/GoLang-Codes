package controller

import (
	"context"
	"encoding/json"
	"fmt"
	model "hello/MONGODB/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

// delete 1 record
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(contect.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got delete with count: ", deleteCount)
}

// delete all records from mongodb

func deleteAllMovie() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

//  get all movies from databse

func getAllMovies() {
	collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primtive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.fatal(err)
		}
		movie = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

//  Actual controller  - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)

}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	inserOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}
func markAsWatched(q http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAMovie(w http.RespnseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(params["id"])
}
