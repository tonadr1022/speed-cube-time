package solve

import (
	"context"
	"errors"

	"github.com/tonadr1022/speed-cube-time/internal/auth"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

var ErrSolveNotFound = errors.New("solve not found")

type Service interface {
	Get(ctx context.Context, id string) (*entity.Solve, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, req *entity.CreateSolvePayload) (*entity.Solve, error)
	Update(ctx context.Context, id string, req *entity.UpdateSolvePayload) error
	UpdateMany(ctx context.Context, req []*entity.UpdateManySolvePayload) error
	Delete(ctx context.Context, id string) error
	DeleteMany(ctx context.Context, ids []string) error
	QueryForUser(ctx context.Context, userId string) ([]*entity.Solve, error)
	Query(ctx context.Context) ([]*entity.Solve, error)
	QueryForSession(ctx context.Context, sessionId string) ([]*entity.Solve, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

func (s service) Get(ctx context.Context, id string) (*entity.Solve, error) {
	return s.repo.Get(ctx, id)
}

func (s service) Create(ctx context.Context, req *entity.CreateSolvePayload) (*entity.Solve, error) {
	solve := &entity.Solve{
		Duration:      req.Duration,
		CubeType:      req.CubeType,
		Scramble:      req.Scramble,
		CubeSessionId: req.CubeSessionId,
		Dnf:           req.Dnf,
		PlusTwo:       req.PlusTwo,
		Notes:         req.Notes,
		UserId:        auth.CurrentUser(ctx).GetID(),
	}
	id, err := s.repo.Create(ctx, solve)
	if err != nil {
		return nil, err
	}
	return s.Get(ctx, id)
}

func (s service) Update(ctx context.Context, id string, req *entity.UpdateSolvePayload) error {
	return s.repo.Update(ctx, id, req)
}

func (s service) UpdateMany(ctx context.Context, req []*entity.UpdateManySolvePayload) error {
	return s.repo.UpdateMany(ctx, req)
}

func (s service) DeleteMany(ctx context.Context, ids []string) error {
	return s.repo.DeleteMany(ctx, ids)
}

func (s service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s service) QueryForSession(ctx context.Context, sessionId string) ([]*entity.Solve, error) {
	return s.repo.QueryBySession(ctx, sessionId)
}

func (s service) Query(ctx context.Context) ([]*entity.Solve, error) {
	return s.repo.Query(ctx)
}

func (s service) QueryForUser(ctx context.Context, userId string) ([]*entity.Solve, error) {
	return s.repo.QueryByUser(ctx, userId)
}
