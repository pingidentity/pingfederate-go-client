# \ExtendedPropertiesApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExtendedProperties**](ExtendedPropertiesApi.md#GetExtendedProperties) | **Get** /extendedProperties | Get the defined Extended Properties.
[**UpdateExtendedProperties**](ExtendedPropertiesApi.md#UpdateExtendedProperties) | **Put** /extendedProperties | Update the Extended Properties.



## GetExtendedProperties

> ExtendedProperties GetExtendedProperties(ctx).Execute()

Get the defined Extended Properties.

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
    resp, r, err := apiClient.ExtendedPropertiesApi.GetExtendedProperties(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtendedPropertiesApi.GetExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExtendedProperties`: ExtendedProperties
    fmt.Fprintf(os.Stdout, "Response from `ExtendedPropertiesApi.GetExtendedProperties`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtendedPropertiesRequest struct via the builder pattern


### Return type

[**ExtendedProperties**](ExtendedProperties.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExtendedProperties

> ExtendedProperties UpdateExtendedProperties(ctx).Body(body).Execute()

Update the Extended Properties.

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
    body := *openapiclient.NewExtendedProperties() // ExtendedProperties | Definition of extended properties.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtendedPropertiesApi.UpdateExtendedProperties(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtendedPropertiesApi.UpdateExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExtendedProperties`: ExtendedProperties
    fmt.Fprintf(os.Stdout, "Response from `ExtendedPropertiesApi.UpdateExtendedProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExtendedPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ExtendedProperties**](ExtendedProperties.md) | Definition of extended properties. | 

### Return type

[**ExtendedProperties**](ExtendedProperties.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

