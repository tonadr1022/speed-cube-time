package auth

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tonadr1022/speed-cube-time/internal/apperrors"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type SettingsRepository interface {
	Create(ctx context.Context, s *entity.CreateSettingsPayload) (string, error)
}

type SessionsRepository interface {
	Create(ctx context.Context, s *entity.Session) (string, error)
}

func NewService(jwtSigningKey string, jwtTokenExpirationMinutes int, repository Repository, settingsRepo SettingsRepository, sessionsRepo SessionsRepository) Service {
	return service{jwtSigningKey, jwtTokenExpirationMinutes, repository, settingsRepo, sessionsRepo}
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Service interface {
	Get(ctx context.Context, id string) (*entity.User, error)
	GetUserAndSettings(ctx context.Context, id string) (*entity.UserAndSettings, error)
	Login(ctx context.Context, p *entity.LoginUserPayload) (LoginResponse, error)
	Register(ctx context.Context, req *entity.RegisterUserPayload) (LoginResponse, error)
	Query(ctx context.Context) ([]*entity.User, error)
	Update(ctx context.Context, req *entity.UpdateUserPayload) error
	Delete(ctx context.Context) error
}

type service struct {
	jwtSigningKey             string
	jwtTokenExpirationMinutes int
	repo                      Repository
	settingsRepo              SettingsRepository
	sessionsRepo              SessionsRepository
}

type Identity interface {
	GetID() string
	GetUsername() string
}

func (s service) GetUserAndSettings(ctx context.Context, id string) (*entity.UserAndSettings, error) {
	return s.repo.GetUserAndSettings(ctx, id)
}

// gets a user stripped of their password
func (s service) Get(ctx context.Context, id string) (*entity.User, error) {
	return s.repo.Get(ctx, id)
}

// logs in a user
func (s service) Login(ctx context.Context, payload *entity.LoginUserPayload) (LoginResponse, error) {
	identity, err := s.authenticate(ctx, payload.Username, payload.Password)
	if err != nil {
		return LoginResponse{}, err
	}
	tokenString, err := s.generateJWT(identity)
	return LoginResponse{tokenString}, err
}

// registers a user and creates settings instance and default cube session
func (s service) Register(ctx context.Context, req *entity.RegisterUserPayload) (LoginResponse, error) {
	// check if user exists
	_, err := s.repo.GetOneByUsername(ctx, req.Username)
	if err == nil {
		return LoginResponse{}, apperrors.ErrAlreadyExists
	}

	// create the user
	hashedPassword, err := s.hashPassword(req.Password)
	if err != nil {
		return LoginResponse{}, err
	}
	userId, err := s.repo.Create(ctx, &entity.User{
		Username: req.Username,
		Password: string(hashedPassword),
	})
	if err != nil {
		return LoginResponse{}, err
	}

	// create default session
	sessionId, err := s.sessionsRepo.Create(ctx, &entity.Session{
		Name:     "Default",
		CubeType: "333",
		UserId:   userId,
	})
	if err != nil {
		// delete user since session creation failed
		deleteErr := s.repo.Delete(ctx, userId)
		if deleteErr != nil {
			return LoginResponse{}, deleteErr
		}
		return LoginResponse{}, err
	}

	// create default settings
	_, err = s.settingsRepo.Create(ctx, &entity.CreateSettingsPayload{UserId: userId, ActiveCubeSessionId: sessionId, Theme: "dark"})
	if err != nil {
		// delete user since settings creation failed
		deleteErr := s.repo.Delete(ctx, userId)
		if deleteErr != nil {
			return LoginResponse{}, deleteErr
		}
		return LoginResponse{}, err
	}

	// login the user
	res, err := s.Login(ctx, &entity.LoginUserPayload{Username: req.Username, Password: req.Password})
	return res, err
}

// queries all users stripped of their passwords
func (s service) Query(ctx context.Context) ([]*entity.User, error) {
	return s.repo.Query(ctx)
}

func (s service) Update(ctx context.Context, req *entity.UpdateUserPayload) error {
	// get current
	userId := CurrentUser(ctx).GetID()
	user, err := s.Get(ctx, userId)
	if err != nil {
		return err
	}

	// check fields
	if req.Password != nil {
		hashedPassword, err := s.hashPassword(*req.Password)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	if req.Username != nil {
		user.Username = *req.Username
	}
	return s.repo.Update(ctx, user)
}

func (s service) Delete(ctx context.Context) error {
	user := CurrentUser(ctx)
	return s.repo.Delete(ctx, user.GetID())
}

func (s service) authenticate(ctx context.Context, username string, password string) (Identity, error) {
	user, err := s.repo.GetOneByUsername(ctx, username)
	if err != nil {
		// semi riskily assume the error is not found
		return nil, apperrors.ErrInvalidCredentials
	}
	if !user.ValidPassword(password) {
		return nil, apperrors.ErrInvalidCredentials
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

func (s service) hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
