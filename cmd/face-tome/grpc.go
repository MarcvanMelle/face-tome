package main

import (
	"context"

	"github.com/MarcvanMelle/face-tome/internal/npcgenerator"
	api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"
)

type grpcServer struct{}

func (s *grpcServer) GetNPC(ctx context.Context, request *api.GetNPCRequest) (*api.GetNPCResponse, error) {
	return npcgenerator.GetNPC(request)
}
