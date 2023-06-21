# \OauthAuthorizationDetailProcessorsApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthorizationDetailProcessor**](OauthAuthorizationDetailProcessorsApi.md#CreateAuthorizationDetailProcessor) | **Post** /oauth/authorizationDetailProcessors | Create an authorization detail processor plugin instance.
[**DeleteAuthorizationDetailProcessor**](OauthAuthorizationDetailProcessorsApi.md#DeleteAuthorizationDetailProcessor) | **Delete** /oauth/authorizationDetailProcessors/{id} | Delete an authorization detail processor plugin instance.
[**GetAuthorizationDetailProcessor**](OauthAuthorizationDetailProcessorsApi.md#GetAuthorizationDetailProcessor) | **Get** /oauth/authorizationDetailProcessors/{id} | Get a specific authorization detail processor plugin instance.
[**GetAuthorizationDetailProcessorPluginDescriptor**](OauthAuthorizationDetailProcessorsApi.md#GetAuthorizationDetailProcessorPluginDescriptor) | **Get** /oauth/authorizationDetailProcessors/descriptors/{id} | Get an authorization detail processor plugin descriptor.
[**GetAuthorizationDetailProcessorPluginDescriptors**](OauthAuthorizationDetailProcessorsApi.md#GetAuthorizationDetailProcessorPluginDescriptors) | **Get** /oauth/authorizationDetailProcessors/descriptors | Get a list of available authorization detail processor plugin descriptors.
[**GetAuthorizationDetailProcessors**](OauthAuthorizationDetailProcessorsApi.md#GetAuthorizationDetailProcessors) | **Get** /oauth/authorizationDetailProcessors | Get a list of authorization detail processor plugin instances.
[**UpdateAuthorizationDetailProcessor**](OauthAuthorizationDetailProcessorsApi.md#UpdateAuthorizationDetailProcessor) | **Put** /oauth/authorizationDetailProcessors/{id} | Update an authorization detail processor plugin instance.



## CreateAuthorizationDetailProcessor

> AuthorizationDetailProcessor CreateAuthorizationDetailProcessor(ctx).Body(body).Execute()

Create an authorization detail processor plugin instance.

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
    body := *openapiclient.NewAuthorizationDetailProcessor("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // AuthorizationDetailProcessor | Configuration for a authorization detail processor plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthorizationDetailProcessorsApi.CreateAuthorizationDetailProcessor(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailProcessorsApi.CreateAuthorizationDetailProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthorizationDetailProcessor`: AuthorizationDetailProcessor
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthorizationDetailProcessorsApi.CreateAuthorizationDetailProcessor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthorizationDetailProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthorizationDetailProcessor**](AuthorizationDetailProcessor.md) | Configuration for a authorization detail processor plugin instance. | 

### Return type

[**AuthorizationDetailProcessor**](AuthorizationDetailProcessor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthorizationDetailProcessor

> DeleteAuthorizationDetailProcessor(ctx, id).Execute()

Delete an authorization detail processor plugin instance.

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
    id := "id_example" // string | ID of an authorization detail processor plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthAuthorizationDetailProcessorsApi.DeleteAuthorizationDetailProcessor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailProcessorsApi.DeleteAuthorizationDetailProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an authorization detail processor plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthorizationDetailProcessorRequest struct via the builder pattern


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


## GetAuthorizationDetailProcessor

> AuthorizationDetailProcessor GetAuthorizationDetailProcessor(ctx, id).Execute()

Get a specific authorization detail processor plugin instance.

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
    id := "id_example" // string | ID of an authorization detail processor plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthorizationDetailProcessorsApi.GetAuthorizationDetailProcessor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailProcessorsApi.GetAuthorizationDetailProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationDetailProcessor`: AuthorizationDetailProcessor
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthorizationDetailProcessorsApi.GetAuthorizationDetailProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an authorization detail processor plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationDetailProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthorizationDetailProcessor**](AuthorizationDetailProcessor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorizationDetailProcessorPluginDescriptor

> AuthorizationDetailProcessorDescriptor GetAuthorizationDetailProcessorPluginDescriptor(ctx, id).Execute()

Get an authorization detail processor plugin descriptor.

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
    id := "id_example" // string | ID of authorization detail processor plugin descriptor.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthorizationDetailProcessorsApi.GetAuthorizationDetailProcessorPluginDescriptor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailProcessorsApi.GetAuthorizationDetailProcessorPluginDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationDetailProcessorPluginDescriptor`: AuthorizationDetailProcessorDescriptor
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthorizationDetailProcessorsApi.GetAuthorizationDetailProcessorPluginDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of authorization detail processor plugin descriptor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationDetailProcessorPluginDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthorizationDetailProcessorDescriptor**](AuthorizationDetailProcessorDescriptor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorizationDetailProcessorPluginDescriptors

> AuthorizationDetailProcessorDescriptors GetAuthorizationDetailProcessorPluginDescriptors(ctx).Execute()

Get a list of available authorization detail processor plugin descriptors.

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
    resp, r, err := apiClient.OauthAuthorizationDetailProcessorsApi.GetAuthorizationDetailProcessorPluginDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailProcessorsApi.GetAuthorizationDetailProcessorPluginDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationDetailProcessorPluginDescriptors`: AuthorizationDetailProcessorDescriptors
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthorizationDetailProcessorsApi.GetAuthorizationDetailProcessorPluginDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationDetailProcessorPluginDescriptorsRequest struct via the builder pattern


### Return type

[**AuthorizationDetailProcessorDescriptors**](AuthorizationDetailProcessorDescriptors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorizationDetailProcessors

> AuthorizationDetailProcessors GetAuthorizationDetailProcessors(ctx).Execute()

Get a list of authorization detail processor plugin instances.

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
    resp, r, err := apiClient.OauthAuthorizationDetailProcessorsApi.GetAuthorizationDetailProcessors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailProcessorsApi.GetAuthorizationDetailProcessors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationDetailProcessors`: AuthorizationDetailProcessors
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthorizationDetailProcessorsApi.GetAuthorizationDetailProcessors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationDetailProcessorsRequest struct via the builder pattern


### Return type

[**AuthorizationDetailProcessors**](AuthorizationDetailProcessors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthorizationDetailProcessor

> AuthorizationDetailProcessor UpdateAuthorizationDetailProcessor(ctx, id).Body(body).Execute()

Update an authorization detail processor plugin instance.

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
    id := "id_example" // string | ID of an authorization detail processor plugin instance.
    body := *openapiclient.NewAuthorizationDetailProcessor("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // AuthorizationDetailProcessor | Configuration for a authorization detail processor plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthorizationDetailProcessorsApi.UpdateAuthorizationDetailProcessor(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailProcessorsApi.UpdateAuthorizationDetailProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthorizationDetailProcessor`: AuthorizationDetailProcessor
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthorizationDetailProcessorsApi.UpdateAuthorizationDetailProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an authorization detail processor plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthorizationDetailProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthorizationDetailProcessor**](AuthorizationDetailProcessor.md) | Configuration for a authorization detail processor plugin instance. | 

### Return type

[**AuthorizationDetailProcessor**](AuthorizationDetailProcessor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

