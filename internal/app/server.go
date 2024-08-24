package app

import (
	"context"
	"fmt"

	"github.com/pawpawchat/profile/api/pb"
	"github.com/pawpawchat/profile/internal/domain/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProfileGRPCServer struct {
	pb.UnimplementedProfileServiceServer
	profileService ProfileService
}

func newProfileGRPCServer(profileService ProfileService) *ProfileGRPCServer {
	return &ProfileGRPCServer{profileService: profileService}
}

type ProfileService interface {
	CreateProfile(context.Context, *model.Profile) error
	GetProfileById(context.Context, int64) (*model.Profile, error)
	GetProfileByUsername(context.Context, string) (*model.Profile, error)
}

func (s *ProfileGRPCServer) Create(ctx context.Context, req *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {
	profile := &model.Profile{
		Biography: model.Biography{
			FirstName:  req.FirstName,
			SecondName: req.SecondName,
		},
	}

	s.profileService.CreateProfile(ctx, profile)

	pbProfile := marshalProfile(profile)

	return &pb.CreateProfileResponse{Profile: pbProfile}, nil
}

func (s *ProfileGRPCServer) GetById(ctx context.Context, req *pb.GetProfileByIdRequest) (*pb.GetProfileByIdResponse, error) {
	profile, err := s.profileService.GetProfileById(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("profile not found, err=%v", err)
	}

	pbProfile := marshalProfile(profile)

	return &pb.GetProfileByIdResponse{Profile: pbProfile}, nil
}

func (s *ProfileGRPCServer) GetByUsername(ctx context.Context, req *pb.GetProfileByUsernameRequest) (*pb.GetProfileByUsernameResponse, error) {
	profile, err := s.profileService.GetProfileByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("profile not found, err=%v", err)
	}

	pbProfile := marshalProfile(profile)
	return &pb.GetProfileByUsernameResponse{Profile: pbProfile}, nil
}

func marshalProfile(p *model.Profile) *pb.Profile {
	var birthday *timestamppb.Timestamp
	if p.Biography.Birthday != nil {
		birthday = timestamppb.New(*p.Biography.Birthday)
	}

	return &pb.Profile{
		Id:              p.ID,
		Username:        p.Username,
		Description:     p.Description,
		NumberOfFriends: p.NumFriends,
		Online:          p.Online,
		LastSeen:        timestamppb.New(p.LastSeen),
		CreatedAt:       timestamppb.New(p.CreatedAt),
		Biography: &pb.Biography{
			FirstName:  p.Biography.FirstName,
			SecondName: p.Biography.SecondName,
			Birthday:   birthday,
		},
	}
}
