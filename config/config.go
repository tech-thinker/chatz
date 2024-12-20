package config

// Environment
type Config struct {
	Provider      string
	WebHookURL    string
	Token         string
	ChannelId     string
	ChatId        string
	ConnectionURL string
	SMTPHost      string
	SMTPPort      string
	UseTLS        bool
	UseSTARTTLS   bool
	SMTPUser      string
	SMTPPassword  string
	SMTPSubject   string
	SMTPFrom      string
	SMTPTo        string
}
