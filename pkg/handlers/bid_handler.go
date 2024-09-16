package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ohhart/tender-restapi/models"
	"github.com/ohhart/tender-restapi/pkg/service"
	"github.com/ohhart/tender-restapi/pkg/utils"
)

func CreateBid(bidService *service.BidService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var bid models.Bid
		if err := c.BodyParser(&bid); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		authorID, err := strconv.Atoi(c.FormValue("authorId"))
		if err != nil {
			authorID, err = strconv.Atoi(c.FormValue("author_id"))
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid author ID"})
			}
		}
		authorIDUint, err := utils.SafeIntToUint(authorID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid author ID"})
		}
		bid.AuthorID = authorIDUint

		tenderID, err := strconv.Atoi(c.FormValue("tenderId"))
		if err != nil {
			tenderID, err = strconv.Atoi(c.FormValue("tender_id"))
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tender ID"})
			}
		}
		tenderIDUint, err := utils.SafeIntToUint(tenderID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tender ID"})
		}
		bid.TenderID = tenderIDUint

		bid.Status = "CREATED"

		if err := bidService.CreateBid(bid); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create bid"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Bid created successfully"})
	}
}

func GetUserBids(bidService *service.BidService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userIDStr := c.Query("userId")
		if userIDStr == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User ID is required"})
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}

		userIDUint, err := utils.SafeIntToUint(userID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}

		bids, err := bidService.ListBids(userIDUint)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve bids"})
		}

		return c.JSON(bids)
	}
}

func GetBidsForTender(bidService *service.BidService) fiber.Handler {
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

		bids, err := bidService.ListBids(tenderIDUint)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve bids"})
		}

		return c.JSON(bids)
	}
}

func GetBidStatus(bidService *service.BidService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bidIDStr := c.Params("bidId")
		bidID, err := strconv.Atoi(bidIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid bid ID"})
		}

		bidIDUint, err := utils.SafeIntToUint(bidID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid bid ID"})
		}

		bid, err := bidService.GetBid(bidIDUint)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Bid not found"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve bid"})
		}

		return c.JSON(bid)
	}
}
func UpdateBidStatus(bidService *service.BidService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var bid models.Bid
		if err := c.BodyParser(&bid); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		if err := bidService.UpdateBidStatus(bid); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update bid"})
		}

		return c.JSON(fiber.Map{"status": "success", "message": "Bid updated successfully"})
	}
}

func DeleteBid(bidService *service.BidService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bidIDStr := c.Params("bidId")
		bidID, err := strconv.Atoi(bidIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid bid ID"})
		}

		bidIDUint, err := utils.SafeIntToUint(bidID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid bid ID"})
		}

		if err := bidService.DeleteBid(bidIDUint); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete bid"})
		}

		return c.JSON(fiber.Map{"status": "success", "message": "Bid deleted successfully"})
	}
}

func SubmitBidFeedback(bidService *service.BidService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bidIDStr := c.Params("bidId")
		bidID, err := strconv.Atoi(bidIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid bid ID"})
		}

		bidIDUint, err := utils.SafeIntToUint(bidID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid bid ID"})
		}

		var feedback models.Review
		if err := c.BodyParser(&feedback); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		if err := bidService.SubmitBidFeedback(bidIDUint, feedback); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to submit bid feedback"})
		}

		return c.JSON(fiber.Map{"status": "success", "message": "Bid feedback submitted successfully"})
	}
}
func SubmitBidDecision(bidService *service.BidService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bidIDStr := c.Params("bidId")
		bidID, err := strconv.Atoi(bidIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid bid ID"})
		}

		bidIDUint, err := utils.SafeIntToUint(bidID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid bid ID"})
		}

		var decisionData struct {
			Decision string `json:"decision"`
		}

		if err := c.BodyParser(&decisionData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		if err := bidService.SubmitBidDecision(bidIDUint, decisionData.Decision); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to submit bid decision"})
		}

		return c.JSON(fiber.Map{"status": "success"})
	}
}

func RollbackBidVersion(bidService *service.BidService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bidIDStr := c.Params("bidId")
		bidID, err := strconv.Atoi(bidIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid bid ID"})
		}

		bidIDUint, err := utils.SafeIntToUint(bidID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid bid ID"})
		}

		versionStr := c.Params("version")
		version, err := strconv.Atoi(versionStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid version"})
		}

		if err := bidService.RollbackBidVersion(bidIDUint, version); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to rollback bid version"})
		}

		return c.JSON(fiber.Map{"status": "success", "message": "Bid version rolled back successfully"})
	}
}

func GetBidReviews(bidService *service.BidService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bidIDStr := c.Params("bidId")
		bidID, err := strconv.Atoi(bidIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid bid ID"})
		}

		bidIDUint, err := utils.SafeIntToUint(bidID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid bid ID"})
		}

		reviews, err := bidService.GetBidReviews(bidIDUint)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve bid reviews"})
		}

		return c.JSON(reviews)
	}
}

func EditBid(bidService *service.BidService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bidID, err := strconv.ParseUint(c.Params("bidId"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid bid ID"})
		}

		var bid models.Bid
		if err := c.BodyParser(&bid); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		bid.ID = bidID

		if err := bidService.EditBid(bid); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to edit bid"})
		}

		return c.JSON(fiber.Map{"status": "success", "message": "Bid edited successfully"})
	}
}
