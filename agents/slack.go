package agents

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/tech-thinker/chatz/config"
)

type slackAgent struct {
    config *config.Config
}

func (agent *slackAgent) Post(message string) (interface{}, error) {
    url := "https://slack.com/api/chat.postMessage"

	payloadStr := fmt.Sprintf(
            `{"channel": "%s","text": "%s"}`,
            agent.config.SlackChannelId, message,
        )

    payload := strings.NewReader(payloadStr)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "tech-thinker/chatz")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", agent.config.SlackToken))

	res, err := http.DefaultClient.Do(req)
    if err!=nil {
        return nil, err
    }

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
    return string(body), err
}

func (agent *slackAgent) Reply(threadId string, message string) (interface{}, error) {
    url := "https://slack.com/api/chat.postMessage"

	payloadStr := fmt.Sprintf(
            `{"channel": "%s", "text": "%s", "thread_ts": "%s"}`,
            agent.config.SlackChannelId, message, threadId,
        )

    payload := strings.NewReader(payloadStr)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "tech-thinker/chatz")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", agent.config.SlackToken))

	res, err := http.DefaultClient.Do(req)
    if err!=nil {
        return nil, err
    }

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
    return string(body), err
}

func NewSlackAgent(config *config.Config) Agent {
    return &slackAgent{config: config}
}
