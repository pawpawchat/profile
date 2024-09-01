package adapter

import (
	"context"

	"github.com/pawpawchat/profile/api/pb"
	"github.com/pawpawchat/profile/internal/domain/model"
	"github.com/pawpawchat/profile/pkg/status"
	"github.com/pawpawchat/profile/pkg/validation"
)

type profileGetter interface {
	GetProfileByID(context.Context, int64) (*model.Profile, error)
	GetProfileByUsername(context.Context, string) (*model.Profile, error)
}

type avatarGetter interface {
	GetAllProfileAvatars(context.Context, int64) (model.Avatars, error)
}

func GetProfileAdapter(ctx context.Context, req *pb.GetProfileRequest, pg profileGetter, ag avatarGetter) (*pb.GetProfileResponse, error) {
	if emptyFields := validation.GetEmptyFields(req); len(emptyFields) != 0 {
		return nil, status.MissingFields(emptyFields)
	}

	var err error
	var profile *model.Profile

	switch by := req.SearchBy.(type) {
	case *pb.GetProfileRequest_Id:
		profile, err = pg.GetProfileByID(ctx, by.Id)
		if err != nil {
			return nil, err
		}

	case *pb.GetProfileRequest_Username:
		profile, err = pg.GetProfileByUsername(ctx, by.Username)
		if err != nil {
			return nil, err
		}
	}

	avatars, err := ag.GetAllProfileAvatars(ctx, profile.ID)
	if err != nil {
		return nil, err
	}

	return &pb.GetProfileResponse{Profile: profile.ToPb(), Avatars: avatars.ToPb()}, nil
}
