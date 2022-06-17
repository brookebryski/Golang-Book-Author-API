package app

import (
	"Golang-Book-Author-API/service"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Author struct {
	Name       string `json:"full_name" xml:"name"`
	Birthplace string `json:"birthplace" xml:"city"`
	Movement   string `json:"movement" xml:"zipcode"`
}

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
