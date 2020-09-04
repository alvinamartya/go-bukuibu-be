package main

import (
	"github.com/gorilla/mux"
	"go-bukuibu-be/controllers"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	apiPrefix := "/api/"

	// user routing
	userPrefix := apiPrefix + "user/"
	router.HandleFunc(userPrefix + "login", controllers.Login).Methods("POST")
	router.HandleFunc(userPrefix + "new", controllers.CreateUser).Methods("POST")
	router.HandleFunc(userPrefix + "{user}", controllers.GetUserById).Methods("GET")

	// port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// running server
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalln(err)
	}
}
