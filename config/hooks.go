// Copyright © 2026 Ping Identity Corporation

package config

import (
	"fmt"
	"io"
	"os"

	"github.com/pingidentity/pingfederate-go-client/v1300/utils/browser"
)

// AuthURLHandler is a function type that handles opening authorization URLs.
// It receives the authorization URL that the user must visit.
// Return an error if the URL cannot be handled.
type AuthURLHandler func(url string) error

// DeviceCodePromptHandler is a function type that handles displaying device code prompts to users.
// It receives the verification URI and user code that must be displayed to the user.
// Return an error if the prompt cannot be displayed.
type DeviceCodePromptHandler func(verificationURI, userCode string) error

// AuthResultPageData contains the data used to render the authorization result page.
// SDK clients can provide custom values to personalize the page while using the default template.
type AuthResultPageData struct {
	// ProjectName is displayed as the HTML page title (e.g., "PingFederate Admin Tools").
	ProjectName string
	// Heading is the main heading displayed on the page (e.g., "Authorization Success").
	Heading string
	// Message is the descriptive text displayed below the heading.
	Message string
}

// AuthorizationCode holds the configuration for the authorization_code grant type (with PKCE).
type AuthorizationCode struct {
	// AuthorizationCodeClientID is the OAuth2 client ID used for the authorization code flow.
	AuthorizationCodeClientID *string
	// AuthorizationCodeRedirectURI configures the local redirect (callback) server port and path.
	AuthorizationCodeRedirectURI AuthorizationCodeRedirectURI
	// AuthorizationCodeScopes are the OAuth2 scopes requested during the flow.
	AuthorizationCodeScopes *[]string
	// OnOpenBrowser is an optional handler for custom browser opening logic.
	// If set, this handler is called instead of automatically opening the system browser,
	// allowing consumers to implement custom flows such as headless operation or alternative UX.
	OnOpenBrowser AuthURLHandler
	// Output configures where the default browser-opening handler writes its progress messages.
	// It is honored only when OnOpenBrowser is nil; once a custom handler is set, that handler is
	// solely responsible for its own output, and Output is ignored. A nil Output causes the
	// default handler to write to os.Stdout, preserving the SDK's historical behavior. Set Output
	// to io.Discard to silence the default handler, or to any other io.Writer to capture or
	// redirect its messages.
	Output io.Writer
	// CustomPageDataSuccess contains the data to display on successful authentication.
	// If nil, default values are used. The SDK template is rendered with these values.
	CustomPageDataSuccess *AuthResultPageData
	// CustomPageDataError contains the data to display on authentication failure.
	// If nil, default values are used. The SDK template is rendered with these values.
	CustomPageDataError *AuthResultPageData
}

// AuthorizationCodeRedirectURI configures the local redirect (callback) server used by the
// authorization_code flow. Only the port and path are configurable; the host is always the
// loopback interface.
type AuthorizationCodeRedirectURI struct {
	Port string
	Path string
}

// DeviceCode holds the configuration for the device_code grant type (with PKCE).
type DeviceCode struct {
	// DeviceCodeClientID is the OAuth2 client ID used for the device code flow.
	DeviceCodeClientID *string
	// DeviceCodeScopes are the OAuth2 scopes requested during the flow.
	DeviceCodeScopes *[]string
	// OnDisplayPrompt is an optional handler for custom device code prompt display.
	// If set, this handler is called instead of the default console output, allowing
	// consumers to implement custom UX such as QR codes, notifications, or headless flows.
	OnDisplayPrompt DeviceCodePromptHandler
	// Output configures where the default device code prompt handler writes its progress
	// messages. It is honored only when OnDisplayPrompt is nil; once a custom handler is set,
	// that handler is solely responsible for its own output, and Output is ignored. A nil Output
	// causes the default handler to write to os.Stdout, preserving the SDK's historical behavior.
	// Set Output to io.Discard to silence the default handler, or to any other io.Writer to
	// capture or redirect its messages.
	Output io.Writer
}

// fprint, fprintf, and fprintln write best-effort progress messages to w, mirroring fmt.Fprint,
// fmt.Fprintf, and fmt.Fprintln respectively. Write failures are intentionally ignored: these are
// user-facing progress messages for an interactive login flow, and a failed write to a
// caller-supplied Output writer should not abort authentication.
func fprint(w io.Writer, a ...any) {
	_, _ = fmt.Fprint(w, a...)
}

func fprintf(w io.Writer, format string, a ...any) {
	_, _ = fmt.Fprintf(w, format, a...)
}

func fprintln(w io.Writer, a ...any) {
	_, _ = fmt.Fprintln(w, a...)
}

// DefaultAuthorizationCodeBrowserHandler is the default handler for opening the authorization URL.
// It attempts to open the system browser automatically and provides fallback instructions if that fails.
// This function implements the AuthURLHandler interface and provides a consistent UX pattern.
// Consumer projects can use this handler as a reference or directly in their own implementations.
// Its progress messages are written to os.Stdout; use DefaultAuthorizationCodeBrowserHandlerTo, or
// set AuthorizationCode.Output (via Configuration.WithAuthorizationCodeOutput), to redirect or
// silence this output.
func DefaultAuthorizationCodeBrowserHandler(authURL string) error {
	return DefaultAuthorizationCodeBrowserHandlerTo(nil)(authURL)
}

// DefaultAuthorizationCodeBrowserHandlerTo returns a browser-opening handler equivalent to
// DefaultAuthorizationCodeBrowserHandler, but that writes its progress messages to w instead of
// os.Stdout. If w is nil, the returned handler writes to os.Stdout, matching the behavior of
// DefaultAuthorizationCodeBrowserHandler. This allows consumers to redirect or silence (using
// io.Discard) the default handler's output without reimplementing its browser-opening logic.
func DefaultAuthorizationCodeBrowserHandlerTo(w io.Writer) AuthURLHandler {
	if w == nil {
		w = os.Stdout
	}
	return func(authURL string) error {
		fprintf(w, "Opening browser for authorization: %s\n", authURL)
		if err := browser.Open(authURL); err != nil {
			fprintf(w, "Warning: Failed to open browser automatically: %v\n", err)
			fprintf(w, "Please open this URL in your browser manually: %s\n", authURL)
		}
		fprintln(w, "Waiting for authorization callback...")
		return nil
	}
}

// DefaultDeviceCodePromptHandler is a simple handler that displays device code prompts.
// This function can be used by consumer projects as a reference implementation or directly.
// It implements the DeviceCodePromptHandler interface pattern.
// Its progress messages are written to os.Stdout; use DefaultDeviceCodePromptHandlerTo, or set
// DeviceCode.Output (via Configuration.WithDeviceCodeOutput), to redirect or silence this output.
func DefaultDeviceCodePromptHandler(verificationURI, userCode string) error {
	return DefaultDeviceCodePromptHandlerTo(nil)(verificationURI, userCode)
}

// DefaultDeviceCodePromptHandlerTo returns a device code prompt handler equivalent to
// DefaultDeviceCodePromptHandler, but that writes its progress messages to w instead of
// os.Stdout. If w is nil, the returned handler writes to os.Stdout, matching the behavior of
// DefaultDeviceCodePromptHandler. This allows consumers to redirect or silence (using io.Discard)
// the default handler's output without reimplementing its prompt-display logic.
func DefaultDeviceCodePromptHandlerTo(w io.Writer) DeviceCodePromptHandler {
	if w == nil {
		w = os.Stdout
	}
	return func(verificationURI, userCode string) error {
		fprint(w, deviceAuthPromptHeader)

		// Determine which URL to use and whether to auto-open browser
		browserAvailable := browser.CanOpen()
		verificationURIComplete := fmt.Sprintf("%s?user_code=%s", verificationURI, userCode)

		// Auto-open browser if available
		if browserAvailable {
			fprint(w, deviceAuthBrowserOpeningMessage)
			fprintf(w, deviceAuthURLLabel, verificationURIComplete)
			if err := browser.Open(verificationURIComplete); err != nil {
				fprintf(w, deviceAuthBrowserFailWarning, err)
			}
			fprint(w, deviceAuthManualInstructionsHeader)
			fprintf(w, deviceAuthManualVisitPrompt, verificationURI)
			fprintf(w, deviceAuthManualCodePrompt, userCode)
		} else {
			// No browser available - show manual instructions
			fprintf(w, deviceAuthCompleteURLPrompt, verificationURIComplete)
			fprint(w, deviceAuthAlternativeInstructionsHeader)
			fprintf(w, deviceAuthManualVisitPrompt, verificationURI)
			fprintf(w, deviceAuthManualCodePrompt, userCode)
		}

		fprint(w, deviceAuthWaitingMessage)
		return nil
	}
}
