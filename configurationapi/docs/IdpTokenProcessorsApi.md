# \IdpTokenProcessorsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTokenProcessor**](IdpTokenProcessorsAPI.md#CreateTokenProcessor) | **Post** /idp/tokenProcessors | Create a new token processor instance.
[**DeleteTokenProcessor**](IdpTokenProcessorsAPI.md#DeleteTokenProcessor) | **Delete** /idp/tokenProcessors/{id} | Delete a token processor instance.
[**GetTokenProcessor**](IdpTokenProcessorsAPI.md#GetTokenProcessor) | **Get** /idp/tokenProcessors/{id} | Find a token processor instance by ID.
[**GetTokenProcessorDescriptors**](IdpTokenProcessorsAPI.md#GetTokenProcessorDescriptors) | **Get** /idp/tokenProcessors/descriptors | Get the list of available token processors.
[**GetTokenProcessorDescriptorsById**](IdpTokenProcessorsAPI.md#GetTokenProcessorDescriptorsById) | **Get** /idp/tokenProcessors/descriptors/{id} | Get the description of a token processor plugin by ID.
[**GetTokenProcessors**](IdpTokenProcessorsAPI.md#GetTokenProcessors) | **Get** /idp/tokenProcessors | Get the list of token processor instances.
[**UpdateTokenProcessor**](IdpTokenProcessorsAPI.md#UpdateTokenProcessor) | **Put** /idp/tokenProcessors/{id} | Update a token processor instance.



## CreateTokenProcessor

> TokenProcessor CreateTokenProcessor(ctx).Body(body).Execute()

Create a new token processor instance.



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
    body := *openapiclient.NewTokenProcessor("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // TokenProcessor | Configuration for a token processor instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpTokenProcessorsAPI.CreateTokenProcessor(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpTokenProcessorsAPI.CreateTokenProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTokenProcessor`: TokenProcessor
    fmt.Fprintf(os.Stdout, "Response from `IdpTokenProcessorsAPI.CreateTokenProcessor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TokenProcessor**](TokenProcessor.md) | Configuration for a token processor instance. | 

### Return type

[**TokenProcessor**](TokenProcessor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTokenProcessor

> DeleteTokenProcessor(ctx, id).Execute()

Delete a token processor instance.



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
    id := "id_example" // string | ID of the token processor instance to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdpTokenProcessorsAPI.DeleteTokenProcessor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpTokenProcessorsAPI.DeleteTokenProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the token processor instance to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenProcessorRequest struct via the builder pattern


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


## GetTokenProcessor

> TokenProcessor GetTokenProcessor(ctx, id).Execute()

Find a token processor instance by ID.



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
    id := "id_example" // string | ID of the token processor instance to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpTokenProcessorsAPI.GetTokenProcessor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpTokenProcessorsAPI.GetTokenProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenProcessor`: TokenProcessor
    fmt.Fprintf(os.Stdout, "Response from `IdpTokenProcessorsAPI.GetTokenProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the token processor instance to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenProcessor**](TokenProcessor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenProcessorDescriptors

> TokenProcessorDescriptors GetTokenProcessorDescriptors(ctx).Execute()

Get the list of available token processors.

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
    resp, r, err := apiClient.IdpTokenProcessorsAPI.GetTokenProcessorDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpTokenProcessorsAPI.GetTokenProcessorDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenProcessorDescriptors`: TokenProcessorDescriptors
    fmt.Fprintf(os.Stdout, "Response from `IdpTokenProcessorsAPI.GetTokenProcessorDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenProcessorDescriptorsRequest struct via the builder pattern


### Return type

[**TokenProcessorDescriptors**](TokenProcessorDescriptors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenProcessorDescriptorsById

> TokenProcessorDescriptor GetTokenProcessorDescriptorsById(ctx, id).Execute()

Get the description of a token processor plugin by ID.



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
    id := "id_example" // string | ID of a token processor descriptor to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpTokenProcessorsAPI.GetTokenProcessorDescriptorsById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpTokenProcessorsAPI.GetTokenProcessorDescriptorsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenProcessorDescriptorsById`: TokenProcessorDescriptor
    fmt.Fprintf(os.Stdout, "Response from `IdpTokenProcessorsAPI.GetTokenProcessorDescriptorsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a token processor descriptor to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenProcessorDescriptorsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenProcessorDescriptor**](TokenProcessorDescriptor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenProcessors

> TokenProcessors GetTokenProcessors(ctx).Execute()

Get the list of token processor instances.

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
    resp, r, err := apiClient.IdpTokenProcessorsAPI.GetTokenProcessors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpTokenProcessorsAPI.GetTokenProcessors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenProcessors`: TokenProcessors
    fmt.Fprintf(os.Stdout, "Response from `IdpTokenProcessorsAPI.GetTokenProcessors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenProcessorsRequest struct via the builder pattern


### Return type

[**TokenProcessors**](TokenProcessors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTokenProcessor

> TokenProcessor UpdateTokenProcessor(ctx, id).Body(body).Execute()

Update a token processor instance.



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
    id := "id_example" // string | ID of token processor instance.
    body := *openapiclient.NewTokenProcessor("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // TokenProcessor | Configuration for updated token processor instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpTokenProcessorsAPI.UpdateTokenProcessor(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpTokenProcessorsAPI.UpdateTokenProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTokenProcessor`: TokenProcessor
    fmt.Fprintf(os.Stdout, "Response from `IdpTokenProcessorsAPI.UpdateTokenProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of token processor instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTokenProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TokenProcessor**](TokenProcessor.md) | Configuration for updated token processor instance. | 

### Return type

[**TokenProcessor**](TokenProcessor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

