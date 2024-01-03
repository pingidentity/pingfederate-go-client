# \OauthIdpAdapterMappingsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdpAdapterMapping**](OauthIdpAdapterMappingsAPI.md#CreateIdpAdapterMapping) | **Post** /oauth/idpAdapterMappings | Create a new IdP adapter mapping.
[**DeleteIdpAdapterMapping**](OauthIdpAdapterMappingsAPI.md#DeleteIdpAdapterMapping) | **Delete** /oauth/idpAdapterMappings/{id} | Delete an IdP adapter mapping.
[**GetIdpAdapterMapping**](OauthIdpAdapterMappingsAPI.md#GetIdpAdapterMapping) | **Get** /oauth/idpAdapterMappings/{id} | Find the IdP adapter mapping by the ID.
[**GetIdpAdapterMappings**](OauthIdpAdapterMappingsAPI.md#GetIdpAdapterMappings) | **Get** /oauth/idpAdapterMappings | Get the list of IdP adapter mappings.
[**UpdateIdpAdapterMapping**](OauthIdpAdapterMappingsAPI.md#UpdateIdpAdapterMapping) | **Put** /oauth/idpAdapterMappings/{id} | Update an IdP adapter mapping.



## CreateIdpAdapterMapping

> IdpAdapterMapping CreateIdpAdapterMapping(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new IdP adapter mapping.



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
    body := *openapiclient.NewIdpAdapterMapping("Id_example", map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}) // IdpAdapterMapping | Configuration for IdP adapter mapping.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthIdpAdapterMappingsAPI.CreateIdpAdapterMapping(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthIdpAdapterMappingsAPI.CreateIdpAdapterMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdpAdapterMapping`: IdpAdapterMapping
    fmt.Fprintf(os.Stdout, "Response from `OauthIdpAdapterMappingsAPI.CreateIdpAdapterMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdpAdapterMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IdpAdapterMapping**](IdpAdapterMapping.md) | Configuration for IdP adapter mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**IdpAdapterMapping**](IdpAdapterMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdpAdapterMapping

> DeleteIdpAdapterMapping(ctx, id).Execute()

Delete an IdP adapter mapping.

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
    id := "id_example" // string | ID of the IdP adapter mapping.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthIdpAdapterMappingsAPI.DeleteIdpAdapterMapping(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthIdpAdapterMappingsAPI.DeleteIdpAdapterMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the IdP adapter mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdpAdapterMappingRequest struct via the builder pattern


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


## GetIdpAdapterMapping

> IdpAdapterMapping GetIdpAdapterMapping(ctx, id).Execute()

Find the IdP adapter mapping by the ID.

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
    id := "id_example" // string | ID of the adapter mapping.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthIdpAdapterMappingsAPI.GetIdpAdapterMapping(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthIdpAdapterMappingsAPI.GetIdpAdapterMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdpAdapterMapping`: IdpAdapterMapping
    fmt.Fprintf(os.Stdout, "Response from `OauthIdpAdapterMappingsAPI.GetIdpAdapterMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the adapter mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpAdapterMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdpAdapterMapping**](IdpAdapterMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdpAdapterMappings

> IdpAdapterMappings GetIdpAdapterMappings(ctx).Execute()

Get the list of IdP adapter mappings.

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
    resp, r, err := apiClient.OauthIdpAdapterMappingsAPI.GetIdpAdapterMappings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthIdpAdapterMappingsAPI.GetIdpAdapterMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdpAdapterMappings`: IdpAdapterMappings
    fmt.Fprintf(os.Stdout, "Response from `OauthIdpAdapterMappingsAPI.GetIdpAdapterMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpAdapterMappingsRequest struct via the builder pattern


### Return type

[**IdpAdapterMappings**](IdpAdapterMappings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdpAdapterMapping

> IdpAdapterMapping UpdateIdpAdapterMapping(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update an IdP adapter mapping.

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
    id := "id_example" // string | ID of the IdP adapter mapping to update.
    body := *openapiclient.NewIdpAdapterMapping("Id_example", map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}) // IdpAdapterMapping | Configuration for IdP adapter mapping.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthIdpAdapterMappingsAPI.UpdateIdpAdapterMapping(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthIdpAdapterMappingsAPI.UpdateIdpAdapterMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdpAdapterMapping`: IdpAdapterMapping
    fmt.Fprintf(os.Stdout, "Response from `OauthIdpAdapterMappingsAPI.UpdateIdpAdapterMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the IdP adapter mapping to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdpAdapterMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IdpAdapterMapping**](IdpAdapterMapping.md) | Configuration for IdP adapter mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**IdpAdapterMapping**](IdpAdapterMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

