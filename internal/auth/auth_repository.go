package auth

import (
	"context"
	"database/sql"

	"github.com/tonadr1022/speed-cube-time/internal/db"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

type Repository interface {
	Get(ctx context.Context, id string) (*entity.User, error)
	GetOneByUsername(ctx context.Context, username string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Create(ctx context.Context, user *entity.User) (string, error)
	Query(ctx context.Context) ([]*entity.User, error)
	QueryNoPassword(ctx context.Context) ([]*entity.User, error)
	Delete(ctx context.Context, id string) error
	GetUserAndSettings(ctx context.Context, id string) (*entity.UserAndSettings, error)
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

func (r repository) Get(ctx context.Context, id string) (*entity.User, error) {
	query := `SELECT id, username, password, created_at FROM users WHERE id = $1`
	user, err := r.getUserByQuery(ctx, query, id)
	return user, err
}

func (r repository) GetUserAndSettings(ctx context.Context, id string) (*entity.UserAndSettings, error) {
	query := `SELECT users.id, users.username, settings.active_cube_session_id, settings.created_at,
    settings.updated_at FROM users JOIN settings ON users.id = settings.user_id WHERE users.id = $1 `
	row := r.DB.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return nil, row.Err()
	}
	user := &entity.UserAndSettings{}
	err := row.Scan(
		&user.UserID,
		&user.Username,
		&user.ActiveCubeSessionID,
		&user.SettingsCreatedAt,
		&user.SettingsUpdatedAt,
	)
	return user, err
}

func (r repository) Update(ctx context.Context, user *entity.User) error {
	_, err := r.DB.ExecContext(ctx, "UPDATE users SET username = $1, password = $2 WHERE id = $1", user.Username, user.Password, user.ID)
	return err
}

func (r repository) GetOneByUsername(ctx context.Context, username string) (*entity.User, error) {
	row := r.DB.QueryRowContext(ctx, "SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1 LIMIT 1", username)
	if row.Err() != nil {
		return nil, row.Err()
	}
	user, err := scanIntoUser(row)
	if err != nil {
		return &entity.User{}, err
	}
	return user, nil
}

func (r repository) getUserByQuery(ctx context.Context, query string, thing string) (*entity.User, error) {
	rows, err := r.DB.QueryContext(ctx, query, thing)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		return scanIntoUser(rows)
	}
	return nil, sql.ErrNoRows
}

func (r repository) Create(ctx context.Context, user *entity.User) (string, error) {
	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	var id string
	err := r.DB.QueryRowContext(ctx, query, user.Username, user.Password).Scan(&id)
	return id, err
}

// TODO query builder?
func (r repository) QueryNoPassword(ctx context.Context) ([]*entity.User, error) {
	query := "SELECT id, username, created_at, updated_at FROM users"
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	users := []*entity.User{}
	for rows.Next() {
		user, err := scanIntoUserReponse(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r repository) Query(ctx context.Context) ([]*entity.User, error) {
	query := `SELECT id, username, password, created_at, updated_at FROM users`
	rows, err := r.DB.QueryContext(ctx, query)
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

func (r repository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
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

func scanIntoUserReponse(scanner db.Scanner) (*entity.User, error) {
	user := &entity.User{}
	err := scanner.Scan(
		&user.ID,
		&user.Username,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	return user, err
}

func scanIntoUser(scanner db.Scanner) (*entity.User, error) {
	user := &entity.User{}
	err := scanner.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	return user, err
}
