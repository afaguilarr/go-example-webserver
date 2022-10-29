package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/afaguilarr/go-example-webserver/app/src/dao"
	"github.com/afaguilarr/go-example-webserver/app/src/http_helpers"
	"github.com/afaguilarr/go-example-webserver/app/src/service"
	"github.com/afaguilarr/go-example-webserver/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gorilla/mux"
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

func testRPC(w http.ResponseWriter, r *http.Request) {
	log.Println("Test RPC called")
	switch r.Method {
	case http.MethodGet:
		addr := "crypto:8080"
		// Set up a connection to the server.
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := proto.NewCryptoClient(conn)
		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp, err := c.Encrypt(ctx, &proto.EncryptRequest{})
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
	db := dao.CreateDBConnection()
	defer db.Close()

	hnHandler := service.NewHelloNameHandler(db)

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
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Something went wrong with the webserver: %s", err)
	}
}
