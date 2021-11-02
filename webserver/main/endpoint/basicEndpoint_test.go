package endpoint

import (
	"net/http"
	"testing"
)

func TestBasicEndpointGetPath(t *testing.T) {
	p := "myPath"
	be := BasicEndpoint{
		Path:     p,
		Function: func(w http.ResponseWriter, r *http.Request) {},
	}
	if be.GetPath() != p {
		t.Errorf("Expected %s but got %s", p, be.GetPath())
	}
}

func TestBasicEndpointGetPathFail(t *testing.T) {
	p := "myPath"
	be := BasicEndpoint{
		Path:     p,
		Function: func(w http.ResponseWriter, r *http.Request) {},
	}
	if be.GetPath() == p {
		t.Errorf("Expected %s but got %s", p, be.GetPath())
	}
}
