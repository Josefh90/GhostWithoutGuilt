package main

import (
	"fmt"
	"log"

	rod "github.com/Josefh90/GhostWithoutGuilt/internal/rod"
	"github.com/Josefh90/GhostWithoutGuilt/internal/routes"
	webhook "github.com/Josefh90/GhostWithoutGuilt/internal/webhooks"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	// This runs before main()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	/* 	accessToken := os.Getenv("WHATSAPP_ACCESS_TOKEN")

	   	if accessToken == "" {
	   		log.Fatal("WHATSAPP_ACCESS_TOKEN environment variable not set")
	   	}
	   	phoneNumberID := os.Getenv("WHATSAPP_PHONE_NUMBER_ID")
	   	if phoneNumberID == "" {
	   		log.Fatal("WHATSAPP_PHONE_NUMBER_ID environment variable not set")
	   	}
	   	// Optional: My phone number for sending messages
	   	myPhoneNumber := os.Getenv("My_PHONE_NUMBER")
	   	if myPhoneNumber == "" {
	   		log.Fatal("My_PHONE_NUMBER environment variable not set")
	   	} */

	//client := whatsapp.NewClient(accessToken, phoneNumberID)

	// Fiber App initialisieren
	rod.Launch()

	app := fiber.New()

	// Middleware: Logging
	app.Use(logger.New())
	routes.SetupRoutes(app)
	//http.HandleFunc("/webhook", webhook.WebhookHandler)

	// Webhook-Route mit Fiber
	app.Post("/webhook", func(c *fiber.Ctx) error {
		return webhook.WebhookHandler(c)
	})

	fmt.Println("Server läuft auf http://localhost:6969")
	log.Fatal(app.Listen(":6969"))
}
