package persistence

import (
	"go.mongodb.org/mongo-driver/bson"
	"musicMaestro/internal/domain"
)

type UserBSON struct {
	UserId    string        `bson:"userId"`
	Name      string        `bson:"name"`
	Image     *domain.Image `bson:"image"`
	Followers int           `bson:"followers"`
	Urls      *domain.Urls  `bson:"urls"`
}

func createEmptyUserBSON() *UserBSON {
	return &UserBSON{}
}

func createUserBSON(usr *domain.User) bson.D {
	return bson.D{
		{"$set", UserBSON{
			UserId:    usr.Id,
			Name:      usr.Name,
			Image:     usr.Image,
			Followers: usr.Followers,
			Urls:      usr.Urls,
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
