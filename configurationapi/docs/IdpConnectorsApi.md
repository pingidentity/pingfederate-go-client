# \IdpConnectorsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIdpConnectorDescriptorById**](IdpConnectorsAPI.md#GetIdpConnectorDescriptorById) | **Get** /idp/connectors/descriptors/{id} | Get the list of available connector descriptors.
[**GetIdpConnectorDescriptors**](IdpConnectorsAPI.md#GetIdpConnectorDescriptors) | **Get** /idp/connectors/descriptors | Get the list of available IdP connector descriptors.



## GetIdpConnectorDescriptorById

> SaasPluginDescriptor GetIdpConnectorDescriptorById(ctx, id).Execute()

Get the list of available connector descriptors.

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
    id := "id_example" // string | the type of connector descriptor to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpConnectorsAPI.GetIdpConnectorDescriptorById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpConnectorsAPI.GetIdpConnectorDescriptorById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdpConnectorDescriptorById`: SaasPluginDescriptor
    fmt.Fprintf(os.Stdout, "Response from `IdpConnectorsAPI.GetIdpConnectorDescriptorById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the type of connector descriptor to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpConnectorDescriptorByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SaasPluginDescriptor**](SaasPluginDescriptor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdpConnectorDescriptors

> SaasPluginDescriptors GetIdpConnectorDescriptors(ctx).Execute()

Get the list of available IdP connector descriptors.

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
    resp, r, err := apiClient.IdpConnectorsAPI.GetIdpConnectorDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpConnectorsAPI.GetIdpConnectorDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdpConnectorDescriptors`: SaasPluginDescriptors
    fmt.Fprintf(os.Stdout, "Response from `IdpConnectorsAPI.GetIdpConnectorDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpConnectorDescriptorsRequest struct via the builder pattern


### Return type

[**SaasPluginDescriptors**](SaasPluginDescriptors.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

