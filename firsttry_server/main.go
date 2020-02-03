package main

import (
	"context"
	"log"
	"net"
	"math/rand"

	"google.golang.org/grpc"
	pb "github.com/samkulkarni20/my-grpc-try/firsttry"
)

const (
	port = ":5000"
)

// server to implement firsttry.FirstTryServer
type server struct {
	pb.UnimplementedFirstTryServer
}

// SayHowdy implements firsttry.FirstTryServer.
func (s *server) SayHowdy(ctx context.Context, in *pb.HowdyRequest) (*pb.HowdyResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HowdyResponse{Message: "Howdy " + in.GetName()}, nil
}

// GetRandomNumber implements firsttry.FirstTryServer.
func (s *server) GetRandomNumber(ctx context.Context, in *pb.RandomNoRequest) (*pb.RandomNoResponse, error) {
	log.Printf("Received random no base: %v", in.GetBase())
	return &pb.RandomNoResponse{RandomNumber: rand.Int31n(in.GetBase())}, nil
}

func main() {
	lis, err:= net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFirstTryServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
