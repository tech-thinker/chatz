package providers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/tech-thinker/chatz/config"
)

type GotifyProvider struct {
	config *config.Config
}

func (agent *GotifyProvider) Post(message string) (interface{}, error) {
	return agent.PostWithPriority(message, "", 0)
}

func (agent *GotifyProvider) PostWithPriority(message string, title string, priority int) (interface{}, error) {
	url := fmt.Sprintf("%s/message", agent.config.GotifyURL)

	if len(title) == 0 {
		title = "Chatz Notification"
	}
	if priority == 0 {
		priority = 5
	}

	payloadStr := fmt.Sprintf(
		`{"message": "%s", "priority": %d, "title": "%s"}`,
		message,
		priority,
		title,
	)

	payload := strings.NewReader(payloadStr)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "tech-thinker/chatz")
	req.Header.Add("X-Gotify-Key", agent.config.GotifyToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	return string(body), err
}

func (agent *GotifyProvider) Reply(threadId string, message string) (interface{}, error) {
	fmt.Println("Reply to gotify not supported yet.")
	return nil, errors.New("reply to gotify not supported yet")
}