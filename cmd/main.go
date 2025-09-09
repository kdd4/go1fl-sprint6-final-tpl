package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.Flags())

	server := server.Create(logger)

	err := server.Server.ListenAndServe()

	if err != nil {
		log.Fatal("Error while starting server")
	}
}
