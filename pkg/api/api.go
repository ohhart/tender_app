package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ohhart/tender-restapi/pkg/handlers"
	"github.com/ohhart/tender-restapi/pkg/service"
)

// SetupRoutes initializes API routes
func SetupRoutes(app *fiber.App, tenderService *service.TenderService, bidService *service.BidService, reviewService *service.ReviewService) {
	app.Get("/api/ping", Ping)

	// Tenders Routes
	app.Get("/api/tenders", handlers.GetTenders(tenderService))
	app.Post("/api/tenders/new", handlers.CreateTender(tenderService))
	app.Get("/api/tenders/my", handlers.GetUserTenders(tenderService))
	app.Get("/api/tenders/:tenderId/status", handlers.GetTenderStatus(tenderService))
	app.Put("/api/tenders/:tenderId/status", handlers.UpdateTenderStatus(tenderService))
	app.Patch("/api/tenders/:tenderId/edit", handlers.EditTender(tenderService))
	app.Put("/api/tenders/:tenderId/rollback/:version", handlers.RollbackTenderVersion(tenderService))

	// Bids Routes
	app.Post("/api/bids/new", handlers.CreateBid(bidService))
	app.Get("/api/bids/my", handlers.GetUserBids(bidService))
	app.Get("/api/bids/:tenderId/list", handlers.GetBidsForTender(bidService))
	app.Get("/api/bids/:bidId/status", handlers.GetBidStatus(bidService))
	app.Put("/api/bids/:bidId/status", handlers.UpdateBidStatus(bidService))
	app.Patch("/api/bids/:bidId/edit", handlers.EditBid(bidService))
	app.Put("/api/bids/:bidId/submit_decision", handlers.SubmitBidDecision(bidService))
	app.Put("/api/bids/:bidId/feedback", handlers.SubmitBidFeedback(bidService))
	app.Put("/api/bids/:bidId/rollback/:version", handlers.RollbackBidVersion(bidService))
	app.Get("/api/bids/:tenderId/reviews", handlers.GetReviewsForTender(reviewService))
}

func Ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello, I'm OK!", "data": nil})
}
