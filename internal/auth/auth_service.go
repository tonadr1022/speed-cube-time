package auth

import (
	"context"
	"database/sql"
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

type LoginResponse struct {
	Token string `json:"token"`
}

type Service interface {
	Login(p *entity.LoginUserPayload) (LoginResponse, error)
	Register(req *entity.RegisterUserPayload) (LoginResponse, error)
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

func (s service) Login(payload *entity.LoginUserPayload) (LoginResponse, error) {
	identity, err := s.authenticate(payload.Username, payload.Password)
	if err != nil {
		return LoginResponse{}, err
	}
	tokenString, err := s.generateJWT(identity)
	return LoginResponse{tokenString}, err
}

func (s service) DeleteUser(ctx context.Context) error {
	user := CurrentUser(ctx)
	if user == nil {
		// should never happen since this is only called with auth
		return fmt.Errorf("internal server error")
	}
	return s.repo.Delete(user.GetID())
}

func (s service) Register(req *entity.RegisterUserPayload) (LoginResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return LoginResponse{}, err
	}

	// check if user exists
	_, err = s.repo.GetOneByUsername(req.Username)
	if err == nil {
		return LoginResponse{}, ErrUserExists
	}

	// create the user
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

	// login the user
	res, err := s.Login(&entity.LoginUserPayload{Username: req.Username, Password: req.Password})
	return res, err
}

func NewService(jwtSigningKey string, jwtTokenExpirationMinutes int, repository Repository) Service {
	return service{jwtSigningKey, jwtTokenExpirationMinutes, repository}
}

func (s service) authenticate(username string, password string) (Identity, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
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
		"exp":      time.Now().Add(time.Duration(s.jwtTokenExpirationMinutes * int(time.Minute))).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
