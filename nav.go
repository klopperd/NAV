package main

import (
	"fmt"
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))

	port := ":3000"
	fmt.Println("Listening on localhost" + port)

	log.Fatal(http.ListenAndServe(port, router))
}
