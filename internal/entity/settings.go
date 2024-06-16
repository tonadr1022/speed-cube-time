package entity

import "time"

type Settings struct {
	ID                  string    `json:"id"`
	ActiveCubeSessionId string    `json:"active_cube_session_id"`
	Theme               string    `json:"theme"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type UpdateSettingsPayload struct {
	ActiveCubeSessionId string `json:"active_cube_session_id" validate:"omitempty,uuid"`
	Theme               string `json:"theme" validate:"omitempty"`
}

type CreateSettingsPayload struct {
	ActiveCubeSessionId string `json:"active_cube_session_id" validate:"required,uuid"`
	UserId              string `json:"user_id,omitempty"`
	Theme               string `json:"theme,omitempty"`
}
