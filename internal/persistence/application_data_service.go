package persistence

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"musicMaestro/internal/domain"
)

type ApplicationDataService struct{}

func (*ApplicationDataService) SaveApplicationData(applicationData *domain.ApplicationData) bool {
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

func createAppDataBSON(applicationData *domain.ApplicationData) bson.D {
	return bson.D{
		{"$set", bson.M{
			"applicationName": "musicMaestro",
			"accessCode":      applicationData.AccessCode,
			"clientId":        applicationData.ClientId,
			"clientSecret":    applicationData.ClientSecret,
			"refreshToken":    applicationData.RefreshToken,
			"tokenExpiration": applicationData.TokenExpiration,
		}},
	}
}

func getApplicationDataCollection(client *mongo.Client) *mongo.Collection {
	return GetCollection("applicationData", client)
}

func NewApplicationDataService() *ApplicationDataService {
	return &ApplicationDataService{}
}
