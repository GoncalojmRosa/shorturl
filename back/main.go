package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/goncalojmrosa/shorturl/database"
	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World!")
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/hello", helloHandler).Methods("GET")

  http.Handle("/", r)

  db := database.Connect

  log.Fatal(http.ListenAndServe(":8080", nil))
}


