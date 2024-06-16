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
	Create(ctx context.Context, s *entity.CreateSettingsPayload) (string, error)
	Update(ctx context.Context, s *entity.Settings) error
	Delete(ctx context.Context, id string) error
	Query(ctx context.Context) ([]*entity.Settings, error)
}

func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

type repository struct {
	DB *sql.DB
}

func (r repository) Query(ctx context.Context) ([]*entity.Settings, error) {
	query := `SELECT id, active_cube_session_id, created_at, updated_at FROM settings`
	rows, err := r.DB.QueryContext(ctx, query)
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

func (r repository) Create(ctx context.Context, s *entity.CreateSettingsPayload) (string, error) {
	query := "INSERT INTO settings (active_cube_session_id, theme, user_id) VALUES ($1, $2, $3) RETURNING id"
	var id string
	err := r.DB.QueryRowContext(ctx, query, s.ActiveCubeSessionId, s.Theme, s.UserId).Scan(&id)
	return id, err
}

func (r repository) Get(ctx context.Context, id string) (*entity.Settings, error) {
	query := "SELECT id, active_cube_session_id, theme, created_at, updated_at FROM settings WHERE user_id = $1"
	return r.getByQuery(ctx, query, id)
}

func (r repository) GetByUserId(ctx context.Context, userId string) (*entity.Settings, error) {
	query := "SELECT id, active_cube_session_id, theme, created_at, updated_at FROM settings WHERE user_id = $1"
	return r.getByQuery(ctx, query, userId)
}

func (r repository) getByQuery(ctx context.Context, query string, thing string) (*entity.Settings, error) {
	rows, err := r.DB.QueryContext(ctx, query, thing)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		return r.scanIntoSettings(rows)
	}
	return nil, sql.ErrNoRows
}

func (r repository) Update(ctx context.Context, s *entity.Settings) error {
	query := "UPDATE settings SET active_cube_session_id = $1, updated_at = $2, theme = $3 WHERE id = $4"
	_, err := r.DB.ExecContext(ctx, query, s.ActiveCubeSessionId, s.UpdatedAt, s.Theme, s.ID)
	return err
}

func (r repository) Delete(ctx context.Context, id string) error {
	result, err := r.DB.ExecContext(ctx, "DELETE FROM settings WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r repository) scanIntoSettings(scanner db.Scanner) (*entity.Settings, error) {
	settings := &entity.Settings{}
	err := scanner.Scan(
		&settings.ID,
		&settings.ActiveCubeSessionId,
		&settings.Theme,
		&settings.CreatedAt,
		&settings.CreatedAt,
	)
	return settings, err
}
