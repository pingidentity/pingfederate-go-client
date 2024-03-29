# \SpAdaptersAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpAdapter**](SpAdaptersAPI.md#CreateSpAdapter) | **Post** /sp/adapters | Create a new SP adapter instance.
[**DeleteSpAdapter**](SpAdaptersAPI.md#DeleteSpAdapter) | **Delete** /sp/adapters/{id} | Delete an SP adapter instance.
[**GetSpAdapter**](SpAdaptersAPI.md#GetSpAdapter) | **Get** /sp/adapters/{id} | Find an SP adapter instance by ID.
[**GetSpAdapterDescriptors**](SpAdaptersAPI.md#GetSpAdapterDescriptors) | **Get** /sp/adapters/descriptors | Get the list of available SP adapter descriptors.
[**GetSpAdapterDescriptorsById**](SpAdaptersAPI.md#GetSpAdapterDescriptorsById) | **Get** /sp/adapters/descriptors/{id} | Get the description of an SP adapter plugin by ID.
[**GetSpAdapterUrlMappings**](SpAdaptersAPI.md#GetSpAdapterUrlMappings) | **Get** /sp/adapters/urlMappings | (Deprecated) List the mappings between URLs and adapter instances.
[**GetSpAdapters**](SpAdaptersAPI.md#GetSpAdapters) | **Get** /sp/adapters | Get the list of configured SP adapter instances.
[**GetSpAdaptersActionById**](SpAdaptersAPI.md#GetSpAdaptersActionById) | **Get** /sp/adapters/{id}/actions/{actionId} | Find an SP adapter instance&#39;s action by ID.
[**GetSpAdaptersActions**](SpAdaptersAPI.md#GetSpAdaptersActions) | **Get** /sp/adapters/{id}/actions | List the actions for an SP adapter instance.
[**InvokeSpAdapterActionWithOptions**](SpAdaptersAPI.md#InvokeSpAdapterActionWithOptions) | **Post** /sp/adapters/{id}/actions/{actionId}/invokeAction | Invokes an action for an SP adapter instance.
[**UpdateSpAdapter**](SpAdaptersAPI.md#UpdateSpAdapter) | **Put** /sp/adapters/{id} | Update an SP adapter instance.
[**UpdateSpAdapterUrlMappings**](SpAdaptersAPI.md#UpdateSpAdapterUrlMappings) | **Put** /sp/adapters/urlMappings | (Deprecated) Update the mappings between URLs and adapters instances.



## CreateSpAdapter

> SpAdapter CreateSpAdapter(ctx).Body(body).Execute()

Create a new SP adapter instance.



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
    body := *openapiclient.NewSpAdapter("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // SpAdapter | Configuration for the SP adapter instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpAdaptersAPI.CreateSpAdapter(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpAdaptersAPI.CreateSpAdapter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSpAdapter`: SpAdapter
    fmt.Fprintf(os.Stdout, "Response from `SpAdaptersAPI.CreateSpAdapter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpAdapterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SpAdapter**](SpAdapter.md) | Configuration for the SP adapter instance. | 

### Return type

[**SpAdapter**](SpAdapter.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSpAdapter

> DeleteSpAdapter(ctx, id).Execute()

Delete an SP adapter instance.



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
    id := "id_example" // string | ID of SP adapter instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpAdaptersAPI.DeleteSpAdapter(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpAdaptersAPI.DeleteSpAdapter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of SP adapter instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpAdapterRequest struct via the builder pattern


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


## GetSpAdapter

> SpAdapter GetSpAdapter(ctx, id).Execute()

Find an SP adapter instance by ID.



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
    id := "id_example" // string | ID of SP adapter instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpAdaptersAPI.GetSpAdapter(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpAdaptersAPI.GetSpAdapter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpAdapter`: SpAdapter
    fmt.Fprintf(os.Stdout, "Response from `SpAdaptersAPI.GetSpAdapter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of SP adapter instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpAdapterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpAdapter**](SpAdapter.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpAdapterDescriptors

> SpAdapterDescriptors GetSpAdapterDescriptors(ctx).Execute()

Get the list of available SP adapter descriptors.

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
    resp, r, err := apiClient.SpAdaptersAPI.GetSpAdapterDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpAdaptersAPI.GetSpAdapterDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpAdapterDescriptors`: SpAdapterDescriptors
    fmt.Fprintf(os.Stdout, "Response from `SpAdaptersAPI.GetSpAdapterDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpAdapterDescriptorsRequest struct via the builder pattern


### Return type

[**SpAdapterDescriptors**](SpAdapterDescriptors.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpAdapterDescriptorsById

> SpAdapterDescriptor GetSpAdapterDescriptorsById(ctx, id).Execute()

Get the description of an SP adapter plugin by ID.



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
    id := "id_example" // string | ID of SP adapter descriptor to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpAdaptersAPI.GetSpAdapterDescriptorsById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpAdaptersAPI.GetSpAdapterDescriptorsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpAdapterDescriptorsById`: SpAdapterDescriptor
    fmt.Fprintf(os.Stdout, "Response from `SpAdaptersAPI.GetSpAdapterDescriptorsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of SP adapter descriptor to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpAdapterDescriptorsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpAdapterDescriptor**](SpAdapterDescriptor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpAdapterUrlMappings

> SpAdapterUrlMappings GetSpAdapterUrlMappings(ctx).Execute()

(Deprecated) List the mappings between URLs and adapter instances.

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
    resp, r, err := apiClient.SpAdaptersAPI.GetSpAdapterUrlMappings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpAdaptersAPI.GetSpAdapterUrlMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpAdapterUrlMappings`: SpAdapterUrlMappings
    fmt.Fprintf(os.Stdout, "Response from `SpAdaptersAPI.GetSpAdapterUrlMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpAdapterUrlMappingsRequest struct via the builder pattern


### Return type

[**SpAdapterUrlMappings**](SpAdapterUrlMappings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpAdapters

> SpAdapters GetSpAdapters(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()

Get the list of configured SP adapter instances.

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
    filter := "filter_example" // string | Filter criteria limits the SP adapters that are returned to only those that match it. The filter criteria is compared to the SP adapter instance name and ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpAdaptersAPI.GetSpAdapters(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpAdaptersAPI.GetSpAdapters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpAdapters`: SpAdapters
    fmt.Fprintf(os.Stdout, "Response from `SpAdaptersAPI.GetSpAdapters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSpAdaptersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve. | 
 **numberPerPage** | **int64** | Number of adapters per page. | 
 **filter** | **string** | Filter criteria limits the SP adapters that are returned to only those that match it. The filter criteria is compared to the SP adapter instance name and ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. | 

### Return type

[**SpAdapters**](SpAdapters.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpAdaptersActionById

> Action GetSpAdaptersActionById(ctx, id, actionId).Execute()

Find an SP adapter instance's action by ID.



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
    id := "id_example" // string | ID of the SP adapter instance to which this action belongs to.
    actionId := "actionId_example" // string | ID of the action.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpAdaptersAPI.GetSpAdaptersActionById(context.Background(), id, actionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpAdaptersAPI.GetSpAdaptersActionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpAdaptersActionById`: Action
    fmt.Fprintf(os.Stdout, "Response from `SpAdaptersAPI.GetSpAdaptersActionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the SP adapter instance to which this action belongs to. | 
**actionId** | **string** | ID of the action. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpAdaptersActionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Action**](Action.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpAdaptersActions

> Actions GetSpAdaptersActions(ctx, id).Execute()

List the actions for an SP adapter instance.



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
    id := "id_example" // string | ID of the SP adapter instance to which this action belongs to.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpAdaptersAPI.GetSpAdaptersActions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpAdaptersAPI.GetSpAdaptersActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpAdaptersActions`: Actions
    fmt.Fprintf(os.Stdout, "Response from `SpAdaptersAPI.GetSpAdaptersActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the SP adapter instance to which this action belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpAdaptersActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Actions**](Actions.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvokeSpAdapterActionWithOptions

> ActionResult InvokeSpAdapterActionWithOptions(ctx, id, actionId).Body(body).Execute()

Invokes an action for an SP adapter instance.



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
    id := "id_example" // string | ID of the SP adapter instance to which this action belongs to.
    actionId := "actionId_example" // string | ID of the action.
    body := *openapiclient.NewActionOptions([]openapiclient.ActionParameter{*openapiclient.NewActionParameter("Name_example")}) // ActionOptions | Action options for action invoked. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpAdaptersAPI.InvokeSpAdapterActionWithOptions(context.Background(), id, actionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpAdaptersAPI.InvokeSpAdapterActionWithOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvokeSpAdapterActionWithOptions`: ActionResult
    fmt.Fprintf(os.Stdout, "Response from `SpAdaptersAPI.InvokeSpAdapterActionWithOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the SP adapter instance to which this action belongs to. | 
**actionId** | **string** | ID of the action. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvokeSpAdapterActionWithOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ActionOptions**](ActionOptions.md) | Action options for action invoked. | 

### Return type

[**ActionResult**](ActionResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpAdapter

> SpAdapter UpdateSpAdapter(ctx, id).Body(body).Execute()

Update an SP adapter instance.



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
    id := "id_example" // string | ID of SP adapter instance.
    body := *openapiclient.NewSpAdapter("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // SpAdapter | Configuration for the SP adapter instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpAdaptersAPI.UpdateSpAdapter(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpAdaptersAPI.UpdateSpAdapter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSpAdapter`: SpAdapter
    fmt.Fprintf(os.Stdout, "Response from `SpAdaptersAPI.UpdateSpAdapter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of SP adapter instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpAdapterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SpAdapter**](SpAdapter.md) | Configuration for the SP adapter instance. | 

### Return type

[**SpAdapter**](SpAdapter.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpAdapterUrlMappings

> SpAdapterUrlMappings UpdateSpAdapterUrlMappings(ctx).Body(body).Execute()

(Deprecated) Update the mappings between URLs and adapters instances.

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
    body := *openapiclient.NewSpAdapterUrlMappings() // SpAdapterUrlMappings | The SP adapter URL mappings to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpAdaptersAPI.UpdateSpAdapterUrlMappings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpAdaptersAPI.UpdateSpAdapterUrlMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSpAdapterUrlMappings`: SpAdapterUrlMappings
    fmt.Fprintf(os.Stdout, "Response from `SpAdaptersAPI.UpdateSpAdapterUrlMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpAdapterUrlMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SpAdapterUrlMappings**](SpAdapterUrlMappings.md) | The SP adapter URL mappings to update. | 

### Return type

[**SpAdapterUrlMappings**](SpAdapterUrlMappings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

