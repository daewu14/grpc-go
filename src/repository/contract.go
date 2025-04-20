package repository

import (
	"context"
	"grpc/src/genproto/orders"
)

type OrderRepository interface {
	Create(ctx context.Context, order *orders.Order) (*orders.Order, error)
	Get(ctx context.Context) ([]*orders.Order, error)
}
