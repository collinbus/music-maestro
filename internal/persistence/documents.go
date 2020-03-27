package persistence

import (
	"go.mongodb.org/mongo-driver/bson"
	"musicMaestro/internal/domain"
)

type UserBSON struct {
	UserId      string
	Name        string
	ImageUrl    string
	ImageData   string
	Followers   int
	InternalUrl string
	ExternalUrl string
	Uri         string
}

func createEmptyUserBSON() *UserBSON {
	return &UserBSON{}
}

func createUserBSON(usr *domain.User) bson.D {
	return bson.D{
		{"$set", bson.M{
			"userId":      usr.Id,
			"name":        usr.Name,
			"imageUrl":    usr.Image.Url,
			"imageData":   usr.Image.Data,
			"followers":   usr.Followers,
			"internalUrl": usr.Urls.Internal,
			"externalUrl": usr.Urls.External,
			"uri":         usr.Urls.Uri,
		}},
	}
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
