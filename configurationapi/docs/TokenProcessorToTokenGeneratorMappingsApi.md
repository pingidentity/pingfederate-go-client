# \TokenProcessorToTokenGeneratorMappingsApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTokenToTokenMapping**](TokenProcessorToTokenGeneratorMappingsApi.md#CreateTokenToTokenMapping) | **Post** /tokenProcessorToTokenGeneratorMappings | Create a new Token Processor to Token Generator Mapping.
[**DeleteTokenToTokenMappingById**](TokenProcessorToTokenGeneratorMappingsApi.md#DeleteTokenToTokenMappingById) | **Delete** /tokenProcessorToTokenGeneratorMappings/{id} | Delete a Token Processor to Token Generator Mapping.
[**GetTokenToTokenMappingById**](TokenProcessorToTokenGeneratorMappingsApi.md#GetTokenToTokenMappingById) | **Get** /tokenProcessorToTokenGeneratorMappings/{id} | Get a Token Processor to Token Generator Mapping.
[**GetTokenToTokenMappings**](TokenProcessorToTokenGeneratorMappingsApi.md#GetTokenToTokenMappings) | **Get** /tokenProcessorToTokenGeneratorMappings | Get the list of Token Processor to Token Generator Mappings.
[**UpdateTokenToTokenMappingById**](TokenProcessorToTokenGeneratorMappingsApi.md#UpdateTokenToTokenMappingById) | **Put** /tokenProcessorToTokenGeneratorMappings/{id} | Update a Token Processor to Token Generator Mapping.



## CreateTokenToTokenMapping

> TokenToTokenMapping CreateTokenToTokenMapping(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new Token Processor to Token Generator Mapping.

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
    body := *openapiclient.NewTokenToTokenMapping(map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}, "SourceId_example", "TargetId_example") // TokenToTokenMapping | Configuration for a new Token Processor to Token Generator Mapping.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenProcessorToTokenGeneratorMappingsApi.CreateTokenToTokenMapping(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenProcessorToTokenGeneratorMappingsApi.CreateTokenToTokenMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTokenToTokenMapping`: TokenToTokenMapping
    fmt.Fprintf(os.Stdout, "Response from `TokenProcessorToTokenGeneratorMappingsApi.CreateTokenToTokenMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenToTokenMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TokenToTokenMapping**](TokenToTokenMapping.md) | Configuration for a new Token Processor to Token Generator Mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**TokenToTokenMapping**](TokenToTokenMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTokenToTokenMappingById

> DeleteTokenToTokenMappingById(ctx, id).Execute()

Delete a Token Processor to Token Generator Mapping.

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
    id := "id_example" // string | ID of Token Processor to Token Generator Mapping to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TokenProcessorToTokenGeneratorMappingsApi.DeleteTokenToTokenMappingById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenProcessorToTokenGeneratorMappingsApi.DeleteTokenToTokenMappingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Token Processor to Token Generator Mapping to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenToTokenMappingByIdRequest struct via the builder pattern


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


## GetTokenToTokenMappingById

> TokenToTokenMapping GetTokenToTokenMappingById(ctx, id).Execute()

Get a Token Processor to Token Generator Mapping.

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
    id := "id_example" // string | ID of Token Processor to Token Generator Mapping to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenProcessorToTokenGeneratorMappingsApi.GetTokenToTokenMappingById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenProcessorToTokenGeneratorMappingsApi.GetTokenToTokenMappingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenToTokenMappingById`: TokenToTokenMapping
    fmt.Fprintf(os.Stdout, "Response from `TokenProcessorToTokenGeneratorMappingsApi.GetTokenToTokenMappingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Token Processor to Token Generator Mapping to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenToTokenMappingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenToTokenMapping**](TokenToTokenMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenToTokenMappings

> TokenToTokenMappings GetTokenToTokenMappings(ctx).Execute()

Get the list of Token Processor to Token Generator Mappings.

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
    resp, r, err := apiClient.TokenProcessorToTokenGeneratorMappingsApi.GetTokenToTokenMappings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenProcessorToTokenGeneratorMappingsApi.GetTokenToTokenMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenToTokenMappings`: TokenToTokenMappings
    fmt.Fprintf(os.Stdout, "Response from `TokenProcessorToTokenGeneratorMappingsApi.GetTokenToTokenMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenToTokenMappingsRequest struct via the builder pattern


### Return type

[**TokenToTokenMappings**](TokenToTokenMappings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTokenToTokenMappingById

> TokenToTokenMapping UpdateTokenToTokenMappingById(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update a Token Processor to Token Generator Mapping.

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
    id := "id_example" // string | ID of Token Processor to Token Generator Mapping to update.
    body := *openapiclient.NewTokenToTokenMapping(map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}, "SourceId_example", "TargetId_example") // TokenToTokenMapping | Configuration for updated Token Processor to Token Generator Mapping.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenProcessorToTokenGeneratorMappingsApi.UpdateTokenToTokenMappingById(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenProcessorToTokenGeneratorMappingsApi.UpdateTokenToTokenMappingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTokenToTokenMappingById`: TokenToTokenMapping
    fmt.Fprintf(os.Stdout, "Response from `TokenProcessorToTokenGeneratorMappingsApi.UpdateTokenToTokenMappingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Token Processor to Token Generator Mapping to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTokenToTokenMappingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TokenToTokenMapping**](TokenToTokenMapping.md) | Configuration for updated Token Processor to Token Generator Mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**TokenToTokenMapping**](TokenToTokenMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

