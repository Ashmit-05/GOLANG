package netflix

import (
	"context"
	"fmt"
	"log"

	netflix "github.com/Ashmit-05/mongodbapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://jacksparrow:blackpearl@test.upsrluc.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect to db
func init() { // special function in Go that is automatically called when the code is first executed
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions) // context is used when interacting with files outside our system, like connecting to mongo db

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connected successfully")

	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("Collection instance is ready")
}

// MONGODB helpers - file

// insert 1 record
func insertOneMovie(movie netflix.Netflix) {
	result, err := collection.InsertOne(context.Background(), movie) // always pass the context when performing any operation on the db

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted the movie with id : ", result.InsertedID)
}

// update 1 record
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count : ", result.ModifiedCount)
}

// delete 1 record
func deleteOneMove(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted with count : ", result.DeletedCount)
}

// delete all records
func deleteAllMovies() {
	result, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted all records with count : ",result.DeletedCount)
}
