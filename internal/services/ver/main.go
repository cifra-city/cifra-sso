package ver

import (
	"context"

	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/cifra-city/cifra-sso/internal/tools/email"
	"github.com/cifra-city/cifra-sso/internal/tools/events"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"github.com/sirupsen/logrus"
)

// VerServer - structure for implementing the gRPC service.
type VerServer struct {
	Queries *data.Queries
	Config  *config.Config // Interface for handling authentication methods.
	Log     *logrus.Logger
	Email   email.Mailman
	Events  events.Events

	ssov1.UnimplementedAuthServer
	VerService
}

// NewVerServer create new RegServer.
func NewVerServer(queries *data.Queries, cfg *config.Config, log *logrus.Logger) *VerServer {
	return &VerServer{
		Queries: queries,
		Config:  cfg,
		Log:     log,
		Email:   *email.NewMailman(cfg, log),
		Events:  *events.NewEvents(cfg, log),
	}
}

// VerService defines the methods for user authentication and authorization.
type VerService interface {
	InquiryForChange(ctx context.Context, in *ssov1.InquiryReq) (*ssov1.InquiryResp, error)
	AccessForChanges(ctx context.Context, in *ssov1.AccessReq) (*ssov1.AccessResp, error)
}
