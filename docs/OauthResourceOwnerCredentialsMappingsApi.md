# \OauthResourceOwnerCredentialsMappingsApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceOwnerCredentialsMapping**](OauthResourceOwnerCredentialsMappingsApi.md#CreateResourceOwnerCredentialsMapping) | **Post** /oauth/resourceOwnerCredentialsMappings | Create a new Resource Owner Credentials mapping.
[**DeleteResourceOwnerCredentialsMapping**](OauthResourceOwnerCredentialsMappingsApi.md#DeleteResourceOwnerCredentialsMapping) | **Delete** /oauth/resourceOwnerCredentialsMappings/{id} | Delete a Resource Owner Credentials mapping.
[**GetResourceOwnerCredentialsMapping**](OauthResourceOwnerCredentialsMappingsApi.md#GetResourceOwnerCredentialsMapping) | **Get** /oauth/resourceOwnerCredentialsMappings/{id} | Find the Resource Owner Credentials mapping by the ID.
[**GetResourceOwnerCredentialsMappings**](OauthResourceOwnerCredentialsMappingsApi.md#GetResourceOwnerCredentialsMappings) | **Get** /oauth/resourceOwnerCredentialsMappings | Get the list of Resource Owner Credentials Grant Mapping.
[**UpdateResourceOwnerCredentialsMapping**](OauthResourceOwnerCredentialsMappingsApi.md#UpdateResourceOwnerCredentialsMapping) | **Put** /oauth/resourceOwnerCredentialsMappings/{id} | Update a Resource Owner Credentials mapping.



## CreateResourceOwnerCredentialsMapping

> ResourceOwnerCredentialsMapping CreateResourceOwnerCredentialsMapping(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new Resource Owner Credentials mapping.



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
    body := *openapiclient.NewResourceOwnerCredentialsMapping("Id_example", map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}) // ResourceOwnerCredentialsMapping | Configuration for Resource Owner Credentials mapping.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthResourceOwnerCredentialsMappingsApi.CreateResourceOwnerCredentialsMapping(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthResourceOwnerCredentialsMappingsApi.CreateResourceOwnerCredentialsMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateResourceOwnerCredentialsMapping`: ResourceOwnerCredentialsMapping
    fmt.Fprintf(os.Stdout, "Response from `OauthResourceOwnerCredentialsMappingsApi.CreateResourceOwnerCredentialsMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceOwnerCredentialsMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ResourceOwnerCredentialsMapping**](ResourceOwnerCredentialsMapping.md) | Configuration for Resource Owner Credentials mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**ResourceOwnerCredentialsMapping**](ResourceOwnerCredentialsMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourceOwnerCredentialsMapping

> DeleteResourceOwnerCredentialsMapping(ctx, id).Execute()

Delete a Resource Owner Credentials mapping.

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
    id := "id_example" // string | ID of the Resource Owner Credentials mapping.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthResourceOwnerCredentialsMappingsApi.DeleteResourceOwnerCredentialsMapping(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthResourceOwnerCredentialsMappingsApi.DeleteResourceOwnerCredentialsMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Resource Owner Credentials mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceOwnerCredentialsMappingRequest struct via the builder pattern


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


## GetResourceOwnerCredentialsMapping

> ResourceOwnerCredentialsMapping GetResourceOwnerCredentialsMapping(ctx, id).Execute()

Find the Resource Owner Credentials mapping by the ID.

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
    id := "id_example" // string | ID of the Resource Owner Credentials mapping.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthResourceOwnerCredentialsMappingsApi.GetResourceOwnerCredentialsMapping(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthResourceOwnerCredentialsMappingsApi.GetResourceOwnerCredentialsMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceOwnerCredentialsMapping`: ResourceOwnerCredentialsMapping
    fmt.Fprintf(os.Stdout, "Response from `OauthResourceOwnerCredentialsMappingsApi.GetResourceOwnerCredentialsMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Resource Owner Credentials mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceOwnerCredentialsMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourceOwnerCredentialsMapping**](ResourceOwnerCredentialsMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceOwnerCredentialsMappings

> ResourceOwnerCredentialsMappings GetResourceOwnerCredentialsMappings(ctx).Execute()

Get the list of Resource Owner Credentials Grant Mapping.

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
    resp, r, err := apiClient.OauthResourceOwnerCredentialsMappingsApi.GetResourceOwnerCredentialsMappings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthResourceOwnerCredentialsMappingsApi.GetResourceOwnerCredentialsMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceOwnerCredentialsMappings`: ResourceOwnerCredentialsMappings
    fmt.Fprintf(os.Stdout, "Response from `OauthResourceOwnerCredentialsMappingsApi.GetResourceOwnerCredentialsMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceOwnerCredentialsMappingsRequest struct via the builder pattern


### Return type

[**ResourceOwnerCredentialsMappings**](ResourceOwnerCredentialsMappings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceOwnerCredentialsMapping

> ResourceOwnerCredentialsMapping UpdateResourceOwnerCredentialsMapping(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update a Resource Owner Credentials mapping.

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
    id := "id_example" // string | ID of the Resource Owner Credentials mapping to update.
    body := *openapiclient.NewResourceOwnerCredentialsMapping("Id_example", map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}) // ResourceOwnerCredentialsMapping | Configuration for Resource Owner Credentials mapping.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthResourceOwnerCredentialsMappingsApi.UpdateResourceOwnerCredentialsMapping(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthResourceOwnerCredentialsMappingsApi.UpdateResourceOwnerCredentialsMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateResourceOwnerCredentialsMapping`: ResourceOwnerCredentialsMapping
    fmt.Fprintf(os.Stdout, "Response from `OauthResourceOwnerCredentialsMappingsApi.UpdateResourceOwnerCredentialsMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Resource Owner Credentials mapping to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceOwnerCredentialsMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ResourceOwnerCredentialsMapping**](ResourceOwnerCredentialsMapping.md) | Configuration for Resource Owner Credentials mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**ResourceOwnerCredentialsMapping**](ResourceOwnerCredentialsMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

