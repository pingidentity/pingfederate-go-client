// Copyright © 2026 Ping Identity Corporation

// Package browser provides utilities for opening URLs in the system's default browser.
// This package can be used by applications that need to detect browser availability
// or open URLs as part of authentication flows or other user interactions.
package browser

import (
	"errors"
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)

const (
	// maxURLLength defines the maximum allowed URL length to prevent DoS
	maxURLLength = 2048
)

// CanOpen checks if we can open a browser in the current environment.
// Returns true if the required browser command exists for the current operating system.
// This is useful for determining whether to attempt automatic browser opening or
// to fall back to displaying instructions to the user.
func CanOpen() bool {
	// Check for the appropriate browser command for each platform
	switch runtime.GOOS {
	case "linux", "freebsd", "openbsd", "netbsd":
		// On Linux/Unix, check if xdg-open command exists
		// xdg-open handles display environment detection internally
		if _, err := exec.LookPath("xdg-open"); err != nil {
			return false
		}
		return true
	case "darwin":
		// macOS: check if 'open' command exists
		if _, err := exec.LookPath("open"); err != nil {
			return false
		}
		return true
	case "windows":
		// Windows: check if 'rundll32' command exists
		if _, err := exec.LookPath("rundll32"); err != nil {
			return false
		}
		return true
	default:
		return false
	}
}

// Open attempts to open a URL in the system's default browser.
// Returns an error if the URL is invalid, the platform is unsupported,
// or the browser command fails to execute.
//
// The URL is validated before opening to ensure it uses http or https scheme
// and does not contain embedded credentials or control characters.
//
// Security considerations:
//   - Only http and https URL schemes are allowed
//   - URLs are validated and sanitized before execution
//   - Commands are executed with proper argument separation to prevent injection
//   - The browser process runs independently and does not block the caller
func Open(urlStr string) error {
	// Validate input is not empty
	if urlStr == "" {
		return errors.New("URL cannot be empty")
	}

	// Check URL length to prevent DoS
	if len(urlStr) > maxURLLength {
		return fmt.Errorf("URL exceeds maximum length of %d characters", maxURLLength)
	}

	// Validate and sanitize URL
	sanitizedURL, err := validateAndSanitizeURL(urlStr)
	if err != nil {
		return err
	}

	// Execute platform-specific browser command with proper argument separation
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux", "freebsd", "openbsd", "netbsd":
		// #nosec G204 -- Command and arguments are fixed; URL is strictly validated and sanitized
		cmd = exec.Command("xdg-open", sanitizedURL)
	case "darwin":
		// #nosec G204 -- Command and arguments are fixed; URL is strictly validated and sanitized
		cmd = exec.Command("open", sanitizedURL)
	case "windows":
		// Use rundll32 to open URLs safely on Windows
		// Arguments are passed separately to prevent injection
		// #nosec G204 -- Command and arguments are fixed; URL is strictly validated and sanitized
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", sanitizedURL)
	default:
		return errors.New("unsupported platform")
	}

	// Start the command and don't wait for it to complete
	// Browser processes should run independently
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to open browser: %w", err)
	}

	// Launch a goroutine to clean up the process without blocking
	go func() {
		// Wait for the process to complete
		// Ignore error - browser process runs independently and we cannot
		// meaningfully handle errors after the main function returns
		_ = cmd.Wait()
	}()

	return nil
}

// validateAndSanitizeURL validates and sanitizes a URL string for safe browser opening.
// Returns the sanitized URL string or an error if validation fails.
func validateAndSanitizeURL(urlStr string) (string, error) {
	// Trim whitespace
	urlStr = strings.TrimSpace(urlStr)

	// Reject obvious shell metacharacters that are not valid in URLs
	// even though we don't invoke a shell, this adds defense-in-depth
	// against malformed inputs being passed to system handlers.
	forbiddenChars := []string{"|", ";", "`", "\n", "\r"}
	for _, ch := range forbiddenChars {
		if strings.Contains(urlStr, ch) {
			return "", errors.New("URL contains invalid characters")
		}
	}

	// Parse URL to validate structure
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("invalid URL format: %w", err)
	}

	// Validate scheme - only allow http and https
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return "", fmt.Errorf("unsupported URL scheme '%s': only http and https are allowed", parsedURL.Scheme)
	}

	// Validate that host is present and non-empty
	if parsedURL.Host == "" || parsedURL.Hostname() == "" {
		return "", errors.New("URL must have a valid host")
	}

	// Reject URLs with embedded credentials for security
	if parsedURL.User != nil {
		return "", errors.New("URLs with embedded credentials are not allowed")
	}

	// Reconstruct the URL from parsed components to ensure proper encoding
	// This prevents injection via malformed URLs
	// Ensure URL is absolute and well-formed
	if !parsedURL.IsAbs() {
		return "", errors.New("URL must be absolute")
	}

	// Re-build the URL string to ensure consistent encoding
	sanitizedURL := parsedURL.String()

	// Additional validation: check for suspicious patterns
	// Reject URLs with null bytes or other control characters
	for _, r := range sanitizedURL {
		if r < 32 || r == 127 {
			return "", errors.New("URL contains invalid control characters")
		}
	}

	return sanitizedURL, nil
}
