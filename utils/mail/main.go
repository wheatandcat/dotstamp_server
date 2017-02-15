package mail

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

// ForgetpasswordTemplate パスワード忘れた本文
type ForgetpasswordTemplate struct {
	URL   string
	Host  string
	Email string
}

// GetForgetpasswordBody パスワード忘れた本文を取得する
func GetForgetpasswordBody(f ForgetpasswordTemplate) []byte {
	t := "パスワードの変更申請を受け付けました。\r\n" +
		"\r\n" +
		"下記のURLから、、パスワード変更手続きをしてください\r\n" +
		"{{.URL}}\r\n" +
		"\r\n" +
		"※このURLは発行から1時間有効です。\r\n" +
		"※1時間以内に複数回パスワードの変更申請を行った場合は、直近で発行されたURLのみ有効になるのでご注意ください。\r\n" +
		"\r\n" +
		"--------------------------------\r\n" +
		"dotstamp :{{.Host}}\r\n" +
		"お問い合わせ：{{.Email}}"

	buffer := new(bytes.Buffer)
	template := template.Must(template.New("forgetPassword").Parse(t))
	template.Execute(buffer, &f)

	return buffer.Bytes()
}

// GetForgetpasswordURL パスワード忘れたURLを取得する
func GetForgetpasswordURL(email string, keyword string) (string, error) {
	e, err := utils.Encrypter([]byte(email))
	if err != nil {
		return "", err
	}
	k, err := utils.Encrypter([]byte(keyword))
	if err != nil {
		return "", err
	}

	return utils.Urlencode(e) + "/" + utils.Urlencode(k), nil
}
