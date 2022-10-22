package http_helpers

import (
	"net/http"
	"testing"
)

type FakeEndpoint struct {
	getPathCalled bool
}

func (fe *FakeEndpoint) Handler(w http.ResponseWriter, r *http.Request) {}

func (fe *FakeEndpoint) GetPath() string {
	fe.getPathCalled = true
	return "true"
}

func TestEndpointHandleIsCallable(t *testing.T) {
	fe := &FakeEndpoint{}
	Handle(fe, http.HandleFunc)
	if !(fe.getPathCalled) {
		t.Errorf("Expected the GetPath function to be called but it wasn't")
	}
}
