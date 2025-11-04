package providers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/tech-thinker/chatz/config"
	"github.com/tech-thinker/chatz/models"
)

// MockTransport allows us to intercept HTTP calls for testing
type MockTransport struct {
	RoundTripFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(req)
}

func TestSlackProvider_Post(t *testing.T) {
	tests := []struct {
		name               string
		message            string
		mockResponse       interface{} // Changed to interface{} to allow strings for malformed JSON
		mockStatusCode     int
		expectError        bool
		mockTransportError error // For network-level errors
	}{
		{
			name:    "successful post",
			message: "test message",
			mockResponse: map[string]interface{}{
				"ok":      true,
				"channel": "C1234567890",
				"ts":      "1503435956.000247",
			},
			mockStatusCode: http.StatusOK,
			expectError:    false,
		},
		{
			name:    "slack api error",
			message: "test message",
			mockResponse: map[string]interface{}{
				"ok":    false,
				"error": "invalid_auth",
			},
			mockStatusCode: http.StatusOK,
			expectError:    false, // HTTP call succeeds, but API returns error
		},
		{
			name:    "http server error",
			message: "test message",
			mockResponse: map[string]interface{}{
				"ok":    false,
				"error": "server_error",
			},
			mockStatusCode: http.StatusInternalServerError,
			expectError:    false, // HTTP client doesn't fail, just returns error response
		},
		{
			name:           "malformed json response",
			message:        "test message",
			mockResponse:   "{invalid json",
			mockStatusCode: http.StatusOK,
			expectError:    false, // Slack provider just returns the response as string
		},
		{
			name:               "network connection error",
			message:            "test message",
			mockTransportError: &http.ProtocolError{ErrorString: "connection refused"},
			expectError:        true,
		},
		{
			name:           "empty message",
			message:        "",
			mockResponse:   map[string]interface{}{"ok": true},
			mockStatusCode: http.StatusOK,
			expectError:    false, // Empty message should still work
		},
		{
			name:           "very long message",
			message:        string(make([]byte, 10000)), // 10KB message
			mockResponse:   map[string]interface{}{"ok": true},
			mockStatusCode: http.StatusOK,
			expectError:    false,
		},
		{
			name:           "message with special characters",
			message:        "test message with émojis 🎉 and spëcial chärs",
			mockResponse:   map[string]interface{}{"ok": true},
			mockStatusCode: http.StatusOK,
			expectError:    false,
		},
		{
			name:           "message with newlines and tabs",
			message:        "line1\nline2\tindented",
			mockResponse:   map[string]interface{}{"ok": true},
			mockStatusCode: http.StatusOK,
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock HTTP client
			mockTransport := &MockTransport{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					// Return network error if specified
					if tt.mockTransportError != nil {
						return nil, tt.mockTransportError
					}

					// Verify request for non-error cases
					if tt.name != "network connection error" {
						if req.Method != "POST" {
							t.Errorf("expected POST request, got %s", req.Method)
						}
						if req.Header.Get("Authorization") != "Bearer test-token" {
							t.Errorf("expected Bearer test-token, got %s", req.Header.Get("Authorization"))
						}
						if req.Header.Get("Content-Type") != "application/json" {
							t.Errorf("expected application/json, got %s", req.Header.Get("Content-Type"))
						}
					}

					// Create mock response body
					var body bytes.Buffer
					if tt.mockResponse != nil {
						switch resp := tt.mockResponse.(type) {
						case map[string]interface{}:
							json.NewEncoder(&body).Encode(resp)
						case string:
							body.WriteString(resp) // For malformed JSON
						}
					}

					return &http.Response{
						StatusCode: tt.mockStatusCode,
						Body:       io.NopCloser(&body),
						Header:     make(http.Header),
					}, nil
				},
			}

			// Temporarily replace the default HTTP client
			originalClient := http.DefaultClient
			http.DefaultClient = &http.Client{Transport: mockTransport}
			defer func() { http.DefaultClient = originalClient }()

			// Create provider with test config
			cfg := &config.Config{
				Token:     "test-token",
				ChannelId: "C1234567890",
			}
			provider := &SlackProvider{config: cfg}

			option := models.Option{}
			result, err := provider.Post(tt.message, option)

			if tt.expectError {
				if err == nil {
					t.Error("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result == nil {
					t.Error("expected result but got nil")
				}
			}
		})
	}
}

func TestSlackProvider_Reply(t *testing.T) {
	// Create a mock HTTP client that returns an error
	mockTransport := &MockTransport{
		RoundTripFunc: func(req *http.Request) (*http.Response, error) {
			// Verify request method and headers
			if req.Method != "POST" {
				t.Errorf("expected POST request, got %s", req.Method)
			}
			if req.Header.Get("Authorization") != "Bearer test-token" {
				t.Errorf("expected Bearer test-token, got %s", req.Header.Get("Authorization"))
			}

			// Return a mock error response
			body := bytes.NewBufferString(`{"ok":false,"error":"invalid_auth"}`)
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(body),
				Header:     make(http.Header),
			}, nil
		},
	}

	// Temporarily replace the default HTTP client
	originalClient := http.DefaultClient
	http.DefaultClient = &http.Client{Transport: mockTransport}
	defer func() { http.DefaultClient = originalClient }()

	cfg := &config.Config{
		Token:     "test-token",
		ChannelId: "C1234567890",
	}
	provider := &SlackProvider{config: cfg}

	option := models.Option{}
	result, err := provider.Reply("thread-123", "reply message", option)

	// Should not error on HTTP call, but API returns error response
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result == nil {
		t.Error("expected result, got nil")
	}
}

func TestNewProvider(t *testing.T) {
	tests := []struct {
		name        string
		provider    string
		expectError bool
	}{
		{
			name:        "valid slack provider",
			provider:    "slack",
			expectError: false, // Slack provider should be registered
		},
		{
			name:        "valid discord provider",
			provider:    "discord",
			expectError: false,
		},
		{
			name:        "invalid provider",
			provider:    "invalid",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &config.Config{Provider: tt.provider}
			provider, err := NewProvider(cfg)

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

			if provider == nil {
				t.Error("expected provider but got nil")
			}
		})
	}
}

func TestSlackProvider_Post_WithOptions(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		option   models.Option
		expected map[string]interface{}
	}{
		{
			name:    "with title option",
			message: "test message",
			option: models.Option{
				Title: stringPtr("Custom Title"),
			},
			expected: map[string]interface{}{
				"ok": true,
			},
		},
		{
			name:    "with subject option",
			message: "test message",
			option: models.Option{
				Subject: stringPtr("Custom Subject"),
			},
			expected: map[string]interface{}{
				"ok": true,
			},
		},
		{
			name:    "with priority option",
			message: "test message",
			option: models.Option{
				Priority: intPtr(5),
			},
			expected: map[string]interface{}{
				"ok": true,
			},
		},
		{
			name:    "with all options",
			message: "test message",
			option: models.Option{
				Title:    stringPtr("Title"),
				Subject:  stringPtr("Subject"),
				Priority: intPtr(8),
			},
			expected: map[string]interface{}{
				"ok": true,
			},
		},
		{
			name:    "with nil options",
			message: "test message",
			option:  models.Option{}, // All nil
			expected: map[string]interface{}{
				"ok": true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTransport := &MockTransport{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					// Verify request method and headers
					if req.Method != "POST" {
						t.Errorf("expected POST request, got %s", req.Method)
					}

					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(bytes.NewBufferString(`{"ok":true}`)),
						Header:     make(http.Header),
					}, nil
				},
			}

			originalClient := http.DefaultClient
			http.DefaultClient = &http.Client{Transport: mockTransport}
			defer func() { http.DefaultClient = originalClient }()

			cfg := &config.Config{
				Token:     "test-token",
				ChannelId: "C1234567890",
			}
			provider := &SlackProvider{config: cfg}

			result, err := provider.Post(tt.message, tt.option)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result == nil {
				t.Error("expected result but got nil")
			}
		})
	}
}

func TestSlackProvider_Setup(t *testing.T) {
	tests := []struct {
		name        string
		config      *config.Config
		expectError bool
	}{
		{
			name: "valid slack config",
			config: &config.Config{
				Token:     "test-token",
				ChannelId: "C1234567890",
			},
			expectError: false,
		},
		{
			name: "missing token",
			config: &config.Config{
				ChannelId: "C1234567890",
			},
			expectError: false, // Slack provider doesn't validate config in setup
		},
		{
			name: "missing channel",
			config: &config.Config{
				Token: "test-token",
			},
			expectError: false, // Slack provider doesn't validate config in setup
		},
		{
			name:        "nil config",
			config:      nil,
			expectError: false, // Slack provider doesn't validate config in setup
		},
		{
			name:        "empty config",
			config:      &config.Config{},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider := &SlackProvider{}
			err := provider.setup(tt.config)

			if tt.expectError && err == nil {
				t.Error("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			// Verify config was set (unless nil config)
			if !tt.expectError && tt.config != nil && provider.config != tt.config {
				t.Error("config was not set properly")
			}
		})
	}
}

// Helper functions for creating pointers
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
