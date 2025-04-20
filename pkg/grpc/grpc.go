package grpcpkg

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

// NewServer creates and initializes a new grpcPkg instance with the given context and list of gRPC services.
func NewServer(ctx context.Context, addr string) *grpcPkg {
	return &grpcPkg{
		services:   make([]GrpcServer, 0),
		ctx:        ctx,
		grpcServer: grpc.NewServer(),
		address:    addr,
	}
}

type grpcPkg struct {
	services   []GrpcServer
	ctx        context.Context
	grpcServer *grpc.Server
	address    string
}

func (g *grpcPkg) Handle(server GrpcServer) {
	g.services = append(g.services, server)
}

func (g *grpcPkg) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", g.address))
	if err != nil {
		log.Fatalf("failed to listen grpcfactory: %v", err)
	}

	for _, s := range g.services {
		s.Register(g.grpcServer)
	}

	log.Print("grpc server start run with total services: ", len(g.services), " on port: ", g.address)
	err = g.grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve grpcfactory server: %v", err)
	}

}

func (g grpcPkg) Shutdown() {
	g.grpcServer.GracefulStop()
}
