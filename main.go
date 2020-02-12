package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"context"

	"github.com/micro/go-micro"

	"equation-client/pkg/api"
)

const (
	address         = "localhost:8081"
	defaultFilename = "test1.json"
)

func parseFile(file string) (*api.SolveRequest, error) {
	var req *api.SolveRequest
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &req)
	return req, err
}

func main() {

	service := micro.NewService(micro.Name("solver"))
	service.Init()

	client := api.NewSolverClient("solver", service.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	req, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.Solve(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	prettyPrint(r)
}

func prettyPrint(c *api.SolveResponse) {
	log.Println("======================================")
	log.Printf("a: %d\n", c.GetA())
	log.Printf("b: %d\n", c.GetB())
	log.Printf("c: %d\n", c.GetC())
	log.Printf("answer: %s\n", c.GetAnswer())
	log.Println("======================================")
	log.Println()
}
