package utils

import (
	"bytes"
	"core/config"
	"core/models"
	"fmt"
	"html/template"
	"math/rand"

	"gopkg.in/gomail.v2"
)

func GenerateOTP() string {
	const otpLength = 6
	const digits = "123456789"

	otp := make([]byte, otpLength)
	for i := range otp {
		otp[i] = digits[rand.Intn(len(digits))]
	}
	return string(otp)
}

func SendMail(templatePath string, data models.SendMail, subject string) {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(&body, data)

	m := gomail.NewMessage()
	m.SetHeader("From", config.GetConfig().PrimaryEmail)
	m.SetHeader("To", data.SendTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, config.GetConfig().PrimaryEmail, config.GetConfig().PrimaryEmailPassword)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func SendMailForInvite(templatePath string, data models.SendMail, subject string) {
	fmt.Println("inside send mail function", data)
	var body bytes.Buffer

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}

	// ✅ Pass only the dynamic map
	if err := t.Execute(&body, data.Data); err != nil {
		fmt.Println("error executing template:", err)
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", config.GetConfig().PrimaryEmail)
	m.SetHeader("To", data.SendTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, config.GetConfig().PrimaryEmail, config.GetConfig().PrimaryEmailPassword)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("error sending mail:", err)
	}
}
