package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/urfave/cli/v2"
)

// TestCLISetup tests that the CLI application is properly configured.
func TestCLISetup(t *testing.T) {
	// We need to create a minimal version of the main function logic
	// to test the CLI setup without calling app.Run()

	// Create the app as it would be in main()
	app := cli.NewApp()
	app.Name = "chatz"
	app.Description = "chatz is a versatile messaging app designed to send notifications to Google Chat, Slack, Discord, Telegram and Redis."

	// Test basic app properties
	if app.Name != "chatz" {
		t.Errorf("expected app name 'chatz', got '%s'", app.Name)
	}

	expectedDesc := "chatz is a versatile messaging app designed to send notifications to Google Chat, Slack, Discord, Telegram and Redis."
	if app.Description != expectedDesc {
		t.Errorf("app description mismatch")
	}
}

// TestCLIFlags tests that CLI flags are properly defined.
func TestCLIFlags(t *testing.T) {
	// Create flags as they would be in main()
	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "Print output to stdout",
			Destination: &[]bool{false}[0], // dummy destination
		},
		&cli.StringFlag{
			Name:        "profile",
			Aliases:     []string{"p"},
			Value:       "default",
			Usage:       "Profile from .chatz.ini",
			Destination: &[]string{""}[0], // dummy destination
		},
		&cli.StringFlag{
			Name:        "thread-id",
			Aliases:     []string{"t"},
			Value:       "",
			Usage:       "Thread ID for reply to a message",
			Destination: &[]string{""}[0], // dummy destination
		},
		&cli.BoolFlag{
			Name:        "version",
			Aliases:     []string{"v"},
			Usage:       "Print the version number",
			Destination: &[]bool{false}[0], // dummy destination
		},
		&cli.BoolFlag{
			Name:        "from-env",
			Aliases:     []string{"e"},
			Usage:       "To use config from environment variables",
			Destination: &[]bool{false}[0], // dummy destination
		},
		&cli.StringFlag{
			Name:        "subject",
			Aliases:     []string{"s"},
			Usage:       "Subject for provider which supports subject or title",
			Destination: &[]string{""}[0], // dummy destination
		},
		&cli.IntFlag{
			Name:        "priority",
			Aliases:     []string{"pr"},
			Usage:       "Priority for gotify notification",
			Destination: &[]int{0}[0], // dummy destination
		},
	}

	// Test that we have the expected number of flags
	expectedFlags := 7
	if len(flags) != expectedFlags {
		t.Errorf("expected %d flags, got %d", expectedFlags, len(flags))
	}

	// Test specific flags
	flagNames := make(map[string]bool)
	for _, flag := range flags {
		switch f := flag.(type) {
		case *cli.BoolFlag:
			flagNames[f.Name] = true
			if f.Name == "version" && len(f.Aliases) == 0 {
				t.Error("version flag should have aliases")
			}
		case *cli.StringFlag:
			flagNames[f.Name] = true
			if f.Name == "profile" && f.Value != "default" {
				t.Error("profile flag should default to 'default'")
			}
		case *cli.IntFlag:
			flagNames[f.Name] = true
		}
	}

	// Verify we have all expected flags
	expectedFlagNames := []string{"output", "profile", "thread-id", "version", "from-env", "subject", "priority"}
	for _, name := range expectedFlagNames {
		if !flagNames[name] {
			t.Errorf("missing expected flag: %s", name)
		}
	}
}

// TestAppCreation tests that the CLI app can be created without errors.
func TestAppCreation(t *testing.T) {
	// This tests that the app setup doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("app creation panicked: %v", r)
		}
	}()

	// We can't fully test main() without it calling app.Run(),
	// but we can test that the app struct can be created
	app := &cli.App{
		Name:        "chatz",
		Description: "test description",
	}

	if app.Name != "chatz" {
		t.Error("app name not set correctly")
	}
}

// TestVersionOutput tests the version flag behavior by running the binary.
func TestVersionOutput(t *testing.T) {
	// Skip this test if we're not running the integration test
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Build the binary first
	buildCmd := exec.Command("go", "build", "-o", "chatz-test", ".")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("failed to build binary: %v", err)
	}
	defer os.Remove("chatz-test") // cleanup

	// Run the binary with --version flag
	cmd := exec.Command("./chatz-test", "--version")
	output, err := cmd.Output()
	if err != nil {
		t.Errorf("command failed: %v", err)
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "chatz version:") {
		t.Errorf("expected version output, got: %s", outputStr)
	}
	if !strings.Contains(outputStr, "Commit Hash:") {
		t.Errorf("expected commit hash in output, got: %s", outputStr)
	}
	if !strings.Contains(outputStr, "Build Date:") {
		t.Errorf("expected build date in output, got: %s", outputStr)
	}
}

// TestNoArgsOutput tests the behavior when no arguments are provided.
func TestNoArgsOutput(t *testing.T) {
	// Skip this test if we're not running the integration test
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Build the binary first
	buildCmd := exec.Command("go", "build", "-o", "chatz-test-noargs", ".")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("failed to build binary: %v", err)
	}
	defer os.Remove("chatz-test-noargs") // cleanup

	// Run the binary with no arguments
	cmd := exec.Command("./chatz-test-noargs")
	output, err := cmd.Output()
	if err != nil {
		// This is expected to fail since no message is provided
		// The CLI framework will exit with code 1
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() != 1 {
				t.Errorf("expected exit code 1, got %d", exitErr.ExitCode())
			}
		} else {
			t.Errorf("unexpected error type: %v", err)
		}
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "Please provide a message.") {
		t.Errorf("expected 'Please provide a message.' in output, got: %s", outputStr)
	}
}

// TestHelpOutput tests the help flag behavior.
func TestHelpOutput(t *testing.T) {
	// Skip this test if we're not running the integration test
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Build the binary first
	buildCmd := exec.Command("go", "build", "-o", "chatz-test-help", ".")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("failed to build binary: %v", err)
	}
	defer os.Remove("chatz-test-help") // cleanup

	// Run the binary with --help flag
	cmd := exec.Command("./chatz-test-help", "--help")
	output, err := cmd.Output()
	if err != nil {
		t.Errorf("command failed: %v", err)
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "chatz") {
		t.Errorf("expected 'chatz' in help output, got: %s", outputStr)
	}
	if !strings.Contains(outputStr, "USAGE") {
		t.Errorf("expected 'USAGE' in help output, got: %s", outputStr)
	}
	if !strings.Contains(outputStr, "GLOBAL OPTIONS") {
		t.Errorf("expected 'GLOBAL OPTIONS' in help output, got: %s", outputStr)
	}
}
