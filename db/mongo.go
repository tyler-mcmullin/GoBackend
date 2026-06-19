package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoClient *mongo.Client

func CreateMongoConnection() error {

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("Error: Failed to set MongoDB URI")
	}

	_, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error: Failed to connect to MongoDB")
	}

	MongoClient = client

	return nil

}
