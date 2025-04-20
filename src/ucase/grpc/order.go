package grpc

import (
	"google.golang.org/grpc"
	grpcpkg "grpc/pkg/grpc"
	"grpc/src/genproto/orders"
)

// NewGrpcOrder creates a new gRPC server for the Order service.
func NewGrpcOrder(factory orders.OrderServiceServer) grpcpkg.GrpcServer {
	return grpcOrder{
		factory: factory,
	}
}

type grpcOrder struct {
	factory orders.OrderServiceServer
}

func (g grpcOrder) Register(server *grpc.Server) {
	orders.RegisterOrderServiceServer(server, g.factory)
}
