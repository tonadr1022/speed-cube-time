package test

import (
	"context"

	"github.com/tonadr1022/speed-cube-time/internal/apperrors"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
)

type MockSettingsRepository struct {
	items []*entity.Settings
}

func (m MockSettingsRepository) Get(ctx context.Context, id string) (*entity.Settings, error) {
	for _, item := range m.items {
		if item.ID == id {
			return item, nil
		}
	}
	return nil, apperrors.ErrNotFound
}

func (m MockSettingsRepository) Update(ctx context.Context, item *entity.Settings) error {
	if item.ActiveCubeSessionId == "error" {
		return ErrRepo
	}
	for i, u := range m.items {
		if u.ID == item.ID {
			m.items[i] = item
			return nil
		}
	}
	return apperrors.ErrNotFound
}

func (m MockSettingsRepository) Create(ctx context.Context, userId string, item *entity.Settings) error {
	if item.ActiveCubeSessionId == "error" {
		return ErrRepo
	}
	m.items = append(m.items, item)
	return nil
}

func (m MockSettingsRepository) Query(ctx context.Context) ([]*entity.Settings, error) {
	return m.items, nil
}

func (m MockSettingsRepository) Delete(ctx context.Context, id string) error {
	for i, j := range m.items {
		if j.ID == id {
			m.items[i] = m.items[len(m.items)-1]
			m.items = m.items[:len(m.items)-1]
			return nil
		}
	}
	return apperrors.ErrNotFound
}
