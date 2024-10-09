package main

import (
	"log"
	"net"

	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	authsrv "github.com/cifra-city/cifra-sso/internal/services/authsrv"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

func main() {
	// Load the configuration.
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger := config.SetupLogger(cfg.Logging.Level, cfg.Logging.Format)
	logger.Info("Starting gRPC server...")

	// Initialize the database connection.
	conn, err := data.NewDBConnection(cfg.Database.URL)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer conn.Close()

	// Create a new Queries object for interacting with the database.
	queries := data.New(conn)

	// Create the AuthServer with the necessary dependencies.
	authServer := authsrv.NewAuthServer(queries, nil)

	// Create a new gRPC server.
	grpcServer := grpc.NewServer()

	// Register the AuthServer implementation with the gRPC server.
	ssov1.RegisterAuthServer(grpcServer, authServer)

	// Start listening on the specified port.
	listener, err := net.Listen("tcp", cfg.Server.Port)
	if err != nil {
		logger.Fatalf("failed to listen on port %s: %v", cfg.Server.Port, err)
	}

	logger.Printf("gRPC server is running on port %s", cfg.Server.Port)

	// Start serving gRPC requests.
	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatalf("failed to serve gRPC server: %v", err)
	}
}
