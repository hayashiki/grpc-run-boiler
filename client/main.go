package main

import (
	"context"
	"crypto/tls"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-run-boiler/protos"
	"log"
	"math/rand"
	"time"
)

var (
	grpcEndpoint = flag.String("grpc_endpoint", "", "the gRPC Endpoint of the Server")
)

func main()  {
	flag.Parse()
	if *grpcEndpoint == "" {
		log.Fatal("[main] unable to start client without gRPC endpoint to server")
	}
	
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		//grpc.WithInsecure(),
	}

	log.Printf("Connecting to gRPC Service [%s]", *grpcEndpoint)
	conn, err := grpc.Dial(*grpcEndpoint, opts...)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := NewClient(conn)
	ctx := context.Background()

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for {
		o1 := r.Float32()
		o2 := r.Float32()
		op := randomOp(r.Float32())
		rqst := &pb.BinaryOperation{
			FirstOperand:  o1,
			SecondOperand: o2,
			Operation:     op,
		}
		resp, err := client.Calculate(ctx, rqst)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("[main] %0.4f %s %0.4f = %0.4f", o1, stringOp(op), o2, resp.GetResult())
		time.Sleep(5 * time.Second)
	}
}

func randomOp(r float32) pb.Operation {
	if r < 0.5 {
		return pb.Operation_ADD
	}
	return pb.Operation_SUBTRACT
}

func stringOp(op pb.Operation) string {
	if op == pb.Operation_ADD {
		return "+"
	}
	return "-"
}




