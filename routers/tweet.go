package routers

import (
	"encoding/json"
	"net/http"
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
