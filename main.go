package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello world called")
	_, err := fmt.Fprint(w, "Hello world")
	if err != nil {
		log.Fatalf("Something went wrong with the 'Hello World': %s", err)
	}
}

func helloGenericName(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello generic name called")
	helloName(w, r, "Name")
}

func helloName(w http.ResponseWriter, r *http.Request, name string) {
	if name == "personaFea" {
		errorHandler(w, r, http.StatusNotFound, "Error 404, name personaFea is not available")
		return
	}
	log.Printf("Hello name called with name %s\n", name)
	_, err := fmt.Fprintf(w, "Hello %s", name)
	if err != nil {
		log.Fatalf("Something went wrong with the 'Hello Name' for name %s: %s", name, err)
	}
}

func main(){
	endpoints := []endpoint{
		basicEndpoint{path: "/", function: helloWorld},
		basicEndpoint{path: "/name", function: helloGenericName},
		endpointWithPattern{
			basePath: "/name/",
			pattern: "(?P<name>[A-Za-z0-9]+)",
			baseFunction: helloGenericName,
			patternFunction: helloName,
		},
	}
	for _, e := range endpoints {
		handle(e)
	}
	log.Println("Server starting")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Something went wrong with the webserver: %s", err)
	}
}
