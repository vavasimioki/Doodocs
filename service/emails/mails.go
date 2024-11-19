package emails

import (
	"aosmanova/doodocs/service"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func IsValidMailLenth(mails string) bool {

	return len(mails) > 0 && len(mails) <= 320
}

func IsValidMailLetters(mail string) bool {
	return emailRegex.MatchString(mail)
}
func GetMails(mails string) ([]string, error) {
	e := strings.Split(mails, ",")
	str := []string{}
	for _, mail := range e {
		mail = strings.TrimSpace(mail)
		if !IsValidMailLenth(mail) {
			return nil, service.IsNotVallidEmail
		}
		if !IsValidMailLetters(mail) {
			return nil, service.IsNotVallidEmail
		}
		str = append(str, mail)

	}

	return str, nil

}
