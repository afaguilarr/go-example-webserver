package main

import (
	"encoding/json"
	"fmt"
	"goWebServer/src/dao"
	"goWebServer/src/http_helpers"
	"goWebServer/src/service"
	"io"
	"log"
	"net/http"

	_ "github.com/lib/pq"
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
		body, err := io.ReadAll(r.Body)
		log.Printf("The request's body is\n%s\n", string(body))
		if err != nil {
			http_helpers.ErrorHandler(w, r, http.StatusBadRequest, "Error 400, couldn't parse world JSON")
			return
		}

		err = json.Unmarshal(body, &wo)
		log.Printf("The parsed entity is\n%s\n", wo)
		if err != nil {
			http_helpers.ErrorHandler(w, r, http.StatusBadRequest, "Error 400, couldn't parse world JSON")
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

func main() {
	db := dao.CreateDBConnection()
	defer db.Close()

	hnHandler := service.NewHelloNameHandler(db)
	endpoints := []http_helpers.Endpoint{
		http_helpers.BasicEndpoint{Path: "/", Function: helloWorld},
		http_helpers.BasicEndpoint{Path: "/name", Function: hnHandler.HelloGenericName},
		http_helpers.EndpointWithPattern{
			BasePath:        "/name/",
			Pattern:         "(?P<name>[A-Za-z0-9]+)",
			BaseFunction:    hnHandler.HelloGenericName,
			PatternFunction: hnHandler.HelloName,
		},
	}
	for _, e := range endpoints {
		http_helpers.Handle(e, http.HandleFunc)
	}
	log.Println("Server starting!")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Something went wrong with the webserver: %s", err)
	}
}
