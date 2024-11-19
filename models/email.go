package models

type EmailsList struct {
	Email []string `json:"email"`
}
type FileDitail struct {
	Contenttype string `json:"content-type"`
	Filename    string `json:"filename"`
}
