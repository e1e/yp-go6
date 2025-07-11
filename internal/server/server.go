package server

import (
	"log"
	"net/http"
	"time"
	"yp-go6/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func Create(logger *log.Logger) Server {

	router := http.NewServeMux()
	router.HandleFunc("/", handlers.GetHtml)
	router.HandleFunc("/upload", handlers.Upload)

	server := http.Server{
		Addr:         "8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 15,
	}

	return Server{
		Logger: logger,
		Server: &server,
	}
}
