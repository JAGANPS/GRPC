package main

import pb "github.com/JAGANPS/GRPC/greet/proto"

type Server struct {
	pb.GreetServiceServer
}
