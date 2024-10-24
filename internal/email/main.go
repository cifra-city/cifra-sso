package email

import (
	"sync"

	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/sirupsen/logrus"
)

type Mailman struct {
	Config        *config.Config
	Log           *logrus.Logger
	mu            *sync.Mutex
	emailListCode map[string]string
}

func NewMailman(cfg *config.Config, log *logrus.Logger) *Mailman {
	return &Mailman{
		Config:        cfg,
		Log:           log,
		mu:            &sync.Mutex{},
		emailListCode: make(map[string]string),
	}
}

type mailman interface {
	AddToEmailList(username string, code string)
	CheckInEmailList(username string, code string)
	SendConfirmationEmail(to string, code string)
	GenerateConfirmationCode()
}
