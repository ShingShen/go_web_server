package gmailsender

import (
	"fmt"
	"net/smtp"
	"server/dto"
	"server/utils/operator"
)

func GmailSender(recipient string, msg string) error {
	var config dto.Config
	configData, _ := operator.LoadJson(config, "config/env/config.json")
	sender := configData.(map[string]interface{})["gmail"].(map[string]interface{})["sender"].(string)
	appPassword := configData.(map[string]interface{})["gmail"].(map[string]interface{})["app_password"].(string)
	host := configData.(map[string]interface{})["gmail"].(map[string]interface{})["smtp_mail"].(string)
	port := int(configData.(map[string]interface{})["gmail"].(map[string]interface{})["port"].(float64))
	addr := fmt.Sprintf("%s:%d", host, port)

	auth := smtp.PlainAuth(
		"",
		sender,
		appPassword,
		host,
	)

	err := smtp.SendMail(
		addr,
		auth,
		sender,
		[]string{recipient},
		[]byte(msg),
	)

	if err != nil {
		return err
	}
	return nil

	// msg := "Subject: My special\nThis is the body of my email. HAHAHA!!!!!"
}
