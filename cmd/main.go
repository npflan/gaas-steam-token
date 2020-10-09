package main

import (
	_ "./proto"
	"log"
	"net"

	"google.golang.org/grpc"

)

func main() {
	// lis, err := net.Listen("tcp", "8080")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// s := grpc.NewServer()
	// pb.RegisterTokenserviceServer(s, &)

	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }



}