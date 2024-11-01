package email

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"gopkg.in/gomail.v2"
)

// SendConfirmationCode отправляет письмо с подтверждением, используя HTML-шаблон из файла.
func (m *Mailman) SendConfirmationCode(to string, code string) error {
	mes := gomail.NewMessage()
	mes.SetHeader("From", m.Cfg.Email.Address)
	mes.SetHeader("To", to)
	mes.SetHeader("Subject", "Email Confirmation")

	// Чтение HTML-шаблона из файла
	templatePath := filepath.Join("email_list.html")
	tmplContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		m.Log.Errorf("Error reading email template: %v", err)
		return err
	}

	// Создаем новый шаблон и парсим HTML
	tmpl, err := template.New("emailTemplate").Parse(string(tmplContent))
	if err != nil {
		m.Log.Errorf("Error parsing email template: %v", err)
		return err
	}

	// Используем bytes.Buffer для хранения результата шаблона
	var bodyContent bytes.Buffer
	err = tmpl.Execute(&bodyContent, map[string]string{"Code": code})
	if err != nil {
		m.Log.Errorf("Error executing template: %v", err)
		return err
	}

	// Устанавливаем HTML-сообщение в качестве тела письма
	mes.SetBody("text/html", bodyContent.String())

	// Настройка SMTP-соединения и отправка письма
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
