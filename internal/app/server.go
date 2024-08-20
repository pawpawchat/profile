package app

import (
	"context"

	"github.com/pawpawchat/profile/api/pb"
)

type ProfileGRPCServer struct {
	pb.UnimplementedProfileServiceServer
}

func NewServer() *ProfileGRPCServer {
	return &ProfileGRPCServer{}
}

func (s *ProfileGRPCServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	return nil, nil
}

func (s *ProfileGRPCServer) GetById(ctx context.Context, req *pb.GetByIdRequest) (*pb.GetByIdResponse, error) {
	return nil, nil
}
