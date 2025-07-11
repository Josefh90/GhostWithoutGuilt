package routes

import (
	"github.com/Josefh90/GhostWithoutGuilt/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes legt alle User-bezogenen Routen an
func SetupMessageRoutes(app *fiber.App) {
	messageGroup := app.Group("/messages")

	//userGroup.Get("/", handlers.GetAllUsers)
	messageGroup.Post("/send", handlers.SendMessage)
	//userGroup.Get("/:id", handlers.GetUserByID)
	//userGroup.Put("/:id", handlers.UpdateUser)
	//	messageGroup.Delete("/:id", handlers.DeleteUser)
}
