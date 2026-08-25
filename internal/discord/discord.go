package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

// large wallet movement
// volume spike
// low wallet balance

// everyoneMention is sent as the message content of every notification so
// the whole channel is pinged.
const everyoneMention = "@everyone"

const (
	ColorBlue   = 0x3498DB
	ColorYellow = 0xF1C40F
	ColorGreen  = 0x2ECC71
	ColorRed    = 0xE74C3C
	ColorPurple = 0x9B59B6
)

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type Footer struct {
	Text string `json:"text,omitempty"`
}

type Embed struct {
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	URL         string  `json:"url,omitempty"`
	Color       int     `json:"color,omitempty"`
	Fields      []Field `json:"fields,omitempty"`
	Footer      *Footer `json:"footer,omitempty"`
	Timestamp   string  `json:"timestamp,omitempty"`
}

// allowedMentions opts the payload into the mentions it is allowed to
// resolve. Without it Discord still parses @everyone, but stating it keeps
// the intent explicit.
type allowedMentions struct {
	Parse []string `json:"parse"`
}

type message struct {
	Username string  `json:"username,omitempty"`
	Content  string  `json:"content,omitempty"`
	Embeds   []Embed `json:"embeds,omitempty"`
	// A mention only notifies from the content field; one placed inside an
	// embed renders as text and pings nobody.
	AllowedMentions *allowedMentions `json:"allowed_mentions,omitempty"`
}

var httpClient = &http.Client{Timeout: 10 * time.Second}

var warnMissingWebhookOnce sync.Once

func SendEmbed(embed Embed) {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookURL == "" {
		warnMissingWebhookOnce.Do(func() {
			log.Println("DISCORD_WEBHOOK_URL not set; Discord notifications disabled")
		})
		return
	}

	if embed.Timestamp == "" {
		embed.Timestamp = time.Now().UTC().Format(time.RFC3339)
	}

	payload, err := json.Marshal(message{
		Username:        "Contract Monitor",
		Content:         everyoneMention,
		Embeds:          []Embed{embed},
		AllowedMentions: &allowedMentions{Parse: []string{"everyone"}},
	})
	if err != nil {
		log.Printf("Failed to marshal Discord payload: %v", err)
		return
	}

	if err := post(webhookURL, payload); err != nil {
		log.Printf("Failed to send Discord notification: %v", err)
	}
}

const maxAttempts = 3

// post delivers the payload, retrying on network errors, rate limits and
// server errors. On 429 it waits for Discord's Retry-After before retrying.
func post(webhookURL string, payload []byte) error {
	var lastErr error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		wait := time.Duration(attempt) * 2 * time.Second

		resp, err := httpClient.Post(webhookURL, "application/json", bytes.NewReader(payload))
		if err != nil {
			lastErr = err
		} else {
			retryAfter := resp.Header.Get("Retry-After")
			resp.Body.Close()

			if resp.StatusCode < 300 {
				return nil
			}
			lastErr = fmt.Errorf("discord webhook returned status %d", resp.StatusCode)

			if resp.StatusCode == http.StatusTooManyRequests {
				if secs, parseErr := strconv.ParseFloat(retryAfter, 64); parseErr == nil && secs > 0 {
					wait = time.Duration(secs * float64(time.Second))
				}
			} else if resp.StatusCode < 500 {
				// Other 4xx responses (bad payload, deleted webhook) won't
				// succeed on retry.
				return lastErr
			}
		}

		if attempt < maxAttempts {
			time.Sleep(wait)
		}
	}
	return lastErr
}
