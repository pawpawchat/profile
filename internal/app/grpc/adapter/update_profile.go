package adapter

import (
	"context"

	"github.com/pawpawchat/profile/api/pb"
	"github.com/pawpawchat/profile/internal/domain/model"
	"github.com/pawpawchat/profile/pkg/status"
	"google.golang.org/grpc/codes"
)

type profileUpdatter interface {
	UpdateProfile(context.Context, *model.UpdateProfileData) error
}

func UpdateProfileAdapter(ctx context.Context, pc profileUpdatter, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	var data model.UpdateProfileData

	if req.ProfileId == 0 {
		return nil, status.MissingFields([]string{"profile_id"})
	}

	data.ID = req.ProfileId
	data.Username = req.Username
	data.Description = req.Description

	if req.Biography != nil {
		data.Biography = new(model.Biography)

		if req.Biography.Birthday != nil {
			birthday := req.Biography.Birthday.AsTime()
			data.Biography.Birthday = &birthday
		}

		data.Biography.FirstName = req.Biography.FirstName
		data.Biography.SecondName = req.Biography.SecondName
	}

	if data.Biography == nil && data.Username == "" && data.Description == "" {
		return nil, status.New(codes.InvalidArgument, "no data to update")
	}

	if err := pc.UpdateProfile(ctx, &data); err != nil {
		return nil, err
	}

	return &pb.UpdateProfileResponse{}, nil
}
