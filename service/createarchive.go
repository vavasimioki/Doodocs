package service

import (
	"fmt"
	"mime/multipart"
)

func GetBody(fileHeader []*multipart.FileHeader) ([]*multipart.FileHeader, error) {

	var files []*multipart.FileHeader
	var content string
	for _, f := range fileHeader {
		content = f.Header.Get("Content-Type")
		if !IsValidContentType(content) { 
			return nil, IncorrectContentType
		}
		fmt.Printf("content %s\n", content)
		files = append(files, f)
	}

	return files, nil

}
