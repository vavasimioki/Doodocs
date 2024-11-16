package app

import (
	"log/slog"
	"net/http"
	"os"
)

type Application struct {
	logger *slog.Logger
}

func NewApplication(log *slog.Logger) *Application {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	return &Application{
		logger: logger,
	}
}

func (a *Application) Handler(fn func(w http.ResponseWriter, r *http.Request, logger *slog.Logger),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, a.logger)
	}
}
