package routers

import (
	"encoding/json"
	"net/http"

	"github.com/IsaiasMorochi/twitter-clone-backend/dao"
	"github.com/IsaiasMorochi/twitter-clone-backend/models"
)

func PostUser(w http.ResponseWriter, r *http.Request) {

	var user models.Users

	// Body es un Object Stream, se lee y se destruye.
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", http.StatusBadRequest)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 4 carateres", http.StatusBadRequest)
		return
	}

	_, foundUser, _ := dao.CheckIfExistsUser(user.Email)
	if foundUser {
		http.Error(w, "Ya existe un usuario registrado con el email ingresado", http.StatusBadRequest)
		return
	}

	_, status, err := dao.Post(user)
	if !status {
		http.Error(w, "No se ha logrado insertar el registro del usuario"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func PutUser(w http.ResponseWriter, r *http.Request) {

	var user models.Users

	// leemos el body y lo asignamos a la variable user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), http.StatusBadRequest)
		return
	}

	status, err := dao.Put(user, IDUser)
	if err != nil {
		http.Error(w, "Ocurrió un error alintentar modificar el registro, reintente nuevamente "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
