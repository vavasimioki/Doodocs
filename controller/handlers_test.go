package controller

import (
	"aosmanova/doodocs/config"
	"bytes"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestArchiveInformation(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, _ := writer.CreateFormFile("file", "test.zip")
	part.Write([]byte("mock zip content"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/api/archive/info", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()

	ArchiveInformation(rr, req, logger)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	if rr.Body.Len() == 0 {
		t.Errorf("expected non-empty response body")
	}
}

func TestArchiveSend(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	cfg := &config.Config{
		Host:     "smtp.gmail.com",
		Port:     "587",
		User:     "test@example.com",
		Password: "password",
	}

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, _ := writer.CreateFormFile("file", "file.docx")
	part.Write([]byte("mock file content"))
	writer.WriteField("emails", "test1@example.com,test2@example.com")
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/api/archive/mail/send", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()

	ArchiveSend(rr, req, logger, cfg)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "OK") {
		t.Errorf("expected response to contain 'OK', got %s", rr.Body.String())
	}
}
