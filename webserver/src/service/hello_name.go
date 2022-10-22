package service

import (
	"database/sql"
	"fmt"
	"goWebServer/src/dao"
	"goWebServer/src/http_helpers"
	"log"
	"net/http"
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
	h.HelloName(w, r, "Name")
}

func (h *HelloNameHandler) HelloName(w http.ResponseWriter, r *http.Request, name string) {
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
	return
}
