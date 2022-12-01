package http_helpers

import (
	"fmt"
	"log"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int, message string) {
	log.Println("Some http error is happening")
	w.WriteHeader(status)
	var err error
	if message != "" {
		_, err = fmt.Fprint(w, message)
	} else if status == http.StatusNotFound {
		log.Printf("Returning 404 Not found error, the path '%s' was requested\n", r.URL.Path)
		_, err = fmt.Fprint(w, "404 not found")
	} else if status == http.StatusBadRequest {
		log.Printf("Returning 400 Bad request error, the path '%s' was requested\n", r.URL.Path)
		_, err = fmt.Fprint(w, "400 bad request")
	} else if status == http.StatusMethodNotAllowed {
		log.Printf("Returning 405 Method Not Allowed error, the path '%s' was requested\n", r.URL.Path)
		_, err = fmt.Fprint(w, "405 method not allowed")
	} else {
		log.Printf("Something went wrong, the path '%s' was requested, returning %v error code\n", r.URL.Path, status)
		_, err = fmt.Fprintf(w, "Error %v, Something went wrong", status)
	}
	if err != nil {
		log.Fatalf("Something went wrong with the error: %s", err)
	}
}
