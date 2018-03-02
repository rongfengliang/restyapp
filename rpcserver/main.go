package main

import (
	"context"
	"log"
	"net"

	pb "github.com/rongfengliang/restyapp/service/pb"
	pb2 "github.com/rongfengliang/restyapp/userservice/pb"
	grpc "google.golang.org/grpc"
)

type server struct{}

type user struct {
}

type userapp struct {
}

func (s *server) Echo(ctx context.Context, in *pb.EchoMessage) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Message: in.Message,
	}, nil
}
func (s *user) Message(ctx context.Context, in *pb.AppMessage) (*pb.AppResponse, error) {
	return &pb.AppResponse{
		Message: in.Message,
	}, nil
}
func (s *userapp) Message(ctx context.Context, in *pb2.AppMessage) (*pb2.AppResponse, error) {
	return &pb2.AppResponse{
		Message: in.Message,
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatal("some wrong")
	}
	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &server{})
	pb.RegisterUserServiceAppServer(s, &user{})
	pb2.RegisterUserServiceAppServer(s, &userapp{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("some wrong")
	}
}
