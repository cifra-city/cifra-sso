package main

import (
	"context"
	"log"
	"net"
	"net/http"

	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
	"github.com/cifra-city/cifra-sso/internal/app/handlers/auth"
	"github.com/cifra-city/cifra-sso/internal/app/handlers/reg"
	"github.com/cifra-city/cifra-sso/internal/app/handlers/verify"
	"github.com/cifra-city/cifra-sso/internal/app/modules/action_permission"
	"github.com/cifra-city/cifra-sso/internal/app/modules/email"
	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger := config.SetupLogger(cfg.Logging.Level, cfg.Logging.Format)
	logger.Info("Starting gRPC and HTTP servers...")

	dbCon, err := data.NewDBConnection(cfg.Database.URL)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer dbCon.Close()

	queries := data.New(dbCon)

	// Create support modules
	emailModule := email.NewMailman(cfg, logger)
	actionPermissionModule := action_permission.NewActionPermission(cfg, logger)

	// Create server instances
	registrationSrv := reg.NewRegServer(queries, cfg, logger)
	authenticationSrv := auth.NewAuthServer(queries, cfg, logger, emailModule, actionPermissionModule)
	verificationSrv := verify.NewVerifyServer(queries, cfg, logger, emailModule, actionPermissionModule)

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	ssov1.RegisterRegServer(grpcServer, registrationSrv)
	ssov1.RegisterAuthServer(grpcServer, authenticationSrv)
	ssov1.RegisterVerifyServer(grpcServer, verificationSrv)

	// gRPC listener
	go func() {
		listener, err := net.Listen("tcp", cfg.Server.Port)
		if err != nil {
			logger.Fatalf("failed to listen on port %s: %v", cfg.Server.Port, err)
		}
		logger.Infof("gRPC server is running on port %s", cfg.Server.Port)
		if err := grpcServer.Serve(listener); err != nil {
			logger.Fatalf("failed to serve gRPC server: %v", err)
		}
	}()

	// HTTP Gateway listener
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Register handlers for each service
	err = ssov1.RegisterRegHandlerFromEndpoint(ctx, mux, cfg.Server.Port, opts)
	if err != nil {
		logger.Fatalf("failed to register reg handler: %v", err)
	}

	err = ssov1.RegisterAuthHandlerFromEndpoint(ctx, mux, cfg.Server.Port, opts)
	if err != nil {
		logger.Fatalf("failed to register auth handler: %v", err)
	}

	err = ssov1.RegisterVerifyHandlerFromEndpoint(ctx, mux, cfg.Server.Port, opts)
	if err != nil {
		logger.Fatalf("failed to register verify handler: %v", err)
	}

	gwAddr := ":8080"
	logger.Infof("HTTP Gateway is running on %s", gwAddr)
	if err := http.ListenAndServe(gwAddr, mux); err != nil {
		logger.Fatalf("failed to start gateway server: %v", err)
	}
}
