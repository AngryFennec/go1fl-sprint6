package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Server *http.Server
	Logger *log.Logger
}

func CreateServer(logger *log.Logger) *Server {
	router := chi.NewRouter()
	router.Get("/", handlers.IndexHandler)
	router.Post("/upload", handlers.UploadHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{
		Logger: logger,
		Server: server,
	}

}
