package main

import (
	"net/http"
)

type basicEndpoint struct{
	path     string
	function func(w http.ResponseWriter, r *http.Request)
}

func (e basicEndpoint) handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != e.path {
		errorHandler(w, r, http.StatusNotFound, "")
		return
	}
	e.function(w, r)
}

func (e basicEndpoint) getPath() string{
	return e.path
}
