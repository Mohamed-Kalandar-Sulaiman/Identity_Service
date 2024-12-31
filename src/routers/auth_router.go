package routers

import (
	"github.com/Mohamed-Kalandar-Sulaiman/Identity_Service/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
    authGroup := app.Group("/auth")

    authGroup.Post("/register", handlers.Register)
    authGroup.Post("/login", handlers.Login)

    app.Get("/user/:id/context", handlers.GetUserContext)
    app.Get("/application/:appName/context", handlers.GetApplicationContext)
    app.Get("/application/:appName/public-key", handlers.GetPublicKey)
}
