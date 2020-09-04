package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func RespondObject(w http.ResponseWriter, data interface{}, statusCode int)  {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatalln(err)
	}
}

func Respond(w http.ResponseWriter, data map[string]interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetIdFromUrl(r *http.Request, pathIndex int) (uint, error) {
	p := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(p[pathIndex])
	if err != nil {
		return 0, err
	} else {
		return uint(id), nil
	}
}
