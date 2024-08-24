package service

import (
	"context"

	"github.com/pawpawchat/profile/internal/domain/model"
)

type ProfileRepository interface {
	Create(context.Context, *model.Profile) error
	GetById(context.Context, int64) (*model.Profile, error)
	GetByUsername(context.Context, string) (*model.Profile, error)
}

type AvatarRepository interface {
	Create(ctx context.Context, avatar *model.Avatar) error
}

type BiographyRepository interface {
	Create(ctx context.Context, bio *model.Biography) error
}

// Service has access to repositories and APIs of other services
type ProfileService struct {
	profileRepository   ProfileRepository
	biographyRepository BiographyRepository
	avatarRepository    AvatarRepository
}

func NewProfileService(pr ProfileRepository, br BiographyRepository, ar AvatarRepository) *ProfileService {
	return &ProfileService{pr, br, ar}
}
