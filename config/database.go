package config

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload" // Autoload the .env file. We already check for it in main.go.
	"github.com/lakicsdomi/argus"
)

// Database initializes the MySQL connection and creates the necessary database and tables
func Database() *sql.DB {
	logger, _ := argus.Init("logs")

	logger.Verbose.Log("DATABASE", "Initializing database connection...")

	// Load user
	user, exist := os.LookupEnv("DB_USER")
	if !exist {
		err := errors.New("DB_USER environment variable not set in .env file")
		logger.Critical.LogErr("DATABASE", "Configuration missing", err)
		log.Fatal(err)
	}

	// Load password
	pass, exist := os.LookupEnv("DB_PASS")
	if !exist {
		err := errors.New("DB_PASS environment variable not set in .env file")
		logger.Critical.LogErr("DATABASE", "Configuration missing", err)
		log.Fatal(err)
	}

	// Load host
	host, exist := os.LookupEnv("DB_HOST")
	if !exist {
		err := errors.New("DB_HOST environment variable not set in .env file")
		logger.Critical.LogErr("DATABASE", "Configuration missing", err)
		log.Fatal(err)
	}

	logger.Verbose.Log("DATABASE", "Database configuration loaded successfully")

	// Construct the Data Source Name (DSN) for MySQL connection
	credentials := fmt.Sprintf("%s:%s@(%s:3306)/?charset=utf8&parseTime=True", user, pass, host) // 3306 is the default MySQL port

	// Open a connection to the MySQL db
	db, err := sql.Open("mysql", credentials)
	if err != nil {
		logger.Critical.LogErr("DATABASE", "Database connection failed", err)
		log.Fatal(err) // Fatal error, so the application should exit
	} else {
		logger.Verbose.Log("DATABASE", "Database Connection Successful!")
	}

	// Create gotodo database if it does not exist
	_, err = db.Exec(`CREATE DATABASE IF NOT EXISTS gotodo`)
	if err != nil {
		logger.Error.LogErr("DATABASE", "Failed to create database", err)
	}

	// Select gotodo database
	_, err = db.Exec(`USE gotodo`)
	if err != nil {
		logger.Error.LogErr("DATABASE", "Failed to select database", err)
	}

	// Create todos table if it does not exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
			id INT AUTO_INCREMENT,
			item TEXT NOT NULL,
			completed BOOLEAN DEFAULT FALSE,
			PRIMARY KEY (id)
		);
	`)
	if err != nil {
		logger.Error.LogErr("DATABASE", "Failed to create todos table", err)
	}

	logger.Verbose.Log("DATABASE", "Database initialized successfully.")
	return db
}
