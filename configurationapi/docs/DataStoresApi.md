# \DataStoresAPI

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataStore**](DataStoresAPI.md#CreateDataStore) | **Post** /dataStores | Create a new data store.
[**DeleteDataStore**](DataStoresAPI.md#DeleteDataStore) | **Delete** /dataStores/{id} | Delete a data store.
[**GetCustomDataStoreDescriptor**](DataStoresAPI.md#GetCustomDataStoreDescriptor) | **Get** /dataStores/descriptors/{id} | Get the description of a custom data store plugin by ID.
[**GetCustomDataStoreDescriptors**](DataStoresAPI.md#GetCustomDataStoreDescriptors) | **Get** /dataStores/descriptors | Get the list of available custom data store descriptors.
[**GetDataStore**](DataStoresAPI.md#GetDataStore) | **Get** /dataStores/{id} | Find data store by ID.
[**GetDataStores**](DataStoresAPI.md#GetDataStores) | **Get** /dataStores | Get list of all data stores.
[**GetDataStoresActionById**](DataStoresAPI.md#GetDataStoresActionById) | **Get** /dataStores/{id}/actions/{actionId} | Find a data store instance&#39;s action by ID.
[**GetDataStoresActions**](DataStoresAPI.md#GetDataStoresActions) | **Get** /dataStores/{id}/actions | List the actions for a data store instance.
[**InvokeActionWithOptions**](DataStoresAPI.md#InvokeActionWithOptions) | **Post** /dataStores/{id}/actions/{actionId}/invokeAction | Invokes an action for a data source instance.
[**UpdateDataStore**](DataStoresAPI.md#UpdateDataStore) | **Put** /dataStores/{id} | Update a data store.



## CreateDataStore

> DataStore CreateDataStore(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new data store.



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
    body := *openapiclient.NewDataStore("Type_example") // DataStore | Configuration for new data store.
    xBypassExternalValidation := true // bool | Connection test will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataStoresAPI.CreateDataStore(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStoresAPI.CreateDataStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDataStore`: DataStore
    fmt.Fprintf(os.Stdout, "Response from `DataStoresAPI.CreateDataStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DataStore**](DataStore.md) | Configuration for new data store. | 
 **xBypassExternalValidation** | **bool** | Connection test will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**DataStore**](DataStore.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDataStore

> DeleteDataStore(ctx, id).Execute()

Delete a data store.



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
    id := "id_example" // string | ID of data store instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DataStoresAPI.DeleteDataStore(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStoresAPI.DeleteDataStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of data store instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataStoreRequest struct via the builder pattern


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


## GetCustomDataStoreDescriptor

> CustomDataStoreDescriptor GetCustomDataStoreDescriptor(ctx, id).Execute()

Get the description of a custom data store plugin by ID.



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
    id := "id_example" // string | ID of custom data store descriptor to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataStoresAPI.GetCustomDataStoreDescriptor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStoresAPI.GetCustomDataStoreDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomDataStoreDescriptor`: CustomDataStoreDescriptor
    fmt.Fprintf(os.Stdout, "Response from `DataStoresAPI.GetCustomDataStoreDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of custom data store descriptor to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomDataStoreDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomDataStoreDescriptor**](CustomDataStoreDescriptor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomDataStoreDescriptors

> CustomDataStoreDescriptors GetCustomDataStoreDescriptors(ctx).Execute()

Get the list of available custom data store descriptors.

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
    resp, r, err := apiClient.DataStoresAPI.GetCustomDataStoreDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStoresAPI.GetCustomDataStoreDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomDataStoreDescriptors`: CustomDataStoreDescriptors
    fmt.Fprintf(os.Stdout, "Response from `DataStoresAPI.GetCustomDataStoreDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomDataStoreDescriptorsRequest struct via the builder pattern


### Return type

[**CustomDataStoreDescriptors**](CustomDataStoreDescriptors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataStore

> DataStore GetDataStore(ctx, id).Execute()

Find data store by ID.



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
    id := "id_example" // string | ID of data store instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataStoresAPI.GetDataStore(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStoresAPI.GetDataStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataStore`: DataStore
    fmt.Fprintf(os.Stdout, "Response from `DataStoresAPI.GetDataStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of data store instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataStore**](DataStore.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataStores

> DataStores GetDataStores(ctx).Execute()

Get list of all data stores.

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
    resp, r, err := apiClient.DataStoresAPI.GetDataStores(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStoresAPI.GetDataStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataStores`: DataStores
    fmt.Fprintf(os.Stdout, "Response from `DataStoresAPI.GetDataStores`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataStoresRequest struct via the builder pattern


### Return type

[**DataStores**](DataStores.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataStoresActionById

> Action GetDataStoresActionById(ctx, id, actionId).Execute()

Find a data store instance's action by ID.



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
    id := "id_example" // string | ID of data store to which these actions belong to.
    actionId := "actionId_example" // string | ID of the action.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataStoresAPI.GetDataStoresActionById(context.Background(), id, actionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStoresAPI.GetDataStoresActionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataStoresActionById`: Action
    fmt.Fprintf(os.Stdout, "Response from `DataStoresAPI.GetDataStoresActionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of data store to which these actions belong to. | 
**actionId** | **string** | ID of the action. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataStoresActionByIdRequest struct via the builder pattern


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


## GetDataStoresActions

> Actions GetDataStoresActions(ctx, id).Execute()

List the actions for a data store instance.



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
    id := "id_example" // string | ID of data store to which these actions belong to.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataStoresAPI.GetDataStoresActions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStoresAPI.GetDataStoresActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataStoresActions`: Actions
    fmt.Fprintf(os.Stdout, "Response from `DataStoresAPI.GetDataStoresActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of data store to which these actions belong to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataStoresActionsRequest struct via the builder pattern


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


## InvokeActionWithOptions

> ActionResult InvokeActionWithOptions(ctx, id, actionId).Body(body).Execute()

Invokes an action for a data source instance.



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
    id := "id_example" // string | ID of data store to which these actions belong to.
    actionId := "actionId_example" // string | ID of the action.
    body := *openapiclient.NewActionOptions([]openapiclient.ActionParameter{*openapiclient.NewActionParameter("Name_example")}) // ActionOptions | Action options for action invoked. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataStoresAPI.InvokeActionWithOptions(context.Background(), id, actionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStoresAPI.InvokeActionWithOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvokeActionWithOptions`: ActionResult
    fmt.Fprintf(os.Stdout, "Response from `DataStoresAPI.InvokeActionWithOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of data store to which these actions belong to. | 
**actionId** | **string** | ID of the action. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvokeActionWithOptionsRequest struct via the builder pattern


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


## UpdateDataStore

> DataStore UpdateDataStore(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update a data store.



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
    id := "id_example" // string | ID of data store instance.
    body := *openapiclient.NewDataStore("Type_example") // DataStore | Configuration for the data store.
    xBypassExternalValidation := true // bool | Connection test will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataStoresAPI.UpdateDataStore(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStoresAPI.UpdateDataStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDataStore`: DataStore
    fmt.Fprintf(os.Stdout, "Response from `DataStoresAPI.UpdateDataStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of data store instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DataStore**](DataStore.md) | Configuration for the data store. | 
 **xBypassExternalValidation** | **bool** | Connection test will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**DataStore**](DataStore.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

