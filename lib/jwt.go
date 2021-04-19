package lib

import (
	"time"

	"github.com/IsaiasMorochi/twitter-clone-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user models.Users) (string, error) {

	myKey := []byte("secret_key")

	payload := jwt.MapClaims{
		"email":      user.Email,
		"name":       user.Name,
		"lastname":   user.LastName,
		"birthday":   user.Birthday,
		"biography":  user.Biography,
		"website":    user.WebSite,
		"_id":        user.ID.Hex(),
		"expiration": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}
