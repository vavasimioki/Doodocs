package controller

import (
	"aosmanova/doodocs/models"
	"archive/zip"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func ArchiveInformation(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		serverError(w, "Internal Server Error")
		logger.Error(err.Error(), "method", r.Method, "uri", r.RequestURI)
		return
	}
	defer file.Close()

	if !strings.HasSuffix(strings.ToLower(header.Filename), ".zip") {
		httpError(w, "bad request", http.StatusBadRequest)
		logger.Error("The file is not ZIP format", "filename", header.Filename, "uri", r.RequestURI)
		return
	}

	archiveInfo, err := GetInfo(file, header)
	if err != nil {

		serverError(w, "Internal Server Error")
		logger.Error(err.Error())
	}

	EncodeOK(w, archiveInfo)

}
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

func CreateArchive(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	w.Write([]byte("create Arch"))
}

func ArchiveSend(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	w.Write([]byte("archive send"))

}
