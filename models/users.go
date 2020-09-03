package models

import (
	"github.com/go-gorm/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type User struct {
	gorm.Model
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (u *User) Validate() (map[string]interface{}, int) {
	// invalid password
	if len(u.Password) < 6 {
		return map[string]interface{}{
			"error": "Password is required",
		}, http.StatusBadRequest
	}

	// get user
	user := User{}
	err := GetDB().Table("users").Where("username = ?", u.Username).First(user).Error

	// error
	if err != nil && user.Username != "" {
		return map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError
	}

	// username already exists
	if user.Username != "" {
		return map[string]interface{}{
			"error": "Username already exists",
		}, http.StatusBadRequest
	}

	// success
	return map[string]interface{}{
		"message": "success",
	}, http.StatusOK
}

func (u *User) Create() (map[string]interface{}, int) {
	if resp, statusCode := u.Validate(); statusCode != http.StatusOK {
		return resp, statusCode
	}

	// create hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError
	}

	u.Password = string(hashedPassword)

	// create new user
	GetDB().Create(u)

	if u.ID <= 0 {
		return map[string]interface{}{
			"error": "Register account is failed",
		}, http.StatusBadRequest
	}


}

func Login(username, password string) map[string]interface{} {

}

func GetUser(id int) map[string]interface{} {

}
