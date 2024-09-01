package adapter

import (
	"context"
	"time"

	"github.com/pawpawchat/profile/api/pb"
	"github.com/pawpawchat/profile/internal/domain/model"
	"github.com/pawpawchat/profile/pkg/status"
	"github.com/pawpawchat/profile/pkg/validation"
)

type profileCreator interface {
	CreateProfile(context.Context, *model.Profile) error
}

func CreateProfileAdapter(ctx context.Context, req *pb.CreateProfileRequest, pc profileCreator) (*pb.CreateProfileResponse, error) {
	if emptyFields := validation.GetEmptyFields(req); len(emptyFields) != 0 {
		return nil, status.MissingFields(emptyFields)
	}

	profile := &model.Profile{
		Biography: model.Biography{
			FirstName:  req.FirstName,
			SecondName: req.SecondName,
		},
		CreatedAt: time.Now(),
	}

	if err := pc.CreateProfile(ctx, profile); err != nil {
		return nil, err
	}

	return &pb.CreateProfileResponse{Profile: profile.ToPb()}, nil
}
