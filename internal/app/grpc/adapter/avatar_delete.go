package adapter

import (
	"context"

	"github.com/pawpawchat/profile/api/pb"
	"github.com/pawpawchat/profile/pkg/status"
	"github.com/pawpawchat/profile/pkg/validation"
)

type avatarDeletter interface {
	DeleteProfileAvatar(context.Context, int64, int64) error
}

func DeleteProfileAvatarAdapter(ctx context.Context, req *pb.DeleteProfileAvatarRequest, pd avatarDeletter) (*pb.DeleteProfileAvatarResponse, error) {
	if emptyFields := validation.GetEmptyFields(req); len(emptyFields) != 0 {
		return nil, status.MissingFields(emptyFields)
	}

	if err := pd.DeleteProfileAvatar(ctx, req.ProfileId, req.AvatarId); err != nil {
		return nil, err
	}

	return nil, nil
}
