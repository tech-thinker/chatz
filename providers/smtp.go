package providers

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strings"

	"github.com/tech-thinker/chatz/config"
)

type SMTPProvider struct {
	config *config.Config
}

func (agent *SMTPProvider) Post(message string) (interface{}, error) {
	return agent.PostWithSubject(message, "")
}

func (agent *SMTPProvider) PostWithSubject(message string, subject string) (interface{}, error) {
	host := agent.config.SMTPHost
	port := agent.config.SMTPPort
	smtpServer := fmt.Sprintf("%s:%s", host, port)
	user := agent.config.SMTPUser
	password := agent.config.SMTPPassword

	if len(subject) == 0 {
		subject = agent.config.SMTPSubject
	}
	from := agent.config.SMTPFrom
	recipients := agent.config.SMTPTo

	if len(from) == 0 {
		from = user
	}

	if len(subject) == 0 {
		subject = "Chatz Notification"
	}

	// Create the message with proper headers
	msg := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		from, recipients, subject, message,
	))

	// Split recipients into a slice
	to := strings.Split(recipients, ",")

	// Handle UseTLS
	if agent.config.UseTLS {
		tlsConfig := &tls.Config{
			ServerName: host,
		}

		// Establish a secure connection using TLS
		conn, err := tls.Dial("tcp", smtpServer, tlsConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to SMTP server using TLS: %w", err)
		}
		defer conn.Close()

		client, err := smtp.NewClient(conn, host)
		if err != nil {
			return nil, fmt.Errorf("failed to create SMTP client over TLS: %w", err)
		}
		defer client.Quit()

		// Authenticate
		auth := smtp.PlainAuth("", user, password, host)
		if err = client.Auth(auth); err != nil {
			return nil, fmt.Errorf("failed to authenticate: %w", err)
		}

		// Send the email
		return sendEmail(client, from, to, msg)
	}

	// Connect without encryption initially
	conn, err := net.Dial("tcp", smtpServer)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SMTP server: %w", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return nil, fmt.Errorf("failed to create SMTP client: %w", err)
	}
	defer client.Quit()

	// Handle UseSTARTTLS
	if agent.config.UseSTARTTLS {
		tlsConfig := &tls.Config{
			ServerName: host,
		}

		// Upgrade to STARTTLS
		if err = client.StartTLS(tlsConfig); err != nil {
			return nil, fmt.Errorf("failed to start TLS: %w", err)
		}
	}

	// Authenticate
	auth := smtp.PlainAuth("", user, password, host)
	if err = client.Auth(auth); err != nil {
		return nil, fmt.Errorf("failed to authenticate: %w", err)
	}

	// Send the email
	return sendEmail(client, from, to, msg)
}

// Helper function to send email
func sendEmail(client *smtp.Client, from string, to []string, msg []byte) (interface{}, error) {
	// Set the sender and recipients
	if err := client.Mail(from); err != nil {
		return nil, fmt.Errorf("failed to set sender: %w", err)
	}
	for _, recipient := range to {
		if err := client.Rcpt(recipient); err != nil {
			return nil, fmt.Errorf("failed to add recipient %s: %w", recipient, err)
		}
	}

	// Write the email body
	writer, err := client.Data()
	if err != nil {
		return nil, fmt.Errorf("failed to open data writer: %w", err)
	}
	_, err = writer.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to write email body: %w", err)
	}
	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close data writer: %w", err)
	}

	return `{"status": "success"}`, nil
}

func (agent *SMTPProvider) Reply(threadId string, message string) (interface{}, error) {
	fmt.Println("Reply to SMTP not supported yet.")
	return nil, errors.New("Reply to SMTP not supported yet.")
}
