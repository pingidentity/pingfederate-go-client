# \SpAuthenticationPolicyContractMappingsAPI

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApcToSpAdapterMapping**](SpAuthenticationPolicyContractMappingsAPI.md#CreateApcToSpAdapterMapping) | **Post** /sp/authenticationPolicyContractMappings | Create a new APC-to-SP Adapter Mapping.
[**DeleteApcToSpAdapterMappingById**](SpAuthenticationPolicyContractMappingsAPI.md#DeleteApcToSpAdapterMappingById) | **Delete** /sp/authenticationPolicyContractMappings/{id} | Delete an APC-to-SP Adapter Mapping.
[**GetApcToSpAdapterMappingById**](SpAuthenticationPolicyContractMappingsAPI.md#GetApcToSpAdapterMappingById) | **Get** /sp/authenticationPolicyContractMappings/{id} | Get an APC-to-SP Adapter Mapping.
[**GetApcToSpAdapterMappings**](SpAuthenticationPolicyContractMappingsAPI.md#GetApcToSpAdapterMappings) | **Get** /sp/authenticationPolicyContractMappings | Get the list of APC-to-SP Adapter Mappings.
[**UpdateApcToSpAdapterMappingById**](SpAuthenticationPolicyContractMappingsAPI.md#UpdateApcToSpAdapterMappingById) | **Put** /sp/authenticationPolicyContractMappings/{id} | Update an APC-to-SP Adapter Mapping.



## CreateApcToSpAdapterMapping

> ApcToSpAdapterMapping CreateApcToSpAdapterMapping(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new APC-to-SP Adapter Mapping.

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
    body := *openapiclient.NewApcToSpAdapterMapping(map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}, "SourceId_example", "TargetId_example") // ApcToSpAdapterMapping | Configuration for a new APC-to-SP Adapter Mapping.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpAuthenticationPolicyContractMappingsAPI.CreateApcToSpAdapterMapping(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpAuthenticationPolicyContractMappingsAPI.CreateApcToSpAdapterMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApcToSpAdapterMapping`: ApcToSpAdapterMapping
    fmt.Fprintf(os.Stdout, "Response from `SpAuthenticationPolicyContractMappingsAPI.CreateApcToSpAdapterMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApcToSpAdapterMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApcToSpAdapterMapping**](ApcToSpAdapterMapping.md) | Configuration for a new APC-to-SP Adapter Mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**ApcToSpAdapterMapping**](ApcToSpAdapterMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApcToSpAdapterMappingById

> DeleteApcToSpAdapterMappingById(ctx, id).Execute()

Delete an APC-to-SP Adapter Mapping.

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
    id := "id_example" // string | ID of APC-to-SP Adapter Mapping to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpAuthenticationPolicyContractMappingsAPI.DeleteApcToSpAdapterMappingById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpAuthenticationPolicyContractMappingsAPI.DeleteApcToSpAdapterMappingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of APC-to-SP Adapter Mapping to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApcToSpAdapterMappingByIdRequest struct via the builder pattern


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


## GetApcToSpAdapterMappingById

> ApcToSpAdapterMapping GetApcToSpAdapterMappingById(ctx, id).Execute()

Get an APC-to-SP Adapter Mapping.

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
    id := "id_example" // string | ID of APC-to-SP Adapter Mapping to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpAuthenticationPolicyContractMappingsAPI.GetApcToSpAdapterMappingById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpAuthenticationPolicyContractMappingsAPI.GetApcToSpAdapterMappingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApcToSpAdapterMappingById`: ApcToSpAdapterMapping
    fmt.Fprintf(os.Stdout, "Response from `SpAuthenticationPolicyContractMappingsAPI.GetApcToSpAdapterMappingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of APC-to-SP Adapter Mapping to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApcToSpAdapterMappingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApcToSpAdapterMapping**](ApcToSpAdapterMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApcToSpAdapterMappings

> ApcToSpAdapterMappings GetApcToSpAdapterMappings(ctx).Execute()

Get the list of APC-to-SP Adapter Mappings.

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
    resp, r, err := apiClient.SpAuthenticationPolicyContractMappingsAPI.GetApcToSpAdapterMappings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpAuthenticationPolicyContractMappingsAPI.GetApcToSpAdapterMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApcToSpAdapterMappings`: ApcToSpAdapterMappings
    fmt.Fprintf(os.Stdout, "Response from `SpAuthenticationPolicyContractMappingsAPI.GetApcToSpAdapterMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApcToSpAdapterMappingsRequest struct via the builder pattern


### Return type

[**ApcToSpAdapterMappings**](ApcToSpAdapterMappings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApcToSpAdapterMappingById

> ApcToSpAdapterMapping UpdateApcToSpAdapterMappingById(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update an APC-to-SP Adapter Mapping.

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
    id := "id_example" // string | ID of APC-to-SP Adapter Mapping to update.
    body := *openapiclient.NewApcToSpAdapterMapping(map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}, "SourceId_example", "TargetId_example") // ApcToSpAdapterMapping | Configuration for updated APC-to-SP Adapter Mapping.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpAuthenticationPolicyContractMappingsAPI.UpdateApcToSpAdapterMappingById(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpAuthenticationPolicyContractMappingsAPI.UpdateApcToSpAdapterMappingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApcToSpAdapterMappingById`: ApcToSpAdapterMapping
    fmt.Fprintf(os.Stdout, "Response from `SpAuthenticationPolicyContractMappingsAPI.UpdateApcToSpAdapterMappingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of APC-to-SP Adapter Mapping to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApcToSpAdapterMappingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApcToSpAdapterMapping**](ApcToSpAdapterMapping.md) | Configuration for updated APC-to-SP Adapter Mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**ApcToSpAdapterMapping**](ApcToSpAdapterMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

