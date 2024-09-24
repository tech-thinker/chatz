package providers

import (
	"errors"

	"github.com/tech-thinker/chatz/config"
	"github.com/tech-thinker/chatz/constants"
)

type Provider interface {
    Post(message string) (interface{}, error)
    Reply(threadId string, message string) (interface{}, error)
}

func NewProvider(env *config.Config) (Provider, error) {
    switch env.Provider {
        case constants.PROVIDER_SLACK:
            return &SlackProvider{config: env}, nil
        case constants.PROVIDER_GOOGLE:
            return &GoogleProvider{config: env}, nil
        case constants.PROVIDER_TELEGRAM:
            return &TelegramProvider{config: env}, nil
        case constants.PROVIDER_DISCORD:
            return &DiscordProvider{config: env}, nil
        case constants.PROVIDER_REDIS:
            return &RedisProvider{config: env}, nil
        default:
            return nil, errors.New("Invalid provider config in ~/.chatz.ini")
    }

}
