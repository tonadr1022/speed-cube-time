package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (u User) GetID() string {
	return u.ID
}

func (u User) ValidPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u User) GetUsername() string {
	return u.Username
}
