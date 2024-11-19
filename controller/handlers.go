package controller

import (
	"aosmanova/doodocs/config"
	"aosmanova/doodocs/service"
	"aosmanova/doodocs/service/emails"
	"encoding/json"
	"errors"
	"log"
	"log/slog"
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

	archiveInfo, err := service.GetInfo(file, header)
	if err != nil {

		serverError(w, "Internal Server Error")
		logger.Error(err.Error())
	}

	EncodeOK(w, archiveInfo)

}

func CreateArchive(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		serverError(w, "Internal Server Error")
		logger.Error(err.Error(), "Error parsing multipart form")
		return
	}
	fileHeader := r.MultipartForm.File["files[]"]
	f, err := service.GetBody(fileHeader)
	if err != nil {
		if errors.Is(err, service.IncorrectContentType) {
			httpError(w, "Not Allowed Content Type of the files", http.StatusUnsupportedMediaType)
			logger.Error(err.Error(), "Not Allowed Content Type of the file")
			return
		}
		httpError(w, "Internal Serve Errr", http.StatusInternalServerError)
		logger.Error(err.Error(), "error parse multipart form")
		return
	}

	zip, err := service.CreateZipArchive(f)
	if err != nil {
		serverError(w, "Internal Server Error")
		logger.Error(err.Error(), "method", r.Method, "uri", r.RequestURI)
		return
	}
	defer os.Remove(zip.Name())
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", `attachment; filename="archive.zip"`)
	w.WriteHeader(http.StatusOK)

	z, err := os.ReadFile(zip.Name())
	if err != nil {
		serverError(w, "Internal Server Error")
		logger.Error(err.Error(), "method", r.Method, "uri", r.RequestURI)
		return
	}

	json.NewEncoder(w).Encode(z)

}

func ArchiveSend(w http.ResponseWriter, r *http.Request, logger *slog.Logger, cfg *config.Config) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Fatal(err)
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		logger.Error(err.Error())
		return
	}

	fileContent, err := service.GetFileContent(file, header)
	if err != nil {
		serverError(w, "Internal Server Error")
		logger.Error(err.Error(), "Error reading file")
		return
	}

	mails, err := emails.GetMails(r.FormValue("emails"))
	if err != nil {
		httpError(w, "Email Not Valid, try again", http.StatusBadRequest)
		logger.Error(err.Error())
		return
	}

	err = emails.SendToMail(fileContent, cfg, mails)
	if err != nil {
		serverError(w, "Internal Server Error")
		logger.Error(err.Error(), "Error sending messages")
		return
	}
	logger.Info("file succesfull send")

	EncodeOK(w, "OK")

}
