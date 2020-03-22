package persistence

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveApplicationData(applicationData *ApplicationData) bool {
	ctx, _ := CreateContext()
	client := EstablishConnection(ctx)
	appDataCollection := getApplicationDataCollection(client)

	filter := bson.D{{"applicationName", "musicMaestro"}}
	updateBSON := createAppDataBSON(applicationData)

	_, err := appDataCollection.UpdateOne(ctx, filter, updateBSON)
	if err != nil {
		println(fmt.Errorf(err.Error()))
		return false
	}

	return true
}

func createAppDataBSON(applicationData *ApplicationData) bson.D {
	return bson.D{
		{"$set", bson.M{
			"applicationName": "musicMaestro",
			"accessCode":      applicationData.accessCode,
			"clientId":        applicationData.clientId,
			"clientSecret":    applicationData.clientSecret,
		}},
	}
}

func getApplicationDataCollection(client *mongo.Client) *mongo.Collection {
	return GetCollection("applicationData", client)
}
