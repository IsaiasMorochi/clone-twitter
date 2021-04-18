package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/IsaiasMorochi/twitter-clone-backend/dao"
	"github.com/IsaiasMorochi/twitter-clone-backend/lib"
	"github.com/IsaiasMorochi/twitter-clone-backend/models"
)

func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	var user models.Users

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña inválidos"+err.Error(), http.StatusBadRequest)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email del usurio es requerido", http.StatusBadRequest)
		return
	}

	document, exists := dao.Login(user.Email, user.Password)
	if !exists {
		http.Error(w, "Usuario y/o contraseña inválidos", http.StatusBadRequest)
		return
	}

	jwtKey, err := lib.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el Token correspondiente"+err.Error(), http.StatusBadRequest)
		return
	}

	response := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	// Grabar cookie desde back-end
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
