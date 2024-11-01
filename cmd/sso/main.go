package main

import (
	"database/sql"
	"log"
	"net"

	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
	"github.com/cifra-city/cifra-sso/internal/app/handlers/auth"
	"github.com/cifra-city/cifra-sso/internal/app/handlers/reg"
	"github.com/cifra-city/cifra-sso/internal/app/handlers/verify"
	"github.com/cifra-city/cifra-sso/internal/app/modules/action_permission"
	"github.com/cifra-city/cifra-sso/internal/app/modules/email"
	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger := config.SetupLogger(cfg.Logging.Level, cfg.Logging.Format)
	logger.Info("Starting gRPC server...")

	//services.AuthService(cfg, logger)
	dbCon, err := data.NewDBConnection(cfg.Database.URL)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer func(dbCon *sql.DB) {
		err := dbCon.Close()
		if err != nil {
			logger.Fatalf("failed to close the database connection: %v", err)
			panic(err)
		}
	}(dbCon)
	queries := data.New(dbCon)

	// Create the support modules
	emailModule := email.NewMailman(cfg, logger)
	actionPermissionModule := action_permission.NewActionPermission(cfg, logger)

	// Create the Server with the necessary dependencies
	authenticationSrv := auth.NewAuthServer(queries, cfg, logger, emailModule, actionPermissionModule)
	registrationSrv := reg.NewRegServer(queries, cfg, logger)
	verificationSrv := verify.NewVerifyServer(queries, cfg, logger, emailModule, actionPermissionModule)

	grpcServer := grpc.NewServer()

	// Register the Server implementation with the gRPC server
	ssov1.RegisterAuthServer(grpcServer, authenticationSrv)
	ssov1.RegisterRegServer(grpcServer, registrationSrv)
	ssov1.RegisterVerifyServer(grpcServer, verificationSrv)

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
