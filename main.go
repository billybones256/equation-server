package main

import (
	"context"
	"equation-server/pkg/api"
	"fmt"
	"github.com/micro/go-micro"
)

type Solver interface {
	Solve(context.Context, *api.SolveRequest, *api.SolveResponse) error
}

type service struct {
	solver Solver
}

//Solve ...
func (s *service) Solve(ctx context.Context, req *api.SolveRequest, res *api.SolveResponse) error {
	result := solve(int(req.GetA()), int(req.GetB()), int(req.GetC()))
	res.A = req.GetA()
	res.B = req.GetB()
	res.C = req.GetC()
	res.Answer = result
	return nil
}

func linear(b, c int) string {
	if b == 0 {
		return "has no roots"
	}
	return "has one root"
}

func quadratic(a, b, c int) string {
	d := b*b - 4*a*c
	if d == 0 {
		return "has one root"
	}
	if d > 0 {
		return "has two roots"
	}
	return "has no roots"
}

//Solve
func solve(a, b, c int) string {
	var result string
	if a == 0 {
		result = linear(b, c)
	} else {
		result = quadratic(a, b, c)
	}
	return result
}

func main() {

	srv := micro.NewService(micro.Name("solver"))
	srv.Init()
	api.RegisterSolverHandler(srv.Server(), &service{})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

	//lis, err := net.Listen("tcp", port)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//s := grpc.NewServer()

	//pb.RegisterShippingServiceServer(s, &service{repo})

	//log.Println("Running on port:", port)
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
}
