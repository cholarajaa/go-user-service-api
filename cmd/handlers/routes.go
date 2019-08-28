package handlers

import (
	"log"
	"mus/internal/web"
	"net/http"
)

// API will create app and handle routes
func API(log *log.Logger) http.Handler {
	// func NewHandler(logger *log.Logger, db *sqlx.DB) *Handlers {
	app := web.NewApp(log)

	// Register helloworld endpoint.
	s := Specimen{log: log}
	app.TreeMux.Handle("GET", "/hello", s.HelloWorld)
	return app
}
