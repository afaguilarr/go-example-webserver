package endpoint

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type EndpointWithPattern struct {
	BasePath        string
	Pattern         string
	BaseFunction    func(http.ResponseWriter, *http.Request)
	PatternFunction func(http.ResponseWriter, *http.Request, string)
}

func (e EndpointWithPattern) GetMatchers() (string, *regexp.Regexp) {
	patternRegexp := fmt.Sprintf(`^%s%s/?$`, e.BasePath, e.Pattern)
	return e.BasePath, regexp.MustCompile(patternRegexp)
}

func (e EndpointWithPattern) Handler(w http.ResponseWriter, r *http.Request) {
	baseMatcher, patternMatcher := e.GetMatchers()
	switch {
	case baseMatcher == r.URL.Path:
		log.Println("Path matched base regexp")
		e.BaseFunction(w, r)
	case patternMatcher.MatchString(r.URL.Path):
		log.Println("Path matched its pattern regexp")
		e.PatternFunction(w, r, patternMatcher.FindStringSubmatch(r.URL.Path)[1])
	default:
		log.Println("Path didn't match any regexp")
		ErrorHandler(w, r, http.StatusNotFound, "")
	}
}

func (e EndpointWithPattern) GetPath() string {
	return e.BasePath
}
