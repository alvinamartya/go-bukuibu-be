package controllers

import (
	"encoding/json"
	"go-bukuibu-be/models"
	"go-bukuibu-be/utils"
	"log"
	"net/http"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	u := &models.User{}

	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		log.Println(err)
		utils.Respond(w, map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError)
	}

	resp, statusCode := u.Create()
	utils.Respond(w, resp, statusCode)
}

var Login = func(w http.ResponseWriter, r *http.Request) {
	u := &models.User{}

	err := json.NewDecoder(r.Body).Decode(u)

	if err != nil {
		log.Println(err)
		utils.Respond(w, map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError)
	}

	resp, statusCode := models.Login(u.Username, u.Password)
	utils.Respond(w, resp, statusCode)
}

var GetUserById = func(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIdFromUrl(r, 3)
	if err != nil {
		log.Println(err)
		utils.Respond(w, map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError)
	}

	resp, statusCode := models.GetUserById(id)
	utils.Respond(w, resp, statusCode)
}
