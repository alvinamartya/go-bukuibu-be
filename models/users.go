package models

import "github.com/go-gorm/gorm"

type User struct {
	gorm.Model
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (u *User) Validate() (map[string]interface{}, bool) {

}

func (u *User) Create() map[string]interface{} {

}

func Login(username, password string) map[string]interface{} {

}

func GetUser(id int) map[string]interface{} {

}
