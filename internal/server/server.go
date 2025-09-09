package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func Create(logger *log.Logger) Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HandlerRoot)
	mux.HandleFunc("/upload", handlers.HandlerUpload)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 15,
	}

	return Server{logger, &server}
}
