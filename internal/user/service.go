package user

import "github.com/tonadr1022/speed-cube-time/internal/models"

type Service interface {
	Get(id string) (*User, error)
	GetByUsername(username string) (*User, error)
}

type User struct {
	*models.User
}

type service struct {
	repo Repository
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) Get(id string) (*User, error) {
	return s.repo.Get(id)
}

func (s service) GetByUsername(username string) (*User, error) {
	return s.repo.GetByUsername(username)
}
