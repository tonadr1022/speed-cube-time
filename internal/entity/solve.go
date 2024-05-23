package entity

import "time"

type Solve struct {
	ID            string    `json:"id"`
	Duration      float32   `json:"duration"`
	Scramble      string    `json:"scramble"`
	CubeType      string    `json:"cube_type" `
	CubeSessionId string    `json:"cube_session_id"`
	Dnf           bool      `json:"dnf"`
	PlusTwo       bool      `json:"plus_two"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	UserId        string    `json:"user_id,omitempty"`
}

type CreateSolvePayload struct {
	Duration      float32 `json:"duration" validate:"required,number"`
	Scramble      string  `json:"scramble" validate:"omitempty"`
	CubeSessionId string  `json:"cube_session_id" validate:"required,uuid"`
	CubeType      string  `json:"cube_type" validate:"required,oneof=222 333 444 555 666 777 333_bf 444_bf 555_bf 333_oh clock megaminx pyraminx skewb square_1"`
	Dnf           bool    `json:"dnf" validate:"omitempty,boolean"`
	PlusTwo       bool    `json:"plus_two" validate:"omitempty,boolean"`
	Notes         string  `json:"notes" validate:"omitempty"`
}

type UpdateManySolvePayload struct {
	UpdateSolvePayload
	ID *string `json:"id" validate:"required,uuid"`
}
type UpdateSolvePayload struct {
	Duration      *float32 `json:"duration,omitempty" validate:"omitempty,number"`
	Scramble      *string  `json:"scramble,omitempty" validate:"omitempty"`
	CubeSessionId *string  `json:"cube_session_id,omitempty" validate:"omitempty,uuid"`
	CubeType      *string  `json:"cube_type,omitempty" validate:"omitempty,oneof=222 333 444 555 666 777 333_bf 444_bf 555_bf 333_oh clock megaminx pyraminx skewb square_1"`
	Dnf           *bool    `json:"dnf,omitempty" validate:"omitempty,boolean"`
	PlusTwo       *bool    `json:"plus_two,omitempty" validate:"omitempty,boolean"`
	Notes         *string  `json:"notes,omitempty" validate:"omitempty"`
}
