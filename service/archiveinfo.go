package service

import (
	"aosmanova/doodocs/models"
	"archive/zip"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func GetInfo(file multipart.File, header *multipart.FileHeader) (*models.ZipInfo, error) {
	tempFile := fmt.Sprintf(header.Filename)

	createFile, err := os.Create(tempFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	defer createFile.Close()
	_, err = io.Copy(createFile, file)
	if err != nil {
		return nil, err
	}
	filesZip, err := zip.OpenReader(tempFile)
	if err != nil {
		return nil, err
	}

	var data []models.FileInfo
	zipSize := header.Size
	zipName := header.Filename
	var filesize_total float64
	var filesCount float64

	for _, file := range filesZip.File {

		if file.CompressedSize64 == 0 {
			continue
		}
		f, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer f.Close()
		buf := make([]byte, 512)
		r, err := f.Read(buf)
		if err != nil {
			return nil, err
		}
		data = append(data, models.FileInfo{
			File_path: file.Name,
			Size:      float64(file.CompressedSize64),
			Mimetype:  http.DetectContentType(buf[:r]),
		})

		filesize_total += float64(file.CompressedSize64)
	}
	if len(data) > 0 {
		filesCount = float64(len(data))
	}
	fmt.Println(filesCount)

	return &models.ZipInfo{
		Filename:     zipName,
		Archive_size: float64(zipSize),
		Total_files:  filesCount,
		Files:        data,
	}, nil

}
