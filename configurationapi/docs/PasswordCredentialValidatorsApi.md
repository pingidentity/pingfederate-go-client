# \PasswordCredentialValidatorsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePasswordCredentialValidator**](PasswordCredentialValidatorsAPI.md#CreatePasswordCredentialValidator) | **Post** /passwordCredentialValidators | Create a new password credential validator instance
[**DeletePasswordCredentialValidator**](PasswordCredentialValidatorsAPI.md#DeletePasswordCredentialValidator) | **Delete** /passwordCredentialValidators/{id} | Delete a password credential validator instance.
[**GetPasswordCredentialValidator**](PasswordCredentialValidatorsAPI.md#GetPasswordCredentialValidator) | **Get** /passwordCredentialValidators/{id} | Find a password credential validator by ID.
[**GetPasswordCredentialValidatorDescriptor**](PasswordCredentialValidatorsAPI.md#GetPasswordCredentialValidatorDescriptor) | **Get** /passwordCredentialValidators/descriptors/{id} | Get the description of a password credential validator by ID.
[**GetPasswordCredentialValidatorDescriptors**](PasswordCredentialValidatorsAPI.md#GetPasswordCredentialValidatorDescriptors) | **Get** /passwordCredentialValidators/descriptors | Get a list of available password credential validator descriptors.
[**GetPasswordCredentialValidators**](PasswordCredentialValidatorsAPI.md#GetPasswordCredentialValidators) | **Get** /passwordCredentialValidators | Get the list of available password credential validators
[**UpdatePasswordCredentialValidator**](PasswordCredentialValidatorsAPI.md#UpdatePasswordCredentialValidator) | **Put** /passwordCredentialValidators/{id} | Update a password credential validator instance.



## CreatePasswordCredentialValidator

> PasswordCredentialValidator CreatePasswordCredentialValidator(ctx).Body(body).Execute()

Create a new password credential validator instance



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
    body := *openapiclient.NewPasswordCredentialValidator("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // PasswordCredentialValidator | Configuration for the new password credential validator instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordCredentialValidatorsAPI.CreatePasswordCredentialValidator(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordCredentialValidatorsAPI.CreatePasswordCredentialValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePasswordCredentialValidator`: PasswordCredentialValidator
    fmt.Fprintf(os.Stdout, "Response from `PasswordCredentialValidatorsAPI.CreatePasswordCredentialValidator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePasswordCredentialValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PasswordCredentialValidator**](PasswordCredentialValidator.md) | Configuration for the new password credential validator instance. | 

### Return type

[**PasswordCredentialValidator**](PasswordCredentialValidator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePasswordCredentialValidator

> DeletePasswordCredentialValidator(ctx, id).Execute()

Delete a password credential validator instance.



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
    id := "id_example" // string | ID of the password credential validator to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PasswordCredentialValidatorsAPI.DeletePasswordCredentialValidator(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordCredentialValidatorsAPI.DeletePasswordCredentialValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the password credential validator to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePasswordCredentialValidatorRequest struct via the builder pattern


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


## GetPasswordCredentialValidator

> PasswordCredentialValidator GetPasswordCredentialValidator(ctx, id).Execute()

Find a password credential validator by ID.



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
    id := "id_example" // string | ID of password credential validator instance to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordCredentialValidatorsAPI.GetPasswordCredentialValidator(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordCredentialValidatorsAPI.GetPasswordCredentialValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordCredentialValidator`: PasswordCredentialValidator
    fmt.Fprintf(os.Stdout, "Response from `PasswordCredentialValidatorsAPI.GetPasswordCredentialValidator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of password credential validator instance to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordCredentialValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PasswordCredentialValidator**](PasswordCredentialValidator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPasswordCredentialValidatorDescriptor

> PasswordCredentialValidatorDescriptor GetPasswordCredentialValidatorDescriptor(ctx, id).Execute()

Get the description of a password credential validator by ID.



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
    id := "id_example" // string | ID of the password credential validator descriptor to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordCredentialValidatorsAPI.GetPasswordCredentialValidatorDescriptor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordCredentialValidatorsAPI.GetPasswordCredentialValidatorDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordCredentialValidatorDescriptor`: PasswordCredentialValidatorDescriptor
    fmt.Fprintf(os.Stdout, "Response from `PasswordCredentialValidatorsAPI.GetPasswordCredentialValidatorDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the password credential validator descriptor to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordCredentialValidatorDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PasswordCredentialValidatorDescriptor**](PasswordCredentialValidatorDescriptor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPasswordCredentialValidatorDescriptors

> PasswordCredentialValidatorDescriptors GetPasswordCredentialValidatorDescriptors(ctx).Execute()

Get a list of available password credential validator descriptors.

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
    resp, r, err := apiClient.PasswordCredentialValidatorsAPI.GetPasswordCredentialValidatorDescriptors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordCredentialValidatorsAPI.GetPasswordCredentialValidatorDescriptors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordCredentialValidatorDescriptors`: PasswordCredentialValidatorDescriptors
    fmt.Fprintf(os.Stdout, "Response from `PasswordCredentialValidatorsAPI.GetPasswordCredentialValidatorDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordCredentialValidatorDescriptorsRequest struct via the builder pattern


### Return type

[**PasswordCredentialValidatorDescriptors**](PasswordCredentialValidatorDescriptors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPasswordCredentialValidators

> PasswordCredentialValidators GetPasswordCredentialValidators(ctx).Execute()

Get the list of available password credential validators

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
    resp, r, err := apiClient.PasswordCredentialValidatorsAPI.GetPasswordCredentialValidators(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordCredentialValidatorsAPI.GetPasswordCredentialValidators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordCredentialValidators`: PasswordCredentialValidators
    fmt.Fprintf(os.Stdout, "Response from `PasswordCredentialValidatorsAPI.GetPasswordCredentialValidators`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordCredentialValidatorsRequest struct via the builder pattern


### Return type

[**PasswordCredentialValidators**](PasswordCredentialValidators.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePasswordCredentialValidator

> PasswordCredentialValidator UpdatePasswordCredentialValidator(ctx, id).Body(body).Execute()

Update a password credential validator instance.



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
    id := "id_example" // string | ID of the password credential validator instance
    body := *openapiclient.NewPasswordCredentialValidator("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // PasswordCredentialValidator | Configuration for the updated password credential validator instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordCredentialValidatorsAPI.UpdatePasswordCredentialValidator(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordCredentialValidatorsAPI.UpdatePasswordCredentialValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePasswordCredentialValidator`: PasswordCredentialValidator
    fmt.Fprintf(os.Stdout, "Response from `PasswordCredentialValidatorsAPI.UpdatePasswordCredentialValidator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the password credential validator instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordCredentialValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PasswordCredentialValidator**](PasswordCredentialValidator.md) | Configuration for the updated password credential validator instance. | 

### Return type

[**PasswordCredentialValidator**](PasswordCredentialValidator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

