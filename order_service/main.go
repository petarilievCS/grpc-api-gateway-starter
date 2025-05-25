package main

import (
	"context"
	"log"
	"net"

	orderpb "grpc-api-gateway/gen/order"

	"google.golang.org/grpc"
)

type orderServer struct {
	orderpb.UnimplementedOrderServiceServer
}

func (s *orderServer) GetUser(ctx context.Context, req *orderpb.GetOrderRequest) (*orderpb.Order, error) {
	log.Printf("Fetching order with ID: %s", req.OrderId)

	// Generate mock user
	mockOrder := orderpb.Order{
		OrderId:  req.OrderId,
		UserId:   "1",
		Name:     "Petar",
		Quantity: 3,
	}
	return &mockOrder, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(s, &orderServer{})

	log.Println("🚀 OrderService running on port 50053")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
