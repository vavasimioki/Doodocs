package main

import (
	"io/ioutil"
	"log/slog"
	"net/http"
	"os"

	"aosmanova/doodocs/app"
	"aosmanova/doodocs/config"
	"aosmanova/doodocs/controller"

	"github.com/go-yaml/yaml"
	"github.com/gorilla/mux"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	cfg, err := LoadConfig()
	if err != nil {
		logger.Error(err.Error(), "Error loading yaml file")
	}

	app := app.NewApplication(logger, cfg)

	r := mux.NewRouter()
	r.HandleFunc("/api/archive/info", app.Handler(controller.ArchiveInformation)).Methods("POST")
	r.HandleFunc("/api/archive/create", app.Handler(controller.CreateArchive)).Methods("POST")
	r.HandleFunc("/api/archive/mail/send", app.HandlerWithConfig(controller.ArchiveSend)).Methods("POST")

	logger.Info("http://localhost:8000")
	err = http.ListenAndServe(":8000", r)
	if err != nil {
		logger.Error(err.Error())
	}

}

func LoadConfig() (*config.Config, error) {
	cnf := &config.Config{}

	c, err := ioutil.ReadFile(config.ConfigFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(c, cnf)
	if err != nil {
		return nil, err
	}
	return cnf, nil
}
