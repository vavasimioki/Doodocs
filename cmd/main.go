package main

import (
	"log"
	"net/http"

	"aosmanova/doodocs/controller"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/archive/info", controller.ArchiveInformation).Methods("POST")
	r.HandleFunc("/api/archive/create", controller.CreateArchive).Methods("POST")
	r.HandleFunc("/api/archive/mail/send", controller.ArchiveSend).Methods("POST")

	log.Printf("http://localhost:8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
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
