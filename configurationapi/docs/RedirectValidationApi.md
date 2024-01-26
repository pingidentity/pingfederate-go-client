# \RedirectValidationAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRedirectValidationSettings**](RedirectValidationAPI.md#GetRedirectValidationSettings) | **Get** /redirectValidation | Retrieve redirect validation settings.
[**UpdateRedirectValidationSettings**](RedirectValidationAPI.md#UpdateRedirectValidationSettings) | **Put** /redirectValidation | Update redirect validation settings.



## GetRedirectValidationSettings

> RedirectValidationSettings GetRedirectValidationSettings(ctx).Execute()

Retrieve redirect validation settings.

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
    resp, r, err := apiClient.RedirectValidationAPI.GetRedirectValidationSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectValidationAPI.GetRedirectValidationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRedirectValidationSettings`: RedirectValidationSettings
    fmt.Fprintf(os.Stdout, "Response from `RedirectValidationAPI.GetRedirectValidationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedirectValidationSettingsRequest struct via the builder pattern


### Return type

[**RedirectValidationSettings**](RedirectValidationSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRedirectValidationSettings

> RedirectValidationSettings UpdateRedirectValidationSettings(ctx).Body(body).Execute()

Update redirect validation settings.



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
    body := *openapiclient.NewRedirectValidationSettings() // RedirectValidationSettings | Redirect validation settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RedirectValidationAPI.UpdateRedirectValidationSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectValidationAPI.UpdateRedirectValidationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRedirectValidationSettings`: RedirectValidationSettings
    fmt.Fprintf(os.Stdout, "Response from `RedirectValidationAPI.UpdateRedirectValidationSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRedirectValidationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RedirectValidationSettings**](RedirectValidationSettings.md) | Redirect validation settings. | 

### Return type

[**RedirectValidationSettings**](RedirectValidationSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

