package services

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/tonadr1022/speed-cube-time/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	DB *sql.DB
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *UserService) CreateUser(m CreateUserRequest) (string, error) {
	user_id := uuid.NewString()
	query := "INSERT INTO users (id,username,password,created_at) VALUES (? ? ? ?)"
	created_at := time.Now()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	_, err = s.DB.Exec(query, user_id, m.Username, hashedPassword, created_at)
	if err != nil {
		return "", nil
	}
	return user_id, nil
}

func (s *UserService) GetUsers() ([]*models.User, error) {
	query := "SELECT id, username, password, created_at FROM users"
	rows, err := s.DB.Query(query)
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

func (s *UserService) GetUserById(id string) (*models.User, error) {
	query := "SELECT id, username, password, created_at FROM users WHERE id = $1"
	return s.getUserByQuery(query, id)
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, password, created_at FROM users WHERE username = $1"
	return s.getUserByQuery(query, username)
}

func (s *UserService) getUserByQuery(query string, args ...interface{}) (*models.User, error) {
	rows, err := s.DB.Query(query, args)
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
	user := models.User{}
	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	)
	return &user, err
}
