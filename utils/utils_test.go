package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tech-thinker/chatz/config"
)

func TestLoadEnv(t *testing.T) {
	tests := []struct {
		name        string
		profile     string
		fromEnv     bool
		setupFunc   func() func() // returns cleanup function
		expectError bool
	}{
		{
			name:    "load from environment variables",
			profile: "test",
			fromEnv: true,
			setupFunc: func() func() {
				originalEnv := map[string]string{}
				envVars := map[string]string{
					"PROVIDER":     "slack",
					"TOKEN":        "test-token",
					"CHANNEL_ID":   "C1234567890",
					"WEB_HOOK_URL": "https://hooks.slack.com/test",
				}

				// Set test environment variables
				for key, value := range envVars {
					if original, exists := os.LookupEnv(key); exists {
						originalEnv[key] = original
					}
					os.Setenv(key, value)
				}

				// Return cleanup function
				return func() {
					for key := range envVars {
						if original, exists := originalEnv[key]; exists {
							os.Setenv(key, original)
						} else {
							os.Unsetenv(key)
						}
					}
				}
			},
			expectError: false,
		},
		{
			name:    "load from config file",
			profile: "testprofile",
			fromEnv: false,
			setupFunc: func() func() {
				// Create temporary directory and config file
				tempDir, err := os.MkdirTemp("", "chatz_test")
				if err != nil {
					t.Fatalf("failed to create temp dir: %v", err)
				}

				configContent := `[testprofile]
PROVIDER=telegram
TOKEN=test-bot-token
CHAT_ID=123456789
WEB_HOOK_URL=https://hooks.telegram.com/test`

				configPath := filepath.Join(tempDir, ".chatz.ini")
				err = os.WriteFile(configPath, []byte(configContent), 0644)
				if err != nil {
					t.Fatalf("failed to write config file: %v", err)
				}

				// Mock home directory
				originalHome := os.Getenv("HOME")
				os.Setenv("HOME", tempDir)

				return func() {
					os.Setenv("HOME", originalHome)
					os.RemoveAll(tempDir)
				}
			},
			expectError: false,
		},
		{
			name:    "config file not found",
			profile: "nonexistent",
			fromEnv: false,
			setupFunc: func() func() {
				// Set home to non-existent directory
				originalHome := os.Getenv("HOME")
				os.Setenv("HOME", "/non/existent/path")

				return func() {
					os.Setenv("HOME", originalHome)
				}
			},
			expectError: true,
		},
		{
			name:    "corrupted config file",
			profile: "test",
			fromEnv: false,
			setupFunc: func() func() {
				tempDir, err := os.MkdirTemp("", "chatz_test")
				if err != nil {
					t.Fatalf("failed to create temp dir: %v", err)
				}

				// Write corrupted INI content
				configContent := `[test]
PROVIDER=slack
TOKEN=test-token
INVALID_LINE_WITHOUT_EQUALS
[invalid_section
CHANNEL_ID=C1234567890`

				configPath := filepath.Join(tempDir, ".chatz.ini")
				err = os.WriteFile(configPath, []byte(configContent), 0644)
				if err != nil {
					t.Fatalf("failed to write config file: %v", err)
				}

				// Mock home directory
				originalHome := os.Getenv("HOME")
				os.Setenv("HOME", tempDir)

				return func() {
					os.Setenv("HOME", originalHome)
					os.RemoveAll(tempDir)
				}
			},
			expectError: true, // viper validates INI format and fails on parsing errors
		},
		{
			name:    "empty config file",
			profile: "test",
			fromEnv: false,
			setupFunc: func() func() {
				tempDir, err := os.MkdirTemp("", "chatz_test")
				if err != nil {
					t.Fatalf("failed to create temp dir: %v", err)
				}

				configPath := filepath.Join(tempDir, ".chatz.ini")
				err = os.WriteFile(configPath, []byte(""), 0644)
				if err != nil {
					t.Fatalf("failed to write config file: %v", err)
				}

				originalHome := os.Getenv("HOME")
				os.Setenv("HOME", tempDir)

				return func() {
					os.Setenv("HOME", originalHome)
					os.RemoveAll(tempDir)
				}
			},
			expectError: false, // Empty file should not cause error
		},
		{
			name:    "config file with permission denied",
			profile: "test",
			fromEnv: false,
			setupFunc: func() func() {
				tempDir, err := os.MkdirTemp("", "chatz_test")
				if err != nil {
					t.Fatalf("failed to create temp dir: %v", err)
				}

				configContent := `[test]
PROVIDER=slack
TOKEN=test-token`

				configPath := filepath.Join(tempDir, ".chatz.ini")
				err = os.WriteFile(configPath, []byte(configContent), 0200) // Write only for owner
				if err != nil {
					t.Fatalf("failed to write config file: %v", err)
				}

				// Try to make directory read-only (this might not work on all systems)
				os.Chmod(tempDir, 0500) // Read and execute only

				originalHome := os.Getenv("HOME")
				os.Setenv("HOME", tempDir)

				return func() {
					os.Setenv("HOME", originalHome)
					os.RemoveAll(tempDir)
				}
			},
			expectError: true, // Should fail due to permission issues
		},
		{
			name:    "missing HOME environment variable",
			profile: "test",
			fromEnv: false,
			setupFunc: func() func() {
				originalHome := os.Getenv("HOME")
				os.Unsetenv("HOME")

				return func() {
					if originalHome != "" {
						os.Setenv("HOME", originalHome)
					}
				}
			},
			expectError: true, // Should fail when HOME is not set
		},
		{
			name:    "empty profile name",
			profile: "",
			fromEnv: false,
			setupFunc: func() func() {
				tempDir, err := os.MkdirTemp("", "chatz_test")
				if err != nil {
					t.Fatalf("failed to create temp dir: %v", err)
				}

				configContent := `[]
PROVIDER=slack
TOKEN=test-token`

				configPath := filepath.Join(tempDir, ".chatz.ini")
				err = os.WriteFile(configPath, []byte(configContent), 0644)
				if err != nil {
					t.Fatalf("failed to write config file: %v", err)
				}

				originalHome := os.Getenv("HOME")
				os.Setenv("HOME", tempDir)

				return func() {
					os.Setenv("HOME", originalHome)
					os.RemoveAll(tempDir)
				}
			},
			expectError: true, // Empty section name in INI format is invalid
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanup := tt.setupFunc()
			defer cleanup()

			result, err := LoadEnv(tt.profile, tt.fromEnv)

			if tt.expectError {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if result == nil {
				t.Error("expected config but got nil")
				return
			}

			// Verify the config was loaded correctly
			if tt.fromEnv {
				if result.Provider != "slack" {
					t.Errorf("expected provider 'slack', got '%s'", result.Provider)
				}
				if result.Token != "test-token" {
					t.Errorf("expected token 'test-token', got '%s'", result.Token)
				}
			} else if tt.profile == "testprofile" {
				if result.Provider != "telegram" {
					t.Errorf("expected provider 'telegram', got '%s'", result.Provider)
				}
				if result.ChatId != "123456789" {
					t.Errorf("expected chat_id '123456789', got '%s'", result.ChatId)
				}
			}
		})
	}
}

func TestLoadEnvFromSystemEnv(t *testing.T) {
	// Test loading from environment variables
	envVars := map[string]string{
		"PROVIDER":          "discord",
		"WEB_HOOK_URL":      "https://discord.com/api/webhooks/test",
		"TOKEN":             "discord-token",
		"CHANNEL_ID":        "discord-channel",
		"CHAT_ID":           "telegram-chat",
		"CONNECTION_URL":    "redis://localhost:6379",
		"SMTP_HOST":         "smtp.gmail.com",
		"SMTP_PORT":         "587",
		"SMTP_USE_TLS":      "false",
		"SMTP_USE_STARTTLS": "true",
		"SMTP_USER":         "test@example.com",
		"SMTP_PASSWORD":     "test-password",
		"SMTP_SUBJECT":      "Test Subject",
		"SMTP_FROM":         "from@example.com",
		"SMTP_TO":           "to@example.com",
		"GOTIFY_URL":        "https://gotify.example.com",
		"GOTIFY_TOKEN":      "gotify-token",
		"GOTIFY_TITLE":      "Test Title",
		"GOTIFY_PRIORITY":   "5",
	}

	// Backup original environment
	originalEnv := make(map[string]string)
	for key := range envVars {
		if original, exists := os.LookupEnv(key); exists {
			originalEnv[key] = original
		}
	}

	// Set test environment
	for key, value := range envVars {
		os.Setenv(key, value)
	}

	// Cleanup function
	defer func() {
		for key := range envVars {
			if original, exists := originalEnv[key]; exists {
				os.Setenv(key, original)
			} else {
				os.Unsetenv(key)
			}
		}
	}()

	result, err := loadEnvFromSystemEnv()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	// Verify all fields are set correctly
	expected := &config.Config{
		Provider:       "discord",
		WebHookURL:     "https://discord.com/api/webhooks/test",
		Token:          "discord-token",
		ChannelId:      "discord-channel",
		ChatId:         "telegram-chat",
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

	if *result != *expected {
		t.Errorf("config mismatch.\nExpected: %+v\nGot: %+v", expected, result)
	}
}
