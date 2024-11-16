package service

import (
	"archive/zip"
	"io"
	"mime/multipart"
	"os"
)

func CreateZipArchive(file []*multipart.FileHeader) (*os.File, error) {
	z, err := os.CreateTemp("", "*.zip")
	if err != nil {
		return nil, err
	}

	defer z.Close()
	zipW := zip.NewWriter(z)

	for _, f := range file {
		f1, err := f.Open()
		if err != nil {
			return nil, err
		}
		defer f1.Close()
		w1, err := zipW.Create(f.Filename)
		if err != nil {
			return nil, err
		}
		defer zipW.Close()
		_, err = io.Copy(w1, f1)
		if err != nil {
			return nil, err
		}

	}
	return z, nil
}
