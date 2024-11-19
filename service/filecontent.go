package service

import (
	"bytes"
	"mime/multipart"
)

func GetFileContent(f multipart.File, h *multipart.FileHeader) ([]byte, error) {
	if !IsValidContentTypeForMail(h.Header.Get("Content-Type")) {
		return nil, IncorrectContentType
	}

	buf := bytes.NewBuffer(nil)
	_, err := buf.ReadFrom(f)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return buf.Bytes(), nil
}
