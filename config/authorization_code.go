// Copyright © 2026 Ping Identity Corporation

package config

import (
	"context"
	"crypto/rand"
	_ "embed"
	"encoding/base64"
	"fmt"
	"html/template"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

//go:embed html/auth_result.html
var authResultHTML string

// authResultTemplate is the pre-parsed HTML template for auth result pages.
var authResultTemplate *template.Template

func init() {
	// Parse template once at package initialization
	var err error
	authResultTemplate, err = template.New("authResult").Parse(authResultHTML)
	if err != nil {
		panic(fmt.Sprintf("failed to parse auth result template: %v", err))
	}
}

const (
	// defaultAuthorizationCodeRedirectURIPort is the default port for the authorization code redirect URI.
	defaultAuthorizationCodeRedirectURIPort = "7464"

	// defaultAuthorizationCodeRedirectURIPath is the default path for the authorization code redirect URI.
	defaultAuthorizationCodeRedirectURIPath = "/callback"

	// defaultAuthorizationCodeRedirectURIPrefix is the default redirect URI prefix (loopback interface).
	defaultAuthorizationCodeRedirectURIPrefix = "http://127.0.0.1:"

	// defaultAuthorizationCodeRedirectURI is the default redirect URI for the authorization code.
	defaultAuthorizationCodeRedirectURI = defaultAuthorizationCodeRedirectURIPrefix + defaultAuthorizationCodeRedirectURIPort + defaultAuthorizationCodeRedirectURIPath

	// defaultProjectName is the default project name displayed on auth result pages.
	defaultProjectName = "Ping Identity Developer Tools"

	// defaultAuthFailedHeading is the default heading displayed on the authentication failure page.
	defaultAuthFailedHeading = "Authorization Failed"

	// defaultAuthFailedMessage is the default message displayed on the authentication failure page.
	defaultAuthFailedMessage = "An error has occurred and authorization was not successful."

	// defaultAuthSuccessHeading is the default heading displayed on the authentication success page.
	defaultAuthSuccessHeading = "Authorization Success"

	// defaultAuthSuccessMessage is the default message displayed on the authentication success page.
	defaultAuthSuccessMessage = "You have successfully authenticated to PingFederate and have authorized API access."

	// contentTypeHTML is the content type for HTML responses.
	contentTypeHTML = "text/html; charset=utf-8"

	// callbackServerReadHeaderTimeout is the timeout for reading HTTP headers on the callback server.
	callbackServerReadHeaderTimeout = 10 * time.Second

	// tokenExchangeTimeout is the timeout for waiting for token exchange to complete.
	tokenExchangeTimeout = 30 * time.Second

	// authSuccessWaitTime is the time to wait after successful auth before returning (allows HTTP response to be sent).
	authSuccessWaitTime = 1 * time.Second

	// serverVerificationRetryDelay is the delay between retries when verifying server startup.
	serverVerificationRetryDelay = 10 * time.Millisecond

	// serverVerificationMaxRetries is the maximum number of retries when verifying server startup.
	serverVerificationMaxRetries = 10

	// serverVerificationDialTimeout is the timeout for dial attempts when verifying server startup.
	serverVerificationDialTimeout = 50 * time.Millisecond

	// httpChannelBufferSize is the buffer size for HTTP callback channels.
	httpChannelBufferSize = 1

	// doneChannelBufferSize is the buffer size for done notification channels.
	doneChannelBufferSize = 1

	// httpStatusUnauthorized is the HTTP status code for unauthorized requests.
	httpStatusUnauthorized = http.StatusUnauthorized

	// urlQueryParamError is the URL query parameter name for error code.
	urlQueryParamError = "error"

	// urlQueryParamErrorDescription is the URL query parameter name for error description.
	urlQueryParamErrorDescription = "error_description"

	// urlQueryParamCode is the URL query parameter name for authorization code.
	urlQueryParamCode = "code"

	// urlPathPrefix is the prefix character for URL paths.
	urlPathPrefix = "/"

	// networkProtocolTCP is the network protocol for TCP connections.
	networkProtocolTCP = "tcp"

	// networkPortPrefix is the prefix character for port numbers.
	networkPortPrefix = ":"
)

// GetDefaultAuthorizationCodeRedirectURIPort returns the default authorization code redirect URI port.
func GetDefaultAuthorizationCodeRedirectURIPort() string {
	return defaultAuthorizationCodeRedirectURIPort
}

// GetDefaultAuthorizationCodeRedirectURIPath returns the default authorization code redirect URI path.
func GetDefaultAuthorizationCodeRedirectURIPath() string {
	return defaultAuthorizationCodeRedirectURIPath
}

// GetDefaultAuthorizationCodeRedirectURI returns the default authorization code redirect URI.
func GetDefaultAuthorizationCodeRedirectURI() string {
	return defaultAuthorizationCodeRedirectURI
}

// generateState creates a cryptographically secure random state parameter for OAuth2 CSRF protection.
// It returns a base64-encoded random string of 32 bytes, or an error if random generation fails.
func generateState() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed to generate state parameter: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// AuthorizationCodeTokenSource returns an oauth2.TokenSource using the authorization code grant type
// with PKCE. It starts a local HTTP callback server on the loopback interface, invokes the browser
// handler to send the user to the authorization endpoint, waits for the redirect, and exchanges the
// returned authorization code for an access token.
func (a *AuthorizationCode) AuthorizationCodeTokenSource(ctx context.Context, endpoint oauth2.Endpoint) (oauth2.TokenSource, error) {
	if a.AuthorizationCodeClientID == nil || *a.AuthorizationCodeClientID == "" {
		return nil, fmt.Errorf("client ID is required for authorization code grant type")
	}

	slog.Debug("Using authorization code token source with provided client ID", "client ID", *a.AuthorizationCodeClientID)
	redirectURIPath := GetDefaultAuthorizationCodeRedirectURIPath()
	redirectURIPort := GetDefaultAuthorizationCodeRedirectURIPort()

	if a.AuthorizationCodeRedirectURI.Port != "" {
		redirectURIPort = a.AuthorizationCodeRedirectURI.Port
	}

	if a.AuthorizationCodeRedirectURI.Path != "" {
		redirectURIPath = a.AuthorizationCodeRedirectURI.Path
	}

	redirectURI := fmt.Sprintf("%s%s%s", defaultAuthorizationCodeRedirectURIPrefix, redirectURIPort, redirectURIPath)

	var scopes []string
	if a.AuthorizationCodeScopes != nil {
		scopes = *a.AuthorizationCodeScopes
	}

	config := &oauth2.Config{
		ClientID:    *a.AuthorizationCodeClientID,
		Endpoint:    endpoint,
		RedirectURL: redirectURI,
		Scopes:      scopes,
	}

	codeVerifier := oauth2.GenerateVerifier()

	// Generate cryptographically secure state parameter for CSRF protection
	state, err := generateState()
	if err != nil {
		return nil, fmt.Errorf("failed to generate state parameter: %w", err)
	}

	// Start local HTTP server to capture callback
	codeChan := make(chan string, httpChannelBufferSize)
	errChan := make(chan error, httpChannelBufferSize)
	tokenResultChan := make(chan error, httpChannelBufferSize) // nil for success, error for failure
	doneChan := make(chan struct{}, doneChannelBufferSize)     // signals HTTP response has been sent

	server, err := startCallbackServer(redirectURI, state, codeChan, errChan, tokenResultChan, doneChan, a.CustomPageDataSuccess, a.CustomPageDataError)
	if err != nil {
		return nil, fmt.Errorf("failed to start callback server: %w", err)
	}
	defer func() {
		if closeErr := server.Close(); closeErr != nil {
			slog.Warn("Warning: failed to close server", "error", closeErr)
		}
	}()

	// Generate authorization URL with secure state parameter and handle browser opening
	authURL := config.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(codeVerifier))

	// Use the default handler, which honors the configured output (silent when unset).
	handler := DefaultAuthorizationCodeBrowserHandlerTo(a.Output)

	if err := handler(authURL); err != nil {
		return nil, fmt.Errorf("prompt handler failed: %w", err)
	}

	// Wait for authorization code or error
	var code string
	select {
	case code = <-codeChan:
		slog.Info("Authorization code received")
	case err := <-errChan:
		// Wait for the HTTP response to be sent before returning
		select {
		case <-doneChan:
			// Response sent successfully
		case <-time.After(authSuccessWaitTime):
			// Timeout waiting for response to be sent
			slog.Warn("Timeout waiting for error page to be sent")
		}
		return nil, fmt.Errorf("authorization failed: %w", err)
	case <-ctx.Done():
		return nil, fmt.Errorf("authorization cancelled: %w", ctx.Err())
	}

	// Exchange authorization code for token
	tok, err := config.Exchange(ctx, code, oauth2.VerifierOption(codeVerifier))

	if err != nil {
		// Signal failure to show error page
		tokenResultChan <- err

		// Wait for the HTTP response to be sent before returning
		select {
		case <-doneChan:
			// Response sent successfully
		case <-time.After(authSuccessWaitTime):
			// Timeout waiting for response to be sent
			slog.Warn("Timeout waiting for error page to be sent")
		}

		return nil, fmt.Errorf("failed to exchange code for token: %w", err)
	}

	// Signal success to show success page
	tokenResultChan <- nil

	// Wait for the HTTP response to be sent before returning
	select {
	case <-doneChan:
		// Response sent successfully
	case <-time.After(authSuccessWaitTime):
		// Timeout waiting for success page to be sent
		slog.Warn("Timeout waiting for success page to be sent")
	}

	slog.Debug("Successfully obtained access token via authorization code flow")

	// Return a static token source with the token we just obtained
	return oauth2.StaticTokenSource(tok), nil
}

func returnFailedPage(w http.ResponseWriter, errorDetails string, customPageData *AuthResultPageData) error {
	// Create template data
	templateData := struct {
		ProjectName  string
		Heading      string
		Message      string
		ErrorDetails string
		IsSuccess    bool
	}{
		ProjectName:  defaultProjectName,
		Heading:      defaultAuthFailedHeading,
		Message:      defaultAuthFailedMessage,
		ErrorDetails: errorDetails,
		IsSuccess:    false,
	}

	// Override with custom data if provided
	if customPageData != nil {
		if customPageData.ProjectName != "" {
			templateData.ProjectName = customPageData.ProjectName
		}
		if customPageData.Heading != "" {
			templateData.Heading = customPageData.Heading
		}
		if customPageData.Message != "" {
			templateData.Message = customPageData.Message
		}
	}

	return authResultTemplate.Execute(w, templateData)
}

func returnSuccessPage(w http.ResponseWriter, customPageData *AuthResultPageData) error {
	// Create template data
	templateData := struct {
		ProjectName  string
		Heading      string
		Message      string
		ErrorDetails string
		IsSuccess    bool
	}{
		ProjectName:  defaultProjectName,
		Heading:      defaultAuthSuccessHeading,
		Message:      defaultAuthSuccessMessage,
		ErrorDetails: "", // Empty for success
		IsSuccess:    true,
	}

	// Override with custom data if provided
	if customPageData != nil {
		if customPageData.ProjectName != "" {
			templateData.ProjectName = customPageData.ProjectName
		}
		if customPageData.Heading != "" {
			templateData.Heading = customPageData.Heading
		}
		if customPageData.Message != "" {
			templateData.Message = customPageData.Message
		}
	}

	return authResultTemplate.Execute(w, templateData)
}

// startCallbackServer starts a local HTTP server to handle OAuth2 callbacks.
// It validates the state parameter for CSRF protection and signals completion via doneChan.
func startCallbackServer(redirectURI string, expectedState string, codeChan chan<- string, errChan chan<- error, tokenResultChan <-chan error, doneChan chan<- struct{}, customPageDataSuccess *AuthResultPageData, customPageDataError *AuthResultPageData) (*http.Server, error) {
	// Parse the redirect URI to get the port
	parsedURI, err := url.Parse(redirectURI)
	if err != nil {
		return nil, fmt.Errorf("invalid redirect URI: %w", err)
	}

	// Extract port from URI or use default
	port := parsedURI.Port()
	if port == "" {
		port = defaultAuthorizationCodeRedirectURIPort
	}

	// Extract path and ensure it's valid for HTTP mux
	path := parsedURI.Path
	if path == "" {
		path = defaultAuthorizationCodeRedirectURIPath
	}
	if !strings.HasPrefix(path, urlPathPrefix) {
		path = urlPathPrefix + path
	}

	// Test if port is available and keep the listener
	listener, err := net.Listen(networkProtocolTCP, networkPortPrefix+port)
	if err != nil {
		return nil, fmt.Errorf("port %s is not available: %w", port, err)
	}

	// Create HTTP server
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:              networkPortPrefix + port,
		Handler:           mux,
		ReadHeaderTimeout: callbackServerReadHeaderTimeout,
	}

	// Handle callback endpoint
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		// Validate state parameter for CSRF protection
		receivedState := query.Get("state")
		if receivedState != expectedState {
			errChan <- fmt.Errorf("invalid state parameter - possible CSRF attack")

			w.Header().Set("Content-Type", contentTypeHTML)
			w.WriteHeader(http.StatusBadRequest)
			err := returnFailedPage(w, "Invalid state parameter", customPageDataError)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading failed page. Authentication failed.")
			}
			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			}
			close(doneChan)
			return
		}

		// Check for error in callback
		if errCode := query.Get(urlQueryParamError); errCode != "" {
			errDesc := query.Get(urlQueryParamErrorDescription)
			if errDesc == "" {
				errDesc = errCode
			}
			errChan <- fmt.Errorf("authorization error: %s", errDesc)

			w.Header().Set("Content-Type", contentTypeHTML)
			w.WriteHeader(http.StatusBadRequest)

			err := returnFailedPage(w, errDesc, customPageDataError)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading failed page. Authentication failed.")
			}
			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			}
			close(doneChan)
			return
		}

		// Get authorization code
		code := query.Get(urlQueryParamCode)
		if code == "" {
			errChan <- fmt.Errorf("no authorization code received")

			w.Header().Set("Content-Type", contentTypeHTML)
			w.WriteHeader(http.StatusBadRequest)
			err := returnFailedPage(w, "No authorization code received", customPageDataError)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading failed page. Authentication failed.")
			}
			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			}
			close(doneChan)
			return
		}

		// Send code to channel
		codeChan <- code

		// Wait for token exchange to complete before showing page
		select {
		case tokenErr := <-tokenResultChan:
			if tokenErr == nil {
				// Token exchange succeeded
				w.Header().Set("Content-Type", contentTypeHTML)
				w.WriteHeader(http.StatusOK)

				err := returnSuccessPage(w, customPageDataSuccess)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error loading success page. Authentication was successful.\n%s", err)
				}
				if flusher, ok := w.(http.Flusher); ok {
					flusher.Flush()
				}
				close(doneChan)
			} else {
				// Token exchange failed
				w.Header().Set("Content-Type", contentTypeHTML)
				w.WriteHeader(httpStatusUnauthorized)
				err := returnFailedPage(w, fmt.Sprintf("Token exchange failed: %v", tokenErr), customPageDataError)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error loading failed page. Token exchange failed: %v", tokenErr)
				}
				if flusher, ok := w.(http.Flusher); ok {
					flusher.Flush()
				}
				close(doneChan)
			}
		case <-time.After(tokenExchangeTimeout):
			// Token exchange timed out
			w.Header().Set("Content-Type", contentTypeHTML)
			w.WriteHeader(http.StatusInternalServerError)
			err := returnFailedPage(w, "Token exchange timed out", customPageDataError)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading failed page. Token exchange timed out.")
			}
			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			}
			close(doneChan)
		}
	})

	// Start server in background using the existing listener
	serverStarted := make(chan error, httpChannelBufferSize)
	go func() {
		// Use Serve() with the existing listener instead of ListenAndServe()
		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			select {
			case serverStarted <- err:
				// Server failed to start
			default:
				// Server was already marked as started, so this is a runtime error
				errChan <- fmt.Errorf("callback server error: %w", err)
			}
		}
	}()

	// Verify server is actually accepting connections before returning
	maxRetries := serverVerificationMaxRetries
	for i := 0; i < maxRetries; i++ {
		// Check if server failed
		select {
		case err := <-serverStarted:
			return nil, fmt.Errorf("failed to start server: %w", err)
		default:
		}

		// Try to connect
		conn, err := net.DialTimeout(networkProtocolTCP, networkPortPrefix+port, serverVerificationDialTimeout)
		if err == nil {
			if closeErr := conn.Close(); closeErr != nil {
				slog.Warn("Failed to close connection during server verification", "error", closeErr)
			}
			// Server is accepting connections
			return server, nil
		}

		// Wait a bit before retrying
		time.Sleep(serverVerificationRetryDelay)
	}

	return nil, fmt.Errorf("server failed to start accepting connections in time")
}
