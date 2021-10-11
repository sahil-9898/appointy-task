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

var PostCollection = db.ConnectDB().Database("appointy").Collection("posts")

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}
	insertResult, err := PostCollection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult)
	json.NewEncoder(w).Encode(insertResult.InsertedID)
}
func GetPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var id = strings.Split(r.URL.Path[1:], "/")[0]
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	err := PostCollection.FindOne(context.TODO(), bson.D{{"_id:", id}}).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(result) // returns a Map containing document
}
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var id = strings.Split(r.URL.Path[1:], "/")[0]
	filterposts, err := PostCollection.Find(context.TODO(), bson.M{"user": id})
	if err != nil {
		log.Fatal(err)
	}
	var result []bson.M
	if err = filterposts.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}
}
