package dao

import (
	"context"
	"time"

	"github.com/IsaiasMorochi/twitter-clone-backend/config"
	"github.com/IsaiasMorochi/twitter-clone-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PostTweet(tweet models.Tweet) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := config.MongoCnx.Database("clone-twitter")
	collection := db.Collection("tweet")

	register := bson.M{
		"userid":  tweet.UserId,
		"message": tweet.Message,
		"date":    tweet.Date,
	}

	result, err := collection.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}

	objectId, _ := result.InsertedID.(primitive.ObjectID)
	return objectId.String(), true, nil

}
