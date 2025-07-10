package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	webhook "github.com/Josefh90/GhostWithoutGuilt/internal/webhooks"
	"github.com/Josefh90/GhostWithoutGuilt/internal/whatsapp"
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
	accessToken := os.Getenv("WHATSAPP_ACCESS_TOKEN")

	fmt.Println("Access Token:", accessToken)


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
	}


	client := whatsapp.NewClient(accessToken, phoneNumberID)


	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		err := client.SendTextMessage(myPhoneNumber, "Hallo ich bin die Super Josi KI!")

		if err != nil {
			
			http.Error(w, "Fehler beim Senden der Nachricht: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "Nachricht erfolgreich gesendet!")
	})


	http.HandleFunc("/webhook", webhook.WebhookHandler)
	fmt.Println("Server läuft auf http://localhost:6969")
	log.Fatal(http.ListenAndServe(":6969", nil))
}
