package service

import (
	"context"

	"github.com/pawpawchat/profile/internal/domain/model"
)

func (s *ProfileService) GetProfileById(ctx context.Context, id int64) (*model.Profile, error) {
	return s.profileRepository.GetById(ctx, id)
}
