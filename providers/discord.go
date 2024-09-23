package providers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/tech-thinker/chatz/config"
)

type discordProvider struct {
    config *config.Config
}

func (agent *discordProvider) Post(message string) (interface{}, error) {
    url := agent.config.WebHookURL

	payloadStr := fmt.Sprintf(
            `{"content": "%s"}`,
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

func (agent *discordProvider) Reply(threadId string, message string) (interface{}, error) {
    fmt.Println("Reply to discord not supported yet.")
    return nil, errors.New("Reply to discord not supported yet.")
}

func NewDiscordProvider(config *config.Config) Provider {
    return &discordProvider{config: config}
}
