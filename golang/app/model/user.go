package model

import "app/database"

type User struct {
	Id         string `json:"id" param:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func FindUser(u *User) User {
	var user User
	database.DB.Where(u).First(&user)
	return user
}

func CreateUser(user *User) {
    database.DB.Create(user)
}