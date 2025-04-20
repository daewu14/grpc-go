package grpcpkg

import (
	"google.golang.org/grpc"
)

type GrpcServer interface {
	Register(server *grpc.Server)
}

type GrpcClient interface {
	Conn() *grpc.ClientConn
	Close()
}
