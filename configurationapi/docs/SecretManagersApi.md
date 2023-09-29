# \SecretManagersAPI

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecretManager**](SecretManagersAPI.md#CreateSecretManager) | **Post** /secretManagers | Create a secret manager plugin instance.
[**DeleteSecretManager**](SecretManagersAPI.md#DeleteSecretManager) | **Delete** /secretManagers/{id} | Delete a secret manager plugin instance.
[**GetSecretManager**](SecretManagersAPI.md#GetSecretManager) | **Get** /secretManagers/{id} | Get a specific secret manager plugin instance.
[**GetSecretManagerPluginDescriptor**](SecretManagersAPI.md#GetSecretManagerPluginDescriptor) | **Get** /secretManagers/descriptors/{id} | Get a secret manager plugin descriptor.
[**GetSecretManagerPluginDescriptors**](SecretManagersAPI.md#GetSecretManagerPluginDescriptors) | **Get** /secretManagers/descriptors | Get a list of available secret manager plugin descriptors.
[**GetSecretManagers**](SecretManagersAPI.md#GetSecretManagers) | **Get** /secretManagers | Get a list of secret manager plugin instances.
[**GetSecretManagersAction**](SecretManagersAPI.md#GetSecretManagersAction) | **Get** /secretManagers/{id}/actions/{actionId} | Get a secret manager plugin instance&#39;s action by ID.
[**GetSecretManagersActions**](SecretManagersAPI.md#GetSecretManagersActions) | **Get** /secretManagers/{id}/actions | Get a list of actions for a secret manager plugin instance.
[**InvokeSecretManagersActionWithOptions**](SecretManagersAPI.md#InvokeSecretManagersActionWithOptions) | **Post** /secretManagers/{id}/actions/{actionId}/invokeAction | Invokes an action for secret manager plugin instance.
[**UpdateSecretManager**](SecretManagersAPI.md#UpdateSecretManager) | **Put** /secretManagers/{id} | Update a secret manager plugin instance.



## CreateSecretManager

> SecretManager CreateSecretManager(ctx).Body(body).Execute()

Create a secret manager plugin instance.

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
    body := *openapiclient.NewSecretManager("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // SecretManager | Configuration for a secret manager plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretManagersAPI.CreateSecretManager(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretManagersAPI.CreateSecretManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecretManager`: SecretManager
    fmt.Fprintf(os.Stdout, "Response from `SecretManagersAPI.CreateSecretManager`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecretManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SecretManager**](SecretManager.md) | Configuration for a secret manager plugin instance. | 

### Return type

[**SecretManager**](SecretManager.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecretManager

> DeleteSecretManager(ctx, id).Execute()

Delete a secret manager plugin instance.

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
    id := "id_example" // string | ID of a secret manager plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SecretManagersAPI.DeleteSecretManager(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretManagersAPI.DeleteSecretManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a secret manager plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecretManagerRequest struct via the builder pattern


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


## GetSecretManager

> SecretManager GetSecretManager(ctx, id).Execute()

Get a specific secret manager plugin instance.

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
    id := "id_example" // string | ID of a secret manager plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretManagersAPI.GetSecretManager(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretManagersAPI.GetSecretManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretManager`: SecretManager
    fmt.Fprintf(os.Stdout, "Response from `SecretManagersAPI.GetSecretManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a secret manager plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecretManager**](SecretManager.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretManagerPluginDescriptor

> SecretManagerDescriptor GetSecretManagerPluginDescriptor(ctx, id).Execute()

Get a secret manager plugin descriptor.

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
    id := "id_example" // string | ID of secret manager plugin descriptor.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretManagersAPI.GetSecretManagerPluginDescriptor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretManagersAPI.GetSecretManagerPluginDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretManagerPluginDescriptor`: SecretManagerDescriptor
    fmt.Fprintf(os.Stdout, "Response from `SecretManagersAPI.GetSecretManagerPluginDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of secret manager plugin descriptor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretManagerPluginDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecretManagerDescriptor**](SecretManagerDescriptor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretManagerPluginDescriptors

> SecretManagerDescriptors GetSecretManagerPluginDescriptors(ctx).Execute()

Get a list of available secret manager plugin descriptors.

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
    resp, r, err := apiClient.SecretManagersAPI.GetSecretManagerPluginDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretManagersAPI.GetSecretManagerPluginDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretManagerPluginDescriptors`: SecretManagerDescriptors
    fmt.Fprintf(os.Stdout, "Response from `SecretManagersAPI.GetSecretManagerPluginDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretManagerPluginDescriptorsRequest struct via the builder pattern


### Return type

[**SecretManagerDescriptors**](SecretManagerDescriptors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretManagers

> SecretManagers GetSecretManagers(ctx).Execute()

Get a list of secret manager plugin instances.

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
    resp, r, err := apiClient.SecretManagersAPI.GetSecretManagers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretManagersAPI.GetSecretManagers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretManagers`: SecretManagers
    fmt.Fprintf(os.Stdout, "Response from `SecretManagersAPI.GetSecretManagers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretManagersRequest struct via the builder pattern


### Return type

[**SecretManagers**](SecretManagers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretManagersAction

> Action GetSecretManagersAction(ctx, id, actionId).Execute()

Get a secret manager plugin instance's action by ID.

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
    id := "id_example" // string | ID of a secret manager plugin instance.
    actionId := "actionId_example" // string | ID of the action.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretManagersAPI.GetSecretManagersAction(context.Background(), id, actionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretManagersAPI.GetSecretManagersAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretManagersAction`: Action
    fmt.Fprintf(os.Stdout, "Response from `SecretManagersAPI.GetSecretManagersAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a secret manager plugin instance. | 
**actionId** | **string** | ID of the action. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretManagersActionRequest struct via the builder pattern


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


## GetSecretManagersActions

> Actions GetSecretManagersActions(ctx, id).Execute()

Get a list of actions for a secret manager plugin instance.

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
    id := "id_example" // string | ID of a secret manager plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretManagersAPI.GetSecretManagersActions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretManagersAPI.GetSecretManagersActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretManagersActions`: Actions
    fmt.Fprintf(os.Stdout, "Response from `SecretManagersAPI.GetSecretManagersActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a secret manager plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretManagersActionsRequest struct via the builder pattern


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


## InvokeSecretManagersActionWithOptions

> ActionResult InvokeSecretManagersActionWithOptions(ctx, id, actionId).Body(body).Execute()

Invokes an action for secret manager plugin instance.



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
    id := "id_example" // string | ID of a secret manager plugin instance.
    actionId := "actionId_example" // string | ID of the action.
    body := *openapiclient.NewActionOptions([]openapiclient.ActionParameter{*openapiclient.NewActionParameter("Name_example")}) // ActionOptions | Action options for action invoked. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretManagersAPI.InvokeSecretManagersActionWithOptions(context.Background(), id, actionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretManagersAPI.InvokeSecretManagersActionWithOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvokeSecretManagersActionWithOptions`: ActionResult
    fmt.Fprintf(os.Stdout, "Response from `SecretManagersAPI.InvokeSecretManagersActionWithOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a secret manager plugin instance. | 
**actionId** | **string** | ID of the action. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvokeSecretManagersActionWithOptionsRequest struct via the builder pattern


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


## UpdateSecretManager

> SecretManager UpdateSecretManager(ctx, id).Body(body).Execute()

Update a secret manager plugin instance.

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
    id := "id_example" // string | ID of a secret manager plugin instance.
    body := *openapiclient.NewSecretManager("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // SecretManager | Configuration for a secret manager plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretManagersAPI.UpdateSecretManager(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretManagersAPI.UpdateSecretManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecretManager`: SecretManager
    fmt.Fprintf(os.Stdout, "Response from `SecretManagersAPI.UpdateSecretManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a secret manager plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecretManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SecretManager**](SecretManager.md) | Configuration for a secret manager plugin instance. | 

### Return type

[**SecretManager**](SecretManager.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

