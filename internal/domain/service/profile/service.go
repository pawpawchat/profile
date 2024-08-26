package profile

import (
	"context"

	"github.com/pawpawchat/profile/internal/domain/model"
)

type ProfileRepository interface {
	Create(context.Context, *model.Profile) error
	GetByID(context.Context, int64) (*model.Profile, error)
	GetByUsername(context.Context, string) (*model.Profile, error)
}

type BiographyRepository interface {
	Create(context.Context, *model.Biography) error
}

// Service has access to repositories and APIs of other services
type ProfileService struct {
	profileRepository   ProfileRepository
	biographyRepository BiographyRepository
}

func NewProfileService(pr ProfileRepository, br BiographyRepository) *ProfileService {
	return &ProfileService{pr, br}
}
