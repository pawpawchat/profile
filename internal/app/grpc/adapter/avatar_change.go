package adapter

import (
	"context"

	"github.com/pawpawchat/profile/api/pb"
	"github.com/pawpawchat/profile/pkg/status"
	"github.com/pawpawchat/profile/pkg/validation"
)

type avatarChanger interface {
	ChangeProfileAvatar(context.Context, int64, int64) error
}

func ChangeProfileAvatarAdapter(ctx context.Context, req *pb.ChangeProfileAvatarRequest, ac avatarChanger) (*pb.ChangeProfileAvatarResponse, error) {
	if emptyFields := validation.GetEmptyFields(req); len(emptyFields) != 0 {
		return nil, status.MissingFields(emptyFields)
	}

	if err := ac.ChangeProfileAvatar(ctx, req.ProfileId, req.AvatarId); err != nil {
		return nil, err
	}

	return &pb.ChangeProfileAvatarResponse{}, nil
}
