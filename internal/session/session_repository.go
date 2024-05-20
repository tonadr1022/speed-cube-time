package session

import (
	"database/sql"

	"github.com/tonadr1022/speed-cube-time/internal/db"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

type Repository interface {
	Get(id string) (*entity.Session, error)
	Update(session *entity.Session) error
	Create(userId string, session *entity.Session) error
	// Query() ([]*entity.Session, error)
	Delete(id string) error
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

func (r repository) Update(session *entity.Session) error {
	_, err := r.DB.Exec("UPDATE users SET name = $1 ", session.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r repository) Delete(id string) error {
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

func (r repository) Get(id string) (*entity.Session, error) {
	query := `SELECT id, name, created_at, updated_at FROM sessions WHERE id = $1`
	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		return r.scanIntoSession(rows)
	}
	return nil, sql.ErrNoRows
}

func (r repository) Create(userId string, s *entity.Session) error {
	query := "INSERT INTO sessions (id, name, user_id) VALUES ($1, $2, $3)"
	_, err := r.DB.Exec(query, s.ID, s.Name, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r repository) scanIntoSession(scanner db.Scanner) (*entity.Session, error) {
	session := &entity.Session{}
	err := scanner.Scan(
		&session.ID,
		&session.Name,
		&session.CreatedAt,
		&session.UpdatedAt,
	)
	return session, err
}
