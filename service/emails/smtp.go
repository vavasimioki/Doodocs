package emails

import (
	"aosmanova/doodocs/config"
	"net/smtp"
)

func SendToMail(fileContent []byte, cfg *config.Config, mails []string) error {
	auth := smtp.PlainAuth("", cfg.User, cfg.Password, cfg.Host)

	err := smtp.SendMail(cfg.Host+":"+cfg.Port, auth, cfg.User, mails, fileContent)

	if err != nil {

		return err
	}
	return nil
}
