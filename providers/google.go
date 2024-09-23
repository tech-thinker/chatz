package providers

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/tech-thinker/chatz/config"
)

type googleProvider struct {
    config *config.Config
}

func (agent *googleProvider) Post(message string) (interface{}, error) {
    url := agent.config.WebHookURL

	payloadStr := fmt.Sprintf(
            `{"text": "%s"}`,
            message,
        )

    payload := strings.NewReader(payloadStr)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "tech-thinker/chatz")

	res, err := http.DefaultClient.Do(req)
    if err!=nil {
        return nil, err
    }

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
    return string(body), err
}

func (agent *googleProvider) Reply(threadId string, message string) (interface{}, error) {
    url := fmt.Sprintf("%s&messageReplyOption=REPLY_MESSAGE_FALLBACK_TO_NEW_THREAD",agent.config.WebHookURL)

	payloadStr := fmt.Sprintf(
            `{"text": "%s", "thread": {"name": "%s"}}`,
            message, threadId,
        )

    payload := strings.NewReader(payloadStr)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "tech-thinker/chatz")

	res, err := http.DefaultClient.Do(req)
    if err!=nil {
        return nil, err
    }

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
    return string(body), err
}

func NewGoogleProvider(config *config.Config) Provider {
    return &googleProvider{config: config}
}
