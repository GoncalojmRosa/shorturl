package main

import (
	"fmt"
	"log"

	"github.com/goncalojmrosa/shorturl/api"
	"github.com/goncalojmrosa/shorturl/config"
	"github.com/goncalojmrosa/shorturl/db"
)

func main() {
	mongoClient, err := db.ConnectToMongo(config.Envs.Database_URI)

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(fmt.Sprintf(":%s", config.Envs.Port), mongoClient)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
