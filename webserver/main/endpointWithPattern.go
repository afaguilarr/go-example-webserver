package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type endpointWithPattern struct{
	basePath        string
	pattern         string
	baseFunction    func(http.ResponseWriter, *http.Request)
	patternFunction func(http.ResponseWriter, *http.Request, string)
}

func (e endpointWithPattern) getMatchers() (string, *regexp.Regexp) {
	patternRegexp := fmt.Sprintf(`^%s%s/?$`, e.basePath, e.pattern)
	return e.basePath, regexp.MustCompile(patternRegexp)
}

func (e endpointWithPattern) handler(w http.ResponseWriter, r *http.Request) {
	baseMatcher, patternMatcher := e.getMatchers()
	switch {
	case baseMatcher == r.URL.Path:
		log.Println("Path matched base regexp")
		e.baseFunction(w, r)
	case patternMatcher.MatchString(r.URL.Path):
		log.Println("Path matched its pattern regexp")
		e.patternFunction(w, r, patternMatcher.FindStringSubmatch(r.URL.Path)[1])
	default:
		log.Println("Path didn't match any regexp")
		errorHandler(w, r, http.StatusNotFound, "")
	}
}

func (e endpointWithPattern) getPath() string{
	return e.basePath
}
