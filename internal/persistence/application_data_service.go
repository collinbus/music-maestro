package persistence

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"musicMaestro/internal/domain"
)

type ApplicationDataService struct{}

func (*ApplicationDataService) SaveApplicationData(applicationData *domain.ApplicationData) bool {
	ctx, _ := CreateContext()
	client := EstablishConnection(ctx)
	appDataCollection := getApplicationDataCollection(client)

	filter := bson.D{{"applicationName", "musicMaestro"}}
	updateBSON := createAppDataBSON(applicationData)

	updateOptions := options.Update().SetUpsert(true)
	_, err := appDataCollection.UpdateOne(ctx, filter, updateBSON, updateOptions)
	if err != nil {
		println(fmt.Errorf(err.Error()))
		return false
	}

	return true
}

func (*ApplicationDataService) RetrieveApplicationData() *domain.ApplicationData {
	ctx, _ := CreateContext()
	client := EstablishConnection(ctx)
	appDataCollection := getApplicationDataCollection(client)

	filter := bson.D{{"applicationName", "musicMaestro"}}

	result := appDataCollection.FindOne(ctx, filter)

	applicationData := domain.NewApplicationData("", "", "")
	err := result.Decode(applicationData)

	if err != nil {
		println(fmt.Errorf(err.Error()))
	}

	return applicationData
}

func getApplicationDataCollection(client *mongo.Client) *mongo.Collection {
	return GetCollection("applicationData", client)
}

func NewApplicationDataService() *ApplicationDataService {
	return &ApplicationDataService{}
}
