package dao

import (
	"github.com/IsaiasMorochi/twitter-clone-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (models.Users, bool) {

	user, foundUser, _ := CheckIfExistsUser(email)
	if !foundUser {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return user, false
	}

	return user, true

}
