package authsrv

import (
	"context"

	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// AuthServer - structure for implementing the gRPC service.
type AuthServer struct {
	Queries *data.Queries
	Config  *config.Config // Interface for handling authentication methods.
	Log     *logrus.Logger

	ssov1.UnimplementedAuthServer
	AuthService
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
	Register(ctx context.Context, in *ssov1.RegisterRequest) (*ssov1.Empty, error)
	Login(ctx context.Context, in *ssov1.LoginRequest) (*ssov1.LoginResponse, error)
	Authenticate(ctx context.Context) (uuid.UUID, error)
}
