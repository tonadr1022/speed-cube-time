package session

import (
	"context"

	"github.com/tonadr1022/speed-cube-time/internal/auth"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

type service struct {
	repo Repository
}

type Service interface {
	Get(ctx context.Context, id string) (*entity.Session, error)
	Create(ctx context.Context, req *entity.CreateSessionPayload) (*entity.Session, error)
	Update(ctx context.Context, id string, req *entity.UpdateSessionPayload) error
	Delete(ctx context.Context, id string) error
	DeleteMany(ctx context.Context, ids []string) error
	QueryByUser(ctx context.Context, userId string) ([]*entity.Session, error)
	Query(ctx context.Context) ([]*entity.Session, error)
}

func NewService(repository Repository) Service {
	return service{repo: repository}
}

func (s service) Update(ctx context.Context, id string, req *entity.UpdateSessionPayload) error {
	return s.repo.Update(ctx, id, req)
}

func (s service) Get(ctx context.Context, id string) (*entity.Session, error) {
	session, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s service) Query(ctx context.Context) ([]*entity.Session, error) {
	return s.repo.Query(ctx)
}

func (s service) QueryByUser(ctx context.Context, userId string) ([]*entity.Session, error) {
	return s.repo.QueryByUser(ctx, userId)
}

func (s service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s service) DeleteMany(ctx context.Context, ids []string) error {
	return s.repo.DeleteMany(ctx, ids)
}

func (s service) Create(ctx context.Context, req *entity.CreateSessionPayload) (*entity.Session, error) {
	session := &entity.Session{
		Name:   req.Name,
		UserId: auth.CurrentUser(ctx).GetID(),
	}
	if req.CubeType == "" {
		session.CubeType = "333"
	} else {
		session.CubeType = req.CubeType
	}

	sessionId, err := s.repo.Create(ctx, session)
	if err != nil {
		return nil, err
	}
	return s.Get(ctx, sessionId)
}
