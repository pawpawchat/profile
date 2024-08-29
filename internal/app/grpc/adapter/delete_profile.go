package adapter

import (
	"context"

	"github.com/pawpawchat/profile/api/pb"
)

type ProfileDeletter interface {
	DeleteProfile(context.Context, int64) error
}

func DeleteProfileAdapter(ctx context.Context, req *pb.DeleteProfileRequest, pc ProfileDeletter) (*pb.DeleteProfileResponse, error) {

	return nil, nil
}
