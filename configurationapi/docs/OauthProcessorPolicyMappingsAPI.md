# \OauthProcessorPolicyMappingsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTeppMapping**](OauthProcessorPolicyMappingsAPI.md#CreateTeppMapping) | **Post** /oauth/processorPolicyMappings | Create a new processor policy to persistent grant mapping.
[**DeleteTeppMapping**](OauthProcessorPolicyMappingsAPI.md#DeleteTeppMapping) | **Delete** /oauth/processorPolicyMappings/{id} | Delete a processor policy to persistent grant mapping.
[**GetProcessorPolicyMapping**](OauthProcessorPolicyMappingsAPI.md#GetProcessorPolicyMapping) | **Get** /oauth/processorPolicyMappings/{id} | Find the processor policy to persistent grant mapping by ID.
[**GetProcessorPolicyMappings**](OauthProcessorPolicyMappingsAPI.md#GetProcessorPolicyMappings) | **Get** /oauth/processorPolicyMappings | Get the list of processor policy to persistent grant mappings.
[**UpdateTeppMapping**](OauthProcessorPolicyMappingsAPI.md#UpdateTeppMapping) | **Put** /oauth/processorPolicyMappings/{id} | Update a processor policy to persistent grant mapping.



## CreateTeppMapping

> ProcessorPolicyToPersistentGrantMapping CreateTeppMapping(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new processor policy to persistent grant mapping.



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
    body := *openapiclient.NewProcessorPolicyToPersistentGrantMapping(*openapiclient.NewResourceLink("Id_example"), map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}) // ProcessorPolicyToPersistentGrantMapping | Configuration for processor policy to persistent grant mapping.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthProcessorPolicyMappingsAPI.CreateTeppMapping(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthProcessorPolicyMappingsAPI.CreateTeppMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTeppMapping`: ProcessorPolicyToPersistentGrantMapping
    fmt.Fprintf(os.Stdout, "Response from `OauthProcessorPolicyMappingsAPI.CreateTeppMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeppMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProcessorPolicyToPersistentGrantMapping**](ProcessorPolicyToPersistentGrantMapping.md) | Configuration for processor policy to persistent grant mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**ProcessorPolicyToPersistentGrantMapping**](ProcessorPolicyToPersistentGrantMapping.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeppMapping

> DeleteTeppMapping(ctx, id).Execute()

Delete a processor policy to persistent grant mapping.

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
    id := "id_example" // string | ID of the processor policy to persistent grant mapping.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthProcessorPolicyMappingsAPI.DeleteTeppMapping(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthProcessorPolicyMappingsAPI.DeleteTeppMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the processor policy to persistent grant mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeppMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessorPolicyMapping

> ProcessorPolicyToPersistentGrantMapping GetProcessorPolicyMapping(ctx, id).Execute()

Find the processor policy to persistent grant mapping by ID.

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
    id := "id_example" // string | ID of the processor policy to persistent grant mapping.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthProcessorPolicyMappingsAPI.GetProcessorPolicyMapping(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthProcessorPolicyMappingsAPI.GetProcessorPolicyMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcessorPolicyMapping`: ProcessorPolicyToPersistentGrantMapping
    fmt.Fprintf(os.Stdout, "Response from `OauthProcessorPolicyMappingsAPI.GetProcessorPolicyMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the processor policy to persistent grant mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessorPolicyMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessorPolicyToPersistentGrantMapping**](ProcessorPolicyToPersistentGrantMapping.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessorPolicyMappings

> ProcessorPolicyToPersistentGrantMappings GetProcessorPolicyMappings(ctx).Execute()

Get the list of processor policy to persistent grant mappings.

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
    resp, r, err := apiClient.OauthProcessorPolicyMappingsAPI.GetProcessorPolicyMappings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthProcessorPolicyMappingsAPI.GetProcessorPolicyMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcessorPolicyMappings`: ProcessorPolicyToPersistentGrantMappings
    fmt.Fprintf(os.Stdout, "Response from `OauthProcessorPolicyMappingsAPI.GetProcessorPolicyMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessorPolicyMappingsRequest struct via the builder pattern


### Return type

[**ProcessorPolicyToPersistentGrantMappings**](ProcessorPolicyToPersistentGrantMappings.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeppMapping

> ProcessorPolicyToPersistentGrantMapping UpdateTeppMapping(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update a processor policy to persistent grant mapping.

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
    id := "id_example" // string | ID of the processor policy to persistent grant mapping to update.
    body := *openapiclient.NewProcessorPolicyToPersistentGrantMapping(*openapiclient.NewResourceLink("Id_example"), map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}) // ProcessorPolicyToPersistentGrantMapping | Configuration for a processor policy to persistent grant mapping.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthProcessorPolicyMappingsAPI.UpdateTeppMapping(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthProcessorPolicyMappingsAPI.UpdateTeppMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeppMapping`: ProcessorPolicyToPersistentGrantMapping
    fmt.Fprintf(os.Stdout, "Response from `OauthProcessorPolicyMappingsAPI.UpdateTeppMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the processor policy to persistent grant mapping to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeppMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProcessorPolicyToPersistentGrantMapping**](ProcessorPolicyToPersistentGrantMapping.md) | Configuration for a processor policy to persistent grant mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**ProcessorPolicyToPersistentGrantMapping**](ProcessorPolicyToPersistentGrantMapping.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

