# \AuthenticationSelectorsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthenticationSelector**](AuthenticationSelectorsAPI.md#CreateAuthenticationSelector) | **Post** /authenticationSelectors | Create a new authentication selector instance.
[**DeleteAuthenticationSelector**](AuthenticationSelectorsAPI.md#DeleteAuthenticationSelector) | **Delete** /authenticationSelectors/{id} | Delete an Authentication Selector instance.
[**GetAuthenticationSelector**](AuthenticationSelectorsAPI.md#GetAuthenticationSelector) | **Get** /authenticationSelectors/{id} | Get an Authentication Selector instance by ID.
[**GetAuthenticationSelectorDescriptors**](AuthenticationSelectorsAPI.md#GetAuthenticationSelectorDescriptors) | **Get** /authenticationSelectors/descriptors | Get the list of available Authentication Selector descriptors.
[**GetAuthenticationSelectorDescriptorsById**](AuthenticationSelectorsAPI.md#GetAuthenticationSelectorDescriptorsById) | **Get** /authenticationSelectors/descriptors/{id} | Get the description of an Authentication Selector plugin by ID.
[**GetAuthenticationSelectors**](AuthenticationSelectorsAPI.md#GetAuthenticationSelectors) | **Get** /authenticationSelectors | Get the list of configured Authentication Selector instances.
[**UpdateAuthenticationSelector**](AuthenticationSelectorsAPI.md#UpdateAuthenticationSelector) | **Put** /authenticationSelectors/{id} | Update an authentication selector instance.



## CreateAuthenticationSelector

> AuthenticationSelector CreateAuthenticationSelector(ctx).Body(body).Execute()

Create a new authentication selector instance.



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
    body := *openapiclient.NewAuthenticationSelector("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // AuthenticationSelector | Configuration for a new authentication selector instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationSelectorsAPI.CreateAuthenticationSelector(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationSelectorsAPI.CreateAuthenticationSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthenticationSelector`: AuthenticationSelector
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationSelectorsAPI.CreateAuthenticationSelector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthenticationSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthenticationSelector**](AuthenticationSelector.md) | Configuration for a new authentication selector instance. | 

### Return type

[**AuthenticationSelector**](AuthenticationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthenticationSelector

> DeleteAuthenticationSelector(ctx, id).Execute()

Delete an Authentication Selector instance.



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
    id := "id_example" // string | ID of Authentication Selector to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthenticationSelectorsAPI.DeleteAuthenticationSelector(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationSelectorsAPI.DeleteAuthenticationSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Authentication Selector to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthenticationSelectorRequest struct via the builder pattern


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


## GetAuthenticationSelector

> AuthenticationSelector GetAuthenticationSelector(ctx, id).Execute()

Get an Authentication Selector instance by ID.



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
    id := "id_example" // string | ID of Authentication Selector instance to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationSelectorsAPI.GetAuthenticationSelector(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationSelectorsAPI.GetAuthenticationSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticationSelector`: AuthenticationSelector
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationSelectorsAPI.GetAuthenticationSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Authentication Selector instance to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticationSelector**](AuthenticationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticationSelectorDescriptors

> AuthenticationSelectorDescriptors GetAuthenticationSelectorDescriptors(ctx).Execute()

Get the list of available Authentication Selector descriptors.

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
    resp, r, err := apiClient.AuthenticationSelectorsAPI.GetAuthenticationSelectorDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationSelectorsAPI.GetAuthenticationSelectorDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticationSelectorDescriptors`: AuthenticationSelectorDescriptors
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationSelectorsAPI.GetAuthenticationSelectorDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationSelectorDescriptorsRequest struct via the builder pattern


### Return type

[**AuthenticationSelectorDescriptors**](AuthenticationSelectorDescriptors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticationSelectorDescriptorsById

> AuthenticationSelectorDescriptor GetAuthenticationSelectorDescriptorsById(ctx, id).Execute()

Get the description of an Authentication Selector plugin by ID.



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
    id := "id_example" // string | ID of Authentication Selector descriptor to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationSelectorsAPI.GetAuthenticationSelectorDescriptorsById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationSelectorsAPI.GetAuthenticationSelectorDescriptorsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticationSelectorDescriptorsById`: AuthenticationSelectorDescriptor
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationSelectorsAPI.GetAuthenticationSelectorDescriptorsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Authentication Selector descriptor to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationSelectorDescriptorsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticationSelectorDescriptor**](AuthenticationSelectorDescriptor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticationSelectors

> AuthenticationSelectors GetAuthenticationSelectors(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()

Get the list of configured Authentication Selector instances.

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
    page := int64(56) // int64 | Page number to retrieve. (optional)
    numberPerPage := int64(56) // int64 | Number of selectors per page. (optional)
    filter := "filter_example" // string | Filter criteria limits the authentication selector instances that are returned to only those that match it. The filter criteria is compared to the authentication selector instance name and ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationSelectorsAPI.GetAuthenticationSelectors(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationSelectorsAPI.GetAuthenticationSelectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticationSelectors`: AuthenticationSelectors
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationSelectorsAPI.GetAuthenticationSelectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationSelectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve. | 
 **numberPerPage** | **int64** | Number of selectors per page. | 
 **filter** | **string** | Filter criteria limits the authentication selector instances that are returned to only those that match it. The filter criteria is compared to the authentication selector instance name and ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. | 

### Return type

[**AuthenticationSelectors**](AuthenticationSelectors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthenticationSelector

> AuthenticationSelector UpdateAuthenticationSelector(ctx, id).Body(body).Execute()

Update an authentication selector instance.



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
    id := "id_example" // string | ID of the authentication selector instance.
    body := *openapiclient.NewAuthenticationSelector("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // AuthenticationSelector | Configuration for updated authentication selector instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationSelectorsAPI.UpdateAuthenticationSelector(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationSelectorsAPI.UpdateAuthenticationSelector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthenticationSelector`: AuthenticationSelector
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationSelectorsAPI.UpdateAuthenticationSelector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the authentication selector instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthenticationSelectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthenticationSelector**](AuthenticationSelector.md) | Configuration for updated authentication selector instance. | 

### Return type

[**AuthenticationSelector**](AuthenticationSelector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

