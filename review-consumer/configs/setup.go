package configs

import (
	"context"
	"fmt"
	"log"
	"review-consumer/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	// Set up client options
	clientOptions := options.Client().ApplyURI(EnvMongoURI())

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection by pinging the database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection
}

func InsertOneObj(reqObj models.ReviewReq, collection *mongo.Collection) {

	object := models.NewReview(reqObj.COURSE_ID, reqObj.REVIEWER, reqObj.MESSAGE, reqObj.RATING, reqObj.DATE)
	go collection.InsertOne(context.TODO(), object)
}

func DeleteOneObj(reqObj models.ReviewReq, collection *mongo.Collection) {

	filter := bson.D{{"course_id", reqObj.COURSE_ID}, {"reviewer", reqObj.REVIEWER}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	fmt.Println("deleted", reqObj)
	if err != nil {
		panic(err.Error())
	}
}
