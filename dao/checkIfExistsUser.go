package dao

import (
	"context"
	"time"

	"github.com/IsaiasMorochi/twitter-clone-backend/config"
	"github.com/IsaiasMorochi/twitter-clone-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckIfExistsUser(email string) (models.Users, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := config.MongoCnx.Database("clone-twitter")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.Users

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
