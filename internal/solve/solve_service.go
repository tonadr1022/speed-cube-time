package solve

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/tonadr1022/speed-cube-time/internal/auth"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

var ErrSolveNotFound = errors.New("solve not found")

type Service interface {
	Get(ctx context.Context, id string) (*entity.Solve, error)
	Create(ctx context.Context, req *entity.CreateSolvePayload) (*entity.Solve, error)
	Update(ctx context.Context, id string, req *entity.UpdateSolvePayload) (*entity.Solve, error)
	Delete(ctx context.Context, id string) error
	Query(ctx context.Context) ([]*entity.Solve, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) Get(ctx context.Context, id string) (*entity.Solve, error) {
	solve, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return solve, nil
}

func (s service) Create(ctx context.Context, req *entity.CreateSolvePayload) (*entity.Solve, error) {
	user := auth.CurrentUser(ctx)
	timeNow := time.Now().UTC()
	solve := &entity.Solve{
		ID:        uuid.NewString(),
		CubeType:  req.CubeType,
		Scramble:  req.Scramble,
		Dnf:       req.Dnf,
		PlusTwo:   req.PlusTwo,
		Notes:     req.Notes,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
	sessionId := "active session Id of user"
	s.repo.Create(ctx, user.GetID(), sessionId, solve)

	return nil, nil
}

func (s service) Update(ctx context.Context, id string, req *entity.UpdateSolvePayload) (*entity.Solve, error) {
	return nil, nil
}

func (s service) Delete(ctx context.Context, id string) error {
	return nil
}

func (s service) Query(ctx context.Context) ([]*entity.Solve, error) {
	return nil, nil
}
