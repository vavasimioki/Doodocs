package app

import (
	"aosmanova/doodocs/config"
	"log/slog"
	"net/http"
	"os"
)

type Application struct {
	logger *slog.Logger
	config *config.Config
}

func NewApplication(log *slog.Logger, config *config.Config) *Application {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	return &Application{
		logger: logger,
		config: config,
	}
}

func (a *Application) Handler(fn func(w http.ResponseWriter, r *http.Request, logger *slog.Logger),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, a.logger)
	}
}
func (a *Application) HandlerWithConfig(fn func(w http.ResponseWriter, r *http.Request, logger *slog.Logger, config *config.Config),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, a.logger, a.config)
	}
}
