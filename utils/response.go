package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func HttpResponseObject(w http.ResponseWriter, data interface{}, statusCode int)  {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatalln(err)
	}
}

func HttpResponse(w http.ResponseWriter, data map[string]interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatalln(err)
	}
}
