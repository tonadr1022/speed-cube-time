package session

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/tonadr1022/speed-cube-time/internal/auth"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

type service struct {
	repo Repository
}

type Service interface {
	Get(ctx context.Context, id string) (*entity.Session, error)
	Create(ctx context.Context, req *entity.CreateSessionPayload) (*entity.Session, error)
	Update(ctx context.Context, id string, req *entity.UpdateSessionPayload) (*entity.Session, error)
	Delete(ctx context.Context, id string) error
	Query(ctx context.Context) ([]*entity.Session, error)
}

var (
	ErrSessionNotFound = errors.New("session not found")
	ErrSessionExists   = errors.New("session already exists")
)

func NewService(repository Repository) Service {
	return service{repo: repository}
}

func (s service) Update(ctx context.Context, id string, req *entity.UpdateSessionPayload) (*entity.Session, error) {
	existing, err := s.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	existing.Name = req.Name
	if req.CubeType != "" {
		existing.CubeType = req.CubeType
	}
	existing.UpdatedAt = time.Now()
	return existing, s.repo.Update(ctx, existing)
}

func (s service) Get(ctx context.Context, id string) (*entity.Session, error) {
	session, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s service) Query(ctx context.Context) ([]*entity.Session, error) {
	user := auth.CurrentUser(ctx)
	sessions, err := s.repo.Query(ctx, user.GetID())
	var ret []*entity.Session
	if err != nil {
		return ret, err
	}
	ret = append(ret, sessions...)
	if len(ret) == 0 {
		ret = []*entity.Session{}
		return ret, nil
	}
	return ret, nil
}

func (s service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s service) Create(ctx context.Context, req *entity.CreateSessionPayload) (*entity.Session, error) {
	user := auth.CurrentUser(ctx)

	cubeType := "333"
	if req.CubeType != "" {
		cubeType = req.CubeType
	}
	timeNow := time.Now().UTC()
	session := &entity.Session{
		ID:        uuid.NewString(),
		CubeType:  cubeType,
		Name:      req.Name,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
	return session, s.repo.Create(ctx, user.GetID(), session)
}
