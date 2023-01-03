package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongo_db struct {
	client *mongo.Client
	ctx    context.Context
}

type User struct {
	Name string
}

func setupDB() *mongo_db {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// defer client.Disconnect(ctx)

	databases, err := client.ListDatabaseNames(ctx, bson.M{})

	fmt.Println("DATABASES", databases)

	dbConnection := mongo_db{
		client: client,
		ctx:    ctx,
	}

	return &dbConnection
}
