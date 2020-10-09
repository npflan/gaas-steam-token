package main

import (
	handler "github.com/Cytram/csgo-token/pkg"
	pb "github.com/Cytram/csgo-token/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTokenserviceServer(s, &handler.Server{})
	log.Printf("Start listing on :8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
