package reg

import (
	"context"

	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"github.com/sirupsen/logrus"
)


// RegServer - structure for implementing the gRPC service.
type RegServer struct {
	Queries *data.Queries
	Config  *config.Config // Interface for handling authentication methods.
	Log     *logrus.Logger

	ssov1.UnimplementedAuthServer
	RegService
}

// NewRegServer create new RegServer.
func NewRegServer(queries *data.Queries, cfg *config.Config, log *logrus.Logger) *RegServer {
	return &RegServer{
		Queries: queries,
		Config:  cfg,
		Log:     log,
	}
}

// RegService defines the methods for user authentication and authorization.
type RegService interface {
	Register(ctx context.Context, in *ssov1.RegisterReq) (*ssov1.Empty, error)
}
