# \OauthClientSettingsApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOauthClientSettings**](OauthClientSettingsApi.md#GetOauthClientSettings) | **Get** /oauth/clientSettings | Configure the client settings.
[**UpdateOauthClientSettings**](OauthClientSettingsApi.md#UpdateOauthClientSettings) | **Put** /oauth/clientSettings | Update the client settings.



## GetOauthClientSettings

> ClientSettings GetOauthClientSettings(ctx).Execute()

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
    resp, r, err := apiClient.OauthClientSettingsApi.GetOauthClientSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientSettingsApi.GetOauthClientSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthClientSettings`: ClientSettings
    fmt.Fprintf(os.Stdout, "Response from `OauthClientSettingsApi.GetOauthClientSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthClientSettingsRequest struct via the builder pattern


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


## UpdateOauthClientSettings

> ClientSettings UpdateOauthClientSettings(ctx).Body(body).Execute()

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
    resp, r, err := apiClient.OauthClientSettingsApi.UpdateOauthClientSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientSettingsApi.UpdateOauthClientSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOauthClientSettings`: ClientSettings
    fmt.Fprintf(os.Stdout, "Response from `OauthClientSettingsApi.UpdateOauthClientSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOauthClientSettingsRequest struct via the builder pattern


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

