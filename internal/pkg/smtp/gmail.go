package smtp

import (
	"errors"
	"fmt"
	"log"
	originSmtp "net/smtp"
	"os"
)

type GmailAdminAuth struct {
	username string
	password string
}

func NewGmailAdminAuth(username string, password string) originSmtp.Auth {
	if username == "" && password == "" {
		username, _ = os.LookupEnv("GMAIL_USERNAME")
		password, _ = os.LookupEnv("GMAIL_PASSWORD")
	}
	return &GmailAdminAuth{username, password}
}
func (a *GmailAdminAuth) GetUsername() string {
	return a.username
}
func (a *GmailAdminAuth) Start(server *originSmtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *GmailAdminAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}

type Mail struct {
	Subject string
	To      []string
	Message string
}

func NewMail(subject string, to []string, message string) *Mail {
	return &Mail{
		Subject: subject,
		To:      to,
		Message: message,
	}
}

func (mail *Mail) SendMail(auth originSmtp.Auth) {
	admin, _ := os.LookupEnv("GMAIL_USERNAME")
	// Here we do it all: connect to our server, set up a message and send it
	message := fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", mail.To, mail.Subject, mail.Message)
	msg := []byte(message)
	err := originSmtp.SendMail("smtp.gmail.com:587", auth, admin, mail.To, msg)
	if err != nil {
		log.Fatal(err)
	}
}
