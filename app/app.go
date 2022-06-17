package app

import (
	"Golang-Book-Author-API/domain"
	"Golang-Book-Author-API/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	// ah := AuthorHandlers{service.NewAuthorService(domain.NewAuthorRepositoryStub())}
	ah := AuthorHandlers{service.NewAuthorService(domain.NewAuthorRepositoryDb())}
	// define routes
	router.HandleFunc("/authors", ah.getAllAuthors).Methods(http.MethodGet)

	// starting server
	log.Fatalln(http.ListenAndServe("localhost:8080", router))
}
