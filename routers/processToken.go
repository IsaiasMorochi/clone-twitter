package routers

import (
	"errors"
	"strings"

	"github.com/IsaiasMorochi/twitter-clone-backend/dao"
	"github.com/IsaiasMorochi/twitter-clone-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Variable para almacenar datos de sesion*/
var Email string
var IDUser string

/*Funcion que realiza la validación del token y credencial sean validos*/
func ProcessToken(token string) (*models.Claim, bool, string, error) {

	myKey := []byte("secret_key")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	tokenN, err := jwt.ParseWithClaims(token, claims, func(tokenFunc *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, found, _ := dao.CheckIfExistsUser(claims.Email)
		if found {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}

	if !tokenN.Valid {
		return claims, false, string(""), errors.New("token Inválido")
	}

	return claims, false, string(""), err
}
