package server

import (
	"context"

	"github.com/pawpawchat/profile/api/pb"
	"github.com/pawpawchat/profile/internal/app/grpc/adapter"
	"github.com/pawpawchat/profile/internal/domain/model"
)

type ProfileService interface {
	CreateProfile(context.Context, *model.Profile) error
	DeleteProfile(context.Context, int64) error

	GetProfileByID(context.Context, int64) (*model.Profile, error)
	GetProfileByUsername(context.Context, string) (*model.Profile, error)

	UpdateProfile(context.Context, *model.UpdateProfileData) error
}

type AvatarService interface {
	AddProfileAvatar(context.Context, *model.Avatar) error
	ChangeProfileAvatar(context.Context, int64, int64) error
	GetAllProfileAvatars(context.Context, int64) (model.Avatars, error)
	DeleteProfileAvatar(context.Context, int64, int64) error
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

func (s *ProfileGRPCServer) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	return adapter.UpdateProfileAdapter(ctx, s.profileService, req)
}

func (s *ProfileGRPCServer) AddProfileAvatar(ctx context.Context, req *pb.AddProfileAvatarRequest) (*pb.AddProfileAvatarResponse, error) {
	return adapter.AddProfileAvatar(ctx, req, s.avatarService)
}

func (s *ProfileGRPCServer) ChangeProfileAvatar(ctx context.Context, req *pb.ChangeProfileAvatarRequest) (*pb.ChangeProfileAvatarResponse, error) {
	return adapter.ChangeProfileAvatarAdapter(ctx, req, s.avatarService)
}

func (s *ProfileGRPCServer) DeleteProfileAvatar(ctx context.Context, req *pb.DeleteProfileAvatarRequest) (*pb.DeleteProfileAvatarResponse, error) {
	return adapter.DeleteProfileAvatarAdapter(ctx, req, s.avatarService)
}
