package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Replace the uri string with your MongoDB deployment's connection string.
   mongodb+srv://<username>:<password>@cluster0-zzart.mongodb.net/test?retryWrites=true&w=majority
*/

var HOST = "localhost"
var PORT = 27017

/*mongoCnx es el objeto de conexión a la BD */
var mongoCnx = ConnectionDB()
var clientOptions = options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", HOST, PORT))

/*Función que permite conectar la BD*/
func ConnectionDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión exitosa con la Base de datos.")
	return client
}

/*Función que es el Ping a la BD*/
func CheckConnection() int {
	err := mongoCnx.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
