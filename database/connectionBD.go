package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*mongoConnection: is the objecto to connect DB*/
var clientOption = options.Client().ApplyURI("mongodb+srv://admin:Docker$231997@community.n3qsb.mongodb.net/community?retryWrites=true&w=majority")

/*connectBD is a function to allow a connection with a database*/
func ConnectBD() *mongo.Client {
	client, error := mongo.Connect(context.TODO(), clientOption)
	if error != nil {
		log.Fatal(error)
	}
	error = client.Ping(context.TODO(), nil)
	if error != nil {
		log.Fatal(error)
	}
	log.Println("Connection Sussecsfull with a BD")
	return client
}
