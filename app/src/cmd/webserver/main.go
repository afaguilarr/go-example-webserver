// Package webserver.
//
// This is the public API exposing the features from all other microservices composing this example application.
//
//		Schemes: http, https
//		Host: localhost
//		BasePath: /api
//		Version: 0.0.1
//		License: MIT http://opensource.org/licenses/MIT
//		Contact: Andres Felipe Aguilar Rendon<afaguilarr@unal.edu.co> https://github.com/afaguilarr/go-example-webserver
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
//	 SecurityDefinitions:
//		  bearerAuth:
//		    type: apiKey
//		    in: header
//		    name: Authorization
//		    bearerFormat: JWT # optional, arbitrary value for documentation purposes
//
// swagger:meta
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/afaguilarr/go-example-webserver/app/src/cmd"
	"github.com/afaguilarr/go-example-webserver/app/src/crypto_client"
	"github.com/afaguilarr/go-example-webserver/app/src/dao"
	"github.com/afaguilarr/go-example-webserver/app/src/http_helpers"
	"github.com/afaguilarr/go-example-webserver/app/src/services"
	"github.com/afaguilarr/go-example-webserver/proto"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	helloWorldDBPrefix = "HELLO_WORLD_"
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

func testRPC(w http.ResponseWriter, r *http.Request) {
	log.Println("Test RPC called")
	switch r.Method {
	case http.MethodGet:
		cc := crypto_client.NewCryptoClientHandler(cmd.CryptoHost, cmd.DefaultPort)
		err := cc.CreateConnection()
		if err != nil {
			log.Fatalf("while creating Crypto Connection: %s", err)
		}
		defer func() {
			err := cc.CloseConnection()
			if err != nil {
				log.Fatalf("while closing Crypto Connection: %s", err)
			}
		}()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp, err := cc.Encrypt(ctx, &proto.EncryptRequest{
			Context:          "jiji",
			UnencryptedValue: []byte("jojo"),
		})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		_, err = fmt.Fprintf(w, "Greeting: %s", resp.GetEncryptedValue())
		if err != nil {
			log.Fatalf("Something went wrong with the 'gRPC test endpoint': %s", err)
		}
	default:
		log.Fatalf("Unsupported Method %s", r.Method)
	}
}

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("there was an error loading the env variables: %s", err.Error())
	}

	dbch := dao.NewDBConnectionHandler(helloWorldDBPrefix)
	db, err := dbch.CreateDBConnection()
	if err != nil {
		panic(errors.Wrap(err, "while creating DB connection"))
	}
	defer db.Close()

	hnHandler := services.NewHelloNameHandler(db)

	r := mux.NewRouter()
	r.HandleFunc("/", helloWorld)
	r.HandleFunc("/name", hnHandler.HelloGenericName)
	r.HandleFunc("/name/{name}", hnHandler.HelloName)
	r.HandleFunc("/test/rpc", testRPC)

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}
	log.Println("Server starting!")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Something went wrong with the webserver: %s", err)
	}
}
