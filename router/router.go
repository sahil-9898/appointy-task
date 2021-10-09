package router

import (
	"appointy-task/controllers"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.Homepage)
	mux.HandleFunc("/users", controllers.CreateUser)
	mux.HandleFunc("/users/:id", controllers.GetUser)
	return mux
}
