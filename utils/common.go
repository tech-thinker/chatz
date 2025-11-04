package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/tech-thinker/chatz/config"
)

// LoadEnv loads configuration from environment variables or config file.
func LoadEnv(profile string, fromEnv bool) (*config.Config, error) {
	if fromEnv {
		return loadEnvFromSystemEnv()
	} else {
		return loadEnvFromFile(profile)
	}
}

// loadEnvFromSystemEnv loads configuration from system environment variables.
func loadEnvFromSystemEnv() (*config.Config, error) {
	v := viper.New()
	v.AutomaticEnv()

	// Get values from system environment
	provider := v.GetString("PROVIDER")
	token := v.GetString("TOKEN")
	channelId := v.GetString("CHANNEL_ID")
	webHookURL := v.GetString("WEB_HOOK_URL")
	chatId := v.GetString("CHAT_ID")
	connectionURL := v.GetString("CONNECTION_URL")
	smtpHost := v.GetString("SMTP_HOST")
	smtpPort := v.GetString("SMTP_PORT")
	useTLS := v.GetBool("SMTP_USE_TLS")
	useSTARTTLS := v.GetBool("SMTP_USE_STARTTLS")
	smtpUser := v.GetString("SMTP_USER")
	smtpPassword := v.GetString("SMTP_PASSWORD")
	smtpSubject := v.GetString("SMTP_SUBJECT")
	smtpFrom := v.GetString("SMTP_FROM")
	smtpTo := v.GetString("SMTP_TO")
	gotifyURL := v.GetString("GOTIFY_URL")
	gotifyToken := v.GetString("GOTIFY_TOKEN")
	gotifyTitle := v.GetString("GOTIFY_TITLE")
	gotifyPriority := v.GetInt("GOTIFY_PRIORITY")

	var env config.Config

	env.Provider = provider
	env.WebHookURL = webHookURL
	env.Token = token
	env.ChannelId = channelId
	env.ChatId = chatId
	env.ConnectionURL = connectionURL
	env.SMTPHost = smtpHost
	env.SMTPPort = smtpPort
	env.UseTLS = useTLS
	env.UseSTARTTLS = useSTARTTLS
	env.SMTPUser = smtpUser
	env.SMTPPassword = smtpPassword
	env.SMTPSubject = smtpSubject
	env.SMTPFrom = smtpFrom
	env.SMTPTo = smtpTo
	env.GotifyURL = gotifyURL
	env.GotifyToken = gotifyToken
	env.GotifyTitle = gotifyTitle
	env.GotifyPriority = gotifyPriority

	return &env, nil
}

// loadEnvFromFile loads configuration from the .chatz.ini file in the user's home directory.
func loadEnvFromFile(profile string) (*config.Config, error) {
	// Get the home directory of the user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %s\n", err)
		return nil, err
	}

	// Set the configuration file path
	configPath := filepath.Join(homeDir, ".chatz.ini")

	// Set the file name and type
	viper.SetConfigFile(configPath) // full path to the config file
	viper.SetConfigType("ini")      // or "yaml", "json", etc.

	// Read the configuration file
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return nil, err
	}

	// Get values from the INI file
	provider := viper.GetString(fmt.Sprintf("%s.PROVIDER", profile))
	token := viper.GetString(fmt.Sprintf("%s.TOKEN", profile))
	channelId := viper.GetString(fmt.Sprintf("%s.CHANNEL_ID", profile))
	webHookURL := viper.GetString(fmt.Sprintf("%s.WEB_HOOK_URL", profile))
	chatId := viper.GetString(fmt.Sprintf("%s.CHAT_ID", profile))
	connectionURL := viper.GetString(fmt.Sprintf("%s.CONNECTION_URL", profile))

	smtpHost := viper.GetString(fmt.Sprintf("%s.SMTP_HOST", profile))
	smtpPort := viper.GetString(fmt.Sprintf("%s.SMTP_PORT", profile))
	useTLS := viper.GetBool(fmt.Sprintf("%s.SMTP_USE_TLS", profile))
	useSTARTTLS := viper.GetBool(fmt.Sprintf("%s.SMTP_USE_STARTTLS", profile))
	smtpUser := viper.GetString(fmt.Sprintf("%s.SMTP_USER", profile))
	smtpPassword := viper.GetString(fmt.Sprintf("%s.SMTP_PASSWORD", profile))
	smtpSubject := viper.GetString(fmt.Sprintf("%s.SMTP_SUBJECT", profile))
	smtpFrom := viper.GetString(fmt.Sprintf("%s.SMTP_FROM", profile))
	smtpTo := viper.GetString(fmt.Sprintf("%s.SMTP_TO", profile))

	gotifyURL := viper.GetString(fmt.Sprintf("%s.GOTIFY_URL", profile))
	gotifyToken := viper.GetString(fmt.Sprintf("%s.GOTIFY_TOKEN", profile))
	gotifyTitle := viper.GetString(fmt.Sprintf("%s.GOTIFY_TITLE", profile))
	gotifyPriority := viper.GetInt(fmt.Sprintf("%s.GOTIFY_PRIORITY", profile))

	var env config.Config
	env.Provider = provider
	env.WebHookURL = webHookURL
	env.Token = token
	env.ChannelId = channelId
	env.ChatId = chatId
	env.ConnectionURL = connectionURL
	env.SMTPHost = smtpHost
	env.SMTPPort = smtpPort
	env.UseTLS = useTLS
	env.UseSTARTTLS = useSTARTTLS
	env.SMTPUser = smtpUser
	env.SMTPPassword = smtpPassword
	env.SMTPSubject = smtpSubject
	env.SMTPFrom = smtpFrom
	env.SMTPTo = smtpTo
	env.GotifyURL = gotifyURL
	env.GotifyToken = gotifyToken
	env.GotifyTitle = gotifyTitle
	env.GotifyPriority = gotifyPriority

	return &env, nil
}
