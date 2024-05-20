package entity

import "time"

var CubeTypesSpaceDelimited = "222 333 444 555 666 777 333_bf 444_bf 555_bf 333_oh clock megaminx pyraminx skewb square_1"

type Session struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CubeType  string    `json:"cube_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateSessionPayload struct {
	Name     string `json:"name" validate:"required"`
	CubeType string `json:"cube_type" validate:"oneof=222 333 444 555 666 777 333_bf 444_bf 555_bf 333_oh clock megaminx pyraminx skewb square_1"`
}

type UpdateSessionPayload struct {
	Name     string `json:"name"`
	CubeType string `json:"cube_type" validate:"oneof=222 333 444 555 666 777 333_bf 444_bf 555_bf 333_oh clock megaminx pyraminx skewb square_1"`
}
