package dao

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IsaiasMorochi/twitter-clone-backend/config"
	"github.com/IsaiasMorochi/twitter-clone-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetTweet(UserId string, page int64) ([]*models.ReadTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := config.MongoCnx.Database("clone-twitter")
	collection := db.Collection("tweet")

	// creamos un slide
	var results []*models.ReadTweets

	condition := bson.M{
		"userid": UserId,
	}

	optionsOfPagination := options.Find()
	optionsOfPagination.SetLimit(20)
	optionsOfPagination.SetSort(bson.D{{Key: "date", Value: -1}}) //ordena por el campo fecha y lo trae en orden descendente.
	optionsOfPagination.SetSkip((page - 1) * 20)

	cursor, err := collection.Find(ctx, condition, optionsOfPagination)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	fmt.Println("ok" + err.Error() + " cr")

	for cursor.Next(context.TODO()) {
		var register models.ReadTweets
		err := cursor.Decode(&register)
		if err != nil {
			return results, false
		}
		results = append(results, &register)
	}

	return results, true
}

func DeleteTweet(IDTweet string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := config.MongoCnx.Database("clone-twitter")
	collection := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(IDTweet) // convierte el ID en un OBJECTID

	conditions := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := collection.DeleteOne(ctx, conditions)
	return err
}
