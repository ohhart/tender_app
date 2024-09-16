package handlers

import (
	"strconv"

	"github.com/ohhart/tender-restapi/pkg/service"
	"github.com/ohhart/tender-restapi/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type ReviewHandler struct {
	reviewService *service.ReviewService
}

func NewReviewHandler(reviewService *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{reviewService: reviewService}
}

func GetReviewsForTender(reviewService *service.ReviewService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenderIDStr := c.Params("tenderId")
		tenderID, err := strconv.Atoi(tenderIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tender ID"})
		}

		tenderIDUint, err := utils.SafeIntToUint(tenderID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tender ID"})
		}

		authorUsername := c.Query("authorUsername")

		organizationIDStr := c.Query("organizationId")
		organizationID, err := strconv.Atoi(organizationIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid organization ID"})
		}

		organizationIDUint, err := utils.SafeIntToUint(organizationID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid organization ID"})
		}

		reviews, err := reviewService.GetReviewsForTender(tenderIDUint, authorUsername, organizationIDUint)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get reviews"})
		}

		return c.JSON(reviews)
	}
}
