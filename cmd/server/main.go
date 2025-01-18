package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"os"

	api "github.com/Asterki/go-test/internal/api"

	log "github.com/sirupsen/logrus"
)

func load_env_vars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		requiredEnvVars := []string{"PORT"}
		for _, envVar := range requiredEnvVars {
			if os.Getenv(envVar) == "" {
				log.Fatalf("Environment variable %s is not set", envVar)
			}
		}
	}
}

func main() {
	log.Info("Loading environment variables...")
	load_env_vars()

	log.Info("Starting the application...")

	http.Handle("/", api.Router())

	// Get the port from the environment variables
	port := os.Getenv("PORT")
	log.Infof("Listening on port %s...", port)
  http.ListenAndServe(":"+port, nil)
}
