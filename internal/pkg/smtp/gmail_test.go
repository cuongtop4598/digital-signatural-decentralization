package smtp

import "testing"

func TestSendMail(t *testing.T) {
	auth := NewGmailAdminAuth("", "")
	mail := NewMail("Test sending Gmail", []string{"tttrung01121999@gmail.com"}, "This is your OTP: 123456")
	mail.SendMail(auth)
}
