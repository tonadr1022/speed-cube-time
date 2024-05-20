package solve

import (
	"context"
	"database/sql"

	"github.com/tonadr1022/speed-cube-time/internal/db"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

type Repository interface {
	Get(ctx context.Context, id string) (*entity.Solve, error)
	Create(ctx context.Context, userId string, sessionId string, solve *entity.Solve) error
	Update(ctx context.Context, s *entity.Solve) error
	QueryByUser(ctx context.Context, userId string) ([]*entity.Solve, error)
	QueryBySession(ctx context.Context, sessionId string) ([]*entity.Solve, error)
	Delete(ctx context.Context, id string) error
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

func (r repository) QueryByUser(ctx context.Context, userId string) ([]*entity.Solve, error) {
	query := "SELECT id, duration, scramble, cube_type, dnf, plus_two, notes, created_at, updated_at FROM solves WHERE user_id = $1"
	return r.queryByQuery(ctx, query, userId)
}

func (r repository) QueryBySession(ctx context.Context, sessionId string) ([]*entity.Solve, error) {
	query := "SELECT id, duration, scramble, cube_type, dnf, plus_two, notes, created_at, updated_at FROM solves WHERE session_id = $1"
	return r.queryByQuery(ctx, query, sessionId)
}

func (r repository) queryByQuery(ctx context.Context, query string, thing string) ([]*entity.Solve, error) {
	rows, err := r.DB.Query(query, thing)
	if err != nil {
		return nil, err
	}
	solves := []*entity.Solve{}
	for rows.Next() {
		solve, err := r.scanIntoSolve(rows)
		if err != nil {
			return nil, err
		}
		solves = append(solves, solve)
	}
	return solves, nil
}

func (r repository) Update(ctx context.Context, s *entity.Solve) error {
	query := `UPDATE solves SET duration = $1, scramble = $2, cube_type = $3, dnf = $4, plus_two = $5
    notes = $6, updated_at = $7`
	_, err := r.DB.Exec(query, s.Duration, s.Scramble, s.CubeType, s.Dnf, s.PlusTwo, s.Notes, s.UpdatedAt)
	return err
}

func (r repository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM solves WHERE id = $1`
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrSolveNotFound
	}
	return nil
}

func (r repository) Get(ctx context.Context, id string) (*entity.Solve, error) {
	query := "SELECT id, duration, scramble, cube_type, dnf, plus_two, notes, created_at, updated_at FROM solves WHERE id = $1"
	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		return r.scanIntoSolve(rows)
	}
	return nil, ErrSolveNotFound
}

func (r repository) Create(ctx context.Context, userId string, sessionId string, s *entity.Solve) error {
	query := `INSERT INTO solves (id,duration,scramble,cube_type,dnf,plus_two,notes,user_id,session_id) VALUES
    ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.DB.Exec(query, s.ID, s.Duration, s.Scramble, s.CubeType, s.Dnf, s.PlusTwo, s.Notes, userId, sessionId)
	return err
}

func (r repository) scanIntoSolve(scanner db.Scanner) (*entity.Solve, error) {
	solve := &entity.Solve{}
	err := scanner.Scan(
		&solve.ID,
		&solve.Duration,
		&solve.Scramble,
		&solve.CubeType,
		&solve.Dnf,
		&solve.PlusTwo,
		&solve.Notes,
		&solve.CreatedAt,
		&solve.UpdatedAt,
	)
	return solve, err
}
