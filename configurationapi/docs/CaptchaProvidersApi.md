# \CaptchaProvidersAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCaptchaProvider**](CaptchaProvidersAPI.md#CreateCaptchaProvider) | **Post** /captchaProviders | Create a CAPTCHA provider plugin instance.
[**DeleteCaptchaProvider**](CaptchaProvidersAPI.md#DeleteCaptchaProvider) | **Delete** /captchaProviders/{id} | Delete a CAPTCHA provider plugin instance.
[**GetCaptchaProvider**](CaptchaProvidersAPI.md#GetCaptchaProvider) | **Get** /captchaProviders/{id} | Get a specific CAPTCHA provider plugin instance.
[**GetCaptchaProviderPluginDescriptor**](CaptchaProvidersAPI.md#GetCaptchaProviderPluginDescriptor) | **Get** /captchaProviders/descriptors/{id} | Get a CAPTCHA provider plugin descriptor.
[**GetCaptchaProviderPluginDescriptors**](CaptchaProvidersAPI.md#GetCaptchaProviderPluginDescriptors) | **Get** /captchaProviders/descriptors | Get a list of available CAPTCHA provider plugin descriptors.
[**GetCaptchaProviders**](CaptchaProvidersAPI.md#GetCaptchaProviders) | **Get** /captchaProviders | Get a list of CAPTCHA provider plugin instances.
[**GetCaptchaProvidersSettings**](CaptchaProvidersAPI.md#GetCaptchaProvidersSettings) | **Get** /captchaProviders/settings | Get general CAPTCHA providers settings.
[**UpdateCaptchaProvider**](CaptchaProvidersAPI.md#UpdateCaptchaProvider) | **Put** /captchaProviders/{id} | Update a CAPTCHA provider plugin instance.
[**UpdateCaptchaProvidersSettings**](CaptchaProvidersAPI.md#UpdateCaptchaProvidersSettings) | **Put** /captchaProviders/settings | Update general CAPTCHA providers settings.



## CreateCaptchaProvider

> CaptchaProvider CreateCaptchaProvider(ctx).Body(body).Execute()

Create a CAPTCHA provider plugin instance.

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
    body := *openapiclient.NewCaptchaProvider("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // CaptchaProvider | Configuration for a CAPTCHA provider plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CaptchaProvidersAPI.CreateCaptchaProvider(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaptchaProvidersAPI.CreateCaptchaProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCaptchaProvider`: CaptchaProvider
    fmt.Fprintf(os.Stdout, "Response from `CaptchaProvidersAPI.CreateCaptchaProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCaptchaProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CaptchaProvider**](CaptchaProvider.md) | Configuration for a CAPTCHA provider plugin instance. | 

### Return type

[**CaptchaProvider**](CaptchaProvider.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCaptchaProvider

> DeleteCaptchaProvider(ctx, id).Execute()

Delete a CAPTCHA provider plugin instance.

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
    id := "id_example" // string | ID of a CAPTCHA provider plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CaptchaProvidersAPI.DeleteCaptchaProvider(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaptchaProvidersAPI.DeleteCaptchaProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a CAPTCHA provider plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCaptchaProviderRequest struct via the builder pattern


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


## GetCaptchaProvider

> CaptchaProvider GetCaptchaProvider(ctx, id).Execute()

Get a specific CAPTCHA provider plugin instance.

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
    id := "id_example" // string | ID of a CAPTCHA provider plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CaptchaProvidersAPI.GetCaptchaProvider(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaptchaProvidersAPI.GetCaptchaProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCaptchaProvider`: CaptchaProvider
    fmt.Fprintf(os.Stdout, "Response from `CaptchaProvidersAPI.GetCaptchaProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a CAPTCHA provider plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptchaProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CaptchaProvider**](CaptchaProvider.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptchaProviderPluginDescriptor

> ACAPTCHAProviderPluginDescriptor GetCaptchaProviderPluginDescriptor(ctx, id).Execute()

Get a CAPTCHA provider plugin descriptor.

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
    id := "id_example" // string | ID of CAPTCHA provider plugin descriptor.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CaptchaProvidersAPI.GetCaptchaProviderPluginDescriptor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaptchaProvidersAPI.GetCaptchaProviderPluginDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCaptchaProviderPluginDescriptor`: ACAPTCHAProviderPluginDescriptor
    fmt.Fprintf(os.Stdout, "Response from `CaptchaProvidersAPI.GetCaptchaProviderPluginDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of CAPTCHA provider plugin descriptor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptchaProviderPluginDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ACAPTCHAProviderPluginDescriptor**](ACAPTCHAProviderPluginDescriptor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptchaProviderPluginDescriptors

> ACollectionOfCAPTCHAProviderPluginDescriptors GetCaptchaProviderPluginDescriptors(ctx).Execute()

Get a list of available CAPTCHA provider plugin descriptors.

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
    resp, r, err := apiClient.CaptchaProvidersAPI.GetCaptchaProviderPluginDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaptchaProvidersAPI.GetCaptchaProviderPluginDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCaptchaProviderPluginDescriptors`: ACollectionOfCAPTCHAProviderPluginDescriptors
    fmt.Fprintf(os.Stdout, "Response from `CaptchaProvidersAPI.GetCaptchaProviderPluginDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptchaProviderPluginDescriptorsRequest struct via the builder pattern


### Return type

[**ACollectionOfCAPTCHAProviderPluginDescriptors**](ACollectionOfCAPTCHAProviderPluginDescriptors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptchaProviders

> CaptchaProviders GetCaptchaProviders(ctx).Execute()

Get a list of CAPTCHA provider plugin instances.

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
    resp, r, err := apiClient.CaptchaProvidersAPI.GetCaptchaProviders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaptchaProvidersAPI.GetCaptchaProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCaptchaProviders`: CaptchaProviders
    fmt.Fprintf(os.Stdout, "Response from `CaptchaProvidersAPI.GetCaptchaProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptchaProvidersRequest struct via the builder pattern


### Return type

[**CaptchaProviders**](CaptchaProviders.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptchaProvidersSettings

> CaptchaProvidersSettings GetCaptchaProvidersSettings(ctx).Execute()

Get general CAPTCHA providers settings.

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
    resp, r, err := apiClient.CaptchaProvidersAPI.GetCaptchaProvidersSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaptchaProvidersAPI.GetCaptchaProvidersSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCaptchaProvidersSettings`: CaptchaProvidersSettings
    fmt.Fprintf(os.Stdout, "Response from `CaptchaProvidersAPI.GetCaptchaProvidersSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptchaProvidersSettingsRequest struct via the builder pattern


### Return type

[**CaptchaProvidersSettings**](CaptchaProvidersSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCaptchaProvider

> CaptchaProvider UpdateCaptchaProvider(ctx, id).Body(body).Execute()

Update a CAPTCHA provider plugin instance.

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
    id := "id_example" // string | ID of a CAPTCHA provider plugin instance.
    body := *openapiclient.NewCaptchaProvider("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // CaptchaProvider | Configuration for a CAPTCHA provider plugin instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CaptchaProvidersAPI.UpdateCaptchaProvider(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaptchaProvidersAPI.UpdateCaptchaProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCaptchaProvider`: CaptchaProvider
    fmt.Fprintf(os.Stdout, "Response from `CaptchaProvidersAPI.UpdateCaptchaProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a CAPTCHA provider plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCaptchaProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CaptchaProvider**](CaptchaProvider.md) | Configuration for a CAPTCHA provider plugin instance. | 

### Return type

[**CaptchaProvider**](CaptchaProvider.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCaptchaProvidersSettings

> CaptchaProvidersSettings UpdateCaptchaProvidersSettings(ctx).Body(body).Execute()

Update general CAPTCHA providers settings.

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
    body := *openapiclient.NewCaptchaProvidersSettings() // CaptchaProvidersSettings | CAPTCHA providers settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CaptchaProvidersAPI.UpdateCaptchaProvidersSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaptchaProvidersAPI.UpdateCaptchaProvidersSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCaptchaProvidersSettings`: CaptchaProvidersSettings
    fmt.Fprintf(os.Stdout, "Response from `CaptchaProvidersAPI.UpdateCaptchaProvidersSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCaptchaProvidersSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CaptchaProvidersSettings**](CaptchaProvidersSettings.md) | CAPTCHA providers settings. | 

### Return type

[**CaptchaProvidersSettings**](CaptchaProvidersSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

