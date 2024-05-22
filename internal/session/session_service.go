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
	QueryByUser(ctx context.Context, userId string) ([]*entity.Session, error)
	Query(ctx context.Context) ([]*entity.Session, error)
}

func NewService(repository Repository) Service {
	return service{repo: repository}
}

func (s service) Update(ctx context.Context, id string, req *entity.UpdateSessionPayload) error {
	existing, err := s.Get(ctx, id)
	if err != nil {
		return err
	}
	existing.Name = req.Name
	if req.CubeType != "" {
		existing.CubeType = req.CubeType
	}
	// existing.UpdatedAt = time.Now().UTC()
	_, err = s.repo.Update(ctx, existing)
	return err
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

func (s service) Create(ctx context.Context, req *entity.CreateSessionPayload) (*entity.Session, error) {
	cubeType := "333"
	if req.CubeType != "" {
		cubeType = req.CubeType
	}
	session := &entity.Session{
		CubeType: cubeType,
		Name:     req.Name,
		UserId:   auth.CurrentUser(ctx).GetID(),
	}

	sessionId, err := s.repo.Create(ctx, session)
	if err != nil {
		return nil, err
	}
	return s.Get(ctx, sessionId)
}
