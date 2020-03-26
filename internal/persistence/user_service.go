package persistence

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"musicMaestro/internal/domain"
)

func SaveUser(user *domain.User) bool {
	ctx, _ := CreateContext()
	client := EstablishConnection(ctx)
	collection := getUserCollection(client)

	filter := bson.D{{"userId", user.Id}}
	updateBSON := createUserBSON(user)

	updateOptions := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(ctx, filter, updateBSON, updateOptions)

	if err != nil {
		println(fmt.Errorf(err.Error()))
		return false
	}
	return true
}

func createUserBSON(usr *domain.User) bson.D {
	return bson.D{
		{"$set", bson.M{
			"userId":      usr.Id,
			"name":        usr.Name,
			"imageUrl":    usr.ImageUrl,
			"followers":   usr.Followers,
			"internalUrl": usr.Urls.Internal,
			"externalUrl": usr.Urls.External,
			"uri":         usr.Urls.Uri,
		}},
	}
}

func getUserCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("musicMaestro").Collection("user")
}
