package adapter

import (
	"context"

	"github.com/pawpawchat/profile/api/pb"
)

type profileDeletter interface {
	DeleteProfile(context.Context, int64) error
}

func DeleteProfileAdapter(ctx context.Context, req *pb.DeleteProfileRequest, pc profileDeletter) (*pb.DeleteProfileResponse, error) {

	return nil, nil
}
