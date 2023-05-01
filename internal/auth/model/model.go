package model

import "github.com/dgrijalva/jwt-go/v4"

type User struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	PhotoURL string `json:"photo_url,omitempty"`
}

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}
