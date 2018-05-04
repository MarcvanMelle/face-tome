package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/MarcvanMelle/face-tome/internal/npcgenerator"
	api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = ":5501"
	httpPort = ":5500"
)

type server struct{}

func (s *server) GetNPC(ctx context.Context, request *api.GetNPCRequest) (*api.GetNPCResponse, error) {
	return npcgenerator.GetNPC(request)
}

func main() {
	errChan := make(chan error)
	exitChan := make(chan os.Signal)
	signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM)

	go serveHTTP(errChan)
	go serveGRPC(errChan)

	select {
	case err := <-errChan:
		fmt.Println("Error starting server: ", err)
		return
	case exit := <-exitChan:
		fmt.Println("Exiting: ", exit)
		return
	}
}

func serveGRPC(errChan chan error) {
	grpcListener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("cannot listen to address: %v", err)
	}
	log.Printf("grpc listening on port %s", grpcPort)

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	api.RegisterFaceTomeServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(grpcListener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func serveHTTP(errChan chan error) {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		// w.WriteHeader(http.StatusOK)
		// w.Write([]byte("OK"))
		io.WriteString(w, "hello, world!\n")
	})

	httpListener, err := net.Listen("tcp", httpPort)
	if err != nil {
		errChan <- fmt.Errorf("cannot listen to address: %v", err)
		return
	}
	log.Printf("http listening on %s", httpPort)

	httpServer := &http.Server{}
	if err := httpServer.Serve(httpListener); err != nil {
		errChan <- fmt.Errorf("failed to serve: %v", err)
		return
	}
}
