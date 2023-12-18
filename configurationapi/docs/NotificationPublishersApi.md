# \NotificationPublishersAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationPublisher**](NotificationPublishersAPI.md#CreateNotificationPublisher) | **Post** /notificationPublishers | Create a notification publisher plugin instance.
[**DeleteNotificationPublisher**](NotificationPublishersAPI.md#DeleteNotificationPublisher) | **Delete** /notificationPublishers/{id} | Delete a notification publisher plugin instance.
[**GetNotificationPublisher**](NotificationPublishersAPI.md#GetNotificationPublisher) | **Get** /notificationPublishers/{id} | Get a specific notification publisher plugin instance.
[**GetNotificationPublisherPluginDescriptor**](NotificationPublishersAPI.md#GetNotificationPublisherPluginDescriptor) | **Get** /notificationPublishers/descriptors/{id} | Get the description of a notification publisher plugin descriptor.
[**GetNotificationPublisherPluginDescriptors**](NotificationPublishersAPI.md#GetNotificationPublisherPluginDescriptors) | **Get** /notificationPublishers/descriptors | Get the list of available Notification Publisher Plugin descriptors.
[**GetNotificationPublishers**](NotificationPublishersAPI.md#GetNotificationPublishers) | **Get** /notificationPublishers | Get a list of notification publisher plugin instances.
[**GetNotificationPublishersAction**](NotificationPublishersAPI.md#GetNotificationPublishersAction) | **Get** /notificationPublishers/{id}/actions/{actionId} | Find an notification publisher plugin instance&#39;s action by ID.
[**GetNotificationPublishersActions**](NotificationPublishersAPI.md#GetNotificationPublishersActions) | **Get** /notificationPublishers/{id}/actions | List the actions for a notification publisher plugin instance.
[**GetNotificationPublishersSettings**](NotificationPublishersAPI.md#GetNotificationPublishersSettings) | **Get** /notificationPublishers/settings | Get general notification publisher settings.
[**InvokeNotificationPublishersActionWithOptions**](NotificationPublishersAPI.md#InvokeNotificationPublishersActionWithOptions) | **Post** /notificationPublishers/{id}/actions/{actionId}/invokeAction | Invokes an action for notification publisher plugin instance.
[**UpdateNotificationPublisher**](NotificationPublishersAPI.md#UpdateNotificationPublisher) | **Put** /notificationPublishers/{id} | Update a notification publisher plugin instance.
[**UpdateNotificationPublishersSettings**](NotificationPublishersAPI.md#UpdateNotificationPublishersSettings) | **Put** /notificationPublishers/settings | Update general notification publisher settings.



## CreateNotificationPublisher

> NotificationPublisher CreateNotificationPublisher(ctx).Body(body).Execute()

Create a notification publisher plugin instance.

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
    body := *openapiclient.NewNotificationPublisher("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // NotificationPublisher | Configuration for a notification publisher plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationPublishersAPI.CreateNotificationPublisher(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationPublishersAPI.CreateNotificationPublisher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationPublisher`: NotificationPublisher
    fmt.Fprintf(os.Stdout, "Response from `NotificationPublishersAPI.CreateNotificationPublisher`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationPublisherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NotificationPublisher**](NotificationPublisher.md) | Configuration for a notification publisher plugin instance. | 

### Return type

[**NotificationPublisher**](NotificationPublisher.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationPublisher

> DeleteNotificationPublisher(ctx, id).Execute()

Delete a notification publisher plugin instance.

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
    id := "id_example" // string | ID of a notification publisher plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NotificationPublishersAPI.DeleteNotificationPublisher(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationPublishersAPI.DeleteNotificationPublisher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a notification publisher plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationPublisherRequest struct via the builder pattern


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


## GetNotificationPublisher

> NotificationPublisher GetNotificationPublisher(ctx, id).Execute()

Get a specific notification publisher plugin instance.

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
    id := "id_example" // string | ID of a notification publisher plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationPublishersAPI.GetNotificationPublisher(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationPublishersAPI.GetNotificationPublisher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationPublisher`: NotificationPublisher
    fmt.Fprintf(os.Stdout, "Response from `NotificationPublishersAPI.GetNotificationPublisher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a notification publisher plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationPublisherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationPublisher**](NotificationPublisher.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationPublisherPluginDescriptor

> NotificationPublisherDescriptor GetNotificationPublisherPluginDescriptor(ctx, id).Execute()

Get the description of a notification publisher plugin descriptor.

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
    id := "id_example" // string | ID of notification publisher plugin descriptor.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationPublishersAPI.GetNotificationPublisherPluginDescriptor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationPublishersAPI.GetNotificationPublisherPluginDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationPublisherPluginDescriptor`: NotificationPublisherDescriptor
    fmt.Fprintf(os.Stdout, "Response from `NotificationPublishersAPI.GetNotificationPublisherPluginDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of notification publisher plugin descriptor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationPublisherPluginDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationPublisherDescriptor**](NotificationPublisherDescriptor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationPublisherPluginDescriptors

> NotificationPublisherDescriptors GetNotificationPublisherPluginDescriptors(ctx).Execute()

Get the list of available Notification Publisher Plugin descriptors.

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
    resp, r, err := apiClient.NotificationPublishersAPI.GetNotificationPublisherPluginDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationPublishersAPI.GetNotificationPublisherPluginDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationPublisherPluginDescriptors`: NotificationPublisherDescriptors
    fmt.Fprintf(os.Stdout, "Response from `NotificationPublishersAPI.GetNotificationPublisherPluginDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationPublisherPluginDescriptorsRequest struct via the builder pattern


### Return type

[**NotificationPublisherDescriptors**](NotificationPublisherDescriptors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationPublishers

> NotificationPublishers GetNotificationPublishers(ctx).Execute()

Get a list of notification publisher plugin instances.

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
    resp, r, err := apiClient.NotificationPublishersAPI.GetNotificationPublishers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationPublishersAPI.GetNotificationPublishers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationPublishers`: NotificationPublishers
    fmt.Fprintf(os.Stdout, "Response from `NotificationPublishersAPI.GetNotificationPublishers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationPublishersRequest struct via the builder pattern


### Return type

[**NotificationPublishers**](NotificationPublishers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationPublishersAction

> Action GetNotificationPublishersAction(ctx, id, actionId).Execute()

Find an notification publisher plugin instance's action by ID.

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
    id := "id_example" // string | ID of the notification publisher plugin instance to which these actions belongs to.
    actionId := "actionId_example" // string | ID of the action to get.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationPublishersAPI.GetNotificationPublishersAction(context.Background(), id, actionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationPublishersAPI.GetNotificationPublishersAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationPublishersAction`: Action
    fmt.Fprintf(os.Stdout, "Response from `NotificationPublishersAPI.GetNotificationPublishersAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the notification publisher plugin instance to which these actions belongs to. | 
**actionId** | **string** | ID of the action to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationPublishersActionRequest struct via the builder pattern


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


## GetNotificationPublishersActions

> Actions GetNotificationPublishersActions(ctx, id).Execute()

List the actions for a notification publisher plugin instance.

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
    id := "id_example" // string | ID of the notification publisher plugin instance to which these actions belongs to.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationPublishersAPI.GetNotificationPublishersActions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationPublishersAPI.GetNotificationPublishersActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationPublishersActions`: Actions
    fmt.Fprintf(os.Stdout, "Response from `NotificationPublishersAPI.GetNotificationPublishersActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the notification publisher plugin instance to which these actions belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationPublishersActionsRequest struct via the builder pattern


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


## GetNotificationPublishersSettings

> NotificationPublishersSettings GetNotificationPublishersSettings(ctx).Execute()

Get general notification publisher settings.

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
    resp, r, err := apiClient.NotificationPublishersAPI.GetNotificationPublishersSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationPublishersAPI.GetNotificationPublishersSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationPublishersSettings`: NotificationPublishersSettings
    fmt.Fprintf(os.Stdout, "Response from `NotificationPublishersAPI.GetNotificationPublishersSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationPublishersSettingsRequest struct via the builder pattern


### Return type

[**NotificationPublishersSettings**](NotificationPublishersSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvokeNotificationPublishersActionWithOptions

> ActionResult InvokeNotificationPublishersActionWithOptions(ctx, id, actionId).Body(body).Execute()

Invokes an action for notification publisher plugin instance.

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
    id := "id_example" // string | ID of the notification publisher plugin instance to which these actions belongs to.
    actionId := "actionId_example" // string | ID of the action to get.
    body := *openapiclient.NewActionOptions([]openapiclient.ActionParameter{*openapiclient.NewActionParameter("Name_example")}) // ActionOptions | Action options for action invoked. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationPublishersAPI.InvokeNotificationPublishersActionWithOptions(context.Background(), id, actionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationPublishersAPI.InvokeNotificationPublishersActionWithOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvokeNotificationPublishersActionWithOptions`: ActionResult
    fmt.Fprintf(os.Stdout, "Response from `NotificationPublishersAPI.InvokeNotificationPublishersActionWithOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the notification publisher plugin instance to which these actions belongs to. | 
**actionId** | **string** | ID of the action to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvokeNotificationPublishersActionWithOptionsRequest struct via the builder pattern


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


## UpdateNotificationPublisher

> NotificationPublisher UpdateNotificationPublisher(ctx, id).Body(body).Execute()

Update a notification publisher plugin instance.

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
    id := "id_example" // string | ID of a notification publisher plugin instance.
    body := *openapiclient.NewNotificationPublisher("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // NotificationPublisher | Configuration for a notification publisher plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationPublishersAPI.UpdateNotificationPublisher(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationPublishersAPI.UpdateNotificationPublisher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationPublisher`: NotificationPublisher
    fmt.Fprintf(os.Stdout, "Response from `NotificationPublishersAPI.UpdateNotificationPublisher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a notification publisher plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationPublisherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NotificationPublisher**](NotificationPublisher.md) | Configuration for a notification publisher plugin instance. | 

### Return type

[**NotificationPublisher**](NotificationPublisher.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationPublishersSettings

> NotificationPublishersSettings UpdateNotificationPublishersSettings(ctx).Body(body).Execute()

Update general notification publisher settings.

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
    body := *openapiclient.NewNotificationPublishersSettings() // NotificationPublishersSettings | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationPublishersAPI.UpdateNotificationPublishersSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationPublishersAPI.UpdateNotificationPublishersSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationPublishersSettings`: NotificationPublishersSettings
    fmt.Fprintf(os.Stdout, "Response from `NotificationPublishersAPI.UpdateNotificationPublishersSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationPublishersSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NotificationPublishersSettings**](NotificationPublishersSettings.md) |  | 

### Return type

[**NotificationPublishersSettings**](NotificationPublishersSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

