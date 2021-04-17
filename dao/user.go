package dao

import (
	"context"
	"time"

	"github.com/IsaiasMorochi/twitter-clone-backend/config"
	"github.com/IsaiasMorochi/twitter-clone-backend/lib"
	"github.com/IsaiasMorochi/twitter-clone-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Post(user models.Users) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	// para quitar el espacio del contexto, y este no se quede activo.
	defer cancel()

	db := config.MongoCnx.Database("clone-twitter")
	col := db.Collection("users")

	user.Password, _ = lib.EncryptPassword(user.Password)

	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjectID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjectID.String(), true, nil

}
