// main.go
package main

import (
	pb "grpc-rest-go/hello"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Define GRPC and rest ports
	grpcAddr := ":50051"
	restPort := ":8080"

	// Start gRPC server
	go func() {
		lis, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterHelloServiceServer(s, &helloServer{})
		log.Printf("gRPC server listening on %s", grpcAddr)
		s.Serve(lis)
	}()

	// Start REST gateway
	startGateway(grpcAddr, restPort)
}
