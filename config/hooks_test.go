// Copyright © 2026 Ping Identity Corporation

package config_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/pingidentity/pingfederate-go-client/v1300/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The URLs below deliberately pair a non-http(s) scheme with a loopback host: browser.Open
// rejects the scheme during validation, before it could otherwise shell out to open a real
// browser window, so these tests do NOT open browsers (mirroring the pingcli test approach).
// Loopback is used rather than a domain so nothing outside the machine is referenced.
const (
	testAuthURL         = "ftp://127.0.0.1/authorize"
	testVerificationURI = "ftp://127.0.0.1/device"
	testDeviceUserCode  = "ABCD-1234"
)

func TestDefaultAuthorizationCodeBrowserHandlerTo_WritesToProvidedWriter(t *testing.T) {
	var buf bytes.Buffer

	err := config.DefaultAuthorizationCodeBrowserHandlerTo(&buf)(testAuthURL)
	require.NoError(t, err)

	out := buf.String()
	assert.Contains(t, out, "Opening browser for authorization")
	assert.Contains(t, out, testAuthURL)
	assert.Contains(t, out, "Waiting for authorization callback")
}

func TestDefaultAuthorizationCodeBrowserHandlerTo_StdoutReproducesInteractiveUX(t *testing.T) {
	out := captureStdout(t, func() {
		err := config.DefaultAuthorizationCodeBrowserHandlerTo(os.Stdout)(testAuthURL)
		require.NoError(t, err)
	})

	assert.Contains(t, out, "Opening browser for authorization")
	assert.Contains(t, out, testAuthURL)
	assert.Contains(t, out, "Waiting for authorization callback")
}

func TestDefaultAuthorizationCodeBrowserHandler_NilOutputIsSilent(t *testing.T) {
	out := captureStdout(t, func() {
		err := config.DefaultAuthorizationCodeBrowserHandler(testAuthURL)
		require.NoError(t, err)
	})

	assert.Empty(t, out, "the SDK must not print to stdout by default")
}

func TestDefaultDeviceCodePromptHandlerTo_WritesToProvidedWriter(t *testing.T) {
	var buf bytes.Buffer

	err := config.DefaultDeviceCodePromptHandlerTo(&buf)(testVerificationURI, testDeviceUserCode)
	require.NoError(t, err)

	// Assert only on content common to both the browser-available and no-browser branches, since
	// browser.CanOpen() is environment-dependent.
	out := buf.String()
	assert.Contains(t, out, "Device Authorization Required")
	assert.Contains(t, out, testVerificationURI)
	assert.Contains(t, out, testDeviceUserCode)
	assert.Contains(t, out, "Waiting for authorization")
}

func TestDefaultDeviceCodePromptHandlerTo_StdoutReproducesInteractiveUX(t *testing.T) {
	out := captureStdout(t, func() {
		err := config.DefaultDeviceCodePromptHandlerTo(os.Stdout)(testVerificationURI, testDeviceUserCode)
		require.NoError(t, err)
	})

	// Assert only on content common to both the browser-available and no-browser branches, since
	// browser.CanOpen() is environment-dependent.
	assert.Contains(t, out, "Device Authorization Required")
	assert.Contains(t, out, testVerificationURI)
	assert.Contains(t, out, testDeviceUserCode)
	assert.Contains(t, out, "Waiting for authorization")
}

func TestDefaultDeviceCodePromptHandler_NilOutputIsSilent(t *testing.T) {
	out := captureStdout(t, func() {
		err := config.DefaultDeviceCodePromptHandler(testVerificationURI, testDeviceUserCode)
		require.NoError(t, err)
	})

	assert.Empty(t, out, "the SDK must not print to stdout by default")
}

// captureStdout redirects os.Stdout for the duration of fn and returns everything written to it.
// Not safe to run in parallel with other tests that read or write os.Stdout, so callers must not
// mark their test t.Parallel().
func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	r, w, err := os.Pipe()
	require.NoError(t, err)

	original := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = original }()

	outChan := make(chan string, 1)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		outChan <- buf.String()
	}()

	fn()

	require.NoError(t, w.Close())
	return <-outChan
}
