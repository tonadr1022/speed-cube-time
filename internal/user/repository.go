package user

import (
	"database/sql"

	"github.com/tonadr1022/speed-cube-time/internal/models"
)

type Repository interface {
	Get(id string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Query() ([]*models.User, error)
	Delete(id string) error
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

func (r repository) Get(id string) (*models.User, error) {
	query := `SELECT id, username, password, created_at FROM users WHERE id = $1`
	user, err := r.getUserByQuery(query, id)
	return user, err
}

func (r repository) Create(user *models.User) (*models.User, error) {
	query := `INSERT INTO users (id, username, password, created_at) VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(query, user.ID, user.Username, user.Password, user.CreatedAt)
	if err != nil {
		return &models.User{}, err
	}

	// test for now
	var get_user *models.User
	get_user, err = r.GetByUsername(user.Username)
	if err != nil {
		return &models.User{}, err
	}
	return get_user, nil
}

func (r repository) GetByUsername(username string) (*models.User, error) {
	query := `SELECT id, username, password, created_at FROM users WHERE username = $1`
	return r.getUserByQuery(query, username)
}

func (r repository) Query() ([]*models.User, error) {
	query := `SELECT id, username, password, created_at FROM users`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	users := []*models.User{}
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

func (r repository) getUserByQuery(query string, thing string) (*models.User, error) {
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

func scanIntoUser(rows *sql.Rows) (*models.User, error) {
	user := &models.User{}
	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}
