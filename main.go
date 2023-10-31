package main

import (
	"context"
	"io"
	"log"
	"net"

	userservice "github.com/Pratyush/User/service"
	pb "github.com/Pratyush/User/user/protos"
	"google.golang.org/grpc"
)

type server struct {
	pb.UsersServer
}

func (s *server) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	users := userservice.GetUser()
	return users, nil
}

func (s *server) GetUserById(ctx context.Context, in *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	users := userservice.GetUserById(in.Id)
	return users, nil
}

func (s *server) GetUserByIdStream(stream pb.Users_GetUserByIdStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil // Client closed the stream
		}
		if err != nil {
			log.Printf("Error receiving request: %v", err)
			return err
		}
		user := userservice.GetUserByIdStream(req.Id)
		if err := stream.SendMsg(user); err != nil {
			return err
		}
	}
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUsersServer(s, &server{})

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
