# \AdministrativeApiAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCorsSettings**](AdministrativeApiAPI.md#GetCorsSettings) | **Get** /administrativeApi/corsSettings | Get the settings of the current administrative api cors settings.
[**UpdateCorsSettings**](AdministrativeApiAPI.md#UpdateCorsSettings) | **Put** /administrativeApi/corsSettings | Update the administrative api cors settings.



## GetCorsSettings

> AdministrativeApiCorsSettings GetCorsSettings(ctx).Execute()

Get the settings of the current administrative api cors settings.

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
    resp, r, err := apiClient.AdministrativeApiAPI.GetCorsSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrativeApiAPI.GetCorsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCorsSettings`: AdministrativeApiCorsSettings
    fmt.Fprintf(os.Stdout, "Response from `AdministrativeApiAPI.GetCorsSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCorsSettingsRequest struct via the builder pattern


### Return type

[**AdministrativeApiCorsSettings**](AdministrativeApiCorsSettings.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCorsSettings

> AdministrativeApiCorsSettings UpdateCorsSettings(ctx).Body(body).Execute()

Update the administrative api cors settings.

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
    body := *openapiclient.NewAdministrativeApiCorsSettings() // AdministrativeApiCorsSettings | The settings to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministrativeApiAPI.UpdateCorsSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrativeApiAPI.UpdateCorsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCorsSettings`: AdministrativeApiCorsSettings
    fmt.Fprintf(os.Stdout, "Response from `AdministrativeApiAPI.UpdateCorsSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCorsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AdministrativeApiCorsSettings**](AdministrativeApiCorsSettings.md) | The settings to update | 

### Return type

[**AdministrativeApiCorsSettings**](AdministrativeApiCorsSettings.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

