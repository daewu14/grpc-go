package handler

import (
	"context"
	grpcpkg "grpc/pkg/grpc"
	"grpc/src/grpcfactory"
	"grpc/src/repository"
	"grpc/src/ucase/grpc"
)

// RunGrpc initializes and runs a gRPC server with a rpc service implementation.
func RunGrpc(ctx context.Context) {
	server := grpcpkg.NewServer(ctx, "9000")

	orderRepo := repository.NewOrderRepository()
	orderFactory := grpcfactory.NewOrderGrpcFactory(orderRepo)

	server.Handle(grpc.NewGrpcOrder(orderFactory))

	// Run all registered gRPC services
	server.Run()
}
