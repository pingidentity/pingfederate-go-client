# \OauthJwtBearerGrantProcessorsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJwtBearerGrantProcessor**](OauthJwtBearerGrantProcessorsAPI.md#CreateJwtBearerGrantProcessor) | **Post** /oauth/jwtBearerGrantProcessors | Create a JWT Bearer Grant Processor plugin instance.
[**DeleteJwtBearerGrantProcessor**](OauthJwtBearerGrantProcessorsAPI.md#DeleteJwtBearerGrantProcessor) | **Delete** /oauth/jwtBearerGrantProcessors/{id} | Delete a JWT Bearer Grant Processor plugin instance.
[**GetJwtBearerGrantProcessor**](OauthJwtBearerGrantProcessorsAPI.md#GetJwtBearerGrantProcessor) | **Get** /oauth/jwtBearerGrantProcessors/{id} | Get a specific JWT Bearer Grant Processor plugin instance.
[**GetJwtBearerGrantProcessorDescriptor**](OauthJwtBearerGrantProcessorsAPI.md#GetJwtBearerGrantProcessorDescriptor) | **Get** /oauth/jwtBearerGrantProcessors/descriptors/{id} | Get a JWT Bearer Grant Processor plugin descriptor.
[**GetJwtBearerGrantProcessorDescriptors**](OauthJwtBearerGrantProcessorsAPI.md#GetJwtBearerGrantProcessorDescriptors) | **Get** /oauth/jwtBearerGrantProcessors/descriptors | Get a list of available JWT Bearer Grant Processor plugin descriptors.
[**GetJwtBearerGrantProcessors**](OauthJwtBearerGrantProcessorsAPI.md#GetJwtBearerGrantProcessors) | **Get** /oauth/jwtBearerGrantProcessors | Get a list of JWT Bearer Grant Processor plugin instances.
[**UpdateJwtBearerGrantProcessor**](OauthJwtBearerGrantProcessorsAPI.md#UpdateJwtBearerGrantProcessor) | **Put** /oauth/jwtBearerGrantProcessors/{id} | Update a JWT Bearer Grant Processor plugin instance.



## CreateJwtBearerGrantProcessor

> JwtBearerGrantProcessor CreateJwtBearerGrantProcessor(ctx).Body(body).Execute()

Create a JWT Bearer Grant Processor plugin instance.

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
    body := openapiclient.JwtBearerGrantProcessor{PluginInstance: openapiclient.NewPluginInstance("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration())} // JwtBearerGrantProcessor | Configuration for a JWT Bearer Grant Processor plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthJwtBearerGrantProcessorsAPI.CreateJwtBearerGrantProcessor(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthJwtBearerGrantProcessorsAPI.CreateJwtBearerGrantProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateJwtBearerGrantProcessor`: JwtBearerGrantProcessor
    fmt.Fprintf(os.Stdout, "Response from `OauthJwtBearerGrantProcessorsAPI.CreateJwtBearerGrantProcessor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJwtBearerGrantProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**JwtBearerGrantProcessor**](JwtBearerGrantProcessor.md) | Configuration for a JWT Bearer Grant Processor plugin instance. | 

### Return type

[**JwtBearerGrantProcessor**](JwtBearerGrantProcessor.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJwtBearerGrantProcessor

> DeleteJwtBearerGrantProcessor(ctx, id).Execute()

Delete a JWT Bearer Grant Processor plugin instance.

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
    id := "id_example" // string | ID of a JWT Bearer Grant Processor plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthJwtBearerGrantProcessorsAPI.DeleteJwtBearerGrantProcessor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthJwtBearerGrantProcessorsAPI.DeleteJwtBearerGrantProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a JWT Bearer Grant Processor plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJwtBearerGrantProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJwtBearerGrantProcessor

> JwtBearerGrantProcessor GetJwtBearerGrantProcessor(ctx, id).Execute()

Get a specific JWT Bearer Grant Processor plugin instance.

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
    id := "id_example" // string | ID of a JWT Bearer Grant Processor plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthJwtBearerGrantProcessorsAPI.GetJwtBearerGrantProcessor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthJwtBearerGrantProcessorsAPI.GetJwtBearerGrantProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJwtBearerGrantProcessor`: JwtBearerGrantProcessor
    fmt.Fprintf(os.Stdout, "Response from `OauthJwtBearerGrantProcessorsAPI.GetJwtBearerGrantProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a JWT Bearer Grant Processor plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJwtBearerGrantProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JwtBearerGrantProcessor**](JwtBearerGrantProcessor.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJwtBearerGrantProcessorDescriptor

> GetJwtBearerGrantProcessorDescriptor(ctx, id).Execute()

Get a JWT Bearer Grant Processor plugin descriptor.

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
    id := "id_example" // string | ID of JWT Bearer Grant Processor plugin descriptor.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthJwtBearerGrantProcessorsAPI.GetJwtBearerGrantProcessorDescriptor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthJwtBearerGrantProcessorsAPI.GetJwtBearerGrantProcessorDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of JWT Bearer Grant Processor plugin descriptor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJwtBearerGrantProcessorDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJwtBearerGrantProcessorDescriptors

> GetJwtBearerGrantProcessorDescriptors(ctx).Execute()

Get a list of available JWT Bearer Grant Processor plugin descriptors.

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
    r, err := apiClient.OauthJwtBearerGrantProcessorsAPI.GetJwtBearerGrantProcessorDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthJwtBearerGrantProcessorsAPI.GetJwtBearerGrantProcessorDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetJwtBearerGrantProcessorDescriptorsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJwtBearerGrantProcessors

> JwtBearerGrantProcessorInstances GetJwtBearerGrantProcessors(ctx).Execute()

Get a list of JWT Bearer Grant Processor plugin instances.

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
    resp, r, err := apiClient.OauthJwtBearerGrantProcessorsAPI.GetJwtBearerGrantProcessors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthJwtBearerGrantProcessorsAPI.GetJwtBearerGrantProcessors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJwtBearerGrantProcessors`: JwtBearerGrantProcessorInstances
    fmt.Fprintf(os.Stdout, "Response from `OauthJwtBearerGrantProcessorsAPI.GetJwtBearerGrantProcessors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetJwtBearerGrantProcessorsRequest struct via the builder pattern


### Return type

[**JwtBearerGrantProcessorInstances**](JwtBearerGrantProcessorInstances.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateJwtBearerGrantProcessor

> JwtBearerGrantProcessor UpdateJwtBearerGrantProcessor(ctx, id).Body(body).Execute()

Update a JWT Bearer Grant Processor plugin instance.

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
    id := "id_example" // string | ID of a JWT Bearer Grant Processor plugin instance.
    body := openapiclient.JwtBearerGrantProcessor{PluginInstance: openapiclient.NewPluginInstance("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration())} // JwtBearerGrantProcessor | Configuration for a JWT Bearer Grant Processor plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthJwtBearerGrantProcessorsAPI.UpdateJwtBearerGrantProcessor(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthJwtBearerGrantProcessorsAPI.UpdateJwtBearerGrantProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateJwtBearerGrantProcessor`: JwtBearerGrantProcessor
    fmt.Fprintf(os.Stdout, "Response from `OauthJwtBearerGrantProcessorsAPI.UpdateJwtBearerGrantProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a JWT Bearer Grant Processor plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJwtBearerGrantProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**JwtBearerGrantProcessor**](JwtBearerGrantProcessor.md) | Configuration for a JWT Bearer Grant Processor plugin instance. | 

### Return type

[**JwtBearerGrantProcessor**](JwtBearerGrantProcessor.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

