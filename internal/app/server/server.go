package server

import (
	"context"

	"github.com/pawpawchat/profile/api/pb"
	"github.com/pawpawchat/profile/internal/app/adapter"
	"github.com/pawpawchat/profile/internal/domain/model"
)

type ProfileService interface {
	CreateProfile(context.Context, *model.Profile) error
	DeleteProfile(context.Context, int64) error
	GetProfileByID(context.Context, int64) (*model.Profile, error)
	GetProfileByUsername(context.Context, string) (*model.Profile, error)
}

type AvatarService interface {
	SetProfileAvatar(context.Context, *model.Avatar) error
	GetAllProfileAvatars(context.Context, int64) ([]*model.Avatar, error)
	DeleteProfileAvatar(context.Context, *model.Avatar) error
}

type ProfileGRPCServer struct {
	pb.UnimplementedProfileServiceServer
	profileService ProfileService
	avatarService  AvatarService
}

func NewProfileGRPCServer(profileService ProfileService, avatarService AvatarService) *ProfileGRPCServer {
	return &ProfileGRPCServer{profileService: profileService, avatarService: avatarService}
}

func (s *ProfileGRPCServer) CreateProfile(ctx context.Context, req *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {
	return adapter.CreateProfileAdapter(ctx, req, s.profileService)
}

func (s *ProfileGRPCServer) DeleteProfile(ctx context.Context, req *pb.DeleteProfileRequest) (*pb.DeleteProfileResponse, error) {
	return adapter.DeleteProfileAdapter(ctx, req, s.profileService)
}

func (s *ProfileGRPCServer) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	return adapter.GetProfileAdapter(ctx, req, s.profileService, s.avatarService)
}

func (s *ProfileGRPCServer) SetProfileAvatar(ctx context.Context, req *pb.SetProfileAvatarRequest) (*pb.SetProfileAvatarResponse, error) {
	return adapter.SetProfileAvatar(ctx, req, s.avatarService)
}

func (s *ProfileGRPCServer) DeleteProfileAvatar(ctx context.Context, req *pb.DeleteProfileAvatarRequest) (*pb.DeleteProfileAvatarResponse, error) {
	return adapter.DeleteProfileAvatarAdapter(ctx, req, s.avatarService)
}
