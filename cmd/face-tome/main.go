package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/MarcvanMelle/face-tome/configs"
	api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// ServerCloser is an interface designed to unify server closing across all diferent types of servers, i.e. grpc, http
type ServerCloser interface {
	Shutdown(ctx context.Context) error
}

// GRPCCloser is the struct used to coerce grpc servers into the ServerCloser interface
type GRPCCloser struct {
	Server *grpc.Server
}

// Shutdown is the grpc server closing method
func (c GRPCCloser) Shutdown(ctx context.Context) error {
	doneChan := make(chan struct{})
	go func() {
		c.Server.GracefulStop()
		doneChan <- struct{}{}
	}()

	select {
	case <-doneChan:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func serveHTTP(srv *http.Server, l net.Listener, errChan chan error) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	if err := srv.Serve(l); err != nil {
		errChan <- fmt.Errorf("failed to serve: %s", err)
		return
	}
}

func serveGRPC(srv *grpc.Server, l net.Listener, errChan chan error) {
	api.RegisterFaceTomeServer(srv, &grpcServer{})
	reflection.Register(srv)
	if err := srv.Serve(l); err != nil {
		errChan <- fmt.Errorf("failed to serve: %s", err)
		return
	}
}

func main() {
	errChan := make(chan error)
	sigChan := make(chan os.Signal)
	exitCode := 0
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	httpAddr := fmt.Sprintf(":%d", configs.Config.RestPost)
	httpListener, err := net.Listen("tcp", httpAddr)
	if err != nil {
		errChan <- fmt.Errorf("cannot listen to address: %s", err)
	}
	log.Printf("http listening on %d", configs.Config.RestPost)
	httpSrv := &http.Server{}

	grpcAddr := fmt.Sprintf(":%d", configs.Config.GRPCPort)
	grpcListener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Printf("cannot listen to address: %s", err)
	}
	log.Printf("grpc listening on port %d", configs.Config.GRPCPort)
	grpcSrv := grpc.NewServer()

	go serveHTTP(httpSrv, httpListener, errChan)
	go serveGRPC(grpcSrv, grpcListener, errChan)

	select {
	case err := <-errChan:
		exitCode = 1
		log.Println("Error starting server: ", err)
	case sig := <-sigChan:
		if sigMsg := sig.String(); sigMsg == "interrupt" || sigMsg == "terminated" {
			log.Printf("Received termination signal '%s'. Shutting down server...", sigMsg)
			gracefulShutdownServers(5*time.Second, httpSrv, GRPCCloser{Server: grpcSrv})
		}
	}

	signal.Stop(sigChan)
	os.Exit(exitCode)
}

func gracefulShutdownServers(timeout time.Duration, closers ...ServerCloser) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var shutdownWG sync.WaitGroup
	shutdownWG.Add(len(closers))
	for _, closer := range closers {
		go func(closer ServerCloser) {
			defer shutdownWG.Done()
			err := closer.Shutdown(ctx)
			if err != nil {
				log.Printf("failed to shut down server: %s", err)
			} else {
				log.Printf("server gracefully stopped")
			}
		}(closer)
	}
	shutdownWG.Wait()
	log.Printf("all servers graceful shutdown complete")
}
