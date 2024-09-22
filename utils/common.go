package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/tech-thinker/chatz/config"
)


func LoadEnv(profile string) (*config.Config, error) {
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
    viper.SetConfigType("ini")       // or "yaml", "json", etc.

    // Read the configuration file
    err = viper.ReadInConfig()
    if err != nil {
        fmt.Printf("Error reading config file: %s\n", err)
        return nil, err
    }

    // Get values from the INI file
    provider := viper.GetString(fmt.Sprintf("%s.PROVIDER", profile))
    slackToken := viper.GetString(fmt.Sprintf("%s.SLACK_TOKEN", profile))
    channelId := viper.GetString(fmt.Sprintf("%s.CHANNEL_ID", profile))
    googleWebHookURL := viper.GetString(fmt.Sprintf("%s.WEB_HOOK_URL", profile))

    var env config.Config
    env.Provider = provider
    env.SlackToken = slackToken
    env.SlackChannelId = channelId
    env.GoogleWebHookURL = googleWebHookURL

    return &env, nil
}

func generateDefaultConfig(configPath string) {
    config := `[default]
PROVIDER=slack
SLACK_TOKEN=
CHANNEL_ID=

[slack]
PROVIDER=slack
SLACK_TOKEN=
CHANNEL_ID=

[google]
PROVIDER=google
WEB_HOOK_URL=
`
    f, err := os.Create(configPath)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    f.WriteString(config)
}
