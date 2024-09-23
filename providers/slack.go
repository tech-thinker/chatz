package providers

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/tech-thinker/chatz/config"
)

type SlackProvider struct {
    config *config.Config
}

func (agent *SlackProvider) Post(message string) (interface{}, error) {
    url := "https://slack.com/api/chat.postMessage"

	payloadStr := fmt.Sprintf(
            `{"channel": "%s","text": "%s"}`,
            agent.config.ChannelId, message,
        )

    payload := strings.NewReader(payloadStr)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "tech-thinker/chatz")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", agent.config.Token))

	res, err := http.DefaultClient.Do(req)
    if err!=nil {
        return nil, err
    }

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
    return string(body), err
}

func (agent *SlackProvider) Reply(threadId string, message string) (interface{}, error) {
    url := "https://slack.com/api/chat.postMessage"

	payloadStr := fmt.Sprintf(
            `{"channel": "%s", "text": "%s", "thread_ts": "%s"}`,
            agent.config.ChannelId, message, threadId,
        )

    payload := strings.NewReader(payloadStr)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "tech-thinker/chatz")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", agent.config.Token))

	res, err := http.DefaultClient.Do(req)
    if err!=nil {
        return nil, err
    }

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
    return string(body), err
}

