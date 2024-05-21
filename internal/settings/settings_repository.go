package settings

import (
	"context"
	"database/sql"

	"github.com/tonadr1022/speed-cube-time/internal/db"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

type Repository interface {
	Get(ctx context.Context, id string) (*entity.Settings, error)
	GetByUserId(ctx context.Context, userId string) (*entity.Settings, error)
	Create(ctx context.Context, userId string, s *entity.Settings) error
	Update(ctx context.Context, id string, s *entity.Settings) error
	Delete(ctx context.Context, id string) error

	Query(ctx context.Context, userId string) ([]*entity.Settings, error)
}

func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

type repository struct {
	DB *sql.DB
}

func (r repository) Query(ctx context.Context, userId string) ([]*entity.Settings, error) {
	query := `SELECT id, active_cube_session_id, created_at, updated_at FROM settings where user_id = $1`
	rows, err := r.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	settingsSlice := []*entity.Settings{}
	for rows.Next() {
		settings, err := r.scanIntoSettings(rows)
		if err != nil {
			return nil, err
		}
		settingsSlice = append(settingsSlice, settings)
	}
	return settingsSlice, nil
}

func (r repository) Create(ctx context.Context, userId string, s *entity.Settings) error {
	query := "INSERT INTO settings (id, active_cube_session_id, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.DB.Exec(query, s.ID, s.ActiveCubeSessionId, userId, s.CreatedAt, s.UpdatedAt)
	return err
}

func (r repository) Get(ctx context.Context, id string) (*entity.Settings, error) {
	query := "SELECT id, active_cube_session_id, created_at, updated_at FROM settings WHERE user_id = $1"
	return r.getByQuery(ctx, query, id)
}

func (r repository) GetByUserId(ctx context.Context, userId string) (*entity.Settings, error) {
	query := "SELECT id, active_cube_session_id, created_at, updated_at FROM settings WHERE user_id = $1"
	return r.getByQuery(ctx, query, userId)
}

func (r repository) getByQuery(ctx context.Context, query string, thing string) (*entity.Settings, error) {
	rows, err := r.DB.Query(query, thing)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		return r.scanIntoSettings(rows)
	}
	return nil, ErrSettingsNotFound
}

func (r repository) Update(ctx context.Context, id string, s *entity.Settings) error {
	query := "UPDATE settings SET active_cube_session_id = $1, updated_at = $3"
	_, err := r.DB.Exec(query, s.ActiveCubeSessionId, s.UpdatedAt)
	return err
}

func (r repository) Delete(ctx context.Context, id string) error {
	result, err := r.DB.Exec("DELETE FROM settings WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrSettingsNotFound
	}
	return nil
}

func (r repository) scanIntoSettings(scanner db.Scanner) (*entity.Settings, error) {
	settings := &entity.Settings{}
	err := scanner.Scan(
		&settings.ID,
		&settings.ActiveCubeSessionId,
		&settings.CreatedAt,
		&settings.CreatedAt,
	)
	return settings, err
}
