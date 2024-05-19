package auth

import (
	"database/sql"

	"github.com/tonadr1022/speed-cube-time/internal/db"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

type Repository interface {
	GetByUsername(username string) (*entity.User, error)
	Get(id string) (*entity.User, error)
	Create(user *entity.User) error
	Query() ([]*entity.User, error)
	Delete(id string) error
}

type repository struct {
	DB *sql.DB
}

func (r repository) Get(id string) (*entity.User, error) {
	query := `SELECT id, username, password, created_at FROM users WHERE id = $1`
	user, err := r.getUserByQuery(query, id)
	return user, err
}

func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

func (r repository) GetByUsername(username string) (*entity.User, error) {
	query := `SELECT id, username, password, created_at FROM users WHERE username = $1`
	return r.getUserByQuery(query, username)
}

func (r repository) getUserByQuery(query string, thing string) (*entity.User, error) {
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

func (r repository) Create(user *entity.User) error {
	query := `INSERT INTO users (id, username, password, created_at) VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(query, user.ID, user.Username, user.Password, user.CreatedAt)
	if err != nil {
		return db.TransformError(err)
	}
	return nil
}

func (r repository) Query() ([]*entity.User, error) {
	query := `SELECT id, username, password, created_at FROM users`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	users := []*entity.User{}
	for rows.Next() {
		user, err := scanIntoUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r repository) Delete(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

func scanIntoUser(rows *sql.Rows) (*entity.User, error) {
	user := &entity.User{}
	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}
