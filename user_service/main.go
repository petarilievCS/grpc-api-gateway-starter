package main

import (
	"context"
	"log"
	"net"

	userpb "grpc-api-gateway/gen/user"

	"google.golang.org/grpc"
)

type userServer struct {
	userpb.UnimplementedUserServiceServer
}

func (s *userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.User, error) {
	log.Printf("Fetching user with ID: %s", req.UserId)

	// Generate mock user
	mockUser := userpb.User{
		UserId: req.UserId,
		Name:   "Petar",
		Email:  "petariliev2002@gmail.com",
	}
	return &mockUser, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &userServer{})

	log.Println("🚀 UserService running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
