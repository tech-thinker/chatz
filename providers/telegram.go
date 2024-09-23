package providers

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/tech-thinker/chatz/config"
)

type telegramProvider struct {
    config *config.Config
}

func (agent *telegramProvider) Post(message string) (interface{}, error) {
    url := fmt.Sprintf(
        `https://api.telegram.org/bot%s/sendMessage`,
        agent.config.Token,
        )

	payloadStr := fmt.Sprintf(
            `{"chat_id": "%s","text": "%s"}`,
            agent.config.ChatId, message,
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

func (agent *telegramProvider) Reply(threadId string, message string) (interface{}, error) {
    url := fmt.Sprintf(
        `https://api.telegram.org/bot%s/sendMessage`,
        agent.config.Token,
        )

	payloadStr := fmt.Sprintf(
            `{"chat_id": "%s", "text": "%s", "reply_to_message_id": "%s"}`,
            agent.config.ChatId, message, threadId,
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

func NewTelegramProvider(config *config.Config) Provider {
    return &telegramProvider{config: config}
}
