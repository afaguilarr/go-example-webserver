package http_helpers

import (
	"log"
	"net/http"
)

type Endpoint interface {
	GetPath() string
	Handler(w http.ResponseWriter, r *http.Request)
}

func Handle(e Endpoint, handleFunc func(string, func(http.ResponseWriter, *http.Request))) {
	log.Printf("Handling '%s' endpoint", e.GetPath())
	handleFunc(e.GetPath(), e.Handler)
}
