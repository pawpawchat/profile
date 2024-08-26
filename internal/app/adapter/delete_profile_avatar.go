package adapter

import (
	"context"

	"github.com/pawpawchat/profile/api/pb"
	"github.com/pawpawchat/profile/internal/domain/model"
)

type ProfileAvatarDeletter interface {
	DeleteProfileAvatar(context.Context, *model.Avatar) error
}

func DeleteProfileAvatarAdapter(ctx context.Context, req *pb.DeleteProfileAvatarRequest, pad ProfileAvatarDeletter) (*pb.DeleteProfileAvatarResponse, error) {
	return nil, nil
}
