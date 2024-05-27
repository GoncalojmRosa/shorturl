package api

import (
	"log"
	"net/http"

	"github.com/goncalojmrosa/shorturl/services/sites"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIServer struct {
	addr string
	db	 *mongo.Client
}

func NewAPIServer(addr string, db *mongo.Client) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	log.Println("Starting server on", s.addr)

	siteStore := sites.NewStore(s.db)

	siteHandler := sites.NewHandler(siteStore)
	siteHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr, subrouter)
}