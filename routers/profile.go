package routers

import (
	"encoding/json"
	"net/http"

	"github.com/IsaiasMorochi/twitter-clone-backend/dao"
)

func ViewProfile(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	profile, err := dao.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
