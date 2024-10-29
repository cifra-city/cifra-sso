package authsrv

import (
	"context"

	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/cifra-city/cifra-sso/internal/email"
	"github.com/cifra-city/cifra-sso/internal/events"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// DefaultPasswordLength - default password length.
const (
	CHANGE_EMAIL    = "CHANGE_EMAIL"
	CHANGE_PASS     = "CHANGE_PASS"
	CHANGE_USERNAME = "CHANGE_USERNAME"
	CONFIRM_EMAIL   = "CONFIRM_EMAIL"
)

// AuthServer - structure for implementing the gRPC service.
type AuthServer struct {
	Queries *data.Queries
	Config  *config.Config // Interface for handling authentication methods.
	Log     *logrus.Logger
	Email   email.Mailman
	Events  events.Events

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

	Register(ctx context.Context, in *ssov1.RegisterReq) (*ssov1.Empty, error)
	Login(ctx context.Context, in *ssov1.LoginReq) (*ssov1.LoginResp, error)
	Logout(ctx context.Context, in *ssov1.Empty) (*ssov1.Empty, error)

	InquiryForChange(ctx context.Context, in *ssov1.InquiryReq) (*ssov1.InquiryResp, error)
	AccessForChanges(ctx context.Context, in *ssov1.AccessReq) (*ssov1.Empty, error)

	ChangePassword(ctx context.Context, in *ssov1.ChangePassReq) (*ssov1.Empty, error)
	ChangeUsername(ctx context.Context, in *ssov1.ChangeUsernameReq) (*ssov1.Empty, error)
	ChangeEmail(ctx context.Context, in *ssov1.ChangeEmailReq) (*ssov1.Empty, error)
}
