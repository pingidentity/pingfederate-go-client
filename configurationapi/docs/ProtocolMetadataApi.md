# \ProtocolMetadataApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLifetimeSettings**](ProtocolMetadataApi.md#GetLifetimeSettings) | **Get** /protocolMetadata/lifetimeSettings | Get metadata cache duration and reload delay for automated reloading.
[**GetSigningSettings**](ProtocolMetadataApi.md#GetSigningSettings) | **Get** /protocolMetadata/signingSettings | Get the certificate ID and algorithm used for metadata signing.
[**UpdateLifetimeSettings**](ProtocolMetadataApi.md#UpdateLifetimeSettings) | **Put** /protocolMetadata/lifetimeSettings | Update metadata cache duration and reload delay for automated reloading.
[**UpdateSigningSettings**](ProtocolMetadataApi.md#UpdateSigningSettings) | **Put** /protocolMetadata/signingSettings | Update the certificate and algorithm for metadata signing.



## GetLifetimeSettings

> MetadataLifetimeSettings GetLifetimeSettings(ctx).Execute()

Get metadata cache duration and reload delay for automated reloading.

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
    resp, r, err := apiClient.ProtocolMetadataApi.GetLifetimeSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMetadataApi.GetLifetimeSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLifetimeSettings`: MetadataLifetimeSettings
    fmt.Fprintf(os.Stdout, "Response from `ProtocolMetadataApi.GetLifetimeSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLifetimeSettingsRequest struct via the builder pattern


### Return type

[**MetadataLifetimeSettings**](MetadataLifetimeSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSigningSettings

> MetadataSigningSettings GetSigningSettings(ctx).Execute()

Get the certificate ID and algorithm used for metadata signing.

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
    resp, r, err := apiClient.ProtocolMetadataApi.GetSigningSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMetadataApi.GetSigningSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSigningSettings`: MetadataSigningSettings
    fmt.Fprintf(os.Stdout, "Response from `ProtocolMetadataApi.GetSigningSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSigningSettingsRequest struct via the builder pattern


### Return type

[**MetadataSigningSettings**](MetadataSigningSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLifetimeSettings

> MetadataLifetimeSettings UpdateLifetimeSettings(ctx).Body(body).Execute()

Update metadata cache duration and reload delay for automated reloading.

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
    body := *openapiclient.NewMetadataLifetimeSettings() // MetadataLifetimeSettings | Metadata lifetime settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProtocolMetadataApi.UpdateLifetimeSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMetadataApi.UpdateLifetimeSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLifetimeSettings`: MetadataLifetimeSettings
    fmt.Fprintf(os.Stdout, "Response from `ProtocolMetadataApi.UpdateLifetimeSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLifetimeSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MetadataLifetimeSettings**](MetadataLifetimeSettings.md) | Metadata lifetime settings. | 

### Return type

[**MetadataLifetimeSettings**](MetadataLifetimeSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSigningSettings

> MetadataSigningSettings UpdateSigningSettings(ctx).Body(body).Execute()

Update the certificate and algorithm for metadata signing.

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
    body := *openapiclient.NewMetadataSigningSettings() // MetadataSigningSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProtocolMetadataApi.UpdateSigningSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtocolMetadataApi.UpdateSigningSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSigningSettings`: MetadataSigningSettings
    fmt.Fprintf(os.Stdout, "Response from `ProtocolMetadataApi.UpdateSigningSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSigningSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MetadataSigningSettings**](MetadataSigningSettings.md) |  | 

### Return type

[**MetadataSigningSettings**](MetadataSigningSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

