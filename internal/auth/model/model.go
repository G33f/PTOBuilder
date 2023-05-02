package model

import "github.com/dgrijalva/jwt-go/v4"

type User struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	PhotoURL string `json:"photo_url,omitempty"`
	Role     string `json:"role,omitempty"`
}

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODMyNjUxNzUuMTQxNjU2MiwiaWF0IjoxNjgzMDA1OTc1LjE0MTY1OCwidXNlcm5hbWUiOiIifQ.nYu0_AgyDqacLKT6Fvdb5ARL6yx1_Crl0sReKhfWZxc
//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODMyNjUyNzcuOTEwNDQ2MiwiaWF0IjoxNjgzMDA2MDc3LjkxMDQ0OCwidXNlcm5hbWUiOiIifQ.L5DM3Fna44cEQnCE-vuuwirtcmTDwKL5Jv2fAMLVwko
