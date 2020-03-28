package persistence

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"musicMaestro/internal/domain"
)

func SaveTracks(tracks []domain.Track) bool {
	client := EstablishConnection(context.TODO())
	collection := getTracksCollection(client)

	var tracksBSON []interface{}
	for i := 0; i < len(tracks); i++ {
		tracksBSON = append(tracksBSON, tracks[i])
	}

	_, err := collection.InsertMany(context.TODO(), tracksBSON)

	if err != nil {
		log.Fatal(err)
	}
	return true
}

func DeleteAllTracks() bool {
	client := EstablishConnection(context.TODO())
	collection := getTracksCollection(client)

	_, err := collection.DeleteMany(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	return true
}

func getTracksCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("musicMaestro").Collection("tracks")
}
