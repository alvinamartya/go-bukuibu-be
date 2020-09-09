package main

import (
	"github.com/alvinamartya/go-bukuibu-be/controllers"
	"github.com/alvinamartya/go-bukuibu-be/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	apiPrefix := "/api/"

	// user routing
	userPrefix := apiPrefix + "user/"
	router.HandleFunc(userPrefix+"login", controllers.Login).Methods("POST")
	router.HandleFunc(userPrefix+"new", controllers.CreateUser).Methods("POST")
	router.HandleFunc(userPrefix+"{user}", controllers.GetUserById).Methods("GET")

	// port
	port, err := utils.GetEnvVar("port")
	if err != nil || port == "" {
		port = "8080"
	}

	// running server
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalln(err)
	}
}
