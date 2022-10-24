package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection URI
const uri = "mongodb://localhost:27017"

func main() {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	databases, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	fmt.Printf("Successfully connected and pinged.\n databases %s", databases)

	database := client.Database("bookstore")
	booksCollection := database.Collection("books")

	cursor, err := booksCollection.Find(context.TODO(), bson.M{})

	var books []bson.M
	err = cursor.All(context.TODO(), &books)

	fmt.Printf("books\n %s", books)
}
