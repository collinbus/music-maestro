package persistence

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const connectionString = "mongodb://localhost:27017"

func EstablishConnection(context context.Context) *mongo.Client {
	client, err := mongo.Connect(context, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err.Error())
	}
	return client
}

func CreateContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func GetCollection(collectionName string, client *mongo.Client) *mongo.Collection {
	return client.Database("musicMaestro").Collection(collectionName)
}
