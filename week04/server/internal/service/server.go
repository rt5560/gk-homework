package service

import (
	"context"

	pb "server/api/server"
)

type ServerService struct {
	pb.UnimplementedServerServer
}

func NewServerService() *ServerService {
	return &ServerService{}
}

func (s *ServerService) CreateServer(ctx context.Context, req *pb.CreateServerRequest) (*pb.CreateServerReply, error) {
	return &pb.CreateServerReply{}, nil
}
func (s *ServerService) UpdateServer(ctx context.Context, req *pb.UpdateServerRequest) (*pb.UpdateServerReply, error) {
	return &pb.UpdateServerReply{}, nil
}
func (s *ServerService) DeleteServer(ctx context.Context, req *pb.DeleteServerRequest) (*pb.DeleteServerReply, error) {
	return &pb.DeleteServerReply{}, nil
}
func (s *ServerService) GetServer(ctx context.Context, req *pb.GetServerRequest) (*pb.GetServerReply, error) {
	return &pb.GetServerReply{}, nil
}
func (s *ServerService) ListServer(ctx context.Context, req *pb.ListServerRequest) (*pb.ListServerReply, error) {
	return &pb.ListServerReply{}, nil
}
