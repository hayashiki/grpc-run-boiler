package main

import (
	"context"

	"google.golang.org/grpc"
	pb "grpc-run-boiler/protos"
)

type Client struct {
	client pb.CalculatorClient
}

// NewClient returns a new Client
func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		client: pb.NewCalculatorClient(conn),
	}
}

// Calculate performs an operation on operands defined by pb.BinaryOperation returning pb.CalculationResult
func (c *Client) Calculate(ctx context.Context, r *pb.BinaryOperation) (*pb.CalculationResult, error) {
	return c.client.Calculate(ctx, r)
}
