package grpcpkg

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// NewClient creates a new gRPC client connection to the specified address.
func NewClient(addr string) GrpcClient {
	conn, err := grpc.NewClient(fmt.Sprintf(":%v", addr), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &grpcClient{
		conn: conn,
	}
}

type grpcClient struct {
	conn *grpc.ClientConn
}

func (g *grpcClient) Conn() *grpc.ClientConn {
	return g.conn
}

func (g *grpcClient) Close() {
	err := g.conn.Close()
	if err != nil {
		log.Print("error closing grpc client connection: ", err)
	}
}
