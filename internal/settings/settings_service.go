package settings

import (
	"context"
	"errors"

	"github.com/tonadr1022/speed-cube-time/internal/auth"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

var ErrSettingsForUserExists = errors.New("settings for user exist")

type Service interface {
	GetByUserId(ctx context.Context, userId string) (*entity.Settings, error)
	Get(ctx context.Context, id string) (*entity.Settings, error)
	Delete(ctx context.Context, id string) error
	Create(ctx context.Context, req *entity.CreateSettingsPayload) (*entity.Settings, error)
	Update(ctx context.Context, id string, req *entity.UpdateSettingsPayload) error
	Query(ctx context.Context) ([]*entity.Settings, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) Query(ctx context.Context) ([]*entity.Settings, error) {
	return s.repo.Query(ctx)
}

func (s service) Get(ctx context.Context, id string) (*entity.Settings, error) {
	return s.repo.Get(ctx, id)
}

func (s service) GetByUserId(ctx context.Context, userId string) (*entity.Settings, error) {
	return s.repo.GetByUserId(ctx, userId)
}

func (s service) Create(ctx context.Context, req *entity.CreateSettingsPayload) (*entity.Settings, error) {
	userId := auth.CurrentUser(ctx).GetID()
	// only one settings instance can exist per user
	_, err := s.repo.GetByUserId(ctx, userId)
	if err == nil {
		return nil, ErrSettingsForUserExists
	}
	req.UserId = userId
	id, err := s.repo.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return s.repo.Get(ctx, id)
}

func (s service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s service) Update(ctx context.Context, id string, req *entity.UpdateSettingsPayload) error {
	// existing, err := s.repo.Get(ctx, id)
	// if err != nil {
	// 	return nil, err
	// }
	// existing.UpdatedAt = time.Now()
	// return existing, s.repo.Update(ctx, id, existing)
	return nil
}
