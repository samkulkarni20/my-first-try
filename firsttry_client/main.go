package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/samkulkarni20/my-grpc-try/firsttry"
	"google.golang.org/grpc"
)

const (
	address = "localhost:5000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFirstTryClient(conn)

	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHowdy(ctx, &pb.HowdyRequest{Name: "Sameer"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	r, err = c.GetRandomNumber(ctx, &pb.RandomNoRequest{Base: int32(51)})
	if err != nil {
		log.Fatalf("could not get random number: %v", err)
	}
	log.Printf("Random number: %d", r.GetRandomNumber())
}
