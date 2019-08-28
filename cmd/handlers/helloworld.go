package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// Specimen to get inputs for the handle
type Specimen struct {
	log *log.Logger
}

// HelloWorld is how we start the programming
func (s *Specimen) HelloWorld(w http.ResponseWriter, r *http.Request, params map[string]string) {
	defer s.log.Println("request processed")
	fmt.Fprintln(w, "Hello World!")
}
