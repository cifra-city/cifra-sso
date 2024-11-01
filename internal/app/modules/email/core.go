package email

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"
)

func (m *Mailman) AddToEmailList(username string, code string) {
	m.mu.Lock()
	m.ListCode[username] = code
	m.mu.Unlock()

	time.AfterFunc(60*time.Second, func() {
		m.mu.Lock()
		defer m.mu.Unlock()
		if _, exists := m.ListCode[username]; exists {
			delete(m.ListCode, username)
		}
	})
}

func (m *Mailman) CheckInEmailList(username string, code string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if storedCode, exists := m.ListCode[username]; exists && storedCode == code {
		delete(m.ListCode, username)
		return true
	}

	return false
}

func (m *Mailman) SendConfirmationEmail(to string, code string) error {
	mes := gomail.NewMessage()

	mes.SetHeader("From", m.Cfg.Email.Address)
	mes.SetHeader("To", to)
	mes.SetHeader("Subject", "Email Confirmation")

	mes.SetBody("text/plain", fmt.Sprintf("Your confirmation code: %s", code))

	smtpPort, err := strconv.Atoi(m.Cfg.Email.SmtpPort)
	if err != nil {
		m.Log.Fatalf("Invalid SMTP port: %v", err)
	}

	d := gomail.NewDialer(m.Cfg.Email.SmtpHost, smtpPort, m.Cfg.Email.Address, m.Cfg.Email.Password)

	if err := d.DialAndSend(mes); err != nil {
		return err
	}

	m.AddToEmailList(to, code)

	return nil
}

func GenerateConfirmationCode() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	code := fmt.Sprintf("%06d", b[0:3])
	return code, nil
}
