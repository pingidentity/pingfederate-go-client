# \OauthAccessTokenMappingsApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMapping**](OauthAccessTokenMappingsApi.md#CreateMapping) | **Post** /oauth/accessTokenMappings | Create a new Access Token Mapping.
[**DeleteMapping**](OauthAccessTokenMappingsApi.md#DeleteMapping) | **Delete** /oauth/accessTokenMappings/{id} | Delete an Access Token Mapping.
[**GetMapping**](OauthAccessTokenMappingsApi.md#GetMapping) | **Get** /oauth/accessTokenMappings/{id} | Find the Access Token Mapping by its ID.
[**GetMappings**](OauthAccessTokenMappingsApi.md#GetMappings) | **Get** /oauth/accessTokenMappings | Get the list of Access Token Mappings.
[**UpdateMapping**](OauthAccessTokenMappingsApi.md#UpdateMapping) | **Put** /oauth/accessTokenMappings/{id} | Update an Access Token Mapping.



## CreateMapping

> AccessTokenMapping CreateMapping(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new Access Token Mapping.



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
    body := *openapiclient.NewAccessTokenMapping("Id_example", *openapiclient.NewAccessTokenMappingContext("Type_example", *openapiclient.NewResourceLink("Id_example")), *openapiclient.NewResourceLink("Id_example"), map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}) // AccessTokenMapping | Configuration for the new Access Token Mapping.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAccessTokenMappingsApi.CreateMapping(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAccessTokenMappingsApi.CreateMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMapping`: AccessTokenMapping
    fmt.Fprintf(os.Stdout, "Response from `OauthAccessTokenMappingsApi.CreateMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccessTokenMapping**](AccessTokenMapping.md) | Configuration for the new Access Token Mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**AccessTokenMapping**](AccessTokenMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMapping

> DeleteMapping(ctx, id).Execute()

Delete an Access Token Mapping.

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
    id := "id_example" // string | ID of the Access Token Mapping.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthAccessTokenMappingsApi.DeleteMapping(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAccessTokenMappingsApi.DeleteMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Access Token Mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMappingRequest struct via the builder pattern


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


## GetMapping

> AccessTokenMapping GetMapping(ctx, id).Execute()

Find the Access Token Mapping by its ID.

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
    id := "id_example" // string | ID of the Access Token Mapping.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAccessTokenMappingsApi.GetMapping(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAccessTokenMappingsApi.GetMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMapping`: AccessTokenMapping
    fmt.Fprintf(os.Stdout, "Response from `OauthAccessTokenMappingsApi.GetMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Access Token Mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessTokenMapping**](AccessTokenMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMappings

> AccessTokenMappings GetMappings(ctx).Execute()

Get the list of Access Token Mappings.

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
    resp, r, err := apiClient.OauthAccessTokenMappingsApi.GetMappings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAccessTokenMappingsApi.GetMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMappings`: AccessTokenMappings
    fmt.Fprintf(os.Stdout, "Response from `OauthAccessTokenMappingsApi.GetMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMappingsRequest struct via the builder pattern


### Return type

[**AccessTokenMappings**](AccessTokenMappings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMapping

> AccessTokenMapping UpdateMapping(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update an Access Token Mapping.



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
    id := "id_example" // string | ID of the Access Token Mapping to update.
    body := *openapiclient.NewAccessTokenMapping("Id_example", *openapiclient.NewAccessTokenMappingContext("Type_example", *openapiclient.NewResourceLink("Id_example")), *openapiclient.NewResourceLink("Id_example"), map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}) // AccessTokenMapping | Configuration for updated mapping.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAccessTokenMappingsApi.UpdateMapping(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAccessTokenMappingsApi.UpdateMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMapping`: AccessTokenMapping
    fmt.Fprintf(os.Stdout, "Response from `OauthAccessTokenMappingsApi.UpdateMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Access Token Mapping to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AccessTokenMapping**](AccessTokenMapping.md) | Configuration for updated mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**AccessTokenMapping**](AccessTokenMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
