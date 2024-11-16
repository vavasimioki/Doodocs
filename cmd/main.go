package main

import (
	"log/slog"
	"net/http"
	"os"

	"aosmanova/doodocs/app"
	"aosmanova/doodocs/controller"

	"github.com/gorilla/mux"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := app.NewApplication(logger)
	r := mux.NewRouter()
	r.HandleFunc("/api/archive/info", app.Handler(controller.ArchiveInformation)).Methods("POST")
	r.HandleFunc("/api/archive/create", app.Handler(controller.CreateArchive)).Methods("POST")
	r.HandleFunc("/api/archive/mail/send", app.Handler(controller.ArchiveSend)).Methods("POST")

	logger.Info("http://localhost:8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		logger.Error(err.Error())
	}

	// srv := &http.Server{
	// 	Addr:         "127.0.0.1:8000",
	// 	WriteTimeout: time.Second * 15,
	// 	ReadTimeout:  time.Second * 15,
	// 	IdleTimeout:  time.Second * 60,
	// 	Handler:      r,
	// }
	// go func() {
	// 	err := srv.ListenAndServe()
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	// err := srv.ListenAndServe()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("server connected %v", err)

}
