package app

import (
	"Golang-Book-Author-API/service"
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
	"net/http"
)

type AuthorHandlers struct {
	service service.AuthorService
}

func (ah *AuthorHandlers) getAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, _ := ah.service.GetAllAuthors()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/json")
		xml.NewEncoder(w).Encode(authors)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(authors)
	}
}

func (ah *AuthorHandlers) getAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["author_id"]

	author, err := ah.service.GetAuthor(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, author)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
