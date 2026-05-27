package main

import (
	"errors"
	"fmt"
	"go-todo/config"
	"go-todo/controllers"
	"go-todo/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/lakicsdomi/argus"
	"github.com/lakicsdomi/argus/dashboard"
)

func main() {

	// Initialize the custom Argus logger
	logger, err := argus.Init("logs")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	// Ensure all files are closed when the app shuts down
	defer logger.CloseAll()

	go dashboard.Serve(logger.Directory, ":9090") // Start the dashboard in a separate goroutine

	logger.Verbose.Log("MAIN", "Starting the Go To-Do application...")

	if err := godotenv.Load(); err != nil { // We could use autoload, but this way we can log the error if the .env file is missing
		// In the database package, we use autoload instead, so the error is only logged
		logger.Error.LogErr("MAIN", "no .env file found", err)
	}

	port, exist := os.LookupEnv("PORT")
	if !exist {
		err := errors.New("PORT not set in .env file")
		logger.Error.LogErr("MAIN", "Environment variable missing", err)
	}
	logger.Verbose.Log("MAIN", "Loaded environment variables from .env file")

	db := config.Database()
	controllers.SetDatabase(db) // DI the database connection into the controllers package

	logger.Verbose.Log("MAIN", fmt.Sprintf("Starting server on port %s...", port))
	err = http.ListenAndServe(":"+port, routes.Init())

	if err != nil {
		logger.Critical.LogErr("MAIN", "Error starting server", err)
	}

	logger.Verbose.Log("MAIN", "Server started successfully on port "+port)
}
