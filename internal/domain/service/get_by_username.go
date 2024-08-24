package service

import (
	"context"

	"github.com/pawpawchat/profile/internal/domain/model"
)

func (s *ProfileService) GetProfileByUsername(ctx context.Context, username string) (*model.Profile, error) {
	return s.profileRepository.GetByUsername(ctx, username)
}
