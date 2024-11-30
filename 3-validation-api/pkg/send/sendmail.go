package sendmail

import (
	configmail "MyDz/3-validation-api/configs"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func СreateMail(MailTo string, MailToken string) {
	PassStr := string(configmail.LoadConfig().Pass.Pass)
	MailFrom := string(configmail.LoadConfig().Addr.Address)
	token := string(MailToken)
	e := email.NewEmail()
	e.From = MailFrom
	e.To = []string{MailTo}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>You link verify! </h1>" + "http://localhost:8081/verify/" + token)
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", MailFrom, PassStr, "smtp.gmail.com"))
	if err != nil {
		log.Fatal(err)
	}
}
