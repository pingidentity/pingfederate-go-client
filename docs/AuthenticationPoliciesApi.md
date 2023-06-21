# \AuthenticationPoliciesApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFragment**](AuthenticationPoliciesApi.md#CreateFragment) | **Post** /authenticationPolicies/fragments | Create an authentication policy fragment.
[**CreatePolicy**](AuthenticationPoliciesApi.md#CreatePolicy) | **Post** /authenticationPolicies/policy | Create a new authentication policy.
[**DeleteFragment**](AuthenticationPoliciesApi.md#DeleteFragment) | **Delete** /authenticationPolicies/fragments/{id} | Delete an authentication policy fragment.
[**DeletePolicy**](AuthenticationPoliciesApi.md#DeletePolicy) | **Delete** /authenticationPolicies/policy/{id} | Delete an authentication policy.
[**GetAuthenticationPolicySettings**](AuthenticationPoliciesApi.md#GetAuthenticationPolicySettings) | **Get** /authenticationPolicies/settings | Get the authentication policies settings.
[**GetDefaultAuthenticationPolicy**](AuthenticationPoliciesApi.md#GetDefaultAuthenticationPolicy) | **Get** /authenticationPolicies/default | Get the default configured authentication policy.
[**GetFragment**](AuthenticationPoliciesApi.md#GetFragment) | **Get** /authenticationPolicies/fragments/{id} | Get an authentication policy fragment by ID.
[**GetFragments**](AuthenticationPoliciesApi.md#GetFragments) | **Get** /authenticationPolicies/fragments | Get all of the authentication policies fragments.
[**GetPolicy**](AuthenticationPoliciesApi.md#GetPolicy) | **Get** /authenticationPolicies/policy/{id} | Get an authentication policy by ID.
[**MovePolicy**](AuthenticationPoliciesApi.md#MovePolicy) | **Post** /authenticationPolicies/policy/{id}/move | Move an authentication policy to a location within the policy tree.
[**UpdateAuthenticationPolicySettings**](AuthenticationPoliciesApi.md#UpdateAuthenticationPolicySettings) | **Put** /authenticationPolicies/settings | Set the authentication policies settings.
[**UpdateDefaultAuthenticationPolicy**](AuthenticationPoliciesApi.md#UpdateDefaultAuthenticationPolicy) | **Put** /authenticationPolicies/default | Set the default authentication policy.
[**UpdateFragment**](AuthenticationPoliciesApi.md#UpdateFragment) | **Put** /authenticationPolicies/fragments/{id} | Update an authentication policy fragment.
[**UpdatePolicy**](AuthenticationPoliciesApi.md#UpdatePolicy) | **Put** /authenticationPolicies/policy/{id} | Update an authentication policy.



## CreateFragment

> AuthenticationPolicyFragment CreateFragment(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create an authentication policy fragment.

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
    body := *openapiclient.NewAuthenticationPolicyFragment() // AuthenticationPolicyFragment | Configuration of the authentication policy fragment.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationPoliciesApi.CreateFragment(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPoliciesApi.CreateFragment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFragment`: AuthenticationPolicyFragment
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationPoliciesApi.CreateFragment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFragmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthenticationPolicyFragment**](AuthenticationPolicyFragment.md) | Configuration of the authentication policy fragment. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**AuthenticationPolicyFragment**](AuthenticationPolicyFragment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicy

> AuthenticationPolicyTree CreatePolicy(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new authentication policy.

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
    body := *openapiclient.NewAuthenticationPolicyTree() // AuthenticationPolicyTree | Configuration of the authentication policy.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationPoliciesApi.CreatePolicy(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPoliciesApi.CreatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicy`: AuthenticationPolicyTree
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationPoliciesApi.CreatePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthenticationPolicyTree**](AuthenticationPolicyTree.md) | Configuration of the authentication policy. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**AuthenticationPolicyTree**](AuthenticationPolicyTree.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFragment

> DeleteFragment(ctx, id).Execute()

Delete an authentication policy fragment.

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
    id := "id_example" // string | ID of the policy fragment to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthenticationPoliciesApi.DeleteFragment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPoliciesApi.DeleteFragment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the policy fragment to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFragmentRequest struct via the builder pattern


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


## DeletePolicy

> DeletePolicy(ctx, id).Execute()

Delete an authentication policy.

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
    id := "id_example" // string | Authentication policy Id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthenticationPoliciesApi.DeletePolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPoliciesApi.DeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Authentication policy Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


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


## GetAuthenticationPolicySettings

> AuthenticationPoliciesSettings GetAuthenticationPolicySettings(ctx).Execute()

Get the authentication policies settings.

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
    resp, r, err := apiClient.AuthenticationPoliciesApi.GetAuthenticationPolicySettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPoliciesApi.GetAuthenticationPolicySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticationPolicySettings`: AuthenticationPoliciesSettings
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationPoliciesApi.GetAuthenticationPolicySettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationPolicySettingsRequest struct via the builder pattern


### Return type

[**AuthenticationPoliciesSettings**](AuthenticationPoliciesSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultAuthenticationPolicy

> AuthenticationPolicy GetDefaultAuthenticationPolicy(ctx).Execute()

Get the default configured authentication policy.

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
    resp, r, err := apiClient.AuthenticationPoliciesApi.GetDefaultAuthenticationPolicy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPoliciesApi.GetDefaultAuthenticationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultAuthenticationPolicy`: AuthenticationPolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationPoliciesApi.GetDefaultAuthenticationPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultAuthenticationPolicyRequest struct via the builder pattern


### Return type

[**AuthenticationPolicy**](AuthenticationPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFragment

> AuthenticationPolicyFragment GetFragment(ctx, id).Execute()

Get an authentication policy fragment by ID.

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
    id := "id_example" // string | ID of the policy fragment to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationPoliciesApi.GetFragment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPoliciesApi.GetFragment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFragment`: AuthenticationPolicyFragment
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationPoliciesApi.GetFragment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the policy fragment to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFragmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticationPolicyFragment**](AuthenticationPolicyFragment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFragments

> AuthenticationPolicyFragments GetFragments(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()

Get all of the authentication policies fragments.

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
    page := int64(56) // int64 | Page number to retrieve. (optional)
    numberPerPage := int64(56) // int64 | Number of fragments per page. (optional)
    filter := "filter_example" // string | Filter criteria limits the fragments that are returned to only those that match it. The filter criteria is compared to the fragment instance name and ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationPoliciesApi.GetFragments(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPoliciesApi.GetFragments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFragments`: AuthenticationPolicyFragments
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationPoliciesApi.GetFragments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFragmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve. | 
 **numberPerPage** | **int64** | Number of fragments per page. | 
 **filter** | **string** | Filter criteria limits the fragments that are returned to only those that match it. The filter criteria is compared to the fragment instance name and ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. | 

### Return type

[**AuthenticationPolicyFragments**](AuthenticationPolicyFragments.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicy

> AuthenticationPolicyTree GetPolicy(ctx, id).Execute()

Get an authentication policy by ID.

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
    id := "id_example" // string | Authentication policy Id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationPoliciesApi.GetPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPoliciesApi.GetPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicy`: AuthenticationPolicyTree
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationPoliciesApi.GetPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Authentication policy Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticationPolicyTree**](AuthenticationPolicyTree.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MovePolicy

> MovePolicy(ctx, id).Body(body).Execute()

Move an authentication policy to a location within the policy tree.

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
    id := "id_example" // string | Authentication policy Id.
    body := *openapiclient.NewMoveItemRequest("Location_example") // MoveItemRequest | Metadata about where to move the policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthenticationPoliciesApi.MovePolicy(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPoliciesApi.MovePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Authentication policy Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMovePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MoveItemRequest**](MoveItemRequest.md) | Metadata about where to move the policy | 

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


## UpdateAuthenticationPolicySettings

> AuthenticationPoliciesSettings UpdateAuthenticationPolicySettings(ctx).Body(body).Execute()

Set the authentication policies settings.

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
    body := *openapiclient.NewAuthenticationPoliciesSettings() // AuthenticationPoliciesSettings | Authentication policies settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationPoliciesApi.UpdateAuthenticationPolicySettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPoliciesApi.UpdateAuthenticationPolicySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthenticationPolicySettings`: AuthenticationPoliciesSettings
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationPoliciesApi.UpdateAuthenticationPolicySettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthenticationPolicySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthenticationPoliciesSettings**](AuthenticationPoliciesSettings.md) | Authentication policies settings. | 

### Return type

[**AuthenticationPoliciesSettings**](AuthenticationPoliciesSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDefaultAuthenticationPolicy

> AuthenticationPolicy UpdateDefaultAuthenticationPolicy(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Set the default authentication policy.

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
    body := *openapiclient.NewAuthenticationPolicy() // AuthenticationPolicy | Default authentication policy.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationPoliciesApi.UpdateDefaultAuthenticationPolicy(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPoliciesApi.UpdateDefaultAuthenticationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDefaultAuthenticationPolicy`: AuthenticationPolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationPoliciesApi.UpdateDefaultAuthenticationPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDefaultAuthenticationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthenticationPolicy**](AuthenticationPolicy.md) | Default authentication policy. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**AuthenticationPolicy**](AuthenticationPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFragment

> AuthenticationPolicyFragment UpdateFragment(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update an authentication policy fragment.

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
    id := "id_example" // string | ID of the policy fragment to  update.
    body := *openapiclient.NewAuthenticationPolicyFragment() // AuthenticationPolicyFragment | Configuration of the authentication policy fragment.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationPoliciesApi.UpdateFragment(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPoliciesApi.UpdateFragment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFragment`: AuthenticationPolicyFragment
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationPoliciesApi.UpdateFragment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the policy fragment to  update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFragmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthenticationPolicyFragment**](AuthenticationPolicyFragment.md) | Configuration of the authentication policy fragment. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**AuthenticationPolicyFragment**](AuthenticationPolicyFragment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicy

> AuthenticationPolicyTree UpdatePolicy(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update an authentication policy.

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
    id := "id_example" // string | Authentication policy Id.
    body := *openapiclient.NewAuthenticationPolicyTree() // AuthenticationPolicyTree | Configuration of the authentication policy.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationPoliciesApi.UpdatePolicy(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPoliciesApi.UpdatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePolicy`: AuthenticationPolicyTree
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationPoliciesApi.UpdatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Authentication policy Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthenticationPolicyTree**](AuthenticationPolicyTree.md) | Configuration of the authentication policy. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**AuthenticationPolicyTree**](AuthenticationPolicyTree.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

