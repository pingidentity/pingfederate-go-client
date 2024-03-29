# \IdentityStoreProvisionersAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentityStoreProvisioner**](IdentityStoreProvisionersAPI.md#CreateIdentityStoreProvisioner) | **Post** /identityStoreProvisioners | Create a new identity store provisioner instance.
[**DeleteIdentityStoreProvisioner**](IdentityStoreProvisionersAPI.md#DeleteIdentityStoreProvisioner) | **Delete** /identityStoreProvisioners/{id} | Delete an identity store provisioner instance
[**GetIdentityStoreProvisioner**](IdentityStoreProvisionersAPI.md#GetIdentityStoreProvisioner) | **Get** /identityStoreProvisioners/{id} | Get an identity store provisioner by ID.
[**GetIdentityStoreProvisionerDescriptorById**](IdentityStoreProvisionersAPI.md#GetIdentityStoreProvisionerDescriptorById) | **Get** /identityStoreProvisioners/descriptors/{id} | Get the descriptor of an identity store provisioner by ID.
[**GetIdentityStoreProvisionerDescriptors**](IdentityStoreProvisionersAPI.md#GetIdentityStoreProvisionerDescriptors) | **Get** /identityStoreProvisioners/descriptors | Get the list of available identity store provisioner descriptors.
[**GetIdentityStoreProvisioners**](IdentityStoreProvisionersAPI.md#GetIdentityStoreProvisioners) | **Get** /identityStoreProvisioners | Get the list of configured identity store provisioner instances.
[**UpdateIdentityStoreProvisioner**](IdentityStoreProvisionersAPI.md#UpdateIdentityStoreProvisioner) | **Put** /identityStoreProvisioners/{id} | Update an identity store provisioner instance



## CreateIdentityStoreProvisioner

> IdentityStoreProvisioner CreateIdentityStoreProvisioner(ctx).Body(body).Execute()

Create a new identity store provisioner instance.

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
    body := *openapiclient.NewIdentityStoreProvisioner("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // IdentityStoreProvisioner | Configuration for the identity store provisioner instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityStoreProvisionersAPI.CreateIdentityStoreProvisioner(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityStoreProvisionersAPI.CreateIdentityStoreProvisioner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdentityStoreProvisioner`: IdentityStoreProvisioner
    fmt.Fprintf(os.Stdout, "Response from `IdentityStoreProvisionersAPI.CreateIdentityStoreProvisioner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityStoreProvisionerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IdentityStoreProvisioner**](IdentityStoreProvisioner.md) | Configuration for the identity store provisioner instance. | 

### Return type

[**IdentityStoreProvisioner**](IdentityStoreProvisioner.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityStoreProvisioner

> DeleteIdentityStoreProvisioner(ctx, id).Execute()

Delete an identity store provisioner instance

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
    id := "id_example" // string | ID of identity store provisioner instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdentityStoreProvisionersAPI.DeleteIdentityStoreProvisioner(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityStoreProvisionersAPI.DeleteIdentityStoreProvisioner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of identity store provisioner instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityStoreProvisionerRequest struct via the builder pattern


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


## GetIdentityStoreProvisioner

> IdentityStoreProvisioner GetIdentityStoreProvisioner(ctx, id).Execute()

Get an identity store provisioner by ID.

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
    id := "id_example" // string | ID of identity store provisioner instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityStoreProvisionersAPI.GetIdentityStoreProvisioner(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityStoreProvisionersAPI.GetIdentityStoreProvisioner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityStoreProvisioner`: IdentityStoreProvisioner
    fmt.Fprintf(os.Stdout, "Response from `IdentityStoreProvisionersAPI.GetIdentityStoreProvisioner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of identity store provisioner instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityStoreProvisionerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityStoreProvisioner**](IdentityStoreProvisioner.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityStoreProvisionerDescriptorById

> IdentityStoreProvisionerDescriptor GetIdentityStoreProvisionerDescriptorById(ctx, id).Execute()

Get the descriptor of an identity store provisioner by ID.

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
    id := "id_example" // string | ID of identity store provisioner descriptor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityStoreProvisionersAPI.GetIdentityStoreProvisionerDescriptorById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityStoreProvisionersAPI.GetIdentityStoreProvisionerDescriptorById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityStoreProvisionerDescriptorById`: IdentityStoreProvisionerDescriptor
    fmt.Fprintf(os.Stdout, "Response from `IdentityStoreProvisionersAPI.GetIdentityStoreProvisionerDescriptorById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of identity store provisioner descriptor | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityStoreProvisionerDescriptorByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityStoreProvisionerDescriptor**](IdentityStoreProvisionerDescriptor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityStoreProvisionerDescriptors

> IdentityStoreProvisionerDescriptors GetIdentityStoreProvisionerDescriptors(ctx).Execute()

Get the list of available identity store provisioner descriptors.

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
    resp, r, err := apiClient.IdentityStoreProvisionersAPI.GetIdentityStoreProvisionerDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityStoreProvisionersAPI.GetIdentityStoreProvisionerDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityStoreProvisionerDescriptors`: IdentityStoreProvisionerDescriptors
    fmt.Fprintf(os.Stdout, "Response from `IdentityStoreProvisionersAPI.GetIdentityStoreProvisionerDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityStoreProvisionerDescriptorsRequest struct via the builder pattern


### Return type

[**IdentityStoreProvisionerDescriptors**](IdentityStoreProvisionerDescriptors.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityStoreProvisioners

> IdentityStoreProvisioners GetIdentityStoreProvisioners(ctx).Execute()

Get the list of configured identity store provisioner instances.

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
    resp, r, err := apiClient.IdentityStoreProvisionersAPI.GetIdentityStoreProvisioners(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityStoreProvisionersAPI.GetIdentityStoreProvisioners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityStoreProvisioners`: IdentityStoreProvisioners
    fmt.Fprintf(os.Stdout, "Response from `IdentityStoreProvisionersAPI.GetIdentityStoreProvisioners`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityStoreProvisionersRequest struct via the builder pattern


### Return type

[**IdentityStoreProvisioners**](IdentityStoreProvisioners.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentityStoreProvisioner

> IdentityStoreProvisioner UpdateIdentityStoreProvisioner(ctx, id).Body(body).Execute()

Update an identity store provisioner instance

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
    id := "id_example" // string | ID of identity store provisioner instance
    body := *openapiclient.NewIdentityStoreProvisioner("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // IdentityStoreProvisioner | Configuration for the identity store provisioner instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityStoreProvisionersAPI.UpdateIdentityStoreProvisioner(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityStoreProvisionersAPI.UpdateIdentityStoreProvisioner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdentityStoreProvisioner`: IdentityStoreProvisioner
    fmt.Fprintf(os.Stdout, "Response from `IdentityStoreProvisionersAPI.UpdateIdentityStoreProvisioner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of identity store provisioner instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentityStoreProvisionerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IdentityStoreProvisioner**](IdentityStoreProvisioner.md) | Configuration for the identity store provisioner instance | 

### Return type

[**IdentityStoreProvisioner**](IdentityStoreProvisioner.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

