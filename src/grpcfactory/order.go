package grpcfactory

import (
	ordersRpc "grpc/src/genproto/orders"
	"grpc/src/repository"
)

// NewOrderGrpcFactory initializes a new instance of orderGrpcFactory and binds it to the given gRPC server.
func NewOrderGrpcFactory(repo repository.OrderRepository) *orderGrpcFactory {
	n := &orderGrpcFactory{
		orderRepo: repo,
	}
	return n
}

type orderGrpcFactory struct {
	ordersRpc.UnimplementedOrderServiceServer
	orderRepo repository.OrderRepository
}
