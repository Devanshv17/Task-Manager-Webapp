package main

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client         *mongo.Client
	userCollection *mongo.Collection
	todoCollection *mongo.Collection
)

func initMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://devanshv22:<password>@cluster0.ccgq7vm.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Set up collections
	userCollection = client.Database("your-database").Collection("users")
	todoCollection = client.Database("your-database").Collection("todos")
}

func closeMongoDB() {
	if client != nil {
		if err := client.Disconnect(nil); err != nil {
			log.Println(err)
		}
	}
}

// Other MongoDB-related functions go here...
