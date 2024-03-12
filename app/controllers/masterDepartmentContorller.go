package controllers

import (
	"strconv"
	"time"

	"github.com/cocoasterr/kalbe_test/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type MasterDepartmentController struct {
	Service services.MasterDepartmentService
}

func NewMasterDepartmentController(service services.MasterDepartmentService) *MasterDepartmentController {
	return &MasterDepartmentController{
		Service: service,
	}
}

func (c *MasterDepartmentController) CreateController(ctx *fiber.Ctx) error {
	payload := make(map[string]interface{})
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	trx, err := c.Service.Repo.BeginTransaction()
	if err != nil {
		logrus.WithError(err).Error("Error beginning transaction")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	err = c.Service.CreateService(ctx.Context(), trx, payload)
	if err != nil {
		logrus.WithError(err).Error("Error creating department")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "department created successfully"})
}

func (c *MasterDepartmentController) IndexController(ctx *fiber.Ctx) error {
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
		logrus.WithError(err).Error("Error retrieving attendance data")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": data, "total": total})
}

func (c *MasterDepartmentController) FindByController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid department ID"})
	}

	data, err := c.Service.FindByService(ctx.Context(), "id", id)
	if err != nil {
		logrus.WithError(err).Error("Error retrieving department data by ID")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": data})
}

func (c *MasterDepartmentController) UpdateController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid department ID"})
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

	err = c.Service.UpdateService(ctx.Context(), trx, payload, id)
	if err != nil {
		logrus.WithError(err).Error("Error updating department")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "department updated successfully"})
}

func (c *MasterDepartmentController) DeleteController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid attendance ID"})
	}

	deletePayload := map[string]interface{}{
		"deleted_at": time.Now(),
	}

	err := c.Service.UpdateService(ctx.Context(), nil, deletePayload, id)
	if err != nil {
		logrus.WithError(err).Error("Error deleting attendance")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "department updated successfully"})
}
