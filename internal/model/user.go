package model

type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"password,omitempty"`
	Password          string `json:"-"`
}
