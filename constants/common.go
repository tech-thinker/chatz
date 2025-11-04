package constants

// ProviderType represents the type of messaging provider.
type ProviderType string

const (
	PROVIDER_SLACK    ProviderType = "slack"
	PROVIDER_DISCORD  ProviderType = "discord"
	PROVIDER_TELEGRAM ProviderType = "telegram"
	PROVIDER_GOOGLE   ProviderType = "google"
	PROVIDER_REDIS    ProviderType = "redis"
	PROVIDER_SMTP     ProviderType = "smtp"
	PROVIDER_GOTIFY   ProviderType = "gotify"
)
