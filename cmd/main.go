package main

import (
	"log"
	"mus/cmd/handlers"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Println("error :", err)
		os.Exit(1)
	}
}

func run() error {

	// logger
	log := log.New(os.Stdout, "AUTH : ", log.Lmicroseconds|log.Lshortfile|log.LstdFlags)

	api := http.Server{
		Addr:         "0.0.0.0:3000",
		Handler:      handlers.API(log),
		ReadTimeout:  5,
		WriteTimeout: 5,
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// server
	go func() {
		log.Printf("Listening and serving on %s, press ctrl+c to cancel", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return err
	}

}
