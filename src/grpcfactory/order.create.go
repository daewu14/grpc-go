package grpcfactory

import (
	"context"
	ordersRpc "grpc/src/genproto/orders"
)

func (f orderGrpcFactory) CreateOrder(ctx context.Context, request *ordersRpc.CreateOrderRequest) (*ordersRpc.CreateOrderResponse, error) {
	order := &ordersRpc.Order{
		CustomerID: request.CustomerID,
		ProductID:  request.ProductID,
		Quantity:   request.Quantity,
	}

	_, err := f.orderRepo.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	println("Order created with ID:", order.OrderID)
	res := &ordersRpc.CreateOrderResponse{
		Status: "success",
		Order:  order,
	}

	return res, nil
}
