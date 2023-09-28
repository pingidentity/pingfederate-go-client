# \IdpToSpAdapterMappingApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdpToSpAdapterMapping**](IdpToSpAdapterMappingApi.md#CreateIdpToSpAdapterMapping) | **Post** /idpToSpAdapterMapping | Create a new IdP-to-SP Adapter mapping.
[**DeleteIdpToSpAdapterMappingsById**](IdpToSpAdapterMappingApi.md#DeleteIdpToSpAdapterMappingsById) | **Delete** /idpToSpAdapterMapping/{id} | Delete an Adapter to Adapter Mapping.
[**GetIdpToSpAdapterMappings**](IdpToSpAdapterMappingApi.md#GetIdpToSpAdapterMappings) | **Get** /idpToSpAdapterMapping | Get list of IdP-to-SP Adapter Mappings.
[**GetIdpToSpAdapterMappingsById**](IdpToSpAdapterMappingApi.md#GetIdpToSpAdapterMappingsById) | **Get** /idpToSpAdapterMapping/{id} | Get an IdP-to-SP Adapter Mapping.
[**UpdateIdpToSpAdapterMapping**](IdpToSpAdapterMappingApi.md#UpdateIdpToSpAdapterMapping) | **Put** /idpToSpAdapterMapping/{id} | Update the specified IdP-to-SP Adapter mapping.



## CreateIdpToSpAdapterMapping

> IdpToSpAdapterMapping CreateIdpToSpAdapterMapping(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new IdP-to-SP Adapter mapping.

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
    body := *openapiclient.NewIdpToSpAdapterMapping(map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}, "SourceId_example", "TargetId_example") // IdpToSpAdapterMapping | Configuration for new IdP-to-SP Adapter Mapping.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpToSpAdapterMappingApi.CreateIdpToSpAdapterMapping(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpToSpAdapterMappingApi.CreateIdpToSpAdapterMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdpToSpAdapterMapping`: IdpToSpAdapterMapping
    fmt.Fprintf(os.Stdout, "Response from `IdpToSpAdapterMappingApi.CreateIdpToSpAdapterMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdpToSpAdapterMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IdpToSpAdapterMapping**](IdpToSpAdapterMapping.md) | Configuration for new IdP-to-SP Adapter Mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**IdpToSpAdapterMapping**](IdpToSpAdapterMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdpToSpAdapterMappingsById

> DeleteIdpToSpAdapterMappingsById(ctx, id).Execute()

Delete an Adapter to Adapter Mapping.

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
    id := "id_example" // string | ID of the IdP-to-SP Adapter Mapping to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdpToSpAdapterMappingApi.DeleteIdpToSpAdapterMappingsById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpToSpAdapterMappingApi.DeleteIdpToSpAdapterMappingsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the IdP-to-SP Adapter Mapping to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdpToSpAdapterMappingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdpToSpAdapterMappings

> IdpToSpAdapterMappings GetIdpToSpAdapterMappings(ctx).Execute()

Get list of IdP-to-SP Adapter Mappings.

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
    resp, r, err := apiClient.IdpToSpAdapterMappingApi.GetIdpToSpAdapterMappings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpToSpAdapterMappingApi.GetIdpToSpAdapterMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdpToSpAdapterMappings`: IdpToSpAdapterMappings
    fmt.Fprintf(os.Stdout, "Response from `IdpToSpAdapterMappingApi.GetIdpToSpAdapterMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpToSpAdapterMappingsRequest struct via the builder pattern


### Return type

[**IdpToSpAdapterMappings**](IdpToSpAdapterMappings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdpToSpAdapterMappingsById

> IdpToSpAdapterMapping GetIdpToSpAdapterMappingsById(ctx, id).Execute()

Get an IdP-to-SP Adapter Mapping.

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
    id := "id_example" // string | ID of IdP-to-SP Adapter Mapping to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpToSpAdapterMappingApi.GetIdpToSpAdapterMappingsById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpToSpAdapterMappingApi.GetIdpToSpAdapterMappingsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdpToSpAdapterMappingsById`: IdpToSpAdapterMapping
    fmt.Fprintf(os.Stdout, "Response from `IdpToSpAdapterMappingApi.GetIdpToSpAdapterMappingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of IdP-to-SP Adapter Mapping to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpToSpAdapterMappingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdpToSpAdapterMapping**](IdpToSpAdapterMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdpToSpAdapterMapping

> IdpToSpAdapterMapping UpdateIdpToSpAdapterMapping(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update the specified IdP-to-SP Adapter mapping.

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
    id := "id_example" // string | ID of source adapter in the IdP-to-SP Adapter Mapping to fetch.
    body := *openapiclient.NewIdpToSpAdapterMapping(map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}, "SourceId_example", "TargetId_example") // IdpToSpAdapterMapping | Configuration for updated IdP-to-SP Adapter Mapping.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpToSpAdapterMappingApi.UpdateIdpToSpAdapterMapping(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpToSpAdapterMappingApi.UpdateIdpToSpAdapterMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdpToSpAdapterMapping`: IdpToSpAdapterMapping
    fmt.Fprintf(os.Stdout, "Response from `IdpToSpAdapterMappingApi.UpdateIdpToSpAdapterMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of source adapter in the IdP-to-SP Adapter Mapping to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdpToSpAdapterMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IdpToSpAdapterMapping**](IdpToSpAdapterMapping.md) | Configuration for updated IdP-to-SP Adapter Mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**IdpToSpAdapterMapping**](IdpToSpAdapterMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

