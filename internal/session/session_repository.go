package session

import (
	"context"
	"database/sql"

	"github.com/tonadr1022/speed-cube-time/internal/db"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

type Repository interface {
	Get(ctx context.Context, d string) (*entity.Session, error)
	GetByName(ctx context.Context, name string) (*entity.Session, error)
	Update(ctx context.Context, session *entity.Session) error
	Create(ctx context.Context, userId string, session *entity.Session) error
	Query(ctx context.Context, userId string) ([]*entity.Session, error)
	Delete(ctx context.Context, id string) error
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

func (r repository) Query(ctx context.Context, userId string) ([]*entity.Session, error) {
	query := `SELECT id, name, cube_type, created_at, updated_at FROM sessions where user_id = $1`
	rows, err := r.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	sessions := []*entity.Session{}
	for rows.Next() {
		session, err := r.scanIntoSession(rows)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}

func (r repository) Update(ctx context.Context, s *entity.Session) error {
	query := "UPDATE sessions SET name = $1, cube_type = $2, updated_at = $3"
	_, err := r.DB.Exec(query, s.Name, s.CubeType, s.UpdatedAt)
	return err
}

func (r repository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM sessions WHERE id = $1`
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrSessionNotFound
	}
	return nil
}

func (r repository) GetByName(ctx context.Context, name string) (*entity.Session, error) {
	query := `SELECT id, name, cube_type, created_at, updated_at FROM sessions WHERE name = $1`
	return r.getByQuery(query, name)
}

func (r repository) Get(ctx context.Context, id string) (*entity.Session, error) {
	query := `SELECT id, name, cube_type, created_at, updated_at FROM sessions WHERE id = $1`
	return r.getByQuery(query, id)
}

func (r repository) getByQuery(query string, thing string) (*entity.Session, error) {
	rows, err := r.DB.Query(query, thing)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		return r.scanIntoSession(rows)
	}
	return nil, ErrSessionNotFound
}

func (r repository) Create(ctx context.Context, userId string, s *entity.Session) error {
	query := "INSERT INTO sessions (id, name, cube_type, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.DB.Exec(query, s.ID, s.Name, s.CubeType, userId, s.CreatedAt, s.UpdatedAt)
	return err
}

func (r repository) scanIntoSession(scanner db.Scanner) (*entity.Session, error) {
	session := &entity.Session{}
	err := scanner.Scan(
		&session.ID,
		&session.Name,
		&session.CubeType,
		&session.CreatedAt,
		&session.UpdatedAt,
	)
	return session, err
}
