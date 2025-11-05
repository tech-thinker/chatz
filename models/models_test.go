package models

import (
	"encoding/json"
	"testing"
)

func TestOption_Struct(t *testing.T) {
	// Test Option struct with nil pointers (zero values)
	var option Option

	if option.Priority != nil {
		t.Errorf("expected Priority to be nil, got %v", option.Priority)
	}
	if option.Title != nil {
		t.Errorf("expected Title to be nil, got %v", option.Title)
	}
	if option.Subject != nil {
		t.Errorf("expected Subject to be nil, got %v", option.Subject)
	}
}

func TestOption_WithValues(t *testing.T) {
	// Test Option struct with actual values
	priority := 5
	title := "Test Title"
	subject := "Test Subject"

	option := Option{
		Priority: &priority,
		Title:    &title,
		Subject:  &subject,
	}

	if option.Priority == nil || *option.Priority != 5 {
		t.Errorf("expected Priority to be 5, got %v", option.Priority)
	}
	if option.Title == nil || *option.Title != "Test Title" {
		t.Errorf("expected Title to be 'Test Title', got %v", option.Title)
	}
	if option.Subject == nil || *option.Subject != "Test Subject" {
		t.Errorf("expected Subject to be 'Test Subject', got %v", option.Subject)
	}
}

func TestOption_JSONMarshal(t *testing.T) {
	tests := []struct {
		name     string
		option   Option
		expected string
	}{
		{
			name:     "empty option",
			option:   Option{},
			expected: `{"priority":null,"title":null,"subject":null}`,
		},
		{
			name: "option with values",
			option: Option{
				Priority: func() *int { i := 10; return &i }(),
				Title:    func() *string { s := "Hello"; return &s }(),
				Subject:  func() *string { s := "World"; return &s }(),
			},
			expected: `{"priority":10,"title":"Hello","subject":"World"}`,
		},
		{
			name: "option with some nil values",
			option: Option{
				Priority: func() *int { i := 1; return &i }(),
				Title:    nil,
				Subject:  func() *string { s := "Test"; return &s }(),
			},
			expected: `{"priority":1,"title":null,"subject":"Test"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.option)
			if err != nil {
				t.Errorf("failed to marshal option: %v", err)
				return
			}

			if string(data) != tt.expected {
				t.Errorf("expected JSON '%s', got '%s'", tt.expected, string(data))
			}
		})
	}
}

func TestOption_JSONUnmarshal(t *testing.T) {
	tests := []struct {
		name     string
		jsonStr  string
		expected Option
	}{
		{
			name:    "empty json",
			jsonStr: `{"priority":null,"title":null,"subject":null}`,
			expected: Option{
				Priority: nil,
				Title:    nil,
				Subject:  nil,
			},
		},
		{
			name:    "json with values",
			jsonStr: `{"priority":7,"title":"Test Title","subject":"Test Subject"}`,
			expected: Option{
				Priority: func() *int { i := 7; return &i }(),
				Title:    func() *string { s := "Test Title"; return &s }(),
				Subject:  func() *string { s := "Test Subject"; return &s }(),
			},
		},
		{
			name:    "json with partial values",
			jsonStr: `{"priority":3,"title":null,"subject":"Only Subject"}`,
			expected: Option{
				Priority: func() *int { i := 3; return &i }(),
				Title:    nil,
				Subject:  func() *string { s := "Only Subject"; return &s }(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var option Option
			err := json.Unmarshal([]byte(tt.jsonStr), &option)
			if err != nil {
				t.Errorf("failed to unmarshal JSON: %v", err)
				return
			}

			// Compare Priority
			if tt.expected.Priority == nil && option.Priority != nil {
				t.Errorf("expected Priority to be nil, got %v", option.Priority)
			}
			if tt.expected.Priority != nil && (option.Priority == nil || *option.Priority != *tt.expected.Priority) {
				t.Errorf("expected Priority %v, got %v", tt.expected.Priority, option.Priority)
			}

			// Compare Title
			if tt.expected.Title == nil && option.Title != nil {
				t.Errorf("expected Title to be nil, got %v", option.Title)
			}
			if tt.expected.Title != nil && (option.Title == nil || *option.Title != *tt.expected.Title) {
				t.Errorf("expected Title %v, got %v", tt.expected.Title, option.Title)
			}

			// Compare Subject
			if tt.expected.Subject == nil && option.Subject != nil {
				t.Errorf("expected Subject to be nil, got %v", option.Subject)
			}
			if tt.expected.Subject != nil && (option.Subject == nil || *option.Subject != *tt.expected.Subject) {
				t.Errorf("expected Subject %v, got %v", tt.expected.Subject, option.Subject)
			}
		})
	}
}

func TestOption_PointerHandling(t *testing.T) {
	// Test proper handling of pointer fields
	var option Option

	// Initially all should be nil
	if option.Priority != nil || option.Title != nil || option.Subject != nil {
		t.Error("expected all fields to be initially nil")
	}

	// Set values using pointers
	priorityVal := 8
	titleVal := "Pointer Test"
	subjectVal := "Pointer Subject"

	option.Priority = &priorityVal
	option.Title = &titleVal
	option.Subject = &subjectVal

	// Verify values
	if option.Priority == nil || *option.Priority != 8 {
		t.Errorf("expected Priority to be 8, got %v", option.Priority)
	}
	if option.Title == nil || *option.Title != "Pointer Test" {
		t.Errorf("expected Title to be 'Pointer Test', got %v", option.Title)
	}
	if option.Subject == nil || *option.Subject != "Pointer Subject" {
		t.Errorf("expected Subject to be 'Pointer Subject', got %v", option.Subject)
	}

	// Test modifying through pointers
	*option.Priority = 9
	*option.Title = "Modified Title"

	if *option.Priority != 9 {
		t.Errorf("expected modified Priority to be 9, got %d", *option.Priority)
	}
	if *option.Title != "Modified Title" {
		t.Errorf("expected modified Title to be 'Modified Title', got '%s'", *option.Title)
	}
}

func TestOption_JSONEdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		jsonStr     string
		expectError bool
		description string
	}{
		{
			name:        "malformed json",
			jsonStr:     `{"priority": 5, "title": "test"`,
			expectError: true,
			description: "JSON unmarshaling should fail with malformed JSON",
		},
		{
			name:        "invalid priority type",
			jsonStr:     `{"priority": "not-a-number", "title": "test"}`,
			expectError: true, // JSON unmarshaling fails on type mismatch for int fields
			description: "JSON unmarshaling fails when trying to unmarshal string into int field",
		},
		{
			name:        "null values",
			jsonStr:     `{"priority": null, "title": null, "subject": null}`,
			expectError: false,
			description: "Null values should be handled correctly",
		},
		{
			name:        "empty json object",
			jsonStr:     `{}`,
			expectError: false,
			description: "Empty JSON object should work",
		},
		{
			name:        "json with extra fields",
			jsonStr:     `{"priority": 1, "title": "test", "subject": "subject", "extra_field": "ignored"}`,
			expectError: false,
			description: "Extra fields in JSON should be ignored",
		},
		{
			name:        "unicode characters",
			jsonStr:     `{"title": "Test with émojis 🎉 and spëcial chärs"}`,
			expectError: false,
			description: "Unicode characters should be handled correctly",
		},
		{
			name:        "escaped characters",
			jsonStr:     `{"title": "Test with \"quotes\" and \\backslashes\\"}`,
			expectError: false,
			description: "Escaped characters should be handled correctly",
		},
		{
			name:        "very large numbers",
			jsonStr:     `{"priority": 999999999999}`,
			expectError: false,
			description: "Large numbers should be handled",
		},
		{
			name:        "negative numbers",
			jsonStr:     `{"priority": -5}`,
			expectError: false,
			description: "Negative numbers should be accepted",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var option Option
			err := json.Unmarshal([]byte(tt.jsonStr), &option)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error for %s, but got none", tt.description)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error for %s: %v", tt.description, err)
				}
			}
		})
	}
}

func TestOption_ExtremeValues(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() Option
		validate func(t *testing.T, option Option)
	}{
		{
			name: "maximum int value",
			setup: func() Option {
				maxInt := 9223372036854775807 // Max int64
				return Option{Priority: &maxInt}
			},
			validate: func(t *testing.T, option Option) {
				if option.Priority == nil || *option.Priority != 9223372036854775807 {
					t.Errorf("expected max int value, got %v", option.Priority)
				}
			},
		},
		{
			name: "minimum int value",
			setup: func() Option {
				minInt := -9223372036854775808 // Min int64
				return Option{Priority: &minInt}
			},
			validate: func(t *testing.T, option Option) {
				if option.Priority == nil || *option.Priority != -9223372036854775808 {
					t.Errorf("expected min int value, got %v", option.Priority)
				}
			},
		},
		{
			name: "zero priority",
			setup: func() Option {
				zero := 0
				return Option{Priority: &zero}
			},
			validate: func(t *testing.T, option Option) {
				if option.Priority == nil || *option.Priority != 0 {
					t.Errorf("expected zero priority, got %v", option.Priority)
				}
			},
		},
		{
			name: "empty strings",
			setup: func() Option {
				empty := ""
				return Option{
					Title:   &empty,
					Subject: &empty,
				}
			},
			validate: func(t *testing.T, option Option) {
				if option.Title == nil || *option.Title != "" {
					t.Errorf("expected empty title, got %v", option.Title)
				}
				if option.Subject == nil || *option.Subject != "" {
					t.Errorf("expected empty subject, got %v", option.Subject)
				}
			},
		},
		{
			name: "strings with only whitespace",
			setup: func() Option {
				whitespace := "   \t\n  "
				return Option{
					Title:   &whitespace,
					Subject: &whitespace,
				}
			},
			validate: func(t *testing.T, option Option) {
				expected := "   \t\n  "
				if option.Title == nil || *option.Title != expected {
					t.Errorf("expected whitespace title, got %v", option.Title)
				}
				if option.Subject == nil || *option.Subject != expected {
					t.Errorf("expected whitespace subject, got %v", option.Subject)
				}
			},
		},
		{
			name: "very long strings",
			setup: func() Option {
				longString := string(make([]byte, 100000)) // 100KB string
				for i := range longString {
					longString = longString[:i] + "a" + longString[i+1:]
				}
				return Option{
					Title:   &longString,
					Subject: &longString,
				}
			},
			validate: func(t *testing.T, option Option) {
				if option.Title == nil || len(*option.Title) != 100000 {
					t.Errorf("expected 100KB title, got length %d", len(*option.Title))
				}
				if option.Subject == nil || len(*option.Subject) != 100000 {
					t.Errorf("expected 100KB subject, got length %d", len(*option.Subject))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			option := tt.setup()
			tt.validate(t, option)
		})
	}
}
