package models

import (
	"github.com/go-gorm/gorm"
	"net/http"
	"time"
)

type Authentication struct {
	gorm.Model
	Id      uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	UserId  uint      `json:"user_id"`
	Token   string    `json:"token"`
	Expired time.Time `json:"expired"`
}

func (auth *Authentication) Create() (map[string]interface{}, int) {
	err := GetDB().Create(auth).Error
	if err != nil {
		return map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError
	}

	return map[string]interface{}{
		"message": "success",
	}, http.StatusCreated
}
