package main

import (
	"log"
	"net/http"
	"os"
	"yp-go6/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	server := server.Create(logger)
	port := ":" + server.Server.Addr

	err := http.ListenAndServe(port, server.Server.Handler)
	if err != nil {
		server.Logger.Fatal(err)
	}
}
