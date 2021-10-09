package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	/*
	   Connect to my cluster
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://sahil:%405596Sahil@cluster0.tsjsz.mongodb.net/appointy?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)
	fmt.Println("Database Connected")
	return client
}
