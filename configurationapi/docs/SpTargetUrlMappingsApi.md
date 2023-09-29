# \SpTargetUrlMappingsAPI

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSpUrlMappings**](SpTargetUrlMappingsAPI.md#GetSpUrlMappings) | **Get** /sp/targetUrlMappings | List the mappings between URLs and adapter or connection instances.
[**UpdateSpUrlMappings**](SpTargetUrlMappingsAPI.md#UpdateSpUrlMappings) | **Put** /sp/targetUrlMappings | Update the mappings between URLs and adapters or connections instances.



## GetSpUrlMappings

> SpUrlMappings GetSpUrlMappings(ctx).Execute()

List the mappings between URLs and adapter or connection instances.

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
    resp, r, err := apiClient.SpTargetUrlMappingsAPI.GetSpUrlMappings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpTargetUrlMappingsAPI.GetSpUrlMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpUrlMappings`: SpUrlMappings
    fmt.Fprintf(os.Stdout, "Response from `SpTargetUrlMappingsAPI.GetSpUrlMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpUrlMappingsRequest struct via the builder pattern


### Return type

[**SpUrlMappings**](SpUrlMappings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpUrlMappings

> SpUrlMappings UpdateSpUrlMappings(ctx).Body(body).Execute()

Update the mappings between URLs and adapters or connections instances.

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
    body := *openapiclient.NewSpUrlMappings() // SpUrlMappings | The SP adapter URL mappings to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpTargetUrlMappingsAPI.UpdateSpUrlMappings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpTargetUrlMappingsAPI.UpdateSpUrlMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSpUrlMappings`: SpUrlMappings
    fmt.Fprintf(os.Stdout, "Response from `SpTargetUrlMappingsAPI.UpdateSpUrlMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpUrlMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SpUrlMappings**](SpUrlMappings.md) | The SP adapter URL mappings to update. | 

### Return type

[**SpUrlMappings**](SpUrlMappings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

