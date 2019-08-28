package web

import (
	"log"

	"github.com/dimfeld/httptreemux"
)

//App is entry point to out application
type App struct {
	*httptreemux.TreeMux
	log *log.Logger
}

// NewApp to handle routes
func NewApp(log *log.Logger) *App {
	app := App{
		TreeMux: httptreemux.New(),
		log:     log,
	}

	return &app
}
