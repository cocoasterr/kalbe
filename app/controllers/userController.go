package controllers

import (
	"fmt"
	"time"

	"github.com/cocoasterr/kalbe_test/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"github.com/wpcodevo/golang-fiber-jwt/initializers"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	Service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		Service: service,
	}
}

func (c *UserController) RegisterUser(ctx *fiber.Ctx) error {
	payload := make(map[string]interface{})
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	password, ok := payload["password"].(string)
	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password is required"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithError(err).Error("Error hashing password")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	payload["password"] = string(hashedPassword)

	trx, err := c.Service.Repo.BeginTransaction()
	if err != nil {
		logrus.WithError(err).Error("Error beginning transaction")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	err = c.Service.CreateService(ctx.Context(), trx, payload)
	if err != nil {
		logrus.WithError(err).Error("Error creating User")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	config, _ := initializers.LoadConfig(".")
	secretKey := []byte(config.JwtSecret)

	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var body request
	if err := ctx.BodyParser(&body); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return err
	}

	query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", body.Username)
	userData, _ := c.Service.FindCustomQueryService(ctx.Context(), query)

	hashedPassword := userData["password"].(string)

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(body.Password)); err != nil {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid username or password",
		})
		return nil
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userData["user_id"]
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	s, err := token.SignedString([]byte(secretKey))
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"token": s, "message": "login successfully"})
}
