package lib

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword es la rutina que me permite encriptar la password recibida */
func EncryptPassword(password string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), costo)
	return string(bytes), err
}
