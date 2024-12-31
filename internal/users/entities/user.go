package users_entities

import pkg_entities "github.com/ladmakhi81/realtime-blogs/pkg/entities"

type User struct {
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	ProfileURL string `json:"profileUrl,omitempty"`
	FirstName  string `json:"firstName,omitempty"`
	LastName   string `json:"lastName,omitempty"`
	pkg_entities.BaseEntity
}

func NewUser(email, password string) *User {
	return &User{Email: email, Password: password}
}
