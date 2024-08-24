package service

import (
	"context"

	"github.com/pawpawchat/profile/internal/domain/model"
)

func (s *ProfileService) CreateProfile(ctx context.Context, profile *model.Profile) error {
	if err := s.profileRepository.Create(ctx, profile); err != nil {
		return err
	}

	profile.Biography.ProfileID = profile.ID
	return s.biographyRepository.Create(ctx, &profile.Biography)
}
