package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client         *mongo.Client
	userCollection *mongo.Collection
	todoCollection *mongo.Collection
)

func initMongoDB(connectionURI, databaseName string) {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionURI).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	var err error
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database(databaseName).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		log.Fatal(err)
	}

	// Set up collections
	userCollection = client.Database(databaseName).Collection("users")
	todoCollection = client.Database(databaseName).Collection("todos")
}

func closeMongoDB() {
	if client != nil {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Println(err)
		}
	}
}
