package avatar

import (
	"context"

	"github.com/pawpawchat/profile/internal/domain/model"
)

type AvatarRepository interface {
	Create(context.Context, *model.Avatar) error
	Delete(context.Context, int64, int64) error

	Select(context.Context, int64, int64) error
	Unselect(context.Context, int64) error

	GetAll(context.Context, int64) ([]*model.Avatar, error)
}

// Service has access to repositories and APIs of other services
type AvatarService struct {
	avatarRepository AvatarRepository
}

func NewAvatarService(ar AvatarRepository) *AvatarService {
	return &AvatarService{ar}
}
