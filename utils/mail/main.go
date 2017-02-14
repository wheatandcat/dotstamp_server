package maill

import (
	"bytes"
	"dotstamp_server/utils"
	"html/template"
	"net/smtp"

	"github.com/astaxie/beego"
)

// Body 本文
type Body struct {
	From    string
	To      string
	Subject string
	Message string
}

// Send メールを送信する
func Send(email string, body []byte) error {
	if utils.IsTest() {
		return nil
	}

	auth := smtp.PlainAuth(
		"",
		beego.AppConfig.String("email"),
		beego.AppConfig.String("emailpass"),
		"smtp.gmail.com",
	)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		beego.AppConfig.String("email"),
		[]string{email},
		body,
	)

	return err
}

// GetBody 本文を取得する
func GetBody(b Body) []byte {
	buffer := new(bytes.Buffer)
	template := template.Must(template.New("emailTemplate").Parse(getBodyTemplate()))
	template.Execute(buffer, &b)

	return buffer.Bytes()
}

// getBodyTemplate 本文テンプレートを使用する
func getBodyTemplate() string {
	return "To: {{.To}}\r\n" +
		"Subject: {{.Subject}}\r\n" +
		"\r\n" +
		"{{.Message}}"
}
