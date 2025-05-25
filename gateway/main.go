package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gatewaypb "grpc-api-gateway/gen/gateway"
	userpb "grpc-api-gateway/gen/user"
)

type gatewayServer struct {
	gatewaypb.UnimplementedAPIGatewayServiceServer
	userClient userpb.UserServiceClient
}

func (g *gatewayServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.User, error) {
	log.Printf("Gateway received request for user ID: %s", req.UserId)
	return g.userClient.GetUser(ctx, req)
}

func main() {
	// Connect to UserService running on port 50051
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	defer conn.Close()

	userClient := userpb.NewUserServiceClient(conn)

	// Start API Gateway server on port 50052
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	gatewaypb.RegisterAPIGatewayServiceServer(s, &gatewayServer{
		userClient: userClient,
	})

	log.Println("🚀 API Gateway running on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
