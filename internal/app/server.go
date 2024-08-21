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
	GetProfilrById(context.Context, uint64) (*model.Profile, error)
}

func (s *ProfileGRPCServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	profile := &model.Profile{
		Biography: model.Biography{
			FirstName:  req.FirstName,
			SecondName: req.SecondName,
		},
	}

	s.profileService.CreateProfile(ctx, profile)

	pbProfile := marshalProfile(profile)

	return &pb.CreateResponse{Profile: pbProfile}, nil
}

func (s *ProfileGRPCServer) GetById(ctx context.Context, req *pb.GetByIdRequest) (*pb.GetByIdResponse, error) {
	profile, err := s.profileService.GetProfilrById(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("profile not found, err=%v", err)
	}

	pbProfile := marshalProfile(profile)

	return &pb.GetByIdResponse{Profile: pbProfile}, nil
}

func marshalProfile(p *model.Profile) *pb.Profile {
	var birthday *timestamppb.Timestamp
	if p.Biography.Birthday != nil {
		birthday = timestamppb.New(*p.Biography.Birthday)
	}

	return &pb.Profile{
		Id:              p.Id,
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
