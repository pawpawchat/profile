package adapter

import (
	"context"

	"github.com/pawpawchat/profile/api/pb"
	"github.com/pawpawchat/profile/internal/domain/model"
	"github.com/pawpawchat/profile/pkg/status"
	"github.com/pawpawchat/profile/pkg/validation"
)

type avatarSetter interface {
	AddProfileAvatar(context.Context, *model.Avatar) error
}

func AddProfileAvatar(ctx context.Context, req *pb.AddProfileAvatarRequest, as avatarSetter) (*pb.AddProfileAvatarResponse, error) {
	if emptyFields := validation.GetEmptyFields(req); len(emptyFields) != 0 {
		return nil, status.MissingFields(emptyFields)
	}

	avatar := &model.Avatar{
		ProfileID: req.ProfileId,
		OrigURL:   req.OrigUrl,
		AddedAt:   req.AddedAt.AsTime(),
	}

	if err := as.AddProfileAvatar(ctx, avatar); err != nil {
		return nil, err
	}

	return &pb.AddProfileAvatarResponse{Avatar: avatar.ToPb()}, nil
}
