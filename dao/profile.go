package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/IsaiasMorochi/twitter-clone-backend/config"
	"github.com/IsaiasMorochi/twitter-clone-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(ID string) (models.Users, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := config.MongoCnx.Database("clone-twitter")
	collection := db.Collection("users")

	var profile models.Users
	objectId, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objectId,
	}

	err := collection.FindOne(ctx, condition).Decode(&profile)
	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return profile, err
	}

	return profile, nil
}
