package router

import (
	"appointy-task/controllers"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", controllers.CreateUser)
	mux.HandleFunc("/users/:id", controllers.GetUser)
	mux.HandleFunc("/posts", controllers.CreatePost)
	mux.HandleFunc("/posts/:id", controllers.GetPost)
	mux.HandleFunc("/posts/users/id", controllers.GetAllPosts)
	return mux
}
