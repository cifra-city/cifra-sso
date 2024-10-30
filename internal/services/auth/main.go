package auth

import (
	"context"

	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/cifra-city/cifra-sso/internal/modules/email"
	"github.com/cifra-city/cifra-sso/internal/modules/events"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"github.com/sirupsen/logrus"
)

// DefaultPasswordLength - default password length.
const (
	CHANGE_EMAIL    = "CHANGE_EMAIL"
	CHANGE_PASS     = "CHANGE_PASS"
	CHANGE_USERNAME = "CHANGE_USERNAME"
)

// AuthServer - structure for implementing the gRPC service.
type AuthServer struct {
	Queries *data.Queries
	Config  *config.Config // Interface for handling authentication methods.
	Log     *logrus.Logger
	Email   *email.Mailman
	Events  *events.Events

	ssov1.UnimplementedAuthServer
	AuthService
}

// NewAuthServer create new AuthServer.
func NewAuthServer(queries *data.Queries, cfg *config.Config, log *logrus.Logger, mailman *email.Mailman, eventsList *events.Events) *AuthServer {
	return &AuthServer{
		Queries: queries,
		Config:  cfg,
		Log:     log,
		Email:   mailman,
		Events:  eventsList,
	}
}

// AuthService defines the methods for user authentication and authorization.
type AuthService interface {
	Login(ctx context.Context, in *ssov1.LoginReq) (*ssov1.LoginResp, error)
	Logout(ctx context.Context, in *ssov1.Empty) (*ssov1.Empty, error)

	ChangePassword(ctx context.Context, in *ssov1.ChangePassReq) (*ssov1.Empty, error)
	ChangeUsername(ctx context.Context, in *ssov1.ChangeUsernameReq) (*ssov1.Empty, error)
	ChangeEmail(ctx context.Context, in *ssov1.ChangeEmailReq) (*ssov1.Empty, error)
}
