package routes

import "github.com/gofiber/fiber/v2"

// SetupRoutes registriert alle Routen-Gruppen
func SetupRoutes(app *fiber.App) {
	// Beispiel: User-Routen einbinden
	SetupMessageRoutes(app)

	// Weitere Routengruppen hier hinzufügen, z.B.
	// SetupProductRoutes(app)
}
