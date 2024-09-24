package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/tech-thinker/chatz/config"
)


func LoadEnv(profile string, fromEnv bool) (*config.Config, error) {
    if fromEnv {
        return loadEnvFromSystemEnv()
    } else {
        return loadEnvFromFile(profile)
    }
}

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

    var env config.Config

    env.Provider = provider
    env.WebHookURL = webHookURL
    env.Token = token
    env.ChannelId = channelId
    env.ChatId = chatId
    env.ConnectionURL = connectionURL

    return &env, nil
}

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
    viper.SetConfigType("ini")       // or "yaml", "json", etc.

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


    var env config.Config
    env.Provider = provider
    env.WebHookURL = webHookURL
    env.Token = token
    env.ChannelId = channelId
    env.ChatId = chatId
    env.ConnectionURL = connectionURL

    return &env, nil
}

