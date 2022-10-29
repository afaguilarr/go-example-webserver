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

func TestErrorHandlerCallableWithNoMessageAndNotFoundStatus(t *testing.T) {
	rWrong := httptest.NewRequest("", "holi://holi.holi", nil)
	w := &FakeResponseWriter{}
	ErrorHandler(w, rWrong, http.StatusNotFound, "")
	if w.StatusCode != http.StatusNotFound {
		t.Errorf("Expected %v but got %v", http.StatusNotFound, w.StatusCode)
	}
}

func TestErrorHandlerCallableWithNoMessageAndBadRequest(t *testing.T) {
	rWrong := httptest.NewRequest("", "holi://holi.holi", nil)
	w := &FakeResponseWriter{}
	ErrorHandler(w, rWrong, http.StatusBadRequest, "")
	if w.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected %v but got %v", http.StatusBadRequest, w.StatusCode)
	}
}

func TestErrorHandlerCallableWithNoMessageAndUnknownStatus(t *testing.T) {
	rWrong := httptest.NewRequest("", "holi://holi.holi", nil)
	w := &FakeResponseWriter{}
	ErrorHandler(w, rWrong, http.StatusAccepted, "")
	if w.StatusCode != http.StatusAccepted {
		t.Errorf("Expected %v but got %v", http.StatusAccepted, w.StatusCode)
	}
}

func TestErrorHandlerCallableWithAMessage(t *testing.T) {
	rWrong := httptest.NewRequest("", "holi://holi.holi", nil)
	w := &FakeResponseWriter{}
	ErrorHandler(w, rWrong, http.StatusAccepted, "ASas12")
	if w.StatusCode != http.StatusAccepted {
		t.Errorf("Expected %v but got %v", http.StatusAccepted, w.StatusCode)
	}
}
