package reg

import (
	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/cifra-city/cifra-sso/internal/domain/services"
	"github.com/sirupsen/logrus"
)

// Server - structure for implementing the gRPC service.
type Server struct {
	Handlers services.Registration
	Queries  *data.Queries
	Config   *config.Config // Interface for handling authentication methods.
	Log      *logrus.Logger

	ssov1.UnimplementedRegServer
}

func NewRegServer(queries *data.Queries, cfg *config.Config, log *logrus.Logger) *Server {
	return &Server{
		Queries: queries,
		Config:  cfg,
		Log:     log,
	}
}
