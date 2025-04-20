package repository

import (
	"context"
	"grpc/src/genproto/orders"
)

// NewOrderRepository defines the interface for order repository operations.
func NewOrderRepository() OrderRepository {
	return &order{}
}

type order struct{}

var ordersDb = make([]*orders.Order, 0)

func (r order) Create(ctx context.Context, order *orders.Order) (*orders.Order, error) {

	orderId := 1
	if len(ordersDb) > 0 {
		orderId = int(ordersDb[len(ordersDb)-1].OrderID) + 1
	}

	order.OrderID = int32(orderId)

	ordersDb = append(ordersDb, order)
	return order, nil
}

func (r order) Get(ctx context.Context) ([]*orders.Order, error) {
	return ordersDb, nil
}
