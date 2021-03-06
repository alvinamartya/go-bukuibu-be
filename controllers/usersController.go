package controllers

import (
	"encoding/json"
	"github.com/alvinamartya/go-bukuibu-be/models"
	"github.com/alvinamartya/go-bukuibu-be/utils"
	"log"
	"net/http"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	u := &models.User{}

	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		log.Println(err)
		utils.HttpResponse(w, map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError)
	}

	respSuccess, respErr, statusCode := u.Create()
	if respSuccess != nil {
		utils.HttpResponseObject(w, respSuccess, statusCode)
	} else {
		utils.HttpResponse(w, respErr, statusCode)
	}
}

var Login = func(w http.ResponseWriter, r *http.Request) {
	u := &models.User{}

	err := json.NewDecoder(r.Body).Decode(u)

	if err != nil {
		log.Println(err)
		utils.HttpResponse(w, map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError)
	}

	respSuccess, respErr, statusCode := models.Login(u.Username, u.Password)
	if respSuccess != nil {
		utils.HttpResponseObject(w, respSuccess, statusCode)
	} else {
		utils.HttpResponse(w, respErr, statusCode)
	}
}

var GetUserById = func(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIdFromUrl(r, 3)
	if err != nil {
		log.Println(err)
		utils.HttpResponse(w, map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError)
	}

	resp, statusCode := models.GetUserById(id)
	utils.HttpResponse(w, resp, statusCode)
}
