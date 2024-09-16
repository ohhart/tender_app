package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ohhart/tender-restapi/migrations"
)

var DB *sqlx.DB

func InitDB() (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"))

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database ping failed: %v", err)
		return nil, err
	}

	log.Println("Connected to the database successfully!")
	DB = db

	migrationsDir := "./migrations"
	if err := migrations.RunMigrations(DB, migrationsDir); err != nil {
		log.Fatalf("Error applying migrations: %v", err)
		return nil, err
	}

	return db, nil
}
