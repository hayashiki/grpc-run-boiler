package main

import (
	"context"
	"fmt"
	pb "grpc-run-boiler/protos"
	"log"
)

var _ pb.CalculatorServer = (*Server)(nil)

type Server struct {
	pb.UnimplementedCalculatorServer
}


func NewServer() *Server {
	return &Server{
	}
}

func (s *Server) Calculate(ctx context.Context, r *pb.BinaryOperation) (*pb.CalculationResult, error) {
	log.Println("[server:Calculate] Started")
	if ctx.Err() == context.Canceled {
		return &pb.CalculationResult{}, fmt.Errorf("client cancelled: abandoning")
	}

	switch r.GetOperation() {
	case pb.Operation_ADD:
		return &pb.CalculationResult{
			Result: r.GetFirstOperand() + r.GetSecondOperand(),
		}, nil
	case pb.Operation_SUBTRACT:
		return &pb.CalculationResult{
			Result: r.GetFirstOperand() - r.GetSecondOperand(),
		}, nil
	default:
		return &pb.CalculationResult{}, fmt.Errorf("undefined operation")
	}

}
