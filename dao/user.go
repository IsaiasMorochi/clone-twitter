package dao

import (
	"context"
	"time"

	"github.com/IsaiasMorochi/twitter-clone-backend/config"
	"github.com/IsaiasMorochi/twitter-clone-backend/lib"
	"github.com/IsaiasMorochi/twitter-clone-backend/models"
	"go.mongodb.org/mongo-driver/bson"
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

func Put(user models.Users, ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := config.MongoCnx.Database("clone-twitter")
	collection := db.Collection("users")

	register := make(map[string]interface{})

	if len(user.Name) > 0 {
		register["name"] = user.Name
	}

	if len(user.LastName) > 0 {
		register["lastname"] = user.LastName
	}

	register["birthday"] = user.Birthday

	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}

	if len(user.Banner) > 0 {
		register["banner"] = user.Banner
	}

	if len(user.Biography) > 0 {
		register["biography"] = user.Biography
	}

	if len(user.Location) > 0 {
		register["location"] = user.Location
	}

	if len(user.WebSite) > 0 {
		register["website"] = user.WebSite
	}

	updateString := bson.M{
		"$set": register,
	}

	objectId, _ := primitive.ObjectIDFromHex(ID)

	// filtro para buscar el ID a actualizar de la BD ($eq = equals)
	filter := bson.M{"_id": bson.M{"$eq": objectId}}

	_, err := collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}
