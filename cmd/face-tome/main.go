package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/MarcvanMelle/face-tome/configs"
	api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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
	grpcAddr := fmt.Sprintf(":%d", configs.Config.GRPCPort)
	grpcListener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("cannot listen to address: %s", err)
	}
	log.Printf("grpc listening on port %d", configs.Config.GRPCPort)

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	api.RegisterFaceTomeServer(s, &grpcServer{})
	reflection.Register(s)
	if err := s.Serve(grpcListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func serveHTTP(errChan chan error) {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	httpAddr := fmt.Sprintf(":%d", configs.Config.RestPost)
	httpListener, err := net.Listen("tcp", httpAddr)
	if err != nil {
		errChan <- fmt.Errorf("cannot listen to address: %s", err)
		return
	}
	log.Printf("http listening on %d", configs.Config.RestPost)

	httpServer := &http.Server{}
	if err := httpServer.Serve(httpListener); err != nil {
		errChan <- fmt.Errorf("failed to serve: %s", err)
		return
	}
}
