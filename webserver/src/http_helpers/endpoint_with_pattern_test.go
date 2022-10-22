package http_helpers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpointWithPatternGetPath(t *testing.T) {
	p := "/myPath/"
	ewp := EndpointWithPattern{BasePath: p}
	if ewp.GetPath() != p {
		t.Errorf("Expected %s but got %s", p, ewp.GetPath())
	}
}

func TestEndpointWithPatternGetMatchers(t *testing.T) {
	p := "/myPath/"
	ewp := EndpointWithPattern{BasePath: p, Pattern: "(?P<name>[A-Za-z0-9]+)"}
	ep, ereg := ewp.GetMatchers()
	if ep != p {
		t.Errorf("Expected %s but got %s", p, ep)
	}
	correctStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	correctStrWSlash := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789/"
	if !(ereg.MatchString(p + correctStr)) {
		t.Errorf("Expected %s to be a match for the regexp", correctStr)
	}
	if !(ereg.MatchString(p + correctStrWSlash)) {
		t.Errorf("Expected %s to be a match for the regexp", correctStrWSlash)
	}
	incorrectStr := "ñ_ñ"
	if ereg.MatchString(p + incorrectStr) {
		t.Errorf("Expected %s to not be a match for the regexp", incorrectStr)
	}
}

func TestEndpointWithPatternHandlerError(t *testing.T) {
	p := "/myPath/"
	ewp := EndpointWithPattern{
		BasePath: p,
		Pattern:  "(?P<name>[A-Za-z0-9]+)",
		BaseFunction: func(w http.ResponseWriter, r *http.Request) {
			t.Errorf("The endpoint base function was called but it shouldn't have been called")
		},
		PatternFunction: func(w http.ResponseWriter, r *http.Request, s string) {
			t.Errorf("The endpoint pattern function was called but it shouldn't have been called")
		},
	}

	paths := []string{
		"holi://holi.holi" + "/notMyPath",        // With a wrong path
		"holi://holi.holi" + "/notMyPath/",       // With a wrong path that has a slash
		"holi://holi.holi" + "/notMyPath/ASas12", // With a wrong path that has a correct pattern
		"holi://holi.holi" + p + "ñ_ñ",           // With a correct path that has an incorrect pattern
	}
	for _, path := range paths {
		rWrong := httptest.NewRequest("", path, nil)
		w := &FakeResponseWriter{}
		ewp.Handler(w, rWrong)
		if w.StatusCode != http.StatusNotFound {
			t.Errorf("Expected %v but got %v", http.StatusNotFound, w.StatusCode)
		}
	}
}

func TestEndpointWithPatternHandlerBaseFunction(t *testing.T) {
	p := "/myPath/"
	ewp := EndpointWithPattern{
		BasePath: p,
		Pattern:  "(?P<name>[A-Za-z0-9]+)",
		BaseFunction: func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
		PatternFunction: func(w http.ResponseWriter, r *http.Request, s string) {
			t.Errorf("The endpoint pattern function was called but it shouldn't have been called")
		},
	}

	rWrong := httptest.NewRequest("", "holi://holi.holi"+p, nil)
	w := &FakeResponseWriter{}
	ewp.Handler(w, rWrong)
	if w.StatusCode != http.StatusOK {
		t.Errorf("Expected %v but got %v", http.StatusOK, w.StatusCode)
	}
}

func TestEndpointWithPatternHandlerPatternFunction(t *testing.T) {
	p := "/myPath/"
	ewp := EndpointWithPattern{
		BasePath: p,
		Pattern:  "(?P<name>[A-Za-z0-9]+)",
		BaseFunction: func(w http.ResponseWriter, r *http.Request) {
			t.Errorf("The endpoint base function was called but it shouldn't have been called")
		},
		PatternFunction: func(w http.ResponseWriter, r *http.Request, s string) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(s))
		},
	}

	patternString := "ASas12"
	rWrong := httptest.NewRequest("", "holi://holi.holi"+p+patternString, nil)
	w := &FakeResponseWriter{}
	ewp.Handler(w, rWrong)
	if w.StatusCode != http.StatusOK {
		t.Errorf("Expected %v but got %v", http.StatusOK, w.StatusCode)
	}
	if w.Content != patternString {
		t.Errorf("Expected %v but got %v", patternString, w.Content)
	}
}
