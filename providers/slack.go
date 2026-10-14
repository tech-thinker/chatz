package providers

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/tech-thinker/chatz/config"
	"github.com/tech-thinker/chatz/constants"
	"github.com/tech-thinker/chatz/models"
)

type SlackProvider struct {
	config *config.Config
}

func (agent *SlackProvider) Post(message string, option models.Option) (any, error) {
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
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	return string(body), err
}

func (agent *SlackProvider) Reply(threadId string, message string, option models.Option) (any, error) {
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
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	return string(body), err
}

func (agent *SlackProvider) setup(env *config.Config) error {
	if env == nil {
		return fmt.Errorf("config is required")
	}
	if env.Token == "" {
		return fmt.Errorf("token is required for Slack provider")
	}
	if env.ChannelId == "" {
		return fmt.Errorf("channel_id is required for Slack provider")
	}
	agent.config = env
	return nil
}

func init() {
	RegisterProvider(constants.PROVIDER_SLACK, new(SlackProvider))
}
