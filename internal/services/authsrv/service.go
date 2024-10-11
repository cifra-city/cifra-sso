package authsrv

import (
	"context"

	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"github.com/sirupsen/logrus"
)

// AuthServer - structure for implementing the gRPC service.
type AuthServer struct {
	ssov1.UnimplementedAuthServer
	Queries *data.Queries
	AuthService
	Config *config.Config // Interface for handling authentication methods.
	Log    *logrus.Logger
}

// NewAuthServer create new AuthServer.
func NewAuthServer(queries *data.Queries, cfg *config.Config, log *logrus.Logger) *AuthServer {
	return &AuthServer{
		Queries: queries,
		Config:  cfg,
		Log:     log,
	}
}

// AuthService defines the methods for user authentication and authorization.
type AuthService interface {
	Register(ctx context.Context, in *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error)
	Login(ctx context.Context, in *ssov1.LoginRequest) (*ssov1.LoginResponse, error)
	IsAdmin(ctx context.Context, in *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error)
}
