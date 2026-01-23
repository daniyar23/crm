package grpc

import (
	"net"

	"google.golang.org/grpc"

	userpb "github.com/daniyar23/crm/proto"
)

func RunGRPCServer(handler *UserGRPCHandler) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	userpb.RegisterUserServiceServer(server, handler)

	return server.Serve(lis)
}
