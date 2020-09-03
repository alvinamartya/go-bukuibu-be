package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-gorm/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
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
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Username string `json:"username"`
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
	errNewUser := GetDB().Create(u).Error
	if errNewUser != nil {
		return map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError
	}

	// user is exists
	if u.ID <= 0 {
		return map[string]interface{}{
			"error": "Register account is failed",
		}, http.StatusBadRequest
	}

	// create auth token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 7 * 24).Unix()
	tokenString, errTokenString := token.SignedString([]byte(os.Getenv("token_password")))
	if errTokenString != nil {
		return map[string]interface{}{
			"error": errTokenString,
		}, http.StatusInternalServerError
	}

	// auth model
	auth := Authentication{
		UserId:  u.Id,
		Token:   tokenString,
		Expired: time.Now().Add(time.Hour * 7 * 24),
	}

	// user response model
	newUser := ResponseUser{
		Id:       u.Id,
		Username: u.Username,
		Name:     u.Name,
	}

	return map[string]interface{}{
		"user": newUser,
		"auth": auth,
	}, http.StatusCreated
}

func Login(username, password string) (map[string]interface{}, int) {
	u := &User{}
	err := GetDB().Table("users").Where("username = ?", username).First(u).Error
	if err != nil {
		log.Println(err)
		return map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError
	}

	if u.Username == "" {
		return map[string]interface{}{
			"error": "user not found",
		}, http.StatusBadRequest
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return map[string]interface{}{
			"error": "Invalid password",
		}, http.StatusBadRequest
	}

	// create auth token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 7 * 24).Unix()
	tokenString, errTokenString := token.SignedString([]byte(os.Getenv("token_password")))
	if errTokenString != nil {
		return map[string]interface{}{
			"error": errTokenString,
		}, http.StatusInternalServerError
	}

	// auth model
	auth := Authentication{
		UserId:  u.Id,
		Token:   tokenString,
		Expired: time.Now().Add(time.Hour * 7 * 24),
	}

	// user response model
	newUser := ResponseUser{
		Username: u.Username,
		Name:     u.Name,
		Id:       u.Id,
	}

	return map[string]interface{}{
		"user": newUser,
		"auth": auth,
	}, http.StatusOK
}

func GetUserById(id uint) (map[string]interface{}, int) {
	u := &User{}
	err := GetDB().Table("users").Where("id = ?", id).First(u).Error
	if err != nil {
		log.Println(err)
		return map[string]interface{}{
			"error": "u not found",
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
