# Device Code Example

This example demonstrates how to use the PingFederate Go Client SDK with the OAuth2 Device
Authorization flow (with PKCE) for authentication on devices with limited input capabilities,
against the PingFederate administrative API.

## Use Case

Device Code flow is ideal for:
- CLI tools and headless applications
- Devices with limited input capabilities (TVs, IoT devices)
- Scenarios where the authenticating device cannot easily host a browser callback

## PingFederate Configuration

1. In the PingFederate administrative console, create an **OAuth Client**.
2. Enable the **Device Authorization** grant type.
3. Note the **Client ID** (no client secret needed).

## Running the Example

### Set Environment Variables

```shell
export PINGFEDERATE_ADMIN_API_URL="https://pingfederate-admin.example.com:9999/pf-admin-api/v1"
export PINGFEDERATE_RUNTIME_URL="https://pingfederate.example.com:9031"
export PINGFEDERATE_CLIENT_ID="your-client-id"
```

### Run the Example

```shell
cd examples/device_code
go run main.go
```

## What to Expect

1. The SDK displays a user code and verification URL.
2. Visit the URL in a browser on any device.
3. Enter the user code shown by the application.
4. Authenticate with your PingFederate credentials.
5. The SDK polls and receives tokens once authentication completes.
6. Tokens are cached in the OS keychain, so a subsequent run reuses (and silently refreshes) them without prompting for login again.
7. The example reads and prints the PingFederate version.

## Token Storage

This example enables keychain caching via `WithStorageName("pingfederate")`. The keychain account
name is derived from the runtime base URL, client ID, and grant type. Disable caching by setting the
storage type to `config.StorageTypeNone`.

## Troubleshooting

**"expired_token" or timeout:**
- Complete the browser login promptly; the user code expires after a short interval.

**Device flow not enabled:**
- Ensure the Device Authorization grant type is enabled on the client.

**Connection errors:**
- Verify `PINGFEDERATE_RUNTIME_URL` points at the runtime engine and `PINGFEDERATE_ADMIN_API_URL` at the admin API.

## Additional Resources

- [OAuth2 Device Authorization Grant (RFC 8628)](https://datatracker.ietf.org/doc/html/rfc8628)
- [PingFederate Documentation](https://docs.pingidentity.com/pingfederate/latest/)
