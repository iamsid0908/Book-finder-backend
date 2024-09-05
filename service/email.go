package service

import (
	"core/config"
	"core/domain"
	"core/models"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	EmailDomain domain.EmailDomain
}

func (c *EmailService) SendOneEmail(param models.Sendsingleemail) error {
	htmlContent := `
		<html>
			<body>
				<div class="title" style="border: 1px solid red; height: 50px; display: flex; align-items: center; justify-content: center; background-color: tomato; color: white;">
					hi sid how are u 
				</div>
				<div class="des" style="border: 1px solid red; height: 100px; display: flex; align-items: center; justify-content: center;">
					this is my description area
				</div>
				<div class="footer" style="border: 1px solid red; height: 50px; display: flex; align-items: center; justify-content: center; background-color: violet; color: white;">
					this is my footer and link area
				</div>
			</body>
		</html>
	`
	m := gomail.NewMessage()
	m.SetHeader("From", config.GetConfig().PrimaryEmail)
	m.SetHeader("To", param.Email)
	m.SetHeader("Subject", param.Subject)
	m.SetBody("text/html", htmlContent)
	d := gomail.NewDialer("smtp.gmail.com", 587, config.GetConfig().PrimaryEmail, config.GetConfig().PrimaryEmailPassword)

	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil

}
