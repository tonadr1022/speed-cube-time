package user

import (
	"database/sql"
)

type Repository interface {
	Get(id string) (*User, error)
	GetByUsername(username string) (*User, error)
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

func (r repository) Get(id string) (*User, error) {
	query := `SELECT id, username, password, created_at FROM "user" WHERE id = $1`
	user, err := r.getUserByQuery(query, id)
	return user, err
}

func (r repository) GetByUsername(username string) (*User, error) {
	query := `SELECT id, username, password, created_at FROM "user" WHERE username = $1`
	return r.getUserByQuery(query, username)
}

func (r repository) getUserByQuery(query string, thing string) (*User, error) {
	rows, err := r.DB.Query(query, thing)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		return scanIntoUser(rows)
	}
	return nil, sql.ErrNoRows
}

func scanIntoUser(rows *sql.Rows) (*User, error) {
	user := &User{}
	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}
