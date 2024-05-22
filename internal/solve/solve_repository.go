package solve

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tonadr1022/speed-cube-time/internal/db"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

type Repository interface {
	Get(ctx context.Context, id string) (*entity.Solve, error)
	Create(ctx context.Context, solve *entity.Solve) (string, error)
	Update(ctx context.Context, s *entity.Solve) error
	QueryByUser(ctx context.Context, userId string) ([]*entity.Solve, error)
	QueryBySession(ctx context.Context, sessionId string) ([]*entity.Solve, error)
	Query(ctx context.Context) ([]*entity.Solve, error)
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int, error)
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

func (r repository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM solves"
	var count int
	err := r.DB.QueryRow(query).Scan(&count)
	return count, err
}

var baseQueryString = "SELECT id, duration, scramble, cube_type, dnf, plus_two, notes, cube_session_id, created_at, updated_at FROM solves"

func (r repository) Query(ctx context.Context) ([]*entity.Solve, error) {
	return r.queryByQuery(ctx, baseQueryString)
}

func (r repository) QueryByUser(ctx context.Context, userId string) ([]*entity.Solve, error) {
	query := baseQueryString + " WHERE user_id = $1"
	return r.queryByQuery(ctx, query, userId)
}

func (r repository) QueryBySession(ctx context.Context, sessionId string) ([]*entity.Solve, error) {
	query := baseQueryString + " WHERE cube_session_id = $1"
	return r.queryByQuery(ctx, query, sessionId)
}

func (r repository) queryByQuery(ctx context.Context, query string, args ...interface{}) ([]*entity.Solve, error) {
	rows, err := r.DB.QueryContext(ctx, query, args...)
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
	query := `UPDATE solves SET duration = $1, scramble = $2, cube_type = $3, dnf = $4, plus_two = $5,
    notes = $6, updated_at = $7 WHERE id = $8`
	_, err := r.DB.ExecContext(ctx, query, s.Duration, s.Scramble, s.CubeType, s.Dnf, s.PlusTwo, s.Notes, s.UpdatedAt, s.ID)
	return err
}

func (r repository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM solves WHERE id = $1`
	result, err := r.DB.ExecContext(ctx, query, id)
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
	query := "SELECT id, duration, scramble, cube_type, dnf, plus_two, notes, cube_session_id, created_at, updated_at FROM solves WHERE id = $1"
	rows, err := r.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		return r.scanIntoSolve(rows)
	}
	return nil, ErrSolveNotFound
}

func (r repository) Create(ctx context.Context, s *entity.Solve) (string, error) {
	query := `
    INSERT INTO solves (duration,scramble,cube_type,dnf,plus_two,notes,user_id,cube_session_id) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
    RETURNING id`
	var id string
	err := r.DB.QueryRowContext(ctx, query, s.Duration, s.Scramble, s.CubeType, s.Dnf, s.PlusTwo,
		s.Notes, s.UserId, s.CubeSessionId).Scan(&id)
	fmt.Printf("cs id : %v \n", s.CubeSessionId)
	return id, err
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
		&solve.CubeSessionId,
		&solve.CreatedAt,
		&solve.UpdatedAt,
	)
	return solve, err
}
