package session

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/tonadr1022/speed-cube-time/internal/auth"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

type service struct {
	repo Repository
}

type Service interface {
	Get(id string) (*entity.Session, error)
	Create(ctx context.Context, req *entity.CreateSessionPayload) (*entity.Session, error)
	Update(ctx context.Context, req *entity.UpdateSessionPayload) (*entity.Session, error)
	Delete(ctx context.Context, id string) error
}

var ErrSessionNotFound = errors.New("session not found")

func NewService(repository Repository) Service {
	return service{repo: repository}
}

func (s service) Update(ctx context.Context, req *entity.UpdateSessionPayload) (*entity.Session, error) {
	return s.repo.Update(&entity.Session{Name: req.Name})
}

func (s service) Get(id string) (*entity.Session, error) {
	session, err := s.repo.Get(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrSessionNotFound
		}
		return nil, err
	}
	return session, nil
}

func (s service) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrSessionNotFound
		}
		return err
	}
	return nil
}

func (s service) Create(ctx context.Context, req *entity.CreateSessionPayload) (*entity.Session, error) {
	user := auth.CurrentUser(ctx)
	if user.GetID() == "" {
		panic("need valid user")
	}
	session := &entity.Session{
		ID:   uuid.NewString(),
		Name: req.Name,
	}
	return session, s.repo.Create(user.GetID(), session)
}
