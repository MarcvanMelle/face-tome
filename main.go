package main

import (
	"context"
	"log"
	"net"

	api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"
	"github.com/MarcvanMelle/face-tome/npcgenerator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":5501"
)

type server struct{}

func (s *server) GetNPC(ctx context.Context, request *api.GetNPCRequest) (*api.GetNPCResponse, error) {
	return npcgenerator.GetNPC(request)
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("cannot listen to address: %v", err)
	}
	log.Printf("listening on port %s", port)

	serveGRPC(listener)
}

func serveGRPC(listener net.Listener) {
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	api.RegisterFaceTomeServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
