package authsrv

import (
	"context"

	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/cifra-city/cifra-sso/internal/email"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// AuthServer - structure for implementing the gRPC service.
type AuthServer struct {
	Queries *data.Queries
	Config  *config.Config // Interface for handling authentication methods.
	Log     *logrus.Logger
	Email   email.Mailman

	ssov1.UnimplementedAuthServer
	AuthService
}

// NewAuthServer create new AuthServer.
func NewAuthServer(queries *data.Queries, cfg *config.Config, log *logrus.Logger) *AuthServer {
	return &AuthServer{
		Queries: queries,
		Config:  cfg,
		Log:     log,
		Email:   *email.NewMailman(cfg, log),
	}
}

// AuthService defines the methods for user authentication and authorization.
type AuthService interface {
	Authenticate(ctx context.Context) (uuid.UUID, error)

	Register(ctx context.Context, in *ssov1.RegisterRequest) (*ssov1.Empty, error)
	Login(ctx context.Context, in *ssov1.LoginRequest) (*ssov1.LoginResponse, error)
	Logout(ctx context.Context, in *ssov1.Empty) (*ssov1.Empty, error)

	ChangePassword(ctx context.Context, in *ssov1.ChangePasswordRequest) (*ssov1.Empty, error)
	ChangeUsername(ctx context.Context, in *ssov1.ChangeUsernameRequest) (*ssov1.Empty, error)
	ChangeEmail(ctx context.Context, in *ssov1.ChangeEmailRequest) (*ssov1.Empty, error)
	UpdateEmailStatus(ctx context.Context, in *ssov1.ChangeEmailStatusRequest) (*ssov1.Empty, error)
}
