package models

import (
	"github.com/pratikgl/bitespeed-fluxkart/pkg/config"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	PhoneNumber    string `json:"phoneNumber"`
	Email          string `json:"email"`
	LinkedId       int    `json:"linkedId"`       // the ID of another Contact linked to this one
	LinkPrecedence string `json:"linkPrecedence"` // either "primary" or "secondary"
}

func init() {
	DB = config.Connect()
	DB.AutoMigrate(&User{})
}

func GetAllUsers() []User {
	var users []User
	DB.Find(&users)
	return users
}

func CreateUser(user *User) *User {
	DB.Create(user)
	return user
}
