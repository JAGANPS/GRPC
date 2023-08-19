package main

import pb "github.com/JAGANPS/GRPC/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
