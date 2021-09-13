package main

import (
	"log"
	"net"
	"rpcapp/pbs"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:3001")

	if err != nil {
		log.Fatal("error : ", err)
	}

	s := pbs.Server{}

	grpcServer := grpc.NewServer()

	pbs.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
