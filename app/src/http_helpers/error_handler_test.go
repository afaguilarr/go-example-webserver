package http_helpers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
