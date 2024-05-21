package settings

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/tonadr1022/speed-cube-time/internal/auth"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

var (
	ErrSettingsNotFound      = errors.New("settings not found")
	ErrSettingsForUserExists = errors.New("settings for user exist")
)

type Service interface {
	GetByUserId(ctx context.Context, userId string) (*entity.Settings, error)
	Get(ctx context.Context, id string) (*entity.Settings, error)
	Delete(ctx context.Context, id string) error
	Create(ctx context.Context, req *entity.CreateSettingsPayload) (*entity.Settings, error)
	Update(ctx context.Context, id string, req *entity.UpdateSettingsPayload) (*entity.Settings, error)
	Query(ctx context.Context) ([]*entity.Settings, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) Query(ctx context.Context) ([]*entity.Settings, error) {
	user := auth.CurrentUser(ctx)
	settings, err := s.repo.Query(ctx, user.GetID())
	var ret []*entity.Settings
	if err != nil {
		return ret, err
	}
	ret = append(ret, settings...)
	if len(ret) == 0 {
		ret = []*entity.Settings{}
		return ret, nil
	}
	return ret, nil
}

func (s service) Get(ctx context.Context, id string) (*entity.Settings, error) {
	return s.repo.Get(ctx, id)
}

func (s service) GetByUserId(ctx context.Context, userId string) (*entity.Settings, error) {
	return s.repo.GetByUserId(ctx, userId)
}

func (s service) Create(ctx context.Context, req *entity.CreateSettingsPayload) (*entity.Settings, error) {
	userId := auth.CurrentUser(ctx).GetID()
	if userId == "" {
		panic("no user id in protected route")
	}

	// only one settings instance can exist per user
	_, err := s.repo.GetByUserId(ctx, userId)
	if err == nil {
		return nil, ErrSettingsForUserExists
	}

	timeNow := time.Now().UTC()
	settings := &entity.Settings{
		ID:                  uuid.NewString(),
		ActiveCubeSessionId: req.ActiveCubeSessionId,
		CreatedAt:           timeNow,
		UpdatedAt:           timeNow,
	}
	return settings, s.repo.Create(ctx, userId, settings)
}

func (s service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s service) Update(ctx context.Context, id string, req *entity.UpdateSettingsPayload) (*entity.Settings, error) {
	existing, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	existing.UpdatedAt = time.Now()
	return existing, s.repo.Update(ctx, id, existing)
}
