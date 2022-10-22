package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"goWebServer/src/http_helpers"
	"io"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "postgres"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "hello_world"
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

func helloGenericName(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello generic name called")
	helloName(w, r, "Name")
}

func helloName(w http.ResponseWriter, r *http.Request, name string) {
	log.Printf("Hello name called with name %s\n", name)
	if name == "personaFea" {
		http_helpers.ErrorHandler(w, r, http.StatusNotFound, "Error 404, name personaFea is not available")
		return
	}
	sqlStatement := `INSERT INTO hello_world (hw_text) VALUES ($1) RETURNING id, hw_text`
	id, hw_text := 0, ""
	err := db.QueryRow(sqlStatement, name).Scan(&id, &hw_text)
	if err != nil {
		panic(err)
	}
	log.Printf("New hello_world record with ID '%v' and text '%s' was inserted\n", id, hw_text)
	_, err = fmt.Fprintf(w, "Hello %s", name)
	if err != nil {
		log.Fatalf("Something went wrong with the 'Hello Name' for name %s: %s", name, err)
	}
}

func createDBConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("The connection to the DB was successful!")
	return db
}

func main() {
	db = createDBConnection()
	defer db.Close()
	endpoints := []http_helpers.Endpoint{
		http_helpers.BasicEndpoint{Path: "/", Function: helloWorld},
		http_helpers.BasicEndpoint{Path: "/name", Function: helloGenericName},
		http_helpers.EndpointWithPattern{
			BasePath:        "/name/",
			Pattern:         "(?P<name>[A-Za-z0-9]+)",
			BaseFunction:    helloGenericName,
			PatternFunction: helloName,
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
