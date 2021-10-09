package router

import (
	"appointy-task/controllers"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", controllers.CreateUser)
	mux.HandleFunc("/users/(?P<id>[a-zA-Z]+)", controllers.GetUser)
	return mux
}
