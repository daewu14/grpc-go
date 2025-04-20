package handler

import (
	"context"
	grpcpkg "grpc/pkg/grpc"
	httppkg "grpc/pkg/http"
	"grpc/src/genproto/orders"
	"grpc/src/ucase/order"
)

func RunHttp(ctx context.Context) {
	server := httppkg.NewServer(ctx, "8000")
	grpcClient9000 := grpcpkg.NewClient("9000")
	defer grpcClient9000.Close()

	grpcOrderClient := orders.NewOrderServiceClient(grpcClient9000.Conn())

	server.Handle("orders", order.NewCreateOrderHandler(grpcOrderClient))

	server.Run()
}
