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
	AddProfileAvatar(context.Context, *model.Avatar) error
}

func AddProfileAvatar(ctx context.Context, req *pb.AddProfileAvatarRequest, as AvatarSetter) (*pb.AddProfileAvatarResponse, error) {
	if emptyFields := validation.GetEmptyFields(req); len(emptyFields) != 0 {
		return nil, status.MissingFields(emptyFields)
	}

	avatar := AddProfileAvatarRequestToModel(req)

	if err := as.AddProfileAvatar(ctx, avatar); err != nil {
		return nil, err
	}

	return modelToAddProfileAvatarResponse(avatar)
}

func AddProfileAvatarRequestToModel(r *pb.AddProfileAvatarRequest) *model.Avatar {
	return &model.Avatar{
		ProfileID: r.ProfileId,
		OrigURL:   r.OrigUrl,
		AddedAt:   r.AddedAt.AsTime(),
	}
}

func modelToAddProfileAvatarResponse(a *model.Avatar) (*pb.AddProfileAvatarResponse, error) {
	return &pb.AddProfileAvatarResponse{
		Avatar: &pb.Avatar{
			AvatarId: a.ID,
			OrigUrl:  a.OrigURL,
			AddedAt:  timestamppb.New(a.AddedAt),
		},
	}, nil
}
