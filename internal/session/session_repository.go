package session

import (
	"context"
	"database/sql"
	"time"

	"github.com/tonadr1022/speed-cube-time/internal/db"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

type Repository interface {
	Get(ctx context.Context, d string) (*entity.Session, error)
	GetByName(ctx context.Context, name string) (*entity.Session, error)
	Update(ctx context.Context, id string, s *entity.UpdateSessionPayload) error
	Create(ctx context.Context, session *entity.Session) (string, error)
	Query(ctx context.Context) ([]*entity.Session, error)
	QueryByUser(ctx context.Context, userId string) ([]*entity.Session, error)
	Delete(ctx context.Context, id string) error
	DeleteMany(ctx context.Context, ids []string) error
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

func (r repository) DeleteMany(ctx context.Context, ids []string) error {
	return db.DeleteMany(r.DB, ctx, "sessions", ids)
}

var baseQuery = "SELECT id, name, cube_type, created_at, updated_at FROM sessions"

func (r repository) Query(ctx context.Context) ([]*entity.Session, error) {
	return r.queryByQuery(ctx, baseQuery)
}

func (r repository) QueryByUser(ctx context.Context, userId string) ([]*entity.Session, error) {
	query := baseQuery + " WHERE user_id = $1"
	return r.queryByQuery(ctx, query, userId)
}

func (r repository) queryByQuery(ctx context.Context, query string, args ...interface{}) ([]*entity.Session, error) {
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	items := []*entity.Session{}
	for rows.Next() {
		item, err := r.scanIntoSession(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r repository) Update(ctx context.Context, id string, s *entity.UpdateSessionPayload) error {
	query := `UPDATE sessions AS s SET 
    name = COALESCE(c.name, s.name), 
    cube_type = COALESCE(c.cube_type, s.cube_type), 
    updated_at = c.updated_at 
    FROM (VALUES ($1::uuid, $2::varchar(255), $3::cube_type, $4::timestamp))
    AS c(id,name,cube_type,updated_at) WHERE c.id::uuid = s.id::uuid`
	_, err := r.DB.ExecContext(ctx, query, id, s.Name, s.CubeType, time.Now().UTC())
	return err
}

func (r repository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM sessions WHERE id = $1`
	result, err := r.DB.ExecContext(ctx, query, id)
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

func (r repository) GetByName(ctx context.Context, name string) (*entity.Session, error) {
	query := `SELECT id, name, cube_type, created_at, updated_at FROM sessions WHERE name = $1`
	return r.getByQuery(ctx, query, name)
}

func (r repository) Get(ctx context.Context, id string) (*entity.Session, error) {
	query := `SELECT id, name, cube_type, created_at, updated_at FROM sessions WHERE id = $1`
	return r.getByQuery(ctx, query, id)
}

func (r repository) getByQuery(ctx context.Context, query string, thing string) (*entity.Session, error) {
	rows, err := r.DB.QueryContext(ctx, query, thing)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		return r.scanIntoSession(rows)
	}
	return nil, sql.ErrNoRows
}

func (r repository) Create(ctx context.Context, s *entity.Session) (string, error) {
	query := "INSERT INTO sessions (name, cube_type, user_id) VALUES ($1, $2, $3) RETURNING id"
	var id string
	err := r.DB.QueryRowContext(ctx, query, s.Name, s.CubeType, s.UserId).Scan(&id)
	return id, err
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
