package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ticketing.io/src/config"
)

func Init(conf *config.Config) (*mongo.Database, func()) {
	// Connect to the database

	fmt.Println(conf.MongoUri)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(conf.MongoUri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	
	return client.Database(conf.DatabaseName), func(){
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}
}