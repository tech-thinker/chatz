package agents

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/tech-thinker/chat/config"
)

type googleAgent struct {
    config *config.Config
}

func (agent *googleAgent) Post(message string) (interface{}, error) {
    url := agent.config.GoogleWebHookURL

	payloadStr := fmt.Sprintf(
            `{"text": "%s"}`,
            message,
        )

    payload := strings.NewReader(payloadStr)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "tech-thinker/chat")

	res, err := http.DefaultClient.Do(req)
    if err!=nil {
        return nil, err
    }

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
    return string(body), err
}

func (agent *googleAgent) Reply(threadId string, message string) (interface{}, error) {
    url := fmt.Sprintf("%s&messageReplyOption=REPLY_MESSAGE_FALLBACK_TO_NEW_THREAD",agent.config.GoogleWebHookURL)

	payloadStr := fmt.Sprintf(
            `{"text": "%s", "thread": {"name": "%s"}}`,
            message, threadId,
        )

    payload := strings.NewReader(payloadStr)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "tech-thinker/chat")

	res, err := http.DefaultClient.Do(req)
    if err!=nil {
        return nil, err
    }

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
    return string(body), err
}

func NewGoogleAgent(config *config.Config) Agent {
    return &googleAgent{config: config}
}
