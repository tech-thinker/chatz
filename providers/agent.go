package providers

import (
	"errors"

	"github.com/tech-thinker/chatz/config"
	"github.com/tech-thinker/chatz/constants"
	"github.com/tech-thinker/chatz/models"
)

var providerRegistry = make(map[constants.ProviderType]Provider)

// Provider defines the interface for messaging providers.
type Provider interface {
	setup(env *config.Config) error
	Post(message string, option models.Option) (any, error)
	Reply(threadId string, message string, option models.Option) (any, error)
}

// RegisterProvider registers a provider implementation for a given provider type.
func RegisterProvider(providerType constants.ProviderType, provider Provider) {
	providerRegistry[providerType] = provider
}

// NewProvider creates and initializes a provider based on the configuration.
func NewProvider(config *config.Config) (Provider, error) {
	if provider, ok := providerRegistry[constants.ProviderType(config.Provider)]; ok {
		err := provider.setup(config)
		if err != nil {
			return nil, err
		}
		return provider, nil
	}
	return nil, errors.New("Invalid PROVIDER in Config/Environment.")
}
