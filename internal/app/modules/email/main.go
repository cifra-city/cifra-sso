package email

import (
	"sync"

	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/sirupsen/logrus"
)

type Mailman struct {
	Cfg      *config.Config
	Log      *logrus.Logger
	mu       *sync.Mutex
	ListCode map[string]string
}

func NewMailman(cfg *config.Config, log *logrus.Logger) *Mailman {
	return &Mailman{
		Cfg:      cfg,
		Log:      log,
		mu:       &sync.Mutex{},
		ListCode: make(map[string]string),
	}
}
