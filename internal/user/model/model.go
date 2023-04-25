package model

type User struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	PhotoURL string `json:"photo_url,omitempty"`
}
