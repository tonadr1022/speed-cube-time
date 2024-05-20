package entity

import "time"

type Session struct {
	ID        string    `json:"session_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateSessionPayload struct {
	Name string `json:"name" validate:"required"`
}

type UpdateSessionPayload struct {
	Name string `json:"name" validate:"omitempty"`
}
