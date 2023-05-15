package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Set GITHUB_ACTIONS environment variable to simulate running in GitHub Actions
	os.Setenv("GITHUB_ACTIONS", "true")

	// Run the tests
	exitCode := m.Run()

	// Clean up
	os.Unsetenv("GITHUB_ACTIONS")

	// Exit with the appropriate exit code
	os.Exit(exitCode)
}

func TestIsRunningInGitHubActions(t *testing.T) {
	// Set GITHUB_ACTIONS environment variable to simulate running in GitHub Actions
	os.Setenv("GITHUB_ACTIONS", "true")

	// Run the main function
	main()

	// Add your assertions based on the expected output
	// Here, we assume that the expected output is "IN Github Action"

	// Clean up
	os.Unsetenv("GITHUB_ACTIONS")
}

func TestIsRunningInGitHubActionsAndSkip(t *testing.T) {
	// Set GITHUB_ACTIONS environment variable to simulate running in GitHub Actions
	os.Setenv("GITHUB_ACTIONS", "true")

	if os.Getenv("GITHUB_ACTIONS") == "true" {
		t.SkipNow()
	}

	// Run the main function
	main()

	// Add your assertions based on the expected output
	// Here, we assume that the expected output is "IN Github Action"

	// Clean up
	os.Unsetenv("GITHUB_ACTIONS")
}

func TestIsNotRunningInGitHubActions(t *testing.T) {
	// Set GITHUB_ACTIONS environment variable to simulate not running in GitHub Actions
	os.Setenv("GITHUB_ACTIONS", "")

	// Run the main function
	main()

	// Add your assertions based on the expected output
	// Here, we assume that the expected output is "NOT Github Action"

	// Clean up
	os.Unsetenv("GITHUB_ACTIONS")
}
