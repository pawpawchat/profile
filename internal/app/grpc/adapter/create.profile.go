package adapter

import (
	"context"
	"time"

	"github.com/pawpawchat/profile/api/pb"
	"github.com/pawpawchat/profile/internal/domain/model"
	"github.com/pawpawchat/profile/pkg/status"
	"github.com/pawpawchat/profile/pkg/validation"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProfileCreator interface {
	CreateProfile(context.Context, *model.Profile) error
}

func CreateProfileAdapter(ctx context.Context, req *pb.CreateProfileRequest, pc ProfileCreator) (*pb.CreateProfileResponse, error) {
	if emptyFields := validation.GetEmptyFields(req); len(emptyFields) != 0 {
		return nil, status.MissingFields(emptyFields)
	}

	profile := createProfileRequestToModel(req)
	if err := pc.CreateProfile(ctx, profile); err != nil {
		return nil, err
	}

	return profileModelToCreateProfileResponse(profile)
}

func createProfileRequestToModel(r *pb.CreateProfileRequest) *model.Profile {
	return &model.Profile{
		Biography: model.Biography{
			FirstName:  r.FirstName,
			SecondName: r.SecondName,
		},
		CreatedAt: time.Now(),
	}
}

func profileModelToCreateProfileResponse(p *model.Profile) (*pb.CreateProfileResponse, error) {
	return &pb.CreateProfileResponse{
		Profile: &pb.Profile{
			Id:        p.ID,
			Username:  p.Username,
			CreatedAt: timestamppb.New(p.CreatedAt),
			Biography: &pb.Biography{
				FirstName:  p.Biography.FirstName,
				SecondName: p.Biography.SecondName,
			},
		},
	}, nil
}
