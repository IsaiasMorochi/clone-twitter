package middleware

import (
	"net/http"

	"github.com/IsaiasMorochi/twitter-clone-backend/config"
)

/*Los middleware reciben y lo pasa al siguiente paso.*/
func CheckCnx(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if config.CheckConnection() == 0 {
			http.Error(w, "Conexión perdida con la Base de Datos", 500)
		}
		/*next un nombre variable que identifica nuestro proximo paso.*/
		next.ServeHTTP(w, r)
	}
}
