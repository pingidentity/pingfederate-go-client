# \OauthOutOfBandAuthPluginsAPI

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOOBAuthenticator**](OauthOutOfBandAuthPluginsAPI.md#CreateOOBAuthenticator) | **Post** /oauth/outOfBandAuthPlugins | Create an Out of Band authenticator plugin instance.
[**DeleteOOBAuthenticator**](OauthOutOfBandAuthPluginsAPI.md#DeleteOOBAuthenticator) | **Delete** /oauth/outOfBandAuthPlugins/{id} | Delete an Out of Band authenticator plugin instance.
[**GetOOBAction**](OauthOutOfBandAuthPluginsAPI.md#GetOOBAction) | **Get** /oauth/outOfBandAuthPlugins/{id}/actions/{actionId} | Find an Out of Band authenticator plugin instance&#39;s action by ID.
[**GetOOBActions**](OauthOutOfBandAuthPluginsAPI.md#GetOOBActions) | **Get** /oauth/outOfBandAuthPlugins/{id}/actions | List of actions for an Out of Band authenticator plugin instance.
[**GetOOBAuthPluginDescriptor**](OauthOutOfBandAuthPluginsAPI.md#GetOOBAuthPluginDescriptor) | **Get** /oauth/outOfBandAuthPlugins/descriptors/{id} | Get the descriptor of an Out of Band authenticator plugin.
[**GetOOBAuthPluginDescriptors**](OauthOutOfBandAuthPluginsAPI.md#GetOOBAuthPluginDescriptors) | **Get** /oauth/outOfBandAuthPlugins/descriptors | Get the list of available Out of Band authenticator plugin descriptors.
[**GetOOBAuthenticator**](OauthOutOfBandAuthPluginsAPI.md#GetOOBAuthenticator) | **Get** /oauth/outOfBandAuthPlugins/{id} | Get a specific Out of Band authenticator plugin instance.
[**GetOOBAuthenticators**](OauthOutOfBandAuthPluginsAPI.md#GetOOBAuthenticators) | **Get** /oauth/outOfBandAuthPlugins | Get a list of Out of Band authenticator plugin instances.
[**InvokeOOBActionWithOptions**](OauthOutOfBandAuthPluginsAPI.md#InvokeOOBActionWithOptions) | **Post** /oauth/outOfBandAuthPlugins/{id}/actions/{actionId}/invokeAction | Invokes an action for Out of Band authenticator plugin instance.
[**UpdateOOBAuthenticator**](OauthOutOfBandAuthPluginsAPI.md#UpdateOOBAuthenticator) | **Put** /oauth/outOfBandAuthPlugins/{id} | Update an Out of Band authenticator plugin instance.



## CreateOOBAuthenticator

> OutOfBandAuthenticator CreateOOBAuthenticator(ctx).Body(body).Execute()

Create an Out of Band authenticator plugin instance.

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
    body := *openapiclient.NewOutOfBandAuthenticator("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // OutOfBandAuthenticator | Configuration for an Out of Band authenticator plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthOutOfBandAuthPluginsAPI.CreateOOBAuthenticator(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthOutOfBandAuthPluginsAPI.CreateOOBAuthenticator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOOBAuthenticator`: OutOfBandAuthenticator
    fmt.Fprintf(os.Stdout, "Response from `OauthOutOfBandAuthPluginsAPI.CreateOOBAuthenticator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOOBAuthenticatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OutOfBandAuthenticator**](OutOfBandAuthenticator.md) | Configuration for an Out of Band authenticator plugin instance. | 

### Return type

[**OutOfBandAuthenticator**](OutOfBandAuthenticator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOOBAuthenticator

> DeleteOOBAuthenticator(ctx, id).Execute()

Delete an Out of Band authenticator plugin instance.

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
    id := "id_example" // string | ID of Out of Band authenticator plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthOutOfBandAuthPluginsAPI.DeleteOOBAuthenticator(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthOutOfBandAuthPluginsAPI.DeleteOOBAuthenticator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Out of Band authenticator plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOOBAuthenticatorRequest struct via the builder pattern


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


## GetOOBAction

> Action GetOOBAction(ctx, id, actionId).Execute()

Find an Out of Band authenticator plugin instance's action by ID.

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
    id := "id_example" // string | ID of the Out of Band authenticator plugin instance to which these actions belongs to.
    actionId := "actionId_example" // string | ID of the action.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthOutOfBandAuthPluginsAPI.GetOOBAction(context.Background(), id, actionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthOutOfBandAuthPluginsAPI.GetOOBAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOOBAction`: Action
    fmt.Fprintf(os.Stdout, "Response from `OauthOutOfBandAuthPluginsAPI.GetOOBAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Out of Band authenticator plugin instance to which these actions belongs to. | 
**actionId** | **string** | ID of the action. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOOBActionRequest struct via the builder pattern


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


## GetOOBActions

> Actions GetOOBActions(ctx, id).Execute()

List of actions for an Out of Band authenticator plugin instance.

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
    id := "id_example" // string | ID of the Out of Band authenticator plugin instance to which these actions belongs to.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthOutOfBandAuthPluginsAPI.GetOOBActions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthOutOfBandAuthPluginsAPI.GetOOBActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOOBActions`: Actions
    fmt.Fprintf(os.Stdout, "Response from `OauthOutOfBandAuthPluginsAPI.GetOOBActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Out of Band authenticator plugin instance to which these actions belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOOBActionsRequest struct via the builder pattern


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


## GetOOBAuthPluginDescriptor

> OutOfBandAuthPluginDescriptor GetOOBAuthPluginDescriptor(ctx, id).Execute()

Get the descriptor of an Out of Band authenticator plugin.

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
    id := "id_example" // string | ID of an Out of Band authenticator plugin descriptor.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthOutOfBandAuthPluginsAPI.GetOOBAuthPluginDescriptor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthOutOfBandAuthPluginsAPI.GetOOBAuthPluginDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOOBAuthPluginDescriptor`: OutOfBandAuthPluginDescriptor
    fmt.Fprintf(os.Stdout, "Response from `OauthOutOfBandAuthPluginsAPI.GetOOBAuthPluginDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an Out of Band authenticator plugin descriptor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOOBAuthPluginDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutOfBandAuthPluginDescriptor**](OutOfBandAuthPluginDescriptor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOOBAuthPluginDescriptors

> OutOfBandAuthPluginDescriptors GetOOBAuthPluginDescriptors(ctx).Execute()

Get the list of available Out of Band authenticator plugin descriptors.

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
    resp, r, err := apiClient.OauthOutOfBandAuthPluginsAPI.GetOOBAuthPluginDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthOutOfBandAuthPluginsAPI.GetOOBAuthPluginDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOOBAuthPluginDescriptors`: OutOfBandAuthPluginDescriptors
    fmt.Fprintf(os.Stdout, "Response from `OauthOutOfBandAuthPluginsAPI.GetOOBAuthPluginDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOOBAuthPluginDescriptorsRequest struct via the builder pattern


### Return type

[**OutOfBandAuthPluginDescriptors**](OutOfBandAuthPluginDescriptors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOOBAuthenticator

> OutOfBandAuthenticator GetOOBAuthenticator(ctx, id).Execute()

Get a specific Out of Band authenticator plugin instance.

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
    id := "id_example" // string | ID of Out of Band authenticator plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthOutOfBandAuthPluginsAPI.GetOOBAuthenticator(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthOutOfBandAuthPluginsAPI.GetOOBAuthenticator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOOBAuthenticator`: OutOfBandAuthenticator
    fmt.Fprintf(os.Stdout, "Response from `OauthOutOfBandAuthPluginsAPI.GetOOBAuthenticator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Out of Band authenticator plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOOBAuthenticatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutOfBandAuthenticator**](OutOfBandAuthenticator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOOBAuthenticators

> OutOfBandAuthenticators GetOOBAuthenticators(ctx).Execute()

Get a list of Out of Band authenticator plugin instances.

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
    resp, r, err := apiClient.OauthOutOfBandAuthPluginsAPI.GetOOBAuthenticators(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthOutOfBandAuthPluginsAPI.GetOOBAuthenticators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOOBAuthenticators`: OutOfBandAuthenticators
    fmt.Fprintf(os.Stdout, "Response from `OauthOutOfBandAuthPluginsAPI.GetOOBAuthenticators`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOOBAuthenticatorsRequest struct via the builder pattern


### Return type

[**OutOfBandAuthenticators**](OutOfBandAuthenticators.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvokeOOBActionWithOptions

> ActionResult InvokeOOBActionWithOptions(ctx, id, actionId).Body(body).Execute()

Invokes an action for Out of Band authenticator plugin instance.

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
    id := "id_example" // string | ID of the Out of Band authenticator plugin instance to which these actions belongs to.
    actionId := "actionId_example" // string | ID of the action.
    body := *openapiclient.NewActionOptions([]openapiclient.ActionParameter{*openapiclient.NewActionParameter("Name_example")}) // ActionOptions | Action options for action invoked. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthOutOfBandAuthPluginsAPI.InvokeOOBActionWithOptions(context.Background(), id, actionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthOutOfBandAuthPluginsAPI.InvokeOOBActionWithOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvokeOOBActionWithOptions`: ActionResult
    fmt.Fprintf(os.Stdout, "Response from `OauthOutOfBandAuthPluginsAPI.InvokeOOBActionWithOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Out of Band authenticator plugin instance to which these actions belongs to. | 
**actionId** | **string** | ID of the action. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvokeOOBActionWithOptionsRequest struct via the builder pattern


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


## UpdateOOBAuthenticator

> OutOfBandAuthenticator UpdateOOBAuthenticator(ctx, id).Body(body).Execute()

Update an Out of Band authenticator plugin instance.

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
    id := "id_example" // string | ID of Out of Band authenticator plugin instance.
    body := *openapiclient.NewOutOfBandAuthenticator("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // OutOfBandAuthenticator | Configuration for an Out of Band authenticator plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthOutOfBandAuthPluginsAPI.UpdateOOBAuthenticator(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthOutOfBandAuthPluginsAPI.UpdateOOBAuthenticator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOOBAuthenticator`: OutOfBandAuthenticator
    fmt.Fprintf(os.Stdout, "Response from `OauthOutOfBandAuthPluginsAPI.UpdateOOBAuthenticator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Out of Band authenticator plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOOBAuthenticatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OutOfBandAuthenticator**](OutOfBandAuthenticator.md) | Configuration for an Out of Band authenticator plugin instance. | 

### Return type

[**OutOfBandAuthenticator**](OutOfBandAuthenticator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

