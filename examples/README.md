# Examples

This directory contains examples of how to use the pingfederate-go-client SDK with different
OAuth2 authentication flows against the PingFederate administrative API.

## Available Examples

- **[authorization_code](authorization_code/)**: Interactive authentication using the OAuth2 authorization code flow (with PKCE)
- **[client_credentials](client_credentials/)**: Server-to-server authentication using the client credentials grant type
- **[device_code](device_code/)**: Device authentication flow for limited-input devices or CLI applications

## Endpoint Model

PingFederate derives its OAuth2 endpoints from the **runtime engine base URL** (for example
`https://pingfederate.example.com:9031`). This is distinct from the **administrative API base URL**
(for example `https://pingfederate-admin.example.com:9999/pf-admin-api/v1`) that the generated
client calls. Both are required by the examples.

## Authentication Flow Guide

### Client Credentials

Best for: server-to-server authentication, automation scripts, backend services.

**PingFederate Setup:**
1. Create an OAuth client in PingFederate with the **Client Credentials** grant type enabled.
2. Configure a client secret and note the Client ID and Client Secret.
3. Grant the client the scopes/permissions required for the admin operations you intend to perform.

**Environment Variables:**
```shell
export PINGFEDERATE_ADMIN_API_URL="https://pingfederate-admin.example.com:9999/pf-admin-api/v1"
export PINGFEDERATE_RUNTIME_URL="https://pingfederate.example.com:9031"
export PINGFEDERATE_CLIENT_ID="your-client-id"
export PINGFEDERATE_CLIENT_SECRET="your-client-secret"
```

### Authorization Code (authorization_code example)

Best for: interactive user authentication where a browser is available.

**PingFederate Setup:**
1. Create an OAuth client with the **Authorization Code** grant type enabled.
2. Configure the redirect URI to `http://127.0.0.1:7464/callback` (the SDK default).
3. Note the Client ID (a public client requires no secret for PKCE).

**Environment Variables:**
```shell
export PINGFEDERATE_ADMIN_API_URL="https://pingfederate-admin.example.com:9999/pf-admin-api/v1"
export PINGFEDERATE_RUNTIME_URL="https://pingfederate.example.com:9031"
export PINGFEDERATE_CLIENT_ID="your-client-id"
```

**What to Expect:**
- The SDK starts a local web server on port 7464.
- Your browser opens to the PingFederate login page.
- After authentication, you are redirected back to the local server.
- The SDK captures the authorization code and exchanges it for tokens, then caches them in the OS keychain.

### Device Code (device_code example)

Best for: CLI tools, devices with limited input capabilities, headless applications.

**PingFederate Setup:**
1. Create an OAuth client with the **Device Authorization** grant type enabled.
2. Note the Client ID (no client secret needed).

**Environment Variables:**
```shell
export PINGFEDERATE_ADMIN_API_URL="https://pingfederate-admin.example.com:9999/pf-admin-api/v1"
export PINGFEDERATE_RUNTIME_URL="https://pingfederate.example.com:9031"
export PINGFEDERATE_CLIENT_ID="your-client-id"
```

**What to Expect:**
- The SDK displays a user code and verification URL.
- Visit the URL in a browser on any device and enter the code.
- Authenticate with your PingFederate credentials.
- The SDK polls and receives tokens once authentication completes, then caches them in the OS keychain.

## Token Storage

The interactive flows (authorization_code, device_code) cache tokens in the OS keychain by default
via `WithStorageName(...)`, so a subsequent run reuses (and silently refreshes) an existing token
rather than prompting for login again. The keychain account name is derived from the runtime base
URL, client ID, and grant type. Set the storage type to `config.StorageTypeNone` to disable caching.

## Running the Examples

Navigate to any example directory and run:

```shell
cd client_credentials  # or authorization_code, device_code
go run main.go
```

> [!NOTE]
> The examples disable TLS verification for brevity only. Configure proper trust for real deployments.
