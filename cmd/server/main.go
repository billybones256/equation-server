package main

import (
	"google.golang.org/grpc"
	"grpcequtation/pkg/api"
	"grpcequtation/pkg/solver"
	"log"
	"net"
)

func main() {
	server := grpc.NewServer()
	microService := &solver.GRPCServer{}
	api.RegisterSolverServer(server, microService)
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
