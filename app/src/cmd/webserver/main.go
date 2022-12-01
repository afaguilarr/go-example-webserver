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
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/afaguilarr/go-example-webserver/app/src/cmd"
	"github.com/afaguilarr/go-example-webserver/app/src/crypto_client"
	"github.com/afaguilarr/go-example-webserver/app/src/http_helpers"
	"github.com/afaguilarr/go-example-webserver/proto"
	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

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
		http_helpers.ErrorHandler(w, r, http.StatusMethodNotAllowed, "")
	}
}

// MethodNotAllowed replies to the request with an HTTP 405 method not allowed.
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
}

// MethodNotAllowedHandler returns a simple request handler
// that replies to each request with a “405 method not allowed” reply.
func MethodNotAllowedHandler() http.Handler { return http.HandlerFunc(MethodNotAllowed) }

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("there was an error loading the env variables: %s", err.Error())
	}

	r := mux.NewRouter()
	// Handler for non existent routes
	r.NotFoundHandler = http.NotFoundHandler()
	// Handler for any Method Not Allowed error
	r.MethodNotAllowedHandler = MethodNotAllowedHandler()
	r.HandleFunc("/login", testRPC)
	logOutRoute := r.PathPrefix("/logout").Subrouter()
	logOutRoute.HandleFunc("", uh.RefreshToken)
	logOutRoute.Use(uh.MiddlewareValidateRefreshToken)

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
