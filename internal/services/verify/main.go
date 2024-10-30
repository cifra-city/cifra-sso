package verify

import (
	"context"

	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/cifra-city/cifra-sso/internal/modules/email"
	"github.com/cifra-city/cifra-sso/internal/modules/events"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"github.com/sirupsen/logrus"
)

// VerifyServer - structure for implementing the gRPC service.
type VerifyServer struct {
	Queries *data.Queries
	Config  *config.Config // Interface for handling authentication methods.
	Log     *logrus.Logger
	Email   *email.Mailman
	Events  *events.Events

	ssov1.UnimplementedAuthServer
	VerifyService
}

// NewVerifyServer create new RegServer.
func NewVerifyServer(queries *data.Queries, cfg *config.Config, log *logrus.Logger, mailman *email.Mailman, eventsList *events.Events) *VerifyServer {
	return &VerifyServer{
		Queries: queries,
		Config:  cfg,
		Log:     log,
		Email:   mailman,
		Events:  eventsList,
	}
}

// VerifyService defines the methods for user authentication and authorization.
type VerifyService interface {
	InquiryForChange(ctx context.Context, in *ssov1.InquiryReq) (*ssov1.InquiryResp, error)
	AccessForChanges(ctx context.Context, in *ssov1.AccessReq) (*ssov1.AccessResp, error)
}
