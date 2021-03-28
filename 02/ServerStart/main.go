package main

import (
	"gRPC-Learn/02/Common"
	"gRPC-Learn/02/ProtoEntity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8003")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	ProtoEntity.RegisterCommonServiceServer(grpcServer, &Common.ServerEntity{})
	reflection.Register(grpcServer)
	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
