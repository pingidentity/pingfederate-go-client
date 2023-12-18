# \IdpAdaptersAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdpAdapter**](IdpAdaptersAPI.md#CreateIdpAdapter) | **Post** /idp/adapters | Create a new IdP adapter instance.
[**DeleteIdpAdapter**](IdpAdaptersAPI.md#DeleteIdpAdapter) | **Delete** /idp/adapters/{id} | Delete an IdP adapter instance.
[**GetIdpAdapter**](IdpAdaptersAPI.md#GetIdpAdapter) | **Get** /idp/adapters/{id} | Find an IdP adapter instance by ID.
[**GetIdpAdapterDescriptors**](IdpAdaptersAPI.md#GetIdpAdapterDescriptors) | **Get** /idp/adapters/descriptors | Get the list of available IdP adapter descriptors.
[**GetIdpAdapterDescriptorsById**](IdpAdaptersAPI.md#GetIdpAdapterDescriptorsById) | **Get** /idp/adapters/descriptors/{id} | Get the description of an IdP adapter plugin by ID.
[**GetIdpAdapters**](IdpAdaptersAPI.md#GetIdpAdapters) | **Get** /idp/adapters | Get the list of configured IdP adapter instances.
[**GetIdpAdaptersActionById**](IdpAdaptersAPI.md#GetIdpAdaptersActionById) | **Get** /idp/adapters/{id}/actions/{actionId} | Find an IdP adapter instance&#39;s action by ID.
[**GetIdpAdaptersActions**](IdpAdaptersAPI.md#GetIdpAdaptersActions) | **Get** /idp/adapters/{id}/actions | List the actions for an IdP adapter instance.
[**InvokeIdpAdaptersActionWithOptions**](IdpAdaptersAPI.md#InvokeIdpAdaptersActionWithOptions) | **Post** /idp/adapters/{id}/actions/{actionId}/invokeAction | Invokes an action for an IdP adapter instance.
[**UpdateIdpAdapter**](IdpAdaptersAPI.md#UpdateIdpAdapter) | **Put** /idp/adapters/{id} | Update an IdP adapter instance.



## CreateIdpAdapter

> IdpAdapter CreateIdpAdapter(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new IdP adapter instance.



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
    body := *openapiclient.NewIdpAdapter("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // IdpAdapter | Configuration for the IdP adapter instance.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpAdaptersAPI.CreateIdpAdapter(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpAdaptersAPI.CreateIdpAdapter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdpAdapter`: IdpAdapter
    fmt.Fprintf(os.Stdout, "Response from `IdpAdaptersAPI.CreateIdpAdapter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdpAdapterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IdpAdapter**](IdpAdapter.md) | Configuration for the IdP adapter instance. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**IdpAdapter**](IdpAdapter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdpAdapter

> DeleteIdpAdapter(ctx, id).Execute()

Delete an IdP adapter instance.



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
    id := "id_example" // string | ID of IdP adapter instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdpAdaptersAPI.DeleteIdpAdapter(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpAdaptersAPI.DeleteIdpAdapter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of IdP adapter instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdpAdapterRequest struct via the builder pattern


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


## GetIdpAdapter

> IdpAdapter GetIdpAdapter(ctx, id).Execute()

Find an IdP adapter instance by ID.



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
    id := "id_example" // string | ID of IdP adapter instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpAdaptersAPI.GetIdpAdapter(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpAdaptersAPI.GetIdpAdapter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdpAdapter`: IdpAdapter
    fmt.Fprintf(os.Stdout, "Response from `IdpAdaptersAPI.GetIdpAdapter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of IdP adapter instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpAdapterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdpAdapter**](IdpAdapter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdpAdapterDescriptors

> IdpAdapterDescriptors GetIdpAdapterDescriptors(ctx).Execute()

Get the list of available IdP adapter descriptors.

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
    resp, r, err := apiClient.IdpAdaptersAPI.GetIdpAdapterDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpAdaptersAPI.GetIdpAdapterDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdpAdapterDescriptors`: IdpAdapterDescriptors
    fmt.Fprintf(os.Stdout, "Response from `IdpAdaptersAPI.GetIdpAdapterDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpAdapterDescriptorsRequest struct via the builder pattern


### Return type

[**IdpAdapterDescriptors**](IdpAdapterDescriptors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdpAdapterDescriptorsById

> IdpAdapterDescriptor GetIdpAdapterDescriptorsById(ctx, id).Execute()

Get the description of an IdP adapter plugin by ID.



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
    id := "id_example" // string | ID of IdP adapter descriptor to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpAdaptersAPI.GetIdpAdapterDescriptorsById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpAdaptersAPI.GetIdpAdapterDescriptorsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdpAdapterDescriptorsById`: IdpAdapterDescriptor
    fmt.Fprintf(os.Stdout, "Response from `IdpAdaptersAPI.GetIdpAdapterDescriptorsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of IdP adapter descriptor to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpAdapterDescriptorsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdpAdapterDescriptor**](IdpAdapterDescriptor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdpAdapters

> IdpAdapters GetIdpAdapters(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()

Get the list of configured IdP adapter instances.

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
    numberPerPage := int64(56) // int64 | Number of adapters per page. (optional)
    filter := "filter_example" // string | Filter criteria limits the IdP adapters that are returned to only those that match it. The filter criteria is compared to the IdP adapter instance name and ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpAdaptersAPI.GetIdpAdapters(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpAdaptersAPI.GetIdpAdapters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdpAdapters`: IdpAdapters
    fmt.Fprintf(os.Stdout, "Response from `IdpAdaptersAPI.GetIdpAdapters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpAdaptersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve. | 
 **numberPerPage** | **int64** | Number of adapters per page. | 
 **filter** | **string** | Filter criteria limits the IdP adapters that are returned to only those that match it. The filter criteria is compared to the IdP adapter instance name and ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. | 

### Return type

[**IdpAdapters**](IdpAdapters.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdpAdaptersActionById

> Action GetIdpAdaptersActionById(ctx, id, actionId).Execute()

Find an IdP adapter instance's action by ID.



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
    id := "id_example" // string | ID of the IdP adapter instance to which these actions belongs to.
    actionId := "actionId_example" // string | ID of the action.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpAdaptersAPI.GetIdpAdaptersActionById(context.Background(), id, actionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpAdaptersAPI.GetIdpAdaptersActionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdpAdaptersActionById`: Action
    fmt.Fprintf(os.Stdout, "Response from `IdpAdaptersAPI.GetIdpAdaptersActionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the IdP adapter instance to which these actions belongs to. | 
**actionId** | **string** | ID of the action. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpAdaptersActionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Action**](Action.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdpAdaptersActions

> Actions GetIdpAdaptersActions(ctx, id).Execute()

List the actions for an IdP adapter instance.



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
    id := "id_example" // string | ID of the IdP adapter instance to which these actions belongs to.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpAdaptersAPI.GetIdpAdaptersActions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpAdaptersAPI.GetIdpAdaptersActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdpAdaptersActions`: Actions
    fmt.Fprintf(os.Stdout, "Response from `IdpAdaptersAPI.GetIdpAdaptersActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the IdP adapter instance to which these actions belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpAdaptersActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Actions**](Actions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvokeIdpAdaptersActionWithOptions

> ActionResult InvokeIdpAdaptersActionWithOptions(ctx, id, actionId).Body(body).Execute()

Invokes an action for an IdP adapter instance.



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
    id := "id_example" // string | ID of the IdP adapter instance to which these actions belongs to.
    actionId := "actionId_example" // string | ID of the action.
    body := *openapiclient.NewActionOptions([]openapiclient.ActionParameter{*openapiclient.NewActionParameter("Name_example")}) // ActionOptions | Action options for action invoked. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpAdaptersAPI.InvokeIdpAdaptersActionWithOptions(context.Background(), id, actionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpAdaptersAPI.InvokeIdpAdaptersActionWithOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvokeIdpAdaptersActionWithOptions`: ActionResult
    fmt.Fprintf(os.Stdout, "Response from `IdpAdaptersAPI.InvokeIdpAdaptersActionWithOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the IdP adapter instance to which these actions belongs to. | 
**actionId** | **string** | ID of the action. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvokeIdpAdaptersActionWithOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ActionOptions**](ActionOptions.md) | Action options for action invoked. | 

### Return type

[**ActionResult**](ActionResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdpAdapter

> IdpAdapter UpdateIdpAdapter(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update an IdP adapter instance.



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
    id := "id_example" // string | ID of IdP adapter instance.
    body := *openapiclient.NewIdpAdapter("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // IdpAdapter | Configuration for the IdP adapter instance.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpAdaptersAPI.UpdateIdpAdapter(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpAdaptersAPI.UpdateIdpAdapter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdpAdapter`: IdpAdapter
    fmt.Fprintf(os.Stdout, "Response from `IdpAdaptersAPI.UpdateIdpAdapter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of IdP adapter instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdpAdapterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IdpAdapter**](IdpAdapter.md) | Configuration for the IdP adapter instance. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**IdpAdapter**](IdpAdapter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

