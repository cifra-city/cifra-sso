package email

import (
	"sync"

	"github.com/cifra-city/cifra-sso/internal/config"
	"github.com/sirupsen/logrus"
)

type Mailman struct {
	Cfg           *config.Config
	Log           *logrus.Logger
	mu            *sync.Mutex
	emailListCode map[string]string
}

type mailman interface {
	AddToEmailList(username string, code string)
	CheckInEmailList(username string, code string) bool
	SendConfirmationEmail(to string, code string) error
	GenerateConfirmationCode() (string, error)
}

func NewMailman(cfg *config.Config, log *logrus.Logger) *Mailman {
	return &Mailman{
		Cfg:           cfg,
		Log:           log,
		mu:            &sync.Mutex{},
		emailListCode: make(map[string]string),
	}
}
