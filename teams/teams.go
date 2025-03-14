package teams

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/tomekwlod/notificator"
)

type Teams struct {
	name       string
	webhookURL string
	// prefix     string
}

func New(name, webhookURL string) *Teams {
	if name == "" {
		panic("name is required")
	}
	return &Teams{name: name, webhookURL: webhookURL}
}

func (t *Teams) Send(ctx context.Context, title, message string, intent notificator.Intent) error {
	title = fmt.Sprintf("[%s] %s", t.name, title)

	msg := MsgCard{
		Title:      title,
		Text:       message,
		ThemeColor: string(intent),
	}

	// Create a JSON payload for the message
	b, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshall err: %+v", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post(t.webhookURL, "application/json", bytes.NewBuffer(b))
	// Send a POST request to the webhook URL
	// resp, err := http.Post(t.webhook, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("teams send error: %+v", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("couldn't send the teams message, status: %s", resp.Status)
	}

	return nil
}
