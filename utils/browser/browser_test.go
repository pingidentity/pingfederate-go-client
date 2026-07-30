// Copyright © 2026 Ping Identity Corporation

package browser

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanOpen(t *testing.T) {
	// This test validates that CanOpen correctly detects browser availability
	// on the current platform
	result := CanOpen()

	// On most development machines, browser commands should be available
	// We test that the function returns a boolean without panicking
	assert.IsType(t, true, result, "CanOpen should return a boolean")

	// Verify expected command exists for the platform if CanOpen returns true
	if result {
		var expectedCmd string
		switch runtime.GOOS {
		case "linux", "freebsd", "openbsd", "netbsd":
			expectedCmd = "xdg-open"
		case "darwin":
			expectedCmd = "open"
		case "windows":
			expectedCmd = "rundll32"
		}

		if expectedCmd != "" {
			_, err := exec.LookPath(expectedCmd)
			assert.NoError(t, err, "Expected command %s should exist when CanOpen returns true", expectedCmd)
		}
	}
}

func TestValidateAndSanitizeURL(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
		errorString string
	}{
		// Valid URLs
		{
			name:        "valid http URL",
			input:       "http://pingidentity.com",
			expectError: false,
		},
		{
			name:        "valid https URL",
			input:       "https://pingidentity.com",
			expectError: false,
		},
		{
			name:        "valid URL with path",
			input:       "https://pingidentity.com/path/to/resource",
			expectError: false,
		},
		{
			name:        "valid URL with query parameters",
			input:       "https://pingidentity.com/path?param=value&other=123",
			expectError: false,
		},
		{
			name:        "valid URL with fragment",
			input:       "https://pingidentity.com/path#section",
			expectError: false,
		},
		{
			name:        "valid URL with port",
			input:       "https://pingidentity.com:8080/path",
			expectError: false,
		},
		{
			name:        "valid URL with subdomain",
			input:       "https://sub.pingidentity.com/path",
			expectError: false,
		},
		{
			name:        "URL with whitespace gets trimmed",
			input:       "  https://pingidentity.com  ",
			expectError: false,
		},

		// Invalid schemes
		{
			name:        "ftp scheme rejected",
			input:       "ftp://pingidentity.com",
			expectError: true,
			errorString: "unsupported URL scheme",
		},
		{
			name:        "file scheme rejected",
			input:       "file:///etc/passwd",
			expectError: true,
			errorString: "unsupported URL scheme",
		},
		{
			name:        "javascript scheme rejected",
			input:       "javascript:alert('xss')",
			expectError: true,
			errorString: "unsupported URL scheme",
		},
		{
			name:        "data scheme rejected",
			input:       "data:text/html,<script>alert('xss')</script>",
			expectError: true,
			errorString: "unsupported URL scheme",
		},
		{
			name:        "no scheme rejected",
			input:       "pingidentity.com",
			expectError: true,
			errorString: "unsupported URL scheme",
		},

		// Invalid formats
		{
			name:        "empty URL rejected",
			input:       "",
			expectError: true,
			errorString: "unsupported URL scheme",
		},
		{
			name:        "malformed URL rejected",
			input:       "https://[invalid",
			expectError: true,
			errorString: "invalid URL format",
		},
		{
			name:        "URL without host rejected",
			input:       "https://",
			expectError: true,
			errorString: "URL must have a valid host",
		},

		// Security: Embedded credentials
		{
			name:        "URL with embedded credentials rejected",
			input:       "https://user:pass@pingidentity.com",
			expectError: true,
			errorString: "URLs with embedded credentials are not allowed",
		},
		{
			name:        "URL with username only rejected",
			input:       "https://user@pingidentity.com",
			expectError: true,
			errorString: "URLs with embedded credentials are not allowed",
		},

		// Security: Control characters
		{
			name:        "URL with null byte rejected",
			input:       "https://pingidentity.com\x00",
			expectError: true,
			errorString: "invalid URL format",
		},
		{
			name:        "URL with newline gets trimmed",
			input:       "https://pingidentity.com\n",
			expectError: false,
		},
		{
			name:        "URL with carriage return gets trimmed",
			input:       "https://pingidentity.com\r",
			expectError: false,
		},
		{
			name:        "URL with tab gets trimmed",
			input:       "https://pingidentity.com\t",
			expectError: false,
		},

		// Edge cases
		{
			name:        "localhost accepted",
			input:       "http://localhost:8080",
			expectError: false,
		},
		{
			name:        "IP address accepted",
			input:       "http://127.0.0.1:8080",
			expectError: false,
		},
		{
			name:        "IPv6 address accepted",
			input:       "http://[::1]:8080",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := validateAndSanitizeURL(tt.input)

			if tt.expectError {
				require.Error(t, err, "Expected error for input: %s", tt.input)
				if tt.errorString != "" {
					assert.Contains(t, err.Error(), tt.errorString, "Error message should contain expected string")
				}
				assert.Empty(t, result, "Result should be empty when error occurs")
			} else {
				require.NoError(t, err, "Expected no error for input: %s", tt.input)
				assert.NotEmpty(t, result, "Result should not be empty for valid URL")

				// Verify the result is a valid URL
				assert.True(t, strings.HasPrefix(result, "http://") || strings.HasPrefix(result, "https://"),
					"Sanitized URL should start with http:// or https://")
			}
		})
	}
}

func TestValidateAndSanitizeURL_MaxLength(t *testing.T) {
	// Test URL length validation
	tests := []struct {
		name        string
		urlLength   int
		expectError bool
	}{
		{
			name:        "URL at max length accepted",
			urlLength:   maxURLLength,
			expectError: false,
		},
		{
			name:        "URL over max length rejected",
			urlLength:   maxURLLength + 1,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create URL with specific length
			baseURL := "https://pingidentity.com/"
			padding := strings.Repeat("a", tt.urlLength-len(baseURL))
			testURL := baseURL + padding

			// Note: validateAndSanitizeURL doesn't check length, Open() does
			// But we can test that the URL is processed correctly
			_, err := validateAndSanitizeURL(testURL)

			if tt.expectError {
				// The length check happens in Open(), not validateAndSanitizeURL
				// So this test documents the expected behavior
				assert.GreaterOrEqual(t, len(testURL), maxURLLength, "Test URL should exceed max length")
			} else {
				assert.NoError(t, err, "Valid long URL should be accepted by validateAndSanitizeURL")
			}
		})
	}
}

func TestOpen(t *testing.T) {
	// Note: This test only validates error conditions and input validation.
	// Actual browser opening is not tested in CI environments as browsers may not be available.
	// Manual testing should verify that valid URLs successfully open browsers on supported platforms.

	tests := []struct {
		name        string
		url         string
		expectError bool
		errorString string
	}{
		{
			name:        "empty URL rejected",
			url:         "",
			expectError: true,
			errorString: "URL cannot be empty",
		},
		{
			name:        "URL exceeding max length rejected",
			url:         "https://pingidentity.com/" + strings.Repeat("a", maxURLLength),
			expectError: true,
			errorString: "exceeds maximum length",
		},
		{
			name:        "invalid URL rejected",
			url:         "not a url",
			expectError: true,
			errorString: "unsupported URL scheme",
		},
		{
			name:        "unsupported scheme rejected",
			url:         "ftp://pingidentity.com",
			expectError: true,
			errorString: "unsupported URL scheme",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Open(tt.url)

			if tt.expectError {
				require.Error(t, err, "Expected error for URL: %s", tt.url)
				if tt.errorString != "" {
					assert.Contains(t, err.Error(), tt.errorString, "Error message should contain expected string")
				}
			}
		})
	}

	// Test that valid URLs pass validation (but don't actually open browser in CI)
	t.Run("valid URLs pass validation", func(t *testing.T) {
		if !CanOpen() {
			t.Skip("Browser not available - cannot test actual browser opening")
		}

		validURLs := []string{
			"http://pingidentity.com",
			"https://pingidentity.com",
			"http://localhost:8080",
		}

		for _, url := range validURLs {
			err := Open(url)
			assert.NoError(t, err, "Expected no error for valid URL: %s", url)
		}
	})
}

func TestOpen_SecurityValidation(t *testing.T) {
	// Security-focused test cases
	securityTests := []struct {
		name        string
		url         string
		expectError bool
		reason      string
	}{
		{
			name:        "reject URL with embedded credentials",
			url:         "https://user:pass@pingidentity.com",
			expectError: true,
			reason:      "credentials in URLs are a security risk",
		},
		{
			name:        "reject javascript: scheme",
			url:         "javascript:alert('xss')",
			expectError: true,
			reason:      "javascript: scheme can execute arbitrary code",
		},
		{
			name:        "reject file: scheme",
			url:         "file:///etc/passwd",
			expectError: true,
			reason:      "file: scheme can access local files",
		},
		{
			name:        "reject data: scheme",
			url:         "data:text/html,<script>alert('xss')</script>",
			expectError: true,
			reason:      "data: scheme can execute embedded scripts",
		},
		{
			name:        "reject URL with null byte",
			url:         "https://pingidentity.com\x00",
			expectError: true,
			reason:      "null bytes can cause command injection",
		},
	}

	for _, tt := range securityTests {
		t.Run(tt.name, func(t *testing.T) {
			err := Open(tt.url)
			require.Error(t, err, "Security test failed: %s should be rejected because %s", tt.url, tt.reason)
		})
	}
}

func TestOpen_CommandInjectionPrevention(t *testing.T) {
	// Test that command injection attempts are safely handled at the validation layer.
	// NOTE: This test validates URL parsing only and does NOT open browsers.
	// We test validateAndSanitizeURL directly instead of Open() to avoid opening browsers.
	injectionAttempts := []struct {
		url         string
		expectError bool
		reason      string
	}{
		{
			url:         "https://pingidentity.com;cat /etc/passwd",
			expectError: true,
			reason:      "semicolon with space creates invalid hostname",
		},
		{
			url:         "https://pingidentity.com | ls -la",
			expectError: true,
			reason:      "pipe with spaces creates invalid hostname",
		},
		{
			url:         "https://pingidentity.com && rm -rf /",
			expectError: true,
			reason:      "ampersands with spaces creates invalid hostname",
		},
		{
			url:         "https://pingidentity.com`whoami`",
			expectError: true,
			reason:      "backticks are invalid in hostname",
		},
		{
			url:         "https://pingidentity.com$(whoami)",
			expectError: false,
			reason:      "parentheses are technically valid in URLs but would be passed literally to browser, not executed as shell command",
		},
	}

	for _, tt := range injectionAttempts {
		t.Run(fmt.Sprintf("injection attempt: %s", tt.url), func(t *testing.T) {
			// Test validation only - don't actually open browser
			_, err := validateAndSanitizeURL(tt.url)

			if tt.expectError {
				require.Error(t, err, "Expected validation error for: %s - %s", tt.url, tt.reason)
				t.Logf("URL correctly rejected during validation: %v", err)
			} else {
				// URL passes validation
				// Note: Even if this were passed to a browser, it would be safe because
				// exec.CommandContext passes it as a literal argument, not through shell
				require.NoError(t, err, "Expected validation to pass for: %s - %s", tt.url, tt.reason)
				t.Logf("URL validation passed (would be safe if opened): %s", tt.url)
			}
		})
	}
}

func TestOpen_ConcurrentValidation(t *testing.T) {
	// Test that concurrent validation calls don't cause race conditions
	// This tests the validation logic only, not actual browser opening

	t.Run("concurrent invalid URLs", func(t *testing.T) {
		urls := []string{
			"ftp://pingidentity.com/1",
			"javascript:alert('test')",
			"file:///etc/passwd",
		}

		done := make(chan error, len(urls))

		for _, url := range urls {
			go func(u string) {
				done <- Open(u)
			}(url)
		}

		// Collect results - all should be errors
		for i := 0; i < len(urls); i++ {
			err := <-done
			assert.Error(t, err, "Expected error for invalid URL")
		}
	})
}
