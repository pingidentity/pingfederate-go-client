# Authorization Code Example

This example demonstrates how to use the PingFederate Go Client SDK with the OAuth2 Authorization
Code flow (with PKCE) for interactive, browser-based authentication against the PingFederate
administrative API.

## Use Case

Authorization Code flow is ideal for:
- Interactive user authentication where a browser is available
- Desktop and CLI applications acting on behalf of a user
- Scenarios requiring a human login rather than a shared secret

## PingFederate Configuration

1. In the PingFederate administrative console, create an **OAuth Client**.
2. Enable the **Authorization Code** grant type.
3. Configure the redirect URI to `http://127.0.0.1:7464/callback` (the SDK default).
   - The SDK starts a temporary local web server on port 7464 to capture the callback.
   - Customize the port and path with `WithAuthorizationCodeRedirectURI(...)` if needed.
4. Note the **Client ID** (a public client using PKCE requires no secret).

## Running the Example

### Set Environment Variables

```shell
export PINGFEDERATE_ADMIN_API_URL="https://pingfederate-admin.example.com:9999/pf-admin-api/v1"
export PINGFEDERATE_RUNTIME_URL="https://pingfederate.example.com:9031"
export PINGFEDERATE_CLIENT_ID="your-client-id"
```

### Run the Example

```shell
cd examples/authorization_code
go run main.go
```

## What to Expect

1. The SDK starts a local web server on port 7464.
2. Your browser opens to the PingFederate login page.
3. After authentication, you are redirected back to the local server.
4. The SDK captures the authorization code and exchanges it for tokens.
5. Tokens are cached in the OS keychain, so a subsequent run reuses (and silently refreshes) them without prompting for login again.
6. The example reads and prints the PingFederate version.

## Token Storage

This example enables keychain caching via `WithStorageName("pingfederate")`. Because PingFederate
has no environment ID, the keychain account name is derived from the runtime base URL, client ID,
and grant type. Disable caching by setting the storage type to `config.StorageTypeNone`.

## Troubleshooting

**Browser does not open:**
- The SDK prints the authorization URL — open it manually.

**Redirect URI mismatch:**
- Ensure the client's redirect URI exactly matches `http://127.0.0.1:7464/callback` (or your customized value).

**Connection errors:**
- Verify `PINGFEDERATE_RUNTIME_URL` points at the runtime engine and `PINGFEDERATE_ADMIN_API_URL` at the admin API.

## Additional Resources

- [OAuth2 Authorization Code Flow](https://oauth.net/2/grant-types/authorization-code/)
- [PKCE (RFC 7636)](https://datatracker.ietf.org/doc/html/rfc7636)
- [PingFederate Documentation](https://docs.pingidentity.com/pingfederate/latest/)
