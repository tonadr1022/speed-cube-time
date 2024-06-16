package solve

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tonadr1022/speed-cube-time/internal/db"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

type Repository interface {
	Get(ctx context.Context, id string) (*entity.Solve, error)
	Create(ctx context.Context, solve *entity.Solve) (string, error)
	Update(ctx context.Context, id string, s *entity.UpdateSolvePayload) error
	UpdateMany(ctx context.Context, s []*entity.UpdateManySolvePayload) error
	QueryByUser(ctx context.Context, userId string) ([]*entity.Solve, error)
	QueryBySession(ctx context.Context, sessionId string) ([]*entity.Solve, error)
	Query(ctx context.Context) ([]*entity.Solve, error)
	Delete(ctx context.Context, id string) error
	DeleteMany(ctx context.Context, ids []string) error
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

func (r repository) DeleteMany(ctx context.Context, ids []string) error {
	return db.DeleteMany(r.DB, ctx, "solves", ids)
}

func (r repository) QueryByUser(ctx context.Context, userId string) ([]*entity.Solve, error) {
	query := baseQueryString + " WHERE user_id = $1 ORDER BY created_at DESC"
	return r.queryByQuery(ctx, query, userId)
}

func (r repository) QueryBySession(ctx context.Context, sessionId string) ([]*entity.Solve, error) {
	query := baseQueryString + " WHERE cube_session_id = $1 ORDER BY created_at DESC"
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

func (r repository) Update(ctx context.Context, id string, s *entity.UpdateSolvePayload) error {
	query := `UPDATE solves AS s SET 
        duration = COALESCE(c.duration, s.duration),
        scramble = COALESCE(c.scramble, s.scramble),
        cube_type = COALESCE(c.cube_type, s.cube_type), 
        dnf = COALESCE(c.dnf, s.dnf),
        plus_two = COALESCE(c.plus_two, s.plus_two),
        notes = COALESCE(c.notes, s.notes),
        updated_at = c.updated_at
        FROM (VALUES ($1::uuid,$2::float,$3::text,$4::cube_type,$5::boolean,$6::boolean,$7::text,$8::timestamp))
        AS c(id,duration,scramble,cube_type,dnf,plus_two,notes,updated_at) WHERE c.id::uuid = s.id::uuid
        `
	fmt.Println("before try access")
	_ = s.Duration
	fmt.Println("after try acces")
	_, err := r.DB.ExecContext(ctx, query, id, s.Duration, s.Scramble, s.CubeType, s.Dnf, s.PlusTwo, s.Notes, time.Now().UTC())
	return err
}

func (r repository) UpdateMany(ctx context.Context, s []*entity.UpdateManySolvePayload) error {
	query := `UPDATE solves AS s SET 
        duration = COALESCE(c.duration, s.duration),
        scramble = COALESCE(c.scramble, s.scramble),
        cube_type = COALESCE(c.cube_type, s.cube_type), 
        dnf = COALESCE(c.dnf, s.dnf),
        plus_two = COALESCE(c.plus_two, s.plus_two),
        notes = COALESCE(c.notes, s.notes),
        updated_at = c.updated_at
        FROM (VALUES
        `
	timeNow := time.Now().UTC()
	valueStrings := make([]string, 0, len(s))
	valueArgs := make([]interface{}, 0, len(s)*7)
	for i, solve := range s {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d::uuid,$%d::float,$%d::text,$%d::cube_type,$%d::boolean,$%d::boolean,$%d::text,$%d::timestamp)",
			i*8+1, i*8+2, i*8+3, i*8+4, i*8+5, i*8+6, i*8+7, i*8+8))
		valueArgs = append(valueArgs, solve.ID, solve.Duration, solve.Scramble, solve.CubeType,
			solve.Dnf, solve.PlusTwo, solve.Notes, timeNow)
	}
	query += strings.Join(valueStrings, ",\n")
	query += ") AS c(id,duration, scramble, cube_type, dnf, plus_two, notes, updated_at) WHERE c.id::uuid = s.id::uuid"
	_, err := r.DB.ExecContext(ctx, query, valueArgs...)
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
		return sql.ErrNoRows
	}
	return nil
}

func (r repository) Get(ctx context.Context, id string) (*entity.Solve, error) {
	query := "SELECT id, duration, scramble, cube_type, dnf, plus_two, notes, cube_session_id, created_at, updated_at FROM solves WHERE id = $1 LIMIT 1"
	rows, err := r.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		return r.scanIntoSolve(rows)
	}
	return nil, sql.ErrNoRows
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
