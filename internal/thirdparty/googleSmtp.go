package thirdparty

import(
	"net/smtp"
	"log"
	"shopping/internal/config"
)

func SendVerificationEmail(toAdd string, token string) bool {
	var mailConfig config.MailConfig
	from := mailConfig.WebsiteEmail
	pass := mailConfig.WebsitePassword
	to := toAdd
	subject := "Verify your email address"
	body := "Please click on the link to verify your email address: https://ell.kreativ.com/verify?token=" + token
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("smtp error: %s", err)
		return false
	}
	return true
}