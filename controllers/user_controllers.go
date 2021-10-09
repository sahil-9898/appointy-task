package controllers

import (
	db "appointy-task/config"
	"appointy-task/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UserCollection = db.ConnectDB().Database("appointy").Collection("users")

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to homepage")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}
	insertResult, err := UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult)
	json.NewEncoder(w).Encode(insertResult.InsertedID)
}
func GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var id = strings.Split(r.URL.Path[1:], "/")[0]
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	err := UserCollection.FindOne(context.TODO(), bson.D{{"_id:", id}}).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(result) // returns a Map containing document

}
