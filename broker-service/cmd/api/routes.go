package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func (app *Config) routes() http.Handler {
	server := gin.Default()

	// Set up CORS
	server.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, Access-Control-Allow-Origin",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	server.POST("/", func(c *gin.Context) {
		app.Broker(c.Writer, c.Request)
	})

	return server

}
