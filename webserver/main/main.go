package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type world struct {
	Name string `json:"name"`
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello world called")
	switch r.Method {

	case http.MethodGet:
		log.Println("Hello world GET called")
		_, err := fmt.Fprint(w, "Hello world")
		if err != nil {
			log.Fatalf("Something went wrong with the 'Hello World': %s", err)
		}

	case http.MethodPost:
		log.Println("Hello world POST called")

		var wo world
		body, err := ioutil.ReadAll(r.Body)
		log.Printf("The request's body is\n%s\n", string(body))
		if err != nil {
			errorHandler(w, r, http.StatusBadRequest, "Error 400, couldn't parse world JSON")
			return
		}

		err = json.Unmarshal(body, &wo)
		log.Printf("The parsed entity is\n%s\n", wo)
		if err != nil {
			errorHandler(w, r, http.StatusBadRequest, "Error 400, couldn't parse world JSON")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(wo)
		log.Printf("The response's body is \n%s\n", string(jsonResp))
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		_, err = w.Write(jsonResp)
		if err != nil {
			log.Fatalf("Error happened writing the JSON response. Err: %s", err)
		}
		return
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

func main() {
	endpoints := []endpoint{
		basicEndpoint{path: "/", function: helloWorld},
		basicEndpoint{path: "/name", function: helloGenericName},
		endpointWithPattern{
			basePath:        "/name/",
			pattern:         "(?P<name>[A-Za-z0-9]+)",
			baseFunction:    helloGenericName,
			patternFunction: helloName,
		},
	}
	for _, e := range endpoints {
		handle(e)
	}
	log.Println("Server starting!")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Something went wrong with the webserver: %s", err)
	}
	asd_asd_asd := func() string {
		return "holi"
	}
	log.Println(asd_asd_asd)
}
