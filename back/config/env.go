package config

import (
	"os"

	t "github.com/goncalojmrosa/shorturl/types"
	"github.com/joho/godotenv"
)

var Envs = initConfig()

func initConfig() t.Config {
	godotenv.Load()
	return t.Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		Database_URI: getEnv("DATABASE_URL", "mongodb://localhost:27017"),
	}

}


func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}