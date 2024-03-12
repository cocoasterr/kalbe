package controllers

import (
	"strconv"
	"time"

	"github.com/cocoasterr/kalbe_test/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type MasterPositionController struct {
	Service services.MasterPositionService
}

func NewMasterPositionController(service services.MasterPositionService) *MasterPositionController {
	return &MasterPositionController{
		Service: service,
	}
}

func (c *MasterPositionController) CreateController(ctx *fiber.Ctx) error {
	payload := make(map[string]interface{})
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	trx, err := c.Service.Repo.BeginTransaction()
	if err != nil {
		logrus.WithError(err).Error("Error beginning transaction")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	defer func() {
		if commitErr := c.Service.Repo.CommitTransaction(trx); commitErr != nil {
			logrus.WithError(commitErr).Errorf("Error committing transaction")
		}
		if err != nil {
			rollbackErr := c.Service.Repo.RollbackTransaction(trx)
			if rollbackErr != nil {
				logrus.WithError(rollbackErr).Error("Error rolling back transaction")
			}
		}
	}()
	err = c.Service.CreateService(ctx.Context(), trx, payload)
	if err != nil {
		logrus.WithError(err).Error("Error creating position")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Position created successfully"})
}

func (c *MasterPositionController) IndexController(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil || page <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid page parameter"})
	}

	limit, err := strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil || limit <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid limit parameter"})
	}

	data, total, err := c.Service.IndexService(ctx.Context(), page, limit)
	if err != nil {
		logrus.WithError(err).Error("Error retrieving position data")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": data, "total": total})
}

func (c *MasterPositionController) FindByController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid position ID"})
	}

	data, err := c.Service.FindByService(ctx.Context(), "id", id)
	if err != nil {
		logrus.WithError(err).Error("Error retrieving position data by ID")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": data})
}

func (c *MasterPositionController) UpdateController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid position ID"})
	}

	payload := make(map[string]interface{})
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	trx, err := c.Service.Repo.BeginTransaction()
	if err != nil {
		logrus.WithError(err).Error("Error beginning transaction")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	defer func() {
		if commitErr := c.Service.Repo.CommitTransaction(trx); commitErr != nil {
			logrus.WithError(commitErr).Errorf("Error committing transaction")
		}
		if err != nil {
			rollbackErr := c.Service.Repo.RollbackTransaction(trx)
			if rollbackErr != nil {
				logrus.WithError(rollbackErr).Error("Error rolling back transaction")
			}
		}
	}()

	err = c.Service.UpdateService(ctx.Context(), trx, payload, id)
	if err != nil {
		logrus.WithError(err).Error("Error updating position")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "position updated successfully"})
}

func (c *MasterPositionController) DeleteController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid position ID"})
	}

	deletePayload := map[string]interface{}{
		"deleted_at": time.Now(),
	}

	err := c.Service.UpdateService(ctx.Context(), nil, deletePayload, id)
	if err != nil {
		logrus.WithError(err).Error("Error deleting position")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "position deleted successfully"})
}
