package adapter

import (
	"context"

	"github.com/pawpawchat/profile/api/pb"
	"github.com/pawpawchat/profile/internal/domain/model"
)

type profileAvatarDeletter interface {
	DeleteProfileAvatar(context.Context, *model.Avatar) error
}

func DeleteProfileAvatarAdapter(ctx context.Context, req *pb.DeleteProfileAvatarRequest, pad profileAvatarDeletter) (*pb.DeleteProfileAvatarResponse, error) {
	return nil, nil
}
