package main

import (
	"net/http"
	"time"
)

type Config struct{}

// This will be the main entry point to my broker
func main() {

	app := Config{}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        app.routes(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
