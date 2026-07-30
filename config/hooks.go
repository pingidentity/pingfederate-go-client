// Copyright © 2026 Ping Identity Corporation

package config

import (
	"fmt"

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
}

// DefaultAuthorizationCodeBrowserHandler is the default handler for opening the authorization URL.
// It attempts to open the system browser automatically and provides fallback instructions if that fails.
// This function implements the AuthURLHandler interface and provides a consistent UX pattern.
// Consumer projects can use this handler as a reference or directly in their own implementations.
func DefaultAuthorizationCodeBrowserHandler(authURL string) error {
	fmt.Printf("Opening browser for authorization: %s\n", authURL)
	if err := browser.Open(authURL); err != nil {
		fmt.Printf("Warning: Failed to open browser automatically: %v\n", err)
		fmt.Printf("Please open this URL in your browser manually: %s\n", authURL)
	}
	fmt.Println("Waiting for authorization callback...")
	return nil
}

// DefaultDeviceCodePromptHandler is a simple handler that displays device code prompts.
// This function can be used by consumer projects as a reference implementation or directly.
// It implements the DeviceCodePromptHandler interface pattern.
func DefaultDeviceCodePromptHandler(verificationURI, userCode string) error {
	fmt.Print(deviceAuthPromptHeader)

	// Determine which URL to use and whether to auto-open browser
	browserAvailable := browser.CanOpen()
	verificationURIComplete := fmt.Sprintf("%s?user_code=%s", verificationURI, userCode)

	// Auto-open browser if available
	if browserAvailable {
		fmt.Print(deviceAuthBrowserOpeningMessage)
		fmt.Printf(deviceAuthURLLabel, verificationURIComplete)
		if err := browser.Open(verificationURIComplete); err != nil {
			fmt.Printf(deviceAuthBrowserFailWarning, err)
		}
		fmt.Print(deviceAuthManualInstructionsHeader)
		fmt.Printf(deviceAuthManualVisitPrompt, verificationURI)
		fmt.Printf(deviceAuthManualCodePrompt, userCode)
	} else {
		// No browser available - show manual instructions
		fmt.Printf(deviceAuthCompleteURLPrompt, verificationURIComplete)
		fmt.Print(deviceAuthAlternativeInstructionsHeader)
		fmt.Printf(deviceAuthManualVisitPrompt, verificationURI)
		fmt.Printf(deviceAuthManualCodePrompt, userCode)
	}

	fmt.Print(deviceAuthWaitingMessage)
	return nil
}
