package events

import (
	"sync"

	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/sirupsen/logrus"
)

type Events struct {
	Cfg   *config.Config
	Log   *logrus.Logger
	mu    *sync.Mutex
	Queue map[string]string
}

func NewEvents(cfg *config.Config, log *logrus.Logger) *Events {
	return &Events{
		Cfg:   cfg,
		Log:   log,
		mu:    new(sync.Mutex),
		Queue: make(map[string]string),
	}
}
