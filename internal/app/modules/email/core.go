package email

import (
	"crypto/rand"
	"fmt"
	"time"
)

func (m *Mailman) AddToEmailList(username string, code string) {
	m.mu.Lock()
	m.ListCode[username] = code
	m.mu.Unlock()

	time.AfterFunc(180*time.Second, func() {
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
		m.Log.Infof("Code for user %s is correct and has been used", username)
		return true
	} else if exists {
		m.Log.Infof("Incorrect code for user %s", username)
	} else {
		m.Log.Infof("No code found for user %s", username)
	}

	return false
}

func GenerateConfirmationCode() (string, error) {
	b := make([]byte, 3)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	code := fmt.Sprintf("%06d", int(b[0])<<16|int(b[1])<<8|int(b[2]))
	return code, nil
}
