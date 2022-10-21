package endpoint

import (
	"net/http"
)

type BasicEndpoint struct {
	Path     string
	Function func(w http.ResponseWriter, r *http.Request)
}

func (e BasicEndpoint) Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != e.Path {
		ErrorHandler(w, r, http.StatusNotFound, "")
		return
	}
	e.Function(w, r)
}

func (e BasicEndpoint) GetPath() string {
	return e.Path
}
