package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ohhart/tender-restapi/pkg/api"
	"github.com/ohhart/tender-restapi/pkg/repository"
	"github.com/ohhart/tender-restapi/pkg/service"
)

func connectDB() (*sqlx.DB, error) {
	connStr := os.Getenv("POSTGRES_CONN")
	if connStr == "" {
		log.Fatalf("POSTGRES_CONN is not set")
	}
	log.Printf("Connecting to database with connection string: %s", connStr)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	db, err := connectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	tenderRepo := repository.NewTenderRepository(db)
	bidRepo := repository.NewBidRepository(db)
	reviewRepo := repository.NewReviewRepository(db)

	tenderService := service.NewTenderService(tenderRepo)
	bidService := service.NewBidService(bidRepo, reviewRepo)
	reviewService := service.NewReviewService(*reviewRepo, *bidRepo)

	app := fiber.New()

	api.SetupRoutes(app, tenderService, bidService, reviewService)

	port := os.Getenv("SERVER_ADDRESS")
	if port == "" {
		port = ":8080"
	}

	log.Printf("Server is starting on %s", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
