package grpcfactory

import (
	"context"
	ordersRpc "grpc/src/genproto/orders"
)

func (f orderGrpcFactory) GetOrders(ctx context.Context, request *ordersRpc.GetOrdersRequest) (*ordersRpc.GetOrderResponse, error) {
	o, err := f.orderRepo.Get(ctx)
	if err != nil {
		return nil, err
	}

	res := &ordersRpc.GetOrderResponse{
		Orders: o,
	}

	return res, nil
}
