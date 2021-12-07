package main

import (
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-run-boiler/protos"
	"log"
	"net"
	"os"
)

func main()  {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	grpcEndpoint := fmt.Sprintf(":%s", port)
	log.Printf("gRPC endpoint [%s]", grpcEndpoint)

	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServer(grpcServer, NewServer())

	listen, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting: gRPC Listener %s", grpcEndpoint)
	log.Fatal(grpcServer.Serve(listen))
}
