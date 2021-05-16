package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/IsaiasMorochi/twitter-clone-backend/dao"
	"github.com/IsaiasMorochi/twitter-clone-backend/models"
)

func PostTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar leer el Body "+err.Error(), http.StatusBadRequest)
		return
	}

	register := models.Tweet{
		UserId:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := dao.PostTweet(register)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), http.StatusNotFound)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el Tweet", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parámetro pagina", http.StatusBadRequest)
		return
	}

	pages, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro pagina con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	page := int64(pages)
	result, status := dao.GetTweet(ID, page)
	if !status {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func DeleteTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	err := dao.DeleteTweet(ID, IDUser)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
