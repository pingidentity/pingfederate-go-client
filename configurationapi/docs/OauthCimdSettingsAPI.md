# \OauthCimdSettingsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCimdSettings**](OauthCimdSettingsAPI.md#GetCimdSettings) | **Get** /oauth/cimd/settings | Get the global CIMD settings.
[**UpdateCimdSettings**](OauthCimdSettingsAPI.md#UpdateCimdSettings) | **Put** /oauth/cimd/settings | Update the global CIMD settings.



## GetCimdSettings

> CimdSettings GetCimdSettings(ctx).Execute()

Get the global CIMD settings.

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
    resp, r, err := apiClient.OauthCimdSettingsAPI.GetCimdSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthCimdSettingsAPI.GetCimdSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCimdSettings`: CimdSettings
    fmt.Fprintf(os.Stdout, "Response from `OauthCimdSettingsAPI.GetCimdSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCimdSettingsRequest struct via the builder pattern


### Return type

[**CimdSettings**](CimdSettings.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCimdSettings

> CimdSettings UpdateCimdSettings(ctx).Body(body).Execute()

Update the global CIMD settings.

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
    body := *openapiclient.NewCimdSettings(false, int64(123), int64(123), int64(123), int64(123), int64(123), int64(123)) // CimdSettings | CIMD settings to apply. Replaces the prior settings in full.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthCimdSettingsAPI.UpdateCimdSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthCimdSettingsAPI.UpdateCimdSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCimdSettings`: CimdSettings
    fmt.Fprintf(os.Stdout, "Response from `OauthCimdSettingsAPI.UpdateCimdSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCimdSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CimdSettings**](CimdSettings.md) | CIMD settings to apply. Replaces the prior settings in full. | 

### Return type

[**CimdSettings**](CimdSettings.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

