package auth

import (
	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
	"github.com/cifra-city/cifra-sso/internal/app/modules/action_permission"
	"github.com/cifra-city/cifra-sso/internal/app/modules/email"
	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/cifra-city/cifra-sso/internal/domain/services"
	"github.com/sirupsen/logrus"
)

const (
	ChangeEmail    = "CHANGE_EMAIL"
	ChangePassword = "CHANGE_PASS"
	ChangeUsername = "ChangeUsername"
)

// Server - structure for implementing the gRPC service.
type Server struct {
	Handlers         services.Authentication
	Queries          *data.Queries
	Config           *config.Config
	Log              *logrus.Logger
	Email            *email.Mailman
	ActionPermission *action_permission.ActionPermission

	ssov1.UnimplementedAuthServer
}

func NewAuthServer(queries *data.Queries, cfg *config.Config, log *logrus.Logger, mailman *email.Mailman,
	ActionPermission *action_permission.ActionPermission) *Server {
	return &Server{
		Queries:          queries,
		Config:           cfg,
		Log:              log,
		Email:            mailman,
		ActionPermission: ActionPermission,
	}
}
