package auth

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tonadr1022/speed-cube-time/internal/db"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Service interface {
	Login(username string, password string) (LoginResponse, error)
	Register(req *RegisterUserRequest) (LoginResponse, error)
	DeleteUser(ctx context.Context) error
}

var (
	ErrUserExists         = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type service struct {
	jwtSigningKey             string
	jwtTokenExpirationMinutes int
	repo                      Repository
}

type Identity interface {
	GetID() string
	GetUsername() string
}

func (s service) Login(username string, password string) (LoginResponse, error) {
	identity, err := s.authenticate(username, password)
	if err != nil {
		return LoginResponse{}, err
	}
	// if identity == nil {
	// 	return LoginResponse{}, nil
	// }
	tokenString, err := s.generateJWT(identity)
	return LoginResponse{tokenString}, err
}

func (s service) DeleteUser(ctx context.Context) error {
	user := CurrentUser(ctx)
	if user == nil {
		fmt.Println("delete user")
		return nil
	}
	err := s.repo.Delete(user.GetID())
	if err != nil {
		return err
	}
	return nil
}

func (s service) Register(req *RegisterUserRequest) (LoginResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return LoginResponse{}, err
	}

	user := &entity.User{
		ID:        uuid.NewString(),
		Username:  req.Username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now().UTC(),
	}
	err = s.repo.Create(user)
	if err != nil {
		if errors.Is(err, db.ErrDBRowUniqueConstraint) {
			return LoginResponse{}, ErrUserExists
		}
		return LoginResponse{}, err
	}

	res, err := s.Login(user.Username, req.Password)
	return res, err
}

func NewService(jwtSigningKey string, jwtTokenExpirationMinutes int, repository Repository) Service {
	return service{jwtSigningKey, jwtTokenExpirationMinutes, repository}
}

func (s service) authenticate(username string, password string) (Identity, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if !user.ValidPassword(password) {
		return nil, ErrInvalidCredentials
	}
	return user, nil
}

func (s service) generateJWT(identity Identity) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       identity.GetID(),
		"username": identity.GetUsername(),
		"exp":      time.Now().Add(time.Duration(s.jwtTokenExpirationMinutes)).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
