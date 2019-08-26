package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// router
	http.HandleFunc("/", HelloWorld)

	// logger
	log := log.New(os.Stdout, "AUTH : ", log.Lmicroseconds|log.Lshortfile|log.LstdFlags)
	log.Println("Listening and serving on localhost:3000, press ctrl+c to cancel")

	// server
	http.ListenAndServe("localhost:3000", nil)
}

//Handler

// HelloWorld is how we start the programming
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("request processed")
	fmt.Fprintln(w, "Hello World!")
}
