package webhook

//https://developers.facebook.com/apps/1030673848841803/whatsapp-business/wa-settings/?business_id=2467686813608135&phone_number_id=

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

var verifyToken = os.Getenv("CALLBACK_VERIFY_TOKEN")

func WebhookHandler(c *fiber.Ctx) error {
	log.Println("Webhook-Anfrage erhalten:", c.Method(), c.OriginalURL())

	if c.Method() == fiber.MethodGet {
		// Meta Webhook-Verifizierung
		mode := c.Query("hub.mode")
		token := c.Query("hub.verify_token")
		challenge := c.Query("hub.challenge")

		if mode == "subscribe" && token == "DEIN_VERIFY_TOKEN" {
			log.Println("Webhook verifiziert")
			return c.SendString(challenge)
		} else {
			log.Println("Webhook-Verifizierung fehlgeschlagen")
			return c.SendStatus(fiber.StatusForbidden)
		}
	}

	// POST oder andere Methoden: hier kannst du JSON oder andere Daten verarbeiten
	log.Println("Webhook-POST-Verarbeitung (nicht implementiert)")
	return c.SendStatus(fiber.StatusOK)
}
