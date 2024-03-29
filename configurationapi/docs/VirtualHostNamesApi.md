# \VirtualHostNamesAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVirtualHostNamesSettings**](VirtualHostNamesAPI.md#GetVirtualHostNamesSettings) | **Get** /virtualHostNames | Retrieve virtual host names settings.
[**UpdateVirtualHostNamesSettings**](VirtualHostNamesAPI.md#UpdateVirtualHostNamesSettings) | **Put** /virtualHostNames | Update virtual host names settings.



## GetVirtualHostNamesSettings

> VirtualHostNameSettings GetVirtualHostNamesSettings(ctx).Execute()

Retrieve virtual host names settings.

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
    resp, r, err := apiClient.VirtualHostNamesAPI.GetVirtualHostNamesSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualHostNamesAPI.GetVirtualHostNamesSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualHostNamesSettings`: VirtualHostNameSettings
    fmt.Fprintf(os.Stdout, "Response from `VirtualHostNamesAPI.GetVirtualHostNamesSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualHostNamesSettingsRequest struct via the builder pattern


### Return type

[**VirtualHostNameSettings**](VirtualHostNameSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualHostNamesSettings

> VirtualHostNameSettings UpdateVirtualHostNamesSettings(ctx).Body(body).Execute()

Update virtual host names settings.

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
    body := *openapiclient.NewVirtualHostNameSettings() // VirtualHostNameSettings | Virtual host names settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualHostNamesAPI.UpdateVirtualHostNamesSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualHostNamesAPI.UpdateVirtualHostNamesSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualHostNamesSettings`: VirtualHostNameSettings
    fmt.Fprintf(os.Stdout, "Response from `VirtualHostNamesAPI.UpdateVirtualHostNamesSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualHostNamesSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VirtualHostNameSettings**](VirtualHostNameSettings.md) | Virtual host names settings. | 

### Return type

[**VirtualHostNameSettings**](VirtualHostNameSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

