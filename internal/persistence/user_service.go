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

func GetUser() *domain.User {
	ctx, _ := CreateContext()
	client := EstablishConnection(ctx)
	collection := getUserCollection(client)

	result := collection.FindOne(ctx, bson.D{})

	userBSON := createEmptyUserBSON()
	err := result.Decode(userBSON)

	if err != nil {
		println(fmt.Errorf(err.Error()))
	}

	urls := domain.NewUrls(userBSON.InternalUrl, userBSON.ExternalUrl, userBSON.Uri)
	img := domain.NewImage(userBSON.ImageUrl, userBSON.ImageData)
	usr := domain.NewUser(userBSON.UserId, userBSON.Name, urls, img, userBSON.Followers)
	return usr
}

func getUserCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("musicMaestro").Collection("user")
}
