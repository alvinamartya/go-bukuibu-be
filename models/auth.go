package models

import (
	"gorm.io/gorm"
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

type ResponseAuthentication struct {
	Token   string    `json:"token"`
	Expired time.Time `json:"expired"`
}

func (auth *Authentication) Create() (*ResponseAuthentication, map[string]interface{}, int) {
	temp := &Authentication{}
	err := GetDB().Table("authentications").Where("user_id = ? AND expired >= ?", auth.UserId, time.Now()).First(temp).Error
	if err != nil {
		return nil, map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError
	}

	if temp.Id <= 0 {
		err = GetDB().Create(auth).Error
		if err != nil {
			return nil, map[string]interface{}{
				"error": err,
			}, http.StatusInternalServerError
		}

		newAuth := ResponseAuthentication{
			Token:   auth.Token,
			Expired: auth.Expired,
		}

		return &newAuth, nil, http.StatusCreated
	} else {
		newAuth := ResponseAuthentication{
			Token:   temp.Token,
			Expired: temp.Expired,
		}

		return &newAuth, nil, http.StatusOK
	}
}
