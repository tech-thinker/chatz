package config

// Config holds all configuration settings for chatz providers.
type Config struct {
	Provider       string
	WebHookURL     string
	Token          string
	ChannelId      string
	ChatId         string
	ConnectionURL  string
	SMTPHost       string
	SMTPPort       string
	UseTLS         bool
	UseSTARTTLS    bool
	SMTPUser       string
	SMTPPassword   string
	SMTPSubject    string
	SMTPFrom       string
	SMTPTo         string
	GotifyURL      string
	GotifyToken    string
	GotifyTitle    string
	GotifyPriority int
}
