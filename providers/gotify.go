package providers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/tech-thinker/chatz/config"
	"github.com/tech-thinker/chatz/constants"
	"github.com/tech-thinker/chatz/models"
	"github.com/tech-thinker/chatz/utils"
)

type GotifyProvider struct {
	config *config.Config
}

func (agent *GotifyProvider) Post(message string, option models.Option) (any, error) {
	if option.Title == nil {
		if len(agent.config.GotifyTitle) > 0 {
			option.Title = &agent.config.GotifyTitle
		} else {
			option.Title = utils.NewString("Chatz Notification")
		}
	}

	if option.Priority == nil {
		if agent.config.GotifyPriority > 0 {
			option.Priority = utils.NewInt(agent.config.GotifyPriority)
		} else {
			option.Priority = utils.NewInt(5)
		}
	}

	url := fmt.Sprintf("%s/message", agent.config.GotifyURL)

	payloadStr := fmt.Sprintf(
		`{"message": "%s", "priority": %d, "title": "%s"}`,
		message,
		*option.Priority,
		*option.Title,
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

func (agent *GotifyProvider) Reply(threadId string, message string, option models.Option) (any, error) {
	fmt.Println("Reply to gotify not supported yet.")
	return nil, errors.New("reply to gotify not supported yet")
}

func (agent *GotifyProvider) setup(env *config.Config) error {
	if env == nil {
		return fmt.Errorf("config is required")
	}
	if env.GotifyURL == "" {
		return fmt.Errorf("gotify_url is required for Gotify provider")
	}
	if env.GotifyToken == "" {
		return fmt.Errorf("gotify_token is required for Gotify provider")
	}
	agent.config = env
	return nil
}

func init() {
	RegisterProvider(constants.PROVIDER_GOTIFY, new(GotifyProvider))
}
