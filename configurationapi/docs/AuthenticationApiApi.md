# \AuthenticationApiAPI

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplication**](AuthenticationApiAPI.md#CreateApplication) | **Post** /authenticationApi/applications | Create a new Authentication API Application.
[**DeleteApplication**](AuthenticationApiAPI.md#DeleteApplication) | **Delete** /authenticationApi/applications/{id} | Delete an Authentication API Application.
[**GetApplication**](AuthenticationApiAPI.md#GetApplication) | **Get** /authenticationApi/applications/{id} | Find Authentication API Application by ID.
[**GetAuthenticationApiApplications**](AuthenticationApiAPI.md#GetAuthenticationApiApplications) | **Get** /authenticationApi/applications | Get the collection of Authentication API Applications.
[**GetAuthenticationApiSettings**](AuthenticationApiAPI.md#GetAuthenticationApiSettings) | **Get** /authenticationApi/settings | Get the Authentication API settings.
[**UpdateApplication**](AuthenticationApiAPI.md#UpdateApplication) | **Put** /authenticationApi/applications/{id} | Update an Authentication API Application.
[**UpdateAuthenticationApiSettings**](AuthenticationApiAPI.md#UpdateAuthenticationApiSettings) | **Put** /authenticationApi/settings | Set the Authentication API settings.



## CreateApplication

> AuthnApiApplication CreateApplication(ctx).Body(body).Execute()

Create a new Authentication API Application.

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
    body := *openapiclient.NewAuthnApiApplication("Id_example", "Name_example", "Url_example") // AuthnApiApplication | Configuration for new Authentication API Application.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApiAPI.CreateApplication(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApiAPI.CreateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplication`: AuthnApiApplication
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApiAPI.CreateApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthnApiApplication**](AuthnApiApplication.md) | Configuration for new Authentication API Application. | 

### Return type

[**AuthnApiApplication**](AuthnApiApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplication

> DeleteApplication(ctx, id).Execute()

Delete an Authentication API Application.



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
    id := "id_example" // string | ID of Authentication API Application to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthenticationApiAPI.DeleteApplication(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApiAPI.DeleteApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Authentication API Application to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


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


## GetApplication

> AuthnApiApplication GetApplication(ctx, id).Execute()

Find Authentication API Application by ID.



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
    id := "id_example" // string | ID of the Authentication API Application to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApiAPI.GetApplication(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApiAPI.GetApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplication`: AuthnApiApplication
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApiAPI.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Authentication API Application to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthnApiApplication**](AuthnApiApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticationApiApplications

> AuthnApiApplications GetAuthenticationApiApplications(ctx).Execute()

Get the collection of Authentication API Applications.

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
    resp, r, err := apiClient.AuthenticationApiAPI.GetAuthenticationApiApplications(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApiAPI.GetAuthenticationApiApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticationApiApplications`: AuthnApiApplications
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApiAPI.GetAuthenticationApiApplications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationApiApplicationsRequest struct via the builder pattern


### Return type

[**AuthnApiApplications**](AuthnApiApplications.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticationApiSettings

> AuthnApiSettings GetAuthenticationApiSettings(ctx).Execute()

Get the Authentication API settings.

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
    resp, r, err := apiClient.AuthenticationApiAPI.GetAuthenticationApiSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApiAPI.GetAuthenticationApiSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticationApiSettings`: AuthnApiSettings
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApiAPI.GetAuthenticationApiSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationApiSettingsRequest struct via the builder pattern


### Return type

[**AuthnApiSettings**](AuthnApiSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> AuthnApiApplication UpdateApplication(ctx, id).Body(body).Execute()

Update an Authentication API Application.



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
    id := "id_example" // string | ID of the Authentication API Application to update.
    body := *openapiclient.NewAuthnApiApplication("Id_example", "Name_example", "Url_example") // AuthnApiApplication | Configuration for updated application.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApiAPI.UpdateApplication(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApiAPI.UpdateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplication`: AuthnApiApplication
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApiAPI.UpdateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Authentication API Application to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthnApiApplication**](AuthnApiApplication.md) | Configuration for updated application. | 

### Return type

[**AuthnApiApplication**](AuthnApiApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthenticationApiSettings

> AuthnApiSettings UpdateAuthenticationApiSettings(ctx).Body(body).Execute()

Set the Authentication API settings.

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
    body := *openapiclient.NewAuthnApiSettings() // AuthnApiSettings | Authentication API Settings

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApiAPI.UpdateAuthenticationApiSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApiAPI.UpdateAuthenticationApiSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthenticationApiSettings`: AuthnApiSettings
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApiAPI.UpdateAuthenticationApiSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthenticationApiSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthnApiSettings**](AuthnApiSettings.md) | Authentication API Settings | 

### Return type

[**AuthnApiSettings**](AuthnApiSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

