package users_entities

import pkg_entities "github.com/ladmakhi81/realtime-blogs/pkg/entities"

type User struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	ProfileURL string `json:"profileUrl"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	pkg_entities.BaseEntity
}

func NewUser(email, password string) *User {
	return &User{Email: email, Password: password}
}
