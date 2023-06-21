# \IncomingProxySettingsApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIncomingProxySettings**](IncomingProxySettingsApi.md#GetIncomingProxySettings) | **Get** /incomingProxySettings | Get incoming proxy settings.
[**UpdateIncomingProxySettings**](IncomingProxySettingsApi.md#UpdateIncomingProxySettings) | **Put** /incomingProxySettings | Update incoming proxy settings.



## GetIncomingProxySettings

> IncomingProxySettings GetIncomingProxySettings(ctx).Execute()

Get incoming proxy settings.



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
    resp, r, err := apiClient.IncomingProxySettingsApi.GetIncomingProxySettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncomingProxySettingsApi.GetIncomingProxySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncomingProxySettings`: IncomingProxySettings
    fmt.Fprintf(os.Stdout, "Response from `IncomingProxySettingsApi.GetIncomingProxySettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncomingProxySettingsRequest struct via the builder pattern


### Return type

[**IncomingProxySettings**](IncomingProxySettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIncomingProxySettings

> IncomingProxySettings UpdateIncomingProxySettings(ctx).Body(body).Execute()

Update incoming proxy settings.



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
    body := *openapiclient.NewIncomingProxySettings() // IncomingProxySettings | Incoming proxy settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IncomingProxySettingsApi.UpdateIncomingProxySettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncomingProxySettingsApi.UpdateIncomingProxySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIncomingProxySettings`: IncomingProxySettings
    fmt.Fprintf(os.Stdout, "Response from `IncomingProxySettingsApi.UpdateIncomingProxySettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIncomingProxySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IncomingProxySettings**](IncomingProxySettings.md) | Incoming proxy settings. | 

### Return type

[**IncomingProxySettings**](IncomingProxySettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

