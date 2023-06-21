# \ServiceAuthenticationApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceAuthentication**](ServiceAuthenticationApi.md#GetServiceAuthentication) | **Get** /serviceAuthentication | Get the service authentication settings.
[**UpdateServiceAuthentication**](ServiceAuthenticationApi.md#UpdateServiceAuthentication) | **Put** /serviceAuthentication | Update the service authentication settings.



## GetServiceAuthentication

> ServiceAuthentication GetServiceAuthentication(ctx).Execute()

Get the service authentication settings.

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
    resp, r, err := apiClient.ServiceAuthenticationApi.GetServiceAuthentication(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAuthenticationApi.GetServiceAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceAuthentication`: ServiceAuthentication
    fmt.Fprintf(os.Stdout, "Response from `ServiceAuthenticationApi.GetServiceAuthentication`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAuthenticationRequest struct via the builder pattern


### Return type

[**ServiceAuthentication**](ServiceAuthentication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceAuthentication

> ServiceAuthentication UpdateServiceAuthentication(ctx).Body(body).Execute()

Update the service authentication settings.



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
    body := *openapiclient.NewServiceAuthentication() // ServiceAuthentication | Service authentication settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAuthenticationApi.UpdateServiceAuthentication(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAuthenticationApi.UpdateServiceAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceAuthentication`: ServiceAuthentication
    fmt.Fprintf(os.Stdout, "Response from `ServiceAuthenticationApi.UpdateServiceAuthentication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ServiceAuthentication**](ServiceAuthentication.md) | Service authentication settings. | 

### Return type

[**ServiceAuthentication**](ServiceAuthentication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

