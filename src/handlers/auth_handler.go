package handlers

import (
	"strconv"

	"github.com/Mohamed-Kalandar-Sulaiman/Identity_Service/src/repository"
	utils "github.com/Mohamed-Kalandar-Sulaiman/Identity_Service/src/utilities"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
    type Request struct {
        Email    string `json:"email"`
        Password string `json:"password"`
        Context  string `json:"context"`
    }

    var req Request
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error hashing password"})
    }

    if err := repository.CreateUser(req.Email, string(hashedPassword), req.Context); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func Login(c *fiber.Ctx) error {
    type Request struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    var req Request
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
    }

    user, err := repository.GetUserByEmail(req.Email)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    token, err := utils.GenerateJWT(user.Email)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
    }

    return c.JSON(fiber.Map{"token": token})
}

func GetUserContext(c *fiber.Ctx) error {
    userID, err := strconv.ParseUint(c.Params("id"), 10, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }

    user, err := repository.GetUserContext(uint(userID))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User context not found"})
    }

    return c.JSON(fiber.Map{"context": user.Context})
}

func GetApplicationContext(c *fiber.Ctx) error {
    appName := c.Params("appName")

    app, err := repository.GetApplicationContext(appName)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Application context not found"})
    }

    return c.JSON(fiber.Map{"context": app.Context})
}

func GetPublicKey(c *fiber.Ctx) error {
    appName := c.Params("appName")

    publicKey, err := repository.GetPublicKey(appName)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Public key not found"})
    }

    return c.JSON(fiber.Map{"public_key": publicKey})
}
