package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/goncalojmrosa/shorturl/database"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Client

func allSitesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All sites")
}

func newSiteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New site")
	
}


func Start() {
	r := mux.NewRouter()
	db := database.Connect()
	//ROUTES ENDPOINTS
	r.HandleFunc("/sites", allSitesHandler).Methods("GET")

	r.HandleFunc("/sites", newSiteHandler).Methods("POST")


	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}