package http_helpers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type FakeResponseWriter struct {
	StatusCode int
	Content    string
}

func (frw *FakeResponseWriter) Header() http.Header {
	return http.Header{}
}

func (frw *FakeResponseWriter) Write(ba []byte) (int, error) {
	frw.Content = string(ba)
	return len(ba), nil
}

func (frw *FakeResponseWriter) WriteHeader(statusCode int) {
	frw.StatusCode = statusCode
}

func TestBasicEndpointGetPath(t *testing.T) {
	p := "/myPath"
	be := BasicEndpoint{Path: p}
	if be.GetPath() != p {
		t.Errorf("Expected %s but got %s", p, be.GetPath())
	}
}

func TestBasicEndpointHandlerError(t *testing.T) {
	p := "/myPath"
	be := BasicEndpoint{
		Path: p,
		Function: func(w http.ResponseWriter, r *http.Request) {
			t.Errorf("The endpoint function was called but it shouldn't have been called")
		},
	}
	r := httptest.NewRequest("", "holi://holi.holi"+"/notMyPath", nil)
	w := &FakeResponseWriter{}
	be.Handler(w, r)
	if w.StatusCode != http.StatusNotFound {
		t.Errorf("Expected %v but got %v", http.StatusNotFound, w.StatusCode)
	}
}

func TestBasicEndpointHandler(t *testing.T) {
	p := "/myPath"
	be := BasicEndpoint{
		Path: p,
		Function: func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	}
	r := httptest.NewRequest("", "holi://holi.holi"+p, nil)
	w := &FakeResponseWriter{}
	be.Handler(w, r)
	if w.StatusCode != http.StatusOK {
		t.Errorf("Expected %v but got %v", http.StatusNotFound, w.StatusCode)
	}
}
