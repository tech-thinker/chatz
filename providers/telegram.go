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

type TelegramProvider struct {
	config *config.Config
}

func (agent *TelegramProvider) Post(message string, option models.Option) (any, error) {
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
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	return string(body), err
}

func (agent *TelegramProvider) Reply(threadId string, message string, option models.Option) (any, error) {
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
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	return string(body), err
}

func (agent *TelegramProvider) setup(env *config.Config) error {
	if env == nil {
		return fmt.Errorf("config is required")
	}
	if env.Token == "" {
		return fmt.Errorf("token is required for Telegram provider")
	}
	if env.ChatId == "" {
		return fmt.Errorf("chat_id is required for Telegram provider")
	}
	agent.config = env
	return nil
}

func init() {
	RegisterProvider(constants.PROVIDER_TELEGRAM, new(TelegramProvider))
}
