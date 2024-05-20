package auth

import (
	"database/sql"

	"github.com/tonadr1022/speed-cube-time/internal/db"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

type Repository interface {
	Get(id string) (*entity.User, error)
	GetOneByUsername(username string) (*entity.User, error)
	Update(user *entity.User) error
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

func (r repository) Update(user *entity.User) error {
	_, err := r.DB.Exec("UPDATE users SET username = $1, password = $2", user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r repository) GetOneByUsername(username string) (*entity.User, error) {
	row := r.DB.QueryRow("SELECT id, username, password, created_at FROM users WHERE username = $1 LIMIT 1", username)
	if row.Err() != nil {
		return nil, row.Err()
	}
	user, err := scanIntoUser(row)
	if err != nil {
		return &entity.User{}, err
	}
	return user, nil
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
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrUserNotFound
	}
	return nil
}

func scanIntoUser(scanner db.Scanner) (*entity.User, error) {
	user := &entity.User{}
	err := scanner.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}
