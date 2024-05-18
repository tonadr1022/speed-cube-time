package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/tonadr1022/speed-cube-time/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Get(id string) (*User, error)
	GetByUsername(username string) (*User, error)
	Create(req *CreateUserRequest) (CreateUserReponse, error)
	Query() ([]*User, error)
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

type CreateUserReponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) Get(id string) (*User, error) {
	user, err := s.repo.Get(id)
	if err != nil {
		return &User{}, err
	}
	return &User{user}, nil
}

func (s service) GetByUsername(username string) (*User, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return &User{}, err
	}
	return &User{user}, nil
}

func (s service) Create(req *CreateUserRequest) (CreateUserReponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return CreateUserReponse{}, err
	}
	user := &models.User{
		ID:        uuid.NewString(),
		Username:  req.Username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now().UTC(),
	}

	newUser, err := s.repo.Create(user)
	if err != nil {
		return CreateUserReponse{}, nil
	}

	return CreateUserReponse{UserID: newUser.ID, Username: newUser.Username}, nil
}

func (s service) Query() ([]*User, error) {
	users, err := s.repo.Query()
	var ret []*User
	if err != nil {
		return ret, err
	}
	for _, user := range users {
		ret = append(ret, &User{user})
	}
	return ret, nil
}
