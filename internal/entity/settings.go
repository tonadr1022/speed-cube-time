package entity

import "time"

type Settings struct {
	ID                  string    `json:"id"`
	ActiveCubeSessionId string    `json:"active_cube_session_id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type UpdateSettingsPayload struct {
	ActiveCubeSessionId string `json:"active_cube_session_id" validate:"omitempty,uuid"`
}

type CreateSettingsPayload struct {
	ActiveCubeSessionId string `json:"active_cube_session_id" validate:"required,uuid"`
	UserId              string `json:"user_id,omitempty"`
}
