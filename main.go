package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/namtx/rdb/proto"
	"github.com/namtx/rdb/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port = flag.Int("port", 9090, "Port to listen on.")

func main() {
	fmt.Print("Starting server\n")
	lis, error := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if error != nil {
		log.Fatalf("Failed to listen %v", error)
	}
	grpcServer := grpc.NewServer()
	ds := server.NewServer()
	pb.RegisterRdbServer(grpcServer, ds)

	reflection.Register(grpcServer)

	fmt.Printf("Listening on port %d...\n", *port)
	if error = grpcServer.Serve(lis); error != nil {
		log.Fatalf("Failed to serve: %v", error)
	}

}
