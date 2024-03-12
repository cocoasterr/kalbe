package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cocoasterr/kalbe_test/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AttendanceController struct {
	Service           services.AttendaceService
	ServiceAdditional services.AttendanceServiceAdditional
}

func NewAttendanceController(service services.AttendaceService) *AttendanceController {
	return &AttendanceController{
		Service: service,
	}
}

func (c *AttendanceController) CreateAttendanceController(ctx *fiber.Ctx) error {
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
		logrus.WithError(err).Error("Error creating Attendance")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Attendance created successfully"})
}

func (c *AttendanceController) IndexAttendanceController(ctx *fiber.Ctx) error {
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

func (c *AttendanceController) FindByAttendanceController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Attendance ID"})
	}

	data, err := c.Service.FindByService(ctx.Context(), "id", id)
	if err != nil {
		logrus.WithError(err).Error("Error retrieving Attendance data by ID")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": data})
}

func (c *AttendanceController) UpdateAttendanceController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Attendance ID"})
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
		logrus.WithError(err).Error("Error updating Attendance")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "attendance updated successfully"})
}

func (c *AttendanceController) DeleteAttendanceController(ctx *fiber.Ctx) error {
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

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "attendance updated successfully"})
}

func (c *AttendanceController) Reporting(ctx *fiber.Ctx) error {
	startDate := ctx.Query("start", time.Now().Format("2006-01-02"))
	endDate := ctx.Query("end", time.Now().Add(12*time.Hour).Format("2006-01-02"))

	query := fmt.Sprintf(`
    SELECT
        a.Created_at AS Date,
        e.Employee_code,
        e.Employee_name,
        d.Department_name,
        p.Position_name,
        l.Location_name,
        a.Absent_in,
        a.Absent_out
    FROM
        Attendance a
    JOIN
        Employee e ON a.Employee_id = e.Employee_id
    JOIN
        MasterDepartment d ON e.Department_id = d.Department_id
    JOIN
        MasterPosition p ON e.Position_id = p.Position_id
    JOIN
        MasterLocation l ON a.Location_id = l.Location_id
    WHERE
        a.Absent_in BETWEEN '%s' AND '%s'
    ORDER BY
        a.Created_at
    `, startDate, endDate)
	fmt.Println(query)
	data, err := c.Service.FindCustomQueryService(ctx.Context(), query)
	if err != nil {
		logrus.WithError(err).Error("Error retrieving attendance data")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": data})
}
