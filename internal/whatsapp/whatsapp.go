package whatsapp

//https://developers.facebook.com/docs/whatsapp/cloud-api/overview
//https://developers.facebook.com/apps/1030673848841803/whatsapp-business/wa-dev-console/?business_id=2467686813608135
//https://ngrok.com/downloads/windows
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	AccessToken   string
	PhoneNumberID string
}

func NewClient(accessToken, phoneNumberID string) *Client {
	return &Client{
		AccessToken:   accessToken,
		PhoneNumberID: phoneNumberID,
	}
}

func (c *Client) SendTextMessage(to, message string) error {
	url := fmt.Sprintf("https://graph.facebook.com/v22.0/%s/messages", c.PhoneNumberID)
    log.Println("URL:", url)
	fmt.Println(string(url))

payload := map[string]interface{}{
	"messaging_product": "whatsapp",
	"to":                to,
	"type":              "text",
	"text": map[string]interface{}{
		"preview_url": true,
		"body":        message,
	},
}
	

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	fmt.Println(string(jsonPayload))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
	// Read and print the response body
	bodyBytes, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("unexpected status code: %d\nresponse body: %s", resp.StatusCode, string(bodyBytes))
}

	return nil
}

