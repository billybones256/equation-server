package solver

import (
	"context"
	"grpcequtation/pkg/api"
)

//GRPCServer ...
type GRPCServer struct{}

//Square ...
func (server *GRPCServer) Solve(ctx context.Context, req *api.SolveRequest) (*api.SolveResponse, error) {
	result := solve(int(req.GetA()), int(req.GetB()), int(req.GetC()))
	return &api.SolveResponse{Answer: result}, nil
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
