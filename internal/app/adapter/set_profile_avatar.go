package adapter

import (
	"context"

	"github.com/pawpawchat/profile/api/pb"
	"github.com/pawpawchat/profile/internal/domain/model"
	"github.com/pawpawchat/profile/pkg/status"
	"github.com/pawpawchat/profile/pkg/validation"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AvatarSetter interface {
	SetProfileAvatar(context.Context, *model.Avatar) error
}

func SetProfileAvatar(ctx context.Context, req *pb.SetProfileAvatarRequest, as AvatarSetter) (*pb.SetProfileAvatarResponse, error) {
	if emptyFields := validation.GetZeroFields(req); len(emptyFields) != 0 {
		return nil, status.MissingFields(emptyFields)
	}

	avatar := setProfileAvatarRequestToModel(req)

	if err := as.SetProfileAvatar(ctx, avatar); err != nil {
		return nil, err
	}

	return modelToSetProfileAvatarResponse(avatar)
}

func setProfileAvatarRequestToModel(r *pb.SetProfileAvatarRequest) *model.Avatar {
	return &model.Avatar{
		ProfileID: r.ProfileId,
		OrigURL:   r.OrigUrl,
		AddedAt:   r.AddedAt.AsTime(),
	}
}

func modelToSetProfileAvatarResponse(a *model.Avatar) (*pb.SetProfileAvatarResponse, error) {
	return &pb.SetProfileAvatarResponse{
		Avatar: &pb.Avatar{
			AvatarId: a.ID,
			OrigUrl:  a.OrigURL,
			AddedAt:  timestamppb.New(a.AddedAt),
		},
	}, nil
}
