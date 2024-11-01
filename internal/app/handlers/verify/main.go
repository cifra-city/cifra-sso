package verify

import (
	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
	"github.com/cifra-city/cifra-sso/internal/app/modules/action_permission"
	"github.com/cifra-city/cifra-sso/internal/app/modules/email"
	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/cifra-city/cifra-sso/internal/domain/services"
	"github.com/sirupsen/logrus"
)

// Server - structure for implementing the gRPC service.
type Server struct {
	Handlers         services.Verification
	Queries          *data.Queries
	Config           *config.Config // Interface for handling authentication methods.
	Log              *logrus.Logger
	Email            *email.Mailman
	ActionPermission *action_permission.ActionPermission

	ssov1.UnimplementedVerifyServer
}

func NewVerifyServer(queries *data.Queries, cfg *config.Config, log *logrus.Logger, mailman *email.Mailman, ActionPermission *action_permission.ActionPermission) *Server {
	return &Server{
		Queries:          queries,
		Config:           cfg,
		Log:              log,
		Email:            mailman,
		ActionPermission: ActionPermission,
	}
}
