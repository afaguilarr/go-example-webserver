package services

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/afaguilarr/go-example-webserver/app/src/dao"
	"github.com/afaguilarr/go-example-webserver/app/src/http_helpers"
	"github.com/gorilla/mux"
)

type HelloNameHandler struct {
	DaoNames *dao.DaoNames
}

func NewHelloNameHandler(db *sql.DB) HelloNameHandler {
	dn := dao.DaoNames{
		DB: db,
	}
	return HelloNameHandler{
		DaoNames: &dn,
	}
}

func (h *HelloNameHandler) HelloGenericName(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello generic name called")
	vars := map[string]string{}
	r = mux.SetURLVars(r, vars)
	h.HelloName(w, r)
}

func (h *HelloNameHandler) HelloName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	log.Printf("Hello name called with name %s\n", name)
	if name == "personaFea" {
		http_helpers.ErrorHandler(w, r, http.StatusNotFound, "Error 404, name personaFea is not available")
		return
	}

	_, name, err := h.DaoNames.InsertName(name)
	if err != nil {
		log.Fatalf("Something went wrong with the 'Hello Name' for name %s: %s", name, err)
		http_helpers.ErrorHandler(w, r, http.StatusInternalServerError, "Error 500, an error occurred while inserting the name")
		return
	}

	_, err = fmt.Fprintf(w, "Hello %s", name)
	if err != nil {
		log.Fatalf("Something went wrong with the 'Hello Name' for name %s: %s", name, err)
	}
}