package entity

import "time"

type Solve struct {
	ID        string    `json:"id"`
	Duration  float32   `json:"duration"`
	Scramble  string    `json:"scramble"`
	CubeType  string    `json:"cube_type" `
	Dnf       bool      `json:"dnf"`
	PlusTwo   bool      `json:"plus_two"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateSolvePayload struct {
	Duration float32 `json:"duration" validate:"required,number"`
	Scramble string  `json:"scramble" validate:"omitempty"`
	CubeType string  `json:"cube_type" validate:"required,oneof=222 333 444 555 666 777 333_bf 444_bf 555_bf 333_oh clock megaminx pyraminx skewb square_1"`
	Dnf      bool    `json:"dnf" validate:"omitempty,boolean"`
	PlusTwo  bool    `json:"plus_two" validate:"omitempty,boolean"`
	Notes    string  `json:"notes" validate:"omitempty"`
}

type UpdateSolvePayload struct {
	Duration float32 `json:"duration" validate:"omitempty,number"`
	Scramble string  `json:"scramble" validate:"omitempty"`
	CubeType string  `json:"cube_type" validate:"oneof=222 333 444 555 666 777 333_bf 444_bf 555_bf 333_oh clock megaminx pyraminx skewb square_1"`
	Dnf      bool    `json:"dnf" validate:"omitempty,boolean"`
	PlusTwo  bool    `json:"plus_two" validate:"omitempty,boolean"`
	Notes    string  `json:"notes" validate:"omitempty"`
}
