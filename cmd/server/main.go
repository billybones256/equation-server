package main

import (
	"google.golang.org/grpc"
	"grpcequation/pkg/api"
	"grpcequation/pkg/solver"
	"log"
	"net"
)

const (
	port = ":8081"
)

func main() {
	server := grpc.NewServer()
	microService := &solver.GRPCServer{}
	api.RegisterSolverServer(server, microService)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Running on port:", port)
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
