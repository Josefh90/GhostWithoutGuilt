package webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var verifyToken = os.Getenv("CALLBACK_VERIFY_TOKEN")

func WebhookHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Webhook-Anfrage erhalten:", r.Method, r.URL.Path)
	if r.Method == http.MethodGet {
		// Meta Webhook-Verifizierung
		mode := r.URL.Query().Get("hub.mode")
		token := r.URL.Query().Get("hub.verify_token")
		challenge := r.URL.Query().Get("hub.challenge")

		if mode == "subscribe" && token == verifyToken {
			fmt.Fprint(w, challenge)
			log.Println("Webhook-Verifizierung erfolgreich")
			return
		}

		http.Error(w, "Unauthorized", http.StatusForbidden)
		log.Println("Webhook-Verifizierung fehlgeschlagen")
		return
	}

	if r.Method == http.MethodPost {
		// Eingehende WhatsApp-Nachricht
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "could not read body", http.StatusBadRequest)
			return
		}

		// Debug: Rohdaten anzeigen
		fmt.Println("📥 Eingehende Nachricht:")
		fmt.Println(string(body))

		// Optional: JSON parsen (wenn du strukturierte Daten willst)
		var payload map[string]interface{}
		if err := json.Unmarshal(body, &payload); err != nil {
			log.Println("JSON-Fehler:", err)
		} else {
			log.Printf("Nachricht empfangen von %+v\n", payload)
		}

		w.WriteHeader(http.StatusOK)
		return
	}

	// Andere Methoden nicht erlaubt
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
} 