package action_permission

import (
	"sync"

	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/sirupsen/logrus"
)

type ActionPermission struct {
	Cfg   *config.Config
	Log   *logrus.Logger
	mu    *sync.Mutex
	Queue map[string]string
}

func NewActionPermission(cfg *config.Config, log *logrus.Logger) *ActionPermission {
	return &ActionPermission{
		Cfg:   cfg,
		Log:   log,
		mu:    new(sync.Mutex),
		Queue: make(map[string]string),
	}
}
