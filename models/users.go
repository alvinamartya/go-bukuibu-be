package models

import (
	"github.com/alvinamartya/go-bukuibu-be/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type ResponseUser struct {
	Id       uint                   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Username string                 `json:"username"`
	Name     string                 `json:"name"`
	Auth     ResponseAuthentication `json:"auth"`
}

type ResponseAuthentication struct {
	Token   string `json:"token"`
	Expired string `json:"expired"`
}

func (u *User) Validate() (map[string]interface{}, int) {
	// invalid password
	if len(u.Password) < 6 {
		return map[string]interface{}{
			"error": "Password is required",
		}, http.StatusBadRequest
	}

	// get user
	user := &User{}
	log.Println(u.Username)
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

func (u *User) Create() (*ResponseUser, map[string]interface{}, int) {
	if resp, statusCode := u.Validate(); statusCode != http.StatusOK {
		return nil, resp, statusCode
	}

	// create hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError
	}

	u.Password = string(hashedPassword)

	// create new user
	errNewUser := GetDB().Create(u).Error
	if errNewUser != nil {
		return nil, map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError
	}

	tokenString, err := utils.CreateJWTToken(int64(u.Id), u.Username)
	if err != nil {
		return nil, map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError
	}

	// user response model
	newUser := ResponseUser{
		Username: u.Username,
		Name:     u.Name,
		Id:       u.Id,
		Auth: ResponseAuthentication{
			Token:   tokenString,
			Expired: utils.ConvertTimeToString(time.Now().Add(time.Hour * 7 * 24)),
		},
	}

	return &newUser, nil, http.StatusCreated
}

func Login(username, password string) (*ResponseUser, map[string]interface{}, int) {
	u := &User{}
	err := GetDB().Table("users").Where("username = ?", username).First(u).Error
	if err != nil {
		log.Println(err)
		return nil, map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError
	}

	if u.Username == "" {
		return nil, map[string]interface{}{
			"error": "user not found",
		}, http.StatusBadRequest
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, map[string]interface{}{
			"error": "Invalid password",
		}, http.StatusBadRequest
	}

	tokenString, err := utils.CreateJWTToken(int64(u.Id), u.Username)
	if err != nil {
		return nil, map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError
	}

	// user response model
	newUser := ResponseUser{
		Username: u.Username,
		Name:     u.Name,
		Id:       u.Id,
		Auth: ResponseAuthentication{
			Token:   tokenString,
			Expired: utils.ConvertTimeToString(time.Now().Add(time.Hour * 7 * 24)),
		},
	}

	return &newUser, nil, http.StatusOK
}

func GetUserById(id uint) (map[string]interface{}, int) {
	u := &User{}
	err := GetDB().Table("users").Where("id = ?", id).First(u).Error
	if err != nil {
		log.Println(err)
		return map[string]interface{}{
			"error": "user not found",
		}, http.StatusBadRequest
	} else {
		newUser := ResponseUser{
			Id:       u.Id,
			Name:     u.Name,
			Username: u.Username,
		}

		return map[string]interface{}{
			"user": newUser,
		}, http.StatusOK
	}
}
