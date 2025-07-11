package rod

import (
	"fmt"
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

func Launch() {
	// Starte Chrome sichtbar
	url := launcher.New().
		Headless(false).
		MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://web.whatsapp.com")

	log.Println("🔄 Warte auf WhatsApp-Login...")
	page.MustWaitLoad()

	// Warte bis Seite geladen & "WhatsApp" sichtbar ist
	page.MustElementR("div", "Send and receive messages without keeping your phone online.")
	log.Println("✅ Eingeloggt in WhatsApp Web")

	for {
		log.Println("🔍 Suche ungelesene Chats...")

		// Suche englische Variante von ungelesenen Nachrichten
		unreadSpans, err := page.Elements("span[aria-label*='unread message']")
		if err != nil {
			log.Println("❌ Fehler beim Suchen:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if len(unreadSpans) == 0 || unreadSpans[0] == nil {
			log.Println("📭  Keine ungelesenen Nachrichten.")
		} else {
			log.Printf("📨  %d ungelesene Nachricht(en) gefunden.\n", len(unreadSpans))

			for _, span := range unreadSpans {
				if span != nil {
					err := span.Click(proto.InputMouseButtonLeft, 1)
					if err != nil {
						log.Println("❌ Fehler beim Klicken auf ungelesene Nachricht:", err)
						continue
					}
					log.Println("➡️ Auf ungelesene Nachricht geklickt.")
					//time.Sleep(2 * time.Second) // Chat laden lassen
					page.Timeout(5 * time.Second).MustWait(`span._ao3e`)

					// 🔍 Letzte Nachricht im Chat suchen

					// ❗️Hier ein Beispiel-Wait auf das Element mit dem Text "Orgel IJK"
					// Du kannst auch gezielt mit Klassen arbeiten
					selector := `span._ao3e` // Deine Zielklass
					// Warten bis das Element erscheint
					el := page.MustElement(selector)
					text1 := el.MustText()
					log.Println("📥 Ausgelesener Text:", text1)

					inputField := page.MustElement("div[aria-label='Type a message']")
					inputField.MustClick()
					inputText := fmt.Sprintf("Hallo! %s Ich bin beschäftigt, aber ich werde mich bald melden. Danke für deine Geduld!", text1)

					inputField.MustInput(inputText)

					if text1 == "Pia Schumacher" {
						page.Keyboard.Press(input.Enter)
						inputField.MustFocus()
						log.Println("✅ Nachricht gesendet an:", text1)

					}

					messages, err := page.Elements("span.selectable-text.copyable-text")
					if err != nil || len(messages) == 0 {
						log.Println("❌ Keine Nachrichten gefunden.")
						continue
					}

					// Letzte Nachricht lesen
					lastMsg := messages[len(messages)-1]
					text, err := lastMsg.Text()
					if err != nil {
						log.Println("❌ Fehler beim Auslesen der Nachricht:", err)
						continue
					}

					log.Printf("📩 Neue Nachricht: %s\n", text)
				}
			}
		}

		time.Sleep(10 * time.Second)
	}
}
