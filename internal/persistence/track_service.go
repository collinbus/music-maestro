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

	tracksBSON := createTracksBSON(tracks)

	_, err := collection.InsertMany(context.TODO(), tracksBSON)

	if err != nil {
		log.Fatal(err)
	}
	return true
}

func RetrieveTracks() []domain.Track {
	client := EstablishConnection(context.TODO())
	collection := getTracksCollection(client)

	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	tracks := readElements(cursor)
	return tracks
}

func readElements(cursor *mongo.Cursor) []domain.Track {
	var tracks []domain.Track
	for cursor.Next(context.TODO()) {
		var track domain.Track
		err := cursor.Decode(&track)

		if err != nil {
			log.Fatal(err)
		}

		tracks = append(tracks, track)
	}
	return tracks
}

func createTracksBSON(tracks []domain.Track) []interface{} {
	var tracksBSON []interface{}
	for i := 0; i < len(tracks); i++ {
		tracksBSON = append(tracksBSON, tracks[i])
	}
	return tracksBSON
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
