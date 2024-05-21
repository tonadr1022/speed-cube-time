package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id"`
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

type LoginUserPayload struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=3,max=100"`
}

type UpdateUserPayload struct {
	Username            string `json:"username" validate:"omitempty,min=3,max=50"`
	Password            string `json:"password" validate:"omitempty,min=3,max=100"`
	ActiveCubeSessionId string `json:"active_session_id" validate:"omitempty,uuid"`
}

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=3,max=100"`
}
