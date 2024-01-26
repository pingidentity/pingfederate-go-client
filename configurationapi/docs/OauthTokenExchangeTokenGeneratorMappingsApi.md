# \OauthTokenExchangeTokenGeneratorMappingsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTokenGeneratorMapping**](OauthTokenExchangeTokenGeneratorMappingsAPI.md#CreateTokenGeneratorMapping) | **Post** /oauth/tokenExchange/tokenGeneratorMappings | Create a new Token Exchange Processor policy to Token Generator Mapping.
[**DeleteTokenGeneratorMappingById**](OauthTokenExchangeTokenGeneratorMappingsAPI.md#DeleteTokenGeneratorMappingById) | **Delete** /oauth/tokenExchange/tokenGeneratorMappings/{id} | Delete a Token Exchange Processor policy to Token Generator Mapping.
[**GetTokenGeneratorMappingById**](OauthTokenExchangeTokenGeneratorMappingsAPI.md#GetTokenGeneratorMappingById) | **Get** /oauth/tokenExchange/tokenGeneratorMappings/{id} | Get a Token Exchange Processor policy to Token Generator Mapping.
[**GetTokenGeneratorMappings**](OauthTokenExchangeTokenGeneratorMappingsAPI.md#GetTokenGeneratorMappings) | **Get** /oauth/tokenExchange/tokenGeneratorMappings | Get the list of Token Exchange Processor policy to Token Generator Mappings.
[**UpdateTokenGeneratorMappingById**](OauthTokenExchangeTokenGeneratorMappingsAPI.md#UpdateTokenGeneratorMappingById) | **Put** /oauth/tokenExchange/tokenGeneratorMappings/{id} | Update a Token Exchange Processor policy to Token Generator Mapping.



## CreateTokenGeneratorMapping

> ProcessorPolicyToGeneratorMapping CreateTokenGeneratorMapping(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new Token Exchange Processor policy to Token Generator Mapping.

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
    body := *openapiclient.NewProcessorPolicyToGeneratorMapping(map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}, "SourceId_example", "TargetId_example") // ProcessorPolicyToGeneratorMapping | Configuration for a new Token Exchange Processor policy to Token Generator Mapping.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthTokenExchangeTokenGeneratorMappingsAPI.CreateTokenGeneratorMapping(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeTokenGeneratorMappingsAPI.CreateTokenGeneratorMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTokenGeneratorMapping`: ProcessorPolicyToGeneratorMapping
    fmt.Fprintf(os.Stdout, "Response from `OauthTokenExchangeTokenGeneratorMappingsAPI.CreateTokenGeneratorMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenGeneratorMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProcessorPolicyToGeneratorMapping**](ProcessorPolicyToGeneratorMapping.md) | Configuration for a new Token Exchange Processor policy to Token Generator Mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**ProcessorPolicyToGeneratorMapping**](ProcessorPolicyToGeneratorMapping.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTokenGeneratorMappingById

> DeleteTokenGeneratorMappingById(ctx, id).Execute()

Delete a Token Exchange Processor policy to Token Generator Mapping.

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
    id := "id_example" // string | ID of Token Exchange Processor policy to Token Generator Mapping to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthTokenExchangeTokenGeneratorMappingsAPI.DeleteTokenGeneratorMappingById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeTokenGeneratorMappingsAPI.DeleteTokenGeneratorMappingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Token Exchange Processor policy to Token Generator Mapping to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenGeneratorMappingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenGeneratorMappingById

> ProcessorPolicyToGeneratorMapping GetTokenGeneratorMappingById(ctx, id).Execute()

Get a Token Exchange Processor policy to Token Generator Mapping.

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
    id := "id_example" // string | ID of Token Exchange Processor policy to Token Generator Mapping to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthTokenExchangeTokenGeneratorMappingsAPI.GetTokenGeneratorMappingById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeTokenGeneratorMappingsAPI.GetTokenGeneratorMappingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenGeneratorMappingById`: ProcessorPolicyToGeneratorMapping
    fmt.Fprintf(os.Stdout, "Response from `OauthTokenExchangeTokenGeneratorMappingsAPI.GetTokenGeneratorMappingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Token Exchange Processor policy to Token Generator Mapping to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenGeneratorMappingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessorPolicyToGeneratorMapping**](ProcessorPolicyToGeneratorMapping.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenGeneratorMappings

> ProcessorPolicyToGeneratorMappings GetTokenGeneratorMappings(ctx).Execute()

Get the list of Token Exchange Processor policy to Token Generator Mappings.

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
    resp, r, err := apiClient.OauthTokenExchangeTokenGeneratorMappingsAPI.GetTokenGeneratorMappings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeTokenGeneratorMappingsAPI.GetTokenGeneratorMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenGeneratorMappings`: ProcessorPolicyToGeneratorMappings
    fmt.Fprintf(os.Stdout, "Response from `OauthTokenExchangeTokenGeneratorMappingsAPI.GetTokenGeneratorMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenGeneratorMappingsRequest struct via the builder pattern


### Return type

[**ProcessorPolicyToGeneratorMappings**](ProcessorPolicyToGeneratorMappings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTokenGeneratorMappingById

> ProcessorPolicyToGeneratorMapping UpdateTokenGeneratorMappingById(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update a Token Exchange Processor policy to Token Generator Mapping.

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
    id := "id_example" // string | ID of Token Exchange Processor policy to Token Generator Mapping to update.
    body := *openapiclient.NewProcessorPolicyToGeneratorMapping(map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}, "SourceId_example", "TargetId_example") // ProcessorPolicyToGeneratorMapping | Configuration for updated Token Exchange Processor policy to Token Generator Mapping.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthTokenExchangeTokenGeneratorMappingsAPI.UpdateTokenGeneratorMappingById(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeTokenGeneratorMappingsAPI.UpdateTokenGeneratorMappingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTokenGeneratorMappingById`: ProcessorPolicyToGeneratorMapping
    fmt.Fprintf(os.Stdout, "Response from `OauthTokenExchangeTokenGeneratorMappingsAPI.UpdateTokenGeneratorMappingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Token Exchange Processor policy to Token Generator Mapping to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTokenGeneratorMappingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProcessorPolicyToGeneratorMapping**](ProcessorPolicyToGeneratorMapping.md) | Configuration for updated Token Exchange Processor policy to Token Generator Mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**ProcessorPolicyToGeneratorMapping**](ProcessorPolicyToGeneratorMapping.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

