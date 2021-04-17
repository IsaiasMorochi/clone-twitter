package main

import (
	"log"

	"github.com/IsaiasMorochi/twitter-clone-backend/config"
	"github.com/IsaiasMorochi/twitter-clone-backend/handlers"
)

func main() {

	if config.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la Base de datos.")
		return
	}
	handlers.Manejadores()
}
