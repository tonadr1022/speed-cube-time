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
	"github.com/tonadr1022/speed-cube-time/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

func NewService(jwtSigningKey string, jwtTokenExpirationMinutes int, repository Repository) Service {
	return service{jwtSigningKey, jwtTokenExpirationMinutes, repository}
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Service interface {
	Login(ctx context.Context, p *entity.LoginUserPayload) (LoginResponse, error)
	Register(ctx context.Context, req *entity.RegisterUserPayload) (LoginResponse, error)
	Query(ctx context.Context) ([]*entity.User, error)
	Update(ctx context.Context, req *entity.UpdateUserPayload) error
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
	GetActiveCubeSessionId() string
}

func (s service) Login(ctx context.Context, payload *entity.LoginUserPayload) (LoginResponse, error) {
	identity, err := s.authenticate(ctx, payload.Username, payload.Password)
	if err != nil {
		return LoginResponse{}, err
	}
	tokenString, err := s.generateJWT(identity)
	return LoginResponse{tokenString}, err
}

func (s service) get(ctx context.Context, userId string) (*entity.User, error) {
	return s.repo.Get(ctx, userId)
}

func (s service) Register(ctx context.Context, req *entity.RegisterUserPayload) (LoginResponse, error) {
	hashedPassword, err := s.hashPassword(req.Password)
	if err != nil {
		return LoginResponse{}, err
	}

	// check if user exists
	_, err = s.repo.GetOneByUsername(ctx, req.Username)
	if err == nil {
		return LoginResponse{}, ErrUserExists
	}

	// create the user
	timeNow := time.Now().UTC()
	user := &entity.User{
		ID:        uuid.NewString(),
		Username:  req.Username,
		Password:  string(hashedPassword),
		CreatedAt: timeNow,
	}
	err = s.repo.Create(ctx, user)
	if err != nil {
		return LoginResponse{}, err
	}

	// login the user
	res, err := s.Login(ctx, &entity.LoginUserPayload{Username: req.Username, Password: req.Password})
	return res, err
}

func (s service) Query(ctx context.Context) ([]*entity.User, error) {
	users, err := s.repo.Query(ctx)
	var ret []*entity.User
	if err != nil {
		return ret, err
	}
	ret = append(ret, users...)

	if len(ret) == 0 {
		ret = []*entity.User{}
	}
	return ret, nil
}

func (s service) Update(ctx context.Context, req *entity.UpdateUserPayload) error {
	user, err := s.get(ctx, CurrentUser(ctx).GetID())
	if err != nil {
		return err
	}

	if req.Password != "" {
		hashedPassword, err := s.hashPassword(req.Password)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	if req.ActiveCubeSessionId != "" {
		user.ActiveCubeSessionId = req.ActiveCubeSessionId
	}

	if err != nil {
		return err
	}
	return s.repo.Update(ctx, user)
}

func (s service) DeleteUser(ctx context.Context) error {
	user := CurrentUser(ctx)
	if user == nil {
		// should never happen since this is only called with auth
		return fmt.Errorf("internal server error")
	}
	return s.repo.Delete(ctx, user.GetID())
}

func (s service) authenticate(ctx context.Context, username string, password string) (Identity, error) {
	user, err := s.repo.GetOneByUsername(ctx, username)
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
	fmt.Printf("acc %v\n", identity.GetActiveCubeSessionId())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":                identity.GetID(),
		"username":          identity.GetUsername(),
		"active_session_id": identity.GetActiveCubeSessionId(),
		"exp":               time.Now().Add(time.Duration(s.jwtTokenExpirationMinutes * int(time.Minute))).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s service) hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
