package config

import (
	"reflect"
	"testing"
)

func TestConfig_Struct(t *testing.T) {
	// Test that Config struct can be created and fields can be set
	cfg := &Config{
		Provider:       "slack",
		WebHookURL:     "https://hooks.slack.com/test",
		Token:          "test-token",
		ChannelId:      "C1234567890",
		ChatId:         "123456789",
		ConnectionURL:  "redis://localhost:6379",
		SMTPHost:       "smtp.gmail.com",
		SMTPPort:       "587",
		UseTLS:         false,
		UseSTARTTLS:    true,
		SMTPUser:       "test@example.com",
		SMTPPassword:   "test-password",
		SMTPSubject:    "Test Subject",
		SMTPFrom:       "from@example.com",
		SMTPTo:         "to@example.com",
		GotifyURL:      "https://gotify.example.com",
		GotifyToken:    "gotify-token",
		GotifyTitle:    "Test Title",
		GotifyPriority: 5,
	}

	// Verify all fields are set correctly
	if cfg.Provider != "slack" {
		t.Errorf("expected Provider 'slack', got '%s'", cfg.Provider)
	}
	if cfg.WebHookURL != "https://hooks.slack.com/test" {
		t.Errorf("expected WebHookURL 'https://hooks.slack.com/test', got '%s'", cfg.WebHookURL)
	}
	if cfg.Token != "test-token" {
		t.Errorf("expected Token 'test-token', got '%s'", cfg.Token)
	}
	if cfg.ChannelId != "C1234567890" {
		t.Errorf("expected ChannelId 'C1234567890', got '%s'", cfg.ChannelId)
	}
	if cfg.ChatId != "123456789" {
		t.Errorf("expected ChatId '123456789', got '%s'", cfg.ChatId)
	}
	if cfg.ConnectionURL != "redis://localhost:6379" {
		t.Errorf("expected ConnectionURL 'redis://localhost:6379', got '%s'", cfg.ConnectionURL)
	}
	if cfg.SMTPHost != "smtp.gmail.com" {
		t.Errorf("expected SMTPHost 'smtp.gmail.com', got '%s'", cfg.SMTPHost)
	}
	if cfg.SMTPPort != "587" {
		t.Errorf("expected SMTPPort '587', got '%s'", cfg.SMTPPort)
	}
	if cfg.UseTLS != false {
		t.Errorf("expected UseTLS false, got %v", cfg.UseTLS)
	}
	if cfg.UseSTARTTLS != true {
		t.Errorf("expected UseSTARTTLS true, got %v", cfg.UseSTARTTLS)
	}
	if cfg.SMTPUser != "test@example.com" {
		t.Errorf("expected SMTPUser 'test@example.com', got '%s'", cfg.SMTPUser)
	}
	if cfg.SMTPPassword != "test-password" {
		t.Errorf("expected SMTPPassword 'test-password', got '%s'", cfg.SMTPPassword)
	}
	if cfg.SMTPSubject != "Test Subject" {
		t.Errorf("expected SMTPSubject 'Test Subject', got '%s'", cfg.SMTPSubject)
	}
	if cfg.SMTPFrom != "from@example.com" {
		t.Errorf("expected SMTPFrom 'from@example.com', got '%s'", cfg.SMTPFrom)
	}
	if cfg.SMTPTo != "to@example.com" {
		t.Errorf("expected SMTPTo 'to@example.com', got '%s'", cfg.SMTPTo)
	}
	if cfg.GotifyURL != "https://gotify.example.com" {
		t.Errorf("expected GotifyURL 'https://gotify.example.com', got '%s'", cfg.GotifyURL)
	}
	if cfg.GotifyToken != "gotify-token" {
		t.Errorf("expected GotifyToken 'gotify-token', got '%s'", cfg.GotifyToken)
	}
	if cfg.GotifyTitle != "Test Title" {
		t.Errorf("expected GotifyTitle 'Test Title', got '%s'", cfg.GotifyTitle)
	}
	if cfg.GotifyPriority != 5 {
		t.Errorf("expected GotifyPriority 5, got %d", cfg.GotifyPriority)
	}
}

func TestConfig_ProviderConfigurations(t *testing.T) {
	tests := []struct {
		name     string
		config   Config
		provider string
	}{
		{
			name: "slack configuration",
			config: Config{
				Provider:  "slack",
				Token:     "slack-token",
				ChannelId: "C1234567890",
			},
			provider: "slack",
		},
		{
			name: "discord configuration",
			config: Config{
				Provider:   "discord",
				WebHookURL: "https://discord.com/api/webhooks/test",
			},
			provider: "discord",
		},
		{
			name: "telegram configuration",
			config: Config{
				Provider: "telegram",
				Token:    "telegram-token",
				ChatId:   "123456789",
			},
			provider: "telegram",
		},
		{
			name: "smtp configuration",
			config: Config{
				Provider:     "smtp",
				SMTPHost:     "smtp.gmail.com",
				SMTPPort:     "587",
				UseSTARTTLS:  true,
				SMTPUser:     "user@example.com",
				SMTPPassword: "password",
				SMTPFrom:     "from@example.com",
				SMTPTo:       "to@example.com",
			},
			provider: "smtp",
		},
		{
			name: "redis configuration",
			config: Config{
				Provider:      "redis",
				ConnectionURL: "redis://localhost:6379",
				ChannelId:     "test-channel",
			},
			provider: "redis",
		},
		{
			name: "gotify configuration",
			config: Config{
				Provider:       "gotify",
				GotifyURL:      "https://gotify.example.com",
				GotifyToken:    "gotify-token",
				GotifyPriority: 5,
			},
			provider: "gotify",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.config.Provider != tt.provider {
				t.Errorf("expected Provider '%s', got '%s'", tt.provider, tt.config.Provider)
			}

			// Verify the config struct is properly initialized
			configType := reflect.TypeOf(tt.config)
			configValue := reflect.ValueOf(tt.config)

			for i := 0; i < configType.NumField(); i++ {
				field := configType.Field(i)
				value := configValue.Field(i)

				// Check that fields are accessible (not panicking)
				_ = field.Name
				_ = value.Interface()
			}
		})
	}
}

func TestConfig_ZeroValues(t *testing.T) {
	// Test zero values of Config struct
	var cfg Config

	// Verify zero values
	if cfg.Provider != "" {
		t.Errorf("expected zero value for Provider, got '%s'", cfg.Provider)
	}
	if cfg.WebHookURL != "" {
		t.Errorf("expected zero value for WebHookURL, got '%s'", cfg.WebHookURL)
	}
	if cfg.Token != "" {
		t.Errorf("expected zero value for Token, got '%s'", cfg.Token)
	}
	if cfg.ChannelId != "" {
		t.Errorf("expected zero value for ChannelId, got '%s'", cfg.ChannelId)
	}
	if cfg.UseTLS != false {
		t.Errorf("expected zero value for UseTLS, got %v", cfg.UseTLS)
	}
	if cfg.UseSTARTTLS != false {
		t.Errorf("expected zero value for UseSTARTTLS, got %v", cfg.UseSTARTTLS)
	}
	if cfg.GotifyPriority != 0 {
		t.Errorf("expected zero value for GotifyPriority, got %d", cfg.GotifyPriority)
	}
}

func TestConfig_InvalidDataTypes(t *testing.T) {
	tests := []struct {
		name        string
		config      Config
		expectError bool
		description string
	}{
		{
			name: "invalid port number",
			config: Config{
				SMTPPort: "not-a-number",
			},
			expectError: false, // Config struct doesn't validate data types
			description: "Config struct accepts any string values",
		},
		{
			name: "negative priority",
			config: Config{
				GotifyPriority: -1,
			},
			expectError: false, // Config struct doesn't validate ranges
			description: "Config struct accepts any int values",
		},
		{
			name: "very large priority",
			config: Config{
				GotifyPriority: 999999,
			},
			expectError: false,
			description: "Config struct accepts any int values",
		},
		{
			name: "empty strings",
			config: Config{
				Provider:      "",
				WebHookURL:    "",
				Token:         "",
				ChannelId:     "",
				ChatId:        "",
				ConnectionURL: "",
				SMTPHost:      "",
				SMTPPort:      "",
				SMTPUser:      "",
				SMTPPassword:  "",
				SMTPSubject:   "",
				SMTPFrom:      "",
				SMTPTo:        "",
				GotifyURL:     "",
				GotifyToken:   "",
				GotifyTitle:   "",
			},
			expectError: false,
			description: "All string fields can be empty",
		},
		{
			name: "very long strings",
			config: Config{
				Provider:      string(make([]byte, 10000)),
				WebHookURL:    string(make([]byte, 10000)),
				Token:         string(make([]byte, 10000)),
				ChannelId:     string(make([]byte, 10000)),
				ChatId:        string(make([]byte, 10000)),
				ConnectionURL: string(make([]byte, 10000)),
				SMTPHost:      string(make([]byte, 10000)),
				SMTPPort:      string(make([]byte, 10000)),
				SMTPUser:      string(make([]byte, 10000)),
				SMTPPassword:  string(make([]byte, 10000)),
				SMTPSubject:   string(make([]byte, 10000)),
				SMTPFrom:      string(make([]byte, 10000)),
				SMTPTo:        string(make([]byte, 10000)),
				GotifyURL:     string(make([]byte, 10000)),
				GotifyToken:   string(make([]byte, 10000)),
				GotifyTitle:   string(make([]byte, 10000)),
			},
			expectError: false,
			description: "Config struct accepts very long strings",
		},
		{
			name: "strings with special characters",
			config: Config{
				Provider:      "provider_with_émojis_🎉_and_spëcial_chärs",
				WebHookURL:    "https://example.com/path with spaces?query=value&other=🎉",
				Token:         "token-with-special-chars!@#$%^&*()",
				ChannelId:     "channel_123_!@#",
				ChatId:        "chat-456_!@#",
				ConnectionURL: "redis://user:pass@host:6379/db?param=value&other=🎉",
				SMTPHost:      "smtp.gmail.com",
				SMTPPort:      "587",
				SMTPUser:      "user+tag@example.com",
				SMTPPassword:  "pass!@#$%^&*()",
				SMTPSubject:   "Subject with émojis 🎉",
				SMTPFrom:      "from+tag@example.com",
				SMTPTo:        "to1@example.com,to2@example.com",
				GotifyURL:     "https://gotify.example.com",
				GotifyToken:   "token_!@#$%^&*()",
				GotifyTitle:   "Title with spëcial chärs 🎉",
			},
			expectError: false,
			description: "Config struct accepts strings with special characters",
		},
		{
			name: "malformed URLs",
			config: Config{
				WebHookURL:    "not-a-url",
				ConnectionURL: "invalid-url-format",
				GotifyURL:     "also-not-a-url",
			},
			expectError: false,
			description: "Config struct doesn't validate URL formats",
		},
		{
			name: "invalid email formats",
			config: Config{
				SMTPUser: "not-an-email",
				SMTPFrom: "also-not-an-email",
				SMTPTo:   "invalid-email-1,invalid-email-2",
			},
			expectError: false,
			description: "Config struct doesn't validate email formats",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Config struct doesn't perform validation, so we just test that it can hold the values
			if tt.expectError {
				t.Errorf("unexpected: Config struct should not validate data, but test expects error: %s", tt.description)
			}

			// Verify the config struct can hold these values without panicking
			_ = tt.config.Provider
			_ = tt.config.WebHookURL
			_ = tt.config.Token
			_ = tt.config.ChannelId
			_ = tt.config.ChatId
			_ = tt.config.ConnectionURL
			_ = tt.config.SMTPHost
			_ = tt.config.SMTPPort
			_ = tt.config.UseTLS
			_ = tt.config.UseSTARTTLS
			_ = tt.config.SMTPUser
			_ = tt.config.SMTPPassword
			_ = tt.config.SMTPSubject
			_ = tt.config.SMTPFrom
			_ = tt.config.SMTPTo
			_ = tt.config.GotifyURL
			_ = tt.config.GotifyToken
			_ = tt.config.GotifyTitle
			_ = tt.config.GotifyPriority
		})
	}
}
