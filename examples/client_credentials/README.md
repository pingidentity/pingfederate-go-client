# Client Credentials Example

This example demonstrates how to use the PingFederate Go Client SDK with the OAuth2 Client
Credentials flow for server-to-server authentication against the PingFederate administrative API.

## Use Case

Client Credentials flow is ideal for:
- Server-to-server authentication
- Machine-to-machine (M2M) communication
- Backend services and APIs
- Automation scripts and batch jobs
- Applications that don't require user interaction

## PingFederate Configuration

1. In the PingFederate administrative console, create an **OAuth Client**.
2. Enable the **Client Credentials** grant type.
3. Configure a **client secret** and note the **Client ID** and **Client Secret**.
4. Grant the client the scopes required for the admin operations you intend to perform.

## Running the Example

### Set Environment Variables

```shell
export PINGFEDERATE_ADMIN_API_URL="https://pingfederate-admin.example.com:9999/pf-admin-api/v1"
export PINGFEDERATE_RUNTIME_URL="https://pingfederate.example.com:9031"
export PINGFEDERATE_CLIENT_ID="your-client-id"
export PINGFEDERATE_CLIENT_SECRET="your-client-secret"
```

**Security Note**: Never commit these credentials to version control. Consider using:
- Environment variable files (`.env`) that are `.gitignore`d
- Secret management systems (AWS Secrets Manager, HashiCorp Vault, etc.)
- CI/CD platform secret storage

### Run the Example

```shell
cd examples/client_credentials
go run main.go
```

## What to Expect

1. **Authentication**: The SDK authenticates using the client credentials against the runtime token endpoint.
2. **Token Acquisition**: An access token is obtained from PingFederate.
3. **API Call**: The example reads the PingFederate version from the admin API.
4. **Output**: The version is displayed in the console.

## How It Works

The client credentials flow:
1. The application sends its client ID and secret to the PingFederate token endpoint (derived from the runtime base URL).
2. PingFederate validates the credentials.
3. PingFederate returns an access token.
4. The application uses the token for admin API calls.
5. The token is refreshed automatically by the SDK when it expires (no caching is required for this flow).

## Troubleshooting

**"invalid_client" error:**
- Verify your `PINGFEDERATE_CLIENT_ID` and `PINGFEDERATE_CLIENT_SECRET` are correct.
- Ensure the client exists and the Client Credentials grant type is enabled.

**"insufficient_scope" or authorization errors:**
- The client needs the appropriate scopes/permissions for the admin operations.

**Connection errors:**
- Verify `PINGFEDERATE_RUNTIME_URL` points at the runtime engine and `PINGFEDERATE_ADMIN_API_URL` at the admin API.
- Check network connectivity and that TLS trust is configured for real deployments.

## Security Best Practices

1. **Protect Client Secret**: Treat it like a password — never commit it to source control.
2. **Rotate Credentials**: Regularly rotate client secrets.
3. **Least Privilege**: Grant only the minimum required scopes.
4. **Secure Storage**: Use secure secret management systems in production.
5. **TLS**: Configure proper certificate trust rather than disabling verification (this example disables it for brevity only).

## Additional Resources

- [OAuth2 Client Credentials Flow](https://oauth.net/2/grant-types/client-credentials/)
- [PingFederate Documentation](https://docs.pingidentity.com/pingfederate/latest/)
