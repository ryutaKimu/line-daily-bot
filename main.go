package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type NarrationResponse struct {
	Data  Narration `json:"data"`
	Error string    `json:"error,omitempty"`
}

type Narration struct {
	Episode    int      `json:"episode"`
	Title      string   `json:"title"`
	Narrations []string `json:"narrations"`
}

func fetchNarrationAPI() (NarrationResponse, error) {
	url := "https://fullmetalapi.vercel.app/narrations/random"
	res, err := http.Get(url)
	if err != nil {
		return NarrationResponse{}, fmt.Errorf("get data failed: %w", err)
	}
	defer res.Body.Close()

	var narration NarrationResponse
	data, err := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return NarrationResponse{}, fmt.Errorf("unexpected status %d: %s", res.StatusCode, data)
	}

	if err != nil {
		return NarrationResponse{}, fmt.Errorf("read Data failed: %w", err)
	}
	if err := json.Unmarshal(data, &narration); err != nil {
		return NarrationResponse{}, fmt.Errorf("unmarshal failed:%w", err)
	}

	if narration.Error != "" {
		return NarrationResponse{}, fmt.Errorf("api error: %s", narration.Error)
	}

	return narration, nil
}

func buildMessage(n Narration) string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "鋼の錬金術師 第%d話\n", n.Episode)
	fmt.Fprintf(&sb, "『%s』", n.Title)
	if len(n.Narrations) > 0 {
		sb.WriteString("\n\n")
		sb.WriteString(strings.Join(n.Narrations, "\n"))
	}
	return sb.String()
}

func formatNarattions() (string, error) {
	narration, err := fetchNarrationAPI()
	if err != nil {
		return "", fmt.Errorf("fetchNarrationAPI error:%w", err)
	}
	return buildMessage(narration.Data), nil
}

func main() {
	text, err := formatNarattions()
	if err != nil {
		log.Fatal(err)
	}
	body, _ := json.Marshal(map[string]any{
		"to":       os.Getenv("LINE_TO"),
		"messages": []map[string]string{{"type": "text", "text": text}},
	})

	req, _ := http.NewRequest("POST", "https://api.line.me/v2/bot/message/push", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("LINE_TOKEN"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(res.Body)
		log.Fatalf("送信失敗: status %d: %s", res.StatusCode, b)
	}

}
