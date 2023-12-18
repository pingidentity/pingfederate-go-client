# \OauthClientSettingsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClientSettings**](OauthClientSettingsAPI.md#GetClientSettings) | **Get** /oauth/clientSettings | Configure the client settings.
[**UpdateClientSettings**](OauthClientSettingsAPI.md#UpdateClientSettings) | **Put** /oauth/clientSettings | Update the client settings.



## GetClientSettings

> ClientSettings GetClientSettings(ctx).Execute()

Configure the client settings.

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
    resp, r, err := apiClient.OauthClientSettingsAPI.GetClientSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientSettingsAPI.GetClientSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientSettings`: ClientSettings
    fmt.Fprintf(os.Stdout, "Response from `OauthClientSettingsAPI.GetClientSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientSettingsRequest struct via the builder pattern


### Return type

[**ClientSettings**](ClientSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClientSettings

> ClientSettings UpdateClientSettings(ctx).Body(body).Execute()

Update the client settings.

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
    body := *openapiclient.NewClientSettings() // ClientSettings | Configuration for client settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthClientSettingsAPI.UpdateClientSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientSettingsAPI.UpdateClientSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClientSettings`: ClientSettings
    fmt.Fprintf(os.Stdout, "Response from `OauthClientSettingsAPI.UpdateClientSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClientSettings**](ClientSettings.md) | Configuration for client settings. | 

### Return type

[**ClientSettings**](ClientSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

