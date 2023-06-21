# \ServerSettingsApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCertificate**](ServerSettingsApi.md#DeleteCertificate) | **Delete** /serverSettings/wsTrustStsSettings/issuerCertificates/{id} | Delete a certificate from WS-Trust STS Settings.
[**GetCaptchaSettings**](ServerSettingsApi.md#GetCaptchaSettings) | **Get** /serverSettings/captchaSettings | (Deprecated) Gets the CAPTCHA settings.
[**GetCert**](ServerSettingsApi.md#GetCert) | **Get** /serverSettings/wsTrustStsSettings/issuerCertificates/{id} | Retrieve details of a certificate.
[**GetCerts**](ServerSettingsApi.md#GetCerts) | **Get** /serverSettings/wsTrustStsSettings/issuerCertificates | Get the list of certificates for WS-Trust STS Settings.
[**GetEmailServerSettings**](ServerSettingsApi.md#GetEmailServerSettings) | **Get** /serverSettings/emailServer | (Deprecated) Gets the email server settings
[**GetGeneralSettings**](ServerSettingsApi.md#GetGeneralSettings) | **Get** /serverSettings/generalSettings | Gets the general settings.
[**GetLogSettings**](ServerSettingsApi.md#GetLogSettings) | **Get** /serverSettings/logSettings | Gets the log settings.
[**GetNotificationSettings**](ServerSettingsApi.md#GetNotificationSettings) | **Get** /serverSettings/notifications | Gets the notification settings
[**GetOutBoundProvisioningSettings**](ServerSettingsApi.md#GetOutBoundProvisioningSettings) | **Get** /serverSettings/outboundProvisioning | Get database used for outbound provisioning
[**GetServerSettings**](ServerSettingsApi.md#GetServerSettings) | **Get** /serverSettings | Gets the server settings
[**GetSystemKeys**](ServerSettingsApi.md#GetSystemKeys) | **Get** /serverSettings/systemKeys | Get the system keys.
[**GetWsTrustStsSettings**](ServerSettingsApi.md#GetWsTrustStsSettings) | **Get** /serverSettings/wsTrustStsSettings | Get the current WS-Trust STS Settings.
[**ImportCertificate**](ServerSettingsApi.md#ImportCertificate) | **Post** /serverSettings/wsTrustStsSettings/issuerCertificates | Import a new certificate.
[**RotateSystemKeys**](ServerSettingsApi.md#RotateSystemKeys) | **Post** /serverSettings/systemKeys/rotate | Rotate the system keys.
[**UpdateCaptchaSettings**](ServerSettingsApi.md#UpdateCaptchaSettings) | **Put** /serverSettings/captchaSettings | (Deprecated) Update the CAPTCHA settings.
[**UpdateEmailServerSettings**](ServerSettingsApi.md#UpdateEmailServerSettings) | **Put** /serverSettings/emailServer | (Deprecated) Update the email server settings
[**UpdateGeneralSettings**](ServerSettingsApi.md#UpdateGeneralSettings) | **Put** /serverSettings/generalSettings | Update general settings.
[**UpdateLogSettings**](ServerSettingsApi.md#UpdateLogSettings) | **Put** /serverSettings/logSettings | Update log settings.
[**UpdateNotificationSettings**](ServerSettingsApi.md#UpdateNotificationSettings) | **Put** /serverSettings/notifications | Update the notification settings.
[**UpdateOutBoundProvisioningSettings**](ServerSettingsApi.md#UpdateOutBoundProvisioningSettings) | **Put** /serverSettings/outboundProvisioning | Update database used for outbound provisioning
[**UpdateServerSettings**](ServerSettingsApi.md#UpdateServerSettings) | **Put** /serverSettings | Update the server settings.
[**UpdateSystemKeys**](ServerSettingsApi.md#UpdateSystemKeys) | **Put** /serverSettings/systemKeys | Update the system keys.
[**UpdateWsTrustStsSettings**](ServerSettingsApi.md#UpdateWsTrustStsSettings) | **Put** /serverSettings/wsTrustStsSettings | Update WS-Trust STS Settings.



## DeleteCertificate

> DeleteCertificate(ctx, id).Execute()

Delete a certificate from WS-Trust STS Settings.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {
    id := "id_example" // string | ID of the certificate to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServerSettingsApi.DeleteCertificate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.DeleteCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the certificate to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptchaSettings

> CaptchaSettings GetCaptchaSettings(ctx).Execute()

(Deprecated) Gets the CAPTCHA settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.GetCaptchaSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.GetCaptchaSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCaptchaSettings`: CaptchaSettings
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.GetCaptchaSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptchaSettingsRequest struct via the builder pattern


### Return type

[**CaptchaSettings**](CaptchaSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCert

> IssuerCert GetCert(ctx, id).Execute()

Retrieve details of a certificate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {
    id := "id_example" // string | ID of the certificate to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.GetCert(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.GetCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCert`: IssuerCert
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.GetCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the certificate to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IssuerCert**](IssuerCert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCerts

> IssuerCerts GetCerts(ctx).Execute()

Get the list of certificates for WS-Trust STS Settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.GetCerts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.GetCerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCerts`: IssuerCerts
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.GetCerts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertsRequest struct via the builder pattern


### Return type

[**IssuerCerts**](IssuerCerts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailServerSettings

> EmailServerSettings GetEmailServerSettings(ctx).Execute()

(Deprecated) Gets the email server settings

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.GetEmailServerSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.GetEmailServerSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailServerSettings`: EmailServerSettings
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.GetEmailServerSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailServerSettingsRequest struct via the builder pattern


### Return type

[**EmailServerSettings**](EmailServerSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGeneralSettings

> GeneralSettings GetGeneralSettings(ctx).Execute()

Gets the general settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.GetGeneralSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.GetGeneralSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGeneralSettings`: GeneralSettings
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.GetGeneralSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGeneralSettingsRequest struct via the builder pattern


### Return type

[**GeneralSettings**](GeneralSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogSettings

> LogSettings GetLogSettings(ctx).Execute()

Gets the log settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.GetLogSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.GetLogSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogSettings`: LogSettings
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.GetLogSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogSettingsRequest struct via the builder pattern


### Return type

[**LogSettings**](LogSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationSettings

> NotificationSettings GetNotificationSettings(ctx).Execute()

Gets the notification settings

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.GetNotificationSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.GetNotificationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationSettings`: NotificationSettings
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.GetNotificationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationSettingsRequest struct via the builder pattern


### Return type

[**NotificationSettings**](NotificationSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutBoundProvisioningSettings

> OutboundProvisionDatabase GetOutBoundProvisioningSettings(ctx).Execute()

Get database used for outbound provisioning



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.GetOutBoundProvisioningSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.GetOutBoundProvisioningSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOutBoundProvisioningSettings`: OutboundProvisionDatabase
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.GetOutBoundProvisioningSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutBoundProvisioningSettingsRequest struct via the builder pattern


### Return type

[**OutboundProvisionDatabase**](OutboundProvisionDatabase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerSettings

> ServerSettings GetServerSettings(ctx).Execute()

Gets the server settings

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.GetServerSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.GetServerSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerSettings`: ServerSettings
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.GetServerSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerSettingsRequest struct via the builder pattern


### Return type

[**ServerSettings**](ServerSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemKeys

> SystemKeys GetSystemKeys(ctx).Execute()

Get the system keys.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.GetSystemKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.GetSystemKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemKeys`: SystemKeys
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.GetSystemKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemKeysRequest struct via the builder pattern


### Return type

[**SystemKeys**](SystemKeys.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWsTrustStsSettings

> WsTrustStsSettings GetWsTrustStsSettings(ctx).Execute()

Get the current WS-Trust STS Settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.GetWsTrustStsSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.GetWsTrustStsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWsTrustStsSettings`: WsTrustStsSettings
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.GetWsTrustStsSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWsTrustStsSettingsRequest struct via the builder pattern


### Return type

[**WsTrustStsSettings**](WsTrustStsSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCertificate

> IssuerCert ImportCertificate(ctx).Body(body).Execute()

Import a new certificate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {
    body := *openapiclient.NewX509File("FileData_example") // X509File | File data to import.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.ImportCertificate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.ImportCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportCertificate`: IssuerCert
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.ImportCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**X509File**](X509File.md) | File data to import. | 

### Return type

[**IssuerCert**](IssuerCert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateSystemKeys

> SystemKeys RotateSystemKeys(ctx).Execute()

Rotate the system keys.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.RotateSystemKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.RotateSystemKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RotateSystemKeys`: SystemKeys
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.RotateSystemKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRotateSystemKeysRequest struct via the builder pattern


### Return type

[**SystemKeys**](SystemKeys.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCaptchaSettings

> CaptchaSettings UpdateCaptchaSettings(ctx).Body(body).Execute()

(Deprecated) Update the CAPTCHA settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {
    body := *openapiclient.NewCaptchaSettings() // CaptchaSettings | CAPTCHA settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.UpdateCaptchaSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.UpdateCaptchaSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCaptchaSettings`: CaptchaSettings
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.UpdateCaptchaSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCaptchaSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CaptchaSettings**](CaptchaSettings.md) | CAPTCHA settings. | 

### Return type

[**CaptchaSettings**](CaptchaSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailServerSettings

> EmailServerSettings UpdateEmailServerSettings(ctx).Body(body).ValidationEmail(validationEmail).ValidateOnly(validateOnly).Execute()

(Deprecated) Update the email server settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {
    body := *openapiclient.NewEmailServerSettings("SourceAddr_example", "EmailServer_example", int64(123)) // EmailServerSettings | Configuration for email server settings.
    validationEmail := "validationEmail_example" // string | The email address used to validate the email server settings. (optional)
    validateOnly := true // bool | Only validation will be performed.  Email server settings will not be saved. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.UpdateEmailServerSettings(context.Background()).Body(body).ValidationEmail(validationEmail).ValidateOnly(validateOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.UpdateEmailServerSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEmailServerSettings`: EmailServerSettings
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.UpdateEmailServerSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailServerSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EmailServerSettings**](EmailServerSettings.md) | Configuration for email server settings. | 
 **validationEmail** | **string** | The email address used to validate the email server settings. | 
 **validateOnly** | **bool** | Only validation will be performed.  Email server settings will not be saved. | 

### Return type

[**EmailServerSettings**](EmailServerSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGeneralSettings

> GeneralSettings UpdateGeneralSettings(ctx).Body(body).Execute()

Update general settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {
    body := *openapiclient.NewGeneralSettings() // GeneralSettings | Configuration for general settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.UpdateGeneralSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.UpdateGeneralSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGeneralSettings`: GeneralSettings
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.UpdateGeneralSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGeneralSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GeneralSettings**](GeneralSettings.md) | Configuration for general settings. | 

### Return type

[**GeneralSettings**](GeneralSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogSettings

> LogSettings UpdateLogSettings(ctx).Body(body).Execute()

Update log settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {
    body := *openapiclient.NewLogSettings() // LogSettings | Configuration for log settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.UpdateLogSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.UpdateLogSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogSettings`: LogSettings
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.UpdateLogSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LogSettings**](LogSettings.md) | Configuration for log settings. | 

### Return type

[**LogSettings**](LogSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationSettings

> NotificationSettings UpdateNotificationSettings(ctx).Body(body).Execute()

Update the notification settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {
    body := *openapiclient.NewNotificationSettings() // NotificationSettings | Notification settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.UpdateNotificationSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.UpdateNotificationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationSettings`: NotificationSettings
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.UpdateNotificationSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NotificationSettings**](NotificationSettings.md) | Notification settings. | 

### Return type

[**NotificationSettings**](NotificationSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOutBoundProvisioningSettings

> OutboundProvisionDatabase UpdateOutBoundProvisioningSettings(ctx).Body(body).Execute()

Update database used for outbound provisioning



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {
    body := *openapiclient.NewOutboundProvisionDatabase(*openapiclient.NewResourceLink("Id_example")) // OutboundProvisionDatabase | The Outbound Provision Database settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.UpdateOutBoundProvisioningSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.UpdateOutBoundProvisioningSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOutBoundProvisioningSettings`: OutboundProvisionDatabase
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.UpdateOutBoundProvisioningSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOutBoundProvisioningSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OutboundProvisionDatabase**](OutboundProvisionDatabase.md) | The Outbound Provision Database settings. | 

### Return type

[**OutboundProvisionDatabase**](OutboundProvisionDatabase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerSettings

> ServerSettings UpdateServerSettings(ctx).Body(body).Execute()

Update the server settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {
    body := *openapiclient.NewServerSettings() // ServerSettings | Configuration for server settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.UpdateServerSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.UpdateServerSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServerSettings`: ServerSettings
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.UpdateServerSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ServerSettings**](ServerSettings.md) | Configuration for server settings. | 

### Return type

[**ServerSettings**](ServerSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSystemKeys

> SystemKeys UpdateSystemKeys(ctx).Body(body).Execute()

Update the system keys.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {
    body := *openapiclient.NewSystemKeys(*openapiclient.NewSystemKey(), *openapiclient.NewSystemKey()) // SystemKeys | System keys.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.UpdateSystemKeys(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.UpdateSystemKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSystemKeys`: SystemKeys
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.UpdateSystemKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSystemKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SystemKeys**](SystemKeys.md) | System keys. | 

### Return type

[**SystemKeys**](SystemKeys.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWsTrustStsSettings

> WsTrustStsSettings UpdateWsTrustStsSettings(ctx).Body(body).Execute()

Update WS-Trust STS Settings.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingfederate-go-client"
)

func main() {
    body := *openapiclient.NewWsTrustStsSettings() // WsTrustStsSettings | Configuration for WS-Trust STS Settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSettingsApi.UpdateWsTrustStsSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSettingsApi.UpdateWsTrustStsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWsTrustStsSettings`: WsTrustStsSettings
    fmt.Fprintf(os.Stdout, "Response from `ServerSettingsApi.UpdateWsTrustStsSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWsTrustStsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WsTrustStsSettings**](WsTrustStsSettings.md) | Configuration for WS-Trust STS Settings. | 

### Return type

[**WsTrustStsSettings**](WsTrustStsSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

