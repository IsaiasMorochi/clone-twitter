package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/IsaiasMorochi/twitter-clone-backend/middleware"
	"github.com/IsaiasMorochi/twitter-clone-backend/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el handler y pongo a escuchar el servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middleware.CheckCnx(routers.PostUser)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckCnx(routers.Login)).Methods("POST")
	router.HandleFunc("/view-profile", middleware.CheckCnx(middleware.Validate(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/update-profile", middleware.CheckCnx(middleware.Validate(routers.PutUser))).Methods("PUT")
	router.HandleFunc("/tweet", middleware.CheckCnx(middleware.Validate(routers.PostTweet))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
