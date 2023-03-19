package grpcserver

import (
	"context"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/domain/interactions"
)

type GreetServer struct {
	UnimplementedGreeterServer
}

func (g GreetServer) Greet(ctx context.Context, request *GreetRequest) (*GreetReply, error) {
	return &GreetReply{Message: interactions.Greet(request.Name)}, nil
}

func (g GreetServer) Curse(ctx context.Context, request *CurseRequest) (*CurseReply, error) {
	return &CurseReply{Message: interactions.Curse(request.Name)}, nil
}
