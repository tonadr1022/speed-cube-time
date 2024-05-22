package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserAndSettings struct {
	UserID              string `json:"user_id"`
	Username            string `json:"username"`
	ActiveCubeSessionID string `json:"active_cube_session_id"`
	SettingsCreatedAt   string `json:"settings_created_at"`
	SettingsUpdatedAt   string `json:"settings_updated_at"`
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
	Username *string `json:"username" validate:"omitempty,min=3,max=50"`
	Password *string `json:"password" validate:"omitempty,min=3,max=100"`
}

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=3,max=100"`
}
