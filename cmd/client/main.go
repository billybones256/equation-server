package main

import (
	"encoding/json"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"

	"context"

	pb "grpcequation/pkg/api"
)

const (
	address         = "localhost:8081"
	defaultFilename = "test1.json"
)

func parseFile(file string) (*pb.SolveRequest, error) {
	var req *pb.SolveRequest
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &req)
	return req, err
}

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewSolverClient(conn)

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.Solve(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	prettyPrint(r)
}

func prettyPrint(c *pb.SolveResponse) {
	log.Println("======================================")
	log.Printf("a: %d\n", c.GetA())
	log.Printf("b: %d\n", c.GetB())
	log.Printf("c: %d\n", c.GetC())
	log.Printf("answer: %s\n", c.GetAnswer())
	log.Println("======================================")
	log.Println()
}
