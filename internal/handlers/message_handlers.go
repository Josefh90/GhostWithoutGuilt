package handlers

import (
	"log"
	"os"

	"github.com/Josefh90/GhostWithoutGuilt/internal/whatsapp"
	"github.com/gofiber/fiber/v2"
)

type MessageSendRequest struct {
	Message string `json:"message" validate:"required"`
}

func SendMessage(c *fiber.Ctx) error {
	var req MessageSendRequest

	if err := c.BodyParser(&req); err != nil {
		log.Printf("[SendMessage] JSON Parse Error: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid JSON",
			"details": err.Error(),
		})
	}

	log.Printf("[SendMessage] Request: %+v", req)

	accessToken := os.Getenv("WHATSAPP_ACCESS_TOKEN")
	if accessToken == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "WHATSAPP_ACCESS_TOKEN environment variable not set",
		})
	}

	phoneNumberID := os.Getenv("WHATSAPP_PHONE_NUMBER_ID")
	if phoneNumberID == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "WHATSAPP_PHONE_NUMBER_ID environment variable not set",
		})
	}

	myPhoneNumber := os.Getenv("My_PHONE_NUMBER")
	if myPhoneNumber == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "My_PHONE_NUMBER environment variable not set",
		})
	}

	client := whatsapp.NewClient(accessToken, phoneNumberID)
	err := client.SendTextMessage(myPhoneNumber, req.Message)
	if err != nil {
		log.Printf("[SendMessage] Error sending message: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to send message",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Message sent successfully",
	})
}
