package service

import (
	"context"

	"github.com/pawpawchat/profile/internal/domain/model"
)

// Репозиторий явлвяется адаптер для базы данных
type ProfileRepository interface {
	CreateProfile(ctx context.Context, profile *model.Profile) error
	GetById(ctx context.Context, id uint64) (*model.Profile, error)
}

// Сервис обращается к репозиториям и API зависимых от операции сервисов
type ProfileService struct {
	profileRepository ProfileRepository
}

func NewProfile(profileRepository ProfileRepository) *ProfileService {
	return &ProfileService{profileRepository}
}

func (s *ProfileService) CreateProfile(ctx context.Context, profile *model.Profile) error {
	return s.profileRepository.CreateProfile(ctx, profile)
}

func (s *ProfileService) GetProfilrById(ctx context.Context, id uint64) (*model.Profile, error) {
	return s.profileRepository.GetById(ctx, id)
}
