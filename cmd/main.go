package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stderr, "morse-converter", log.LstdFlags)
	srv := server.CreateServer(logger)
	log.Println("Running server on 8080")

	if err := srv.Server.ListenAndServe(); err != nil {
		logger.Fatalf("%s", err.Error())
	}
}
