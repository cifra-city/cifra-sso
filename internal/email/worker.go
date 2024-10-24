package email

import (
	"crypto/rand"
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
)

func (m *Mailman) AddToEmailList(username string, code string) {
	m.mu.Lock()
	m.emailListCode[username] = code
	m.mu.Unlock()

	time.AfterFunc(60*time.Second, func() {
		m.mu.Lock()
		defer m.mu.Unlock()
		if _, exists := m.emailListCode[username]; exists {
			delete(m.emailListCode, username)
		}
	})
}

func (m *Mailman) CheckInEmailList(username string, code string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if storedCode, exists := m.emailListCode[username]; exists && storedCode == code {
		delete(m.emailListCode, username)
		return true
	}

	return false
}

func (m *Mailman) SendConfirmationEmail(to string, code string) error {
	mes := gomail.NewMessage()

	mes.SetHeader("From", "testemail1488@proton.me")
	mes.SetHeader("To", to)
	mes.SetHeader("Subject", "Email Confirmation")

	mes.SetBody("text/plain", fmt.Sprintf("Your confirmation code: %s", code))

	d := gomail.NewDialer("smtp.example.com", 587, "testemail1488@proton.me", "=APWBAz-%6MCzxz")

	if err := d.DialAndSend(mes); err != nil {
		return err
	}

	m.AddToEmailList(to, code)

	return nil
}

func (m *Mailman) GenerateConfirmationCode() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	code := fmt.Sprintf("%06d", b[0:3])
	return code, nil
}
