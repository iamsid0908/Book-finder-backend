package models

type Sendsingleemail struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
}

type SendMail struct {
	SendTo   string
	UserName string
	OTP      string
	Data     map[string]interface{}
}
