package endpoint

import (
	"fmt"
	"log"
	"net/http"
)

type Endpoint interface {
	GetPath() string
	Handler(w http.ResponseWriter, r *http.Request)
}

func Handle(e Endpoint) {
	log.Printf("Handling '%s' endpoint", e.GetPath())
	http.HandleFunc(e.GetPath(), e.Handler)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int, message string) {
	log.Println("Some http error is happening")
	w.WriteHeader(status)
	var err error
	if message != "" {
		_, err = fmt.Fprint(w, message)
	} else if status == http.StatusNotFound {
		log.Printf("Returning 404 Not found error, the path '%s' was requested\n", r.URL.Path)
		_, err = fmt.Fprint(w, "Error 404, page *Not Found*")
	} else if status == http.StatusBadRequest {
		log.Printf("Returning 400 Bad request error, the path '%s' was requested\n", r.URL.Path)
		_, err = fmt.Fprint(w, "Error 400, Bad request")
	}
	if err != nil {
		log.Fatalf("Something went wrong with the error: %s", err)
	}
}