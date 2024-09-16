package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ohhart/tender-restapi/models"
	"github.com/ohhart/tender-restapi/pkg/service"
	"github.com/ohhart/tender-restapi/pkg/utils"
)

func GetTenders(tenderService *service.TenderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenders, err := tenderService.ListTenders()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to retrieve tenders",
			})
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Tenders retrieved successfully",
			"data":    tenders,
		})
	}
}

func CreateTender(tenderService *service.TenderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tender models.Tender
		if err := c.BodyParser(&tender); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request payload",
			})
		}

		organizationIDStr := c.FormValue("organizationId")
		if organizationIDStr == "" {
			organizationIDStr = c.FormValue("organization_id")
		}
		organizationID, err := strconv.Atoi(organizationIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid organization ID",
			})
		}
		organizationIDUint, err := utils.SafeIntToUint(organizationID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid organization ID",
			})
		}
		tender.OrganizationID = organizationIDUint

		creatorIDStr := c.FormValue("creatorId")
		if creatorIDStr == "" {
			creatorIDStr = c.FormValue("creator_id")
		}
		creatorID, err := strconv.Atoi(creatorIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid creator ID",
			})
		}
		creatorIDUint, err := utils.SafeIntToUint(creatorID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid creator ID",
			})
		}
		tender.CreatorID = creatorIDUint
		tender.Status = "CREATED"

		if err := tenderService.CreateTender(tender); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create tender",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":  "success",
			"message": "Tender created successfully",
		})
	}
}

func GetTenderStatus(tenderService *service.TenderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenderIDStr := c.Params("tenderId")
		tenderID, err := strconv.ParseUint(tenderIDStr, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid tender ID",
			})
		}

		tenderIDUint, err := utils.SafeUint64ToUint(tenderID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid tender ID",
			})
		}

		tender, err := tenderService.GetTender(tenderIDUint)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "Tender not found",
			})
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Tender retrieved successfully",
			"data":    tender,
		})
	}
}

func UpdateTenderStatus(tenderService *service.TenderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenderIDStr := c.Params("tenderId")
		tenderID, err := strconv.Atoi(tenderIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid tender ID",
			})
		}

		tenderIDUint, err := utils.SafeIntToUint(tenderID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid tender ID",
			})
		}

		var updateData struct {
			Status string `json:"status"`
		}

		if err := c.BodyParser(&updateData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request payload",
			})
		}

		if err := tenderService.UpdateTenderStatus(tenderIDUint, updateData.Status); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update tender status",
			})
		}

		return c.JSON(fiber.Map{
			"status": "success",
		})
	}
}

func EditTender(tenderService *service.TenderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tender models.Tender
		if err := c.BodyParser(&tender); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid request payload",
			})
		}

		tenderIDStr := c.Params("tenderId")
		tenderID, err := strconv.ParseUint(tenderIDStr, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid tender ID",
			})
		}

		tender.ID = tenderID

		if err := tenderService.EditTender(tender); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to update tender",
			})
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Tender updated successfully",
			"data":    tender,
		})
	}
}

func DeleteTender(tenderService *service.TenderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenderIDStr := c.Params("tenderId")
		tenderID, err := strconv.ParseUint(tenderIDStr, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid tender ID",
			})
		}

		tenderIDUint, err := utils.SafeUint64ToUint(tenderID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid tender ID",
			})
		}

		if err := tenderService.DeleteTender(tenderIDUint); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to delete tender",
			})
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Tender deleted successfully",
		})
	}
}

func RollbackTenderVersion(tenderService *service.TenderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenderIDStr := c.Params("tenderId")
		tenderID, err := strconv.ParseUint(tenderIDStr, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid tender ID",
			})
		}

		tenderIDUint, err := utils.SafeUint64ToUint(tenderID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid tender ID",
			})
		}

		versionStr := c.Params("version")
		version, err := strconv.Atoi(versionStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid version number",
			})
		}

		if err := tenderService.RollbackTenderVersion(tenderIDUint, version); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to rollback tender version",
			})
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Tender version rolled back successfully",
		})
	}
}

func GetUserTenders(tenderService *service.TenderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Query("username")
		if username == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "username is required",
			})
		}

		tenders, err := tenderService.GetTendersByUsername(username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve user tenders",
			})
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Tenders retrieved successfully",
			"data":    tenders,
		})
	}
}
