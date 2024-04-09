package grpc

import (
	"app/grpc/api"
	"app/grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func RunServerGRPC() {
	listenerGRPC, err := net.Listen("tcp", ":20004")

	if err != nil {
		log.Fatalln(listenerGRPC)
	}

	creds, errKey := credentials.NewServerTLSFromFile(
		"keys/server-file/public.pem",
		"keys/server-file/private.pem",
	)

	if errKey != nil {
		log.Fatalln(errKey)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	proto.RegisterFileServiceServer(grpcServer, api.NewFileGRPC())

	log.Fatalln(grpcServer.Serve(listenerGRPC))
}
