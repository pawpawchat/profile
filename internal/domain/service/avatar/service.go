package avatar

import (
	"context"

	"github.com/pawpawchat/profile/internal/domain/model"
)

type AvatarRepository interface {
	Create(context.Context, *model.Avatar) error
	SetAvatar(context.Context, int64, int64) error
	GetProfileAvatars(context.Context, int64) ([]*model.Avatar, error)
}

// Service has access to repositories and APIs of other services
type AvatarService struct {
	avatarRepository AvatarRepository
}

func NewAvatarService(ar AvatarRepository) *AvatarService {
	return &AvatarService{ar}
}
