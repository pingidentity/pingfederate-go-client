# \SessionAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSourcePolicy**](SessionAPI.md#CreateSourcePolicy) | **Post** /session/authenticationSessionPolicies | Create a new session policy.
[**DeleteSourcePolicy**](SessionAPI.md#DeleteSourcePolicy) | **Delete** /session/authenticationSessionPolicies/{id} | Delete a session policy.
[**GetApplicationPolicy**](SessionAPI.md#GetApplicationPolicy) | **Get** /session/applicationSessionPolicy | Get the application session policy.
[**GetGlobalPolicy**](SessionAPI.md#GetGlobalPolicy) | **Get** /session/authenticationSessionPolicies/global | Get the global authentication session policy.
[**GetSessionSettings**](SessionAPI.md#GetSessionSettings) | **Get** /session/settings | Get general session management settings.
[**GetSourcePolicies**](SessionAPI.md#GetSourcePolicies) | **Get** /session/authenticationSessionPolicies | Get list of session policies.
[**GetSourcePolicy**](SessionAPI.md#GetSourcePolicy) | **Get** /session/authenticationSessionPolicies/{id} | Find session policy by ID.
[**UpdateApplicationPolicy**](SessionAPI.md#UpdateApplicationPolicy) | **Put** /session/applicationSessionPolicy | Update the application session policy.
[**UpdateGlobalPolicy**](SessionAPI.md#UpdateGlobalPolicy) | **Put** /session/authenticationSessionPolicies/global | Update the global authentication session policy.
[**UpdateSessionSettings**](SessionAPI.md#UpdateSessionSettings) | **Put** /session/settings | Update general session management settings.
[**UpdateSourcePolicy**](SessionAPI.md#UpdateSourcePolicy) | **Put** /session/authenticationSessionPolicies/{id} | Update a session policy.



## CreateSourcePolicy

> AuthenticationSessionPolicy CreateSourcePolicy(ctx).Body(body).Execute()

Create a new session policy.



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
    body := *openapiclient.NewAuthenticationSessionPolicy(*openapiclient.NewAuthenticationSource("Type_example", *openapiclient.NewResourceLink("Id_example")), false) // AuthenticationSessionPolicy | Configuration for new policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.CreateSourcePolicy(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.CreateSourcePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSourcePolicy`: AuthenticationSessionPolicy
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.CreateSourcePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourcePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthenticationSessionPolicy**](AuthenticationSessionPolicy.md) | Configuration for new policy. | 

### Return type

[**AuthenticationSessionPolicy**](AuthenticationSessionPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSourcePolicy

> DeleteSourcePolicy(ctx, id).Execute()

Delete a session policy.



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
    id := "id_example" // string | ID of session policy to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SessionAPI.DeleteSourcePolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.DeleteSourcePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of session policy to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSourcePolicyRequest struct via the builder pattern


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


## GetApplicationPolicy

> ApplicationSessionPolicy GetApplicationPolicy(ctx).Execute()

Get the application session policy.

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
    resp, r, err := apiClient.SessionAPI.GetApplicationPolicy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.GetApplicationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationPolicy`: ApplicationSessionPolicy
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.GetApplicationPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationPolicyRequest struct via the builder pattern


### Return type

[**ApplicationSessionPolicy**](ApplicationSessionPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalPolicy

> GlobalAuthenticationSessionPolicy GetGlobalPolicy(ctx).Execute()

Get the global authentication session policy.

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
    resp, r, err := apiClient.SessionAPI.GetGlobalPolicy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.GetGlobalPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalPolicy`: GlobalAuthenticationSessionPolicy
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.GetGlobalPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalPolicyRequest struct via the builder pattern


### Return type

[**GlobalAuthenticationSessionPolicy**](GlobalAuthenticationSessionPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionSettings

> SessionSettings GetSessionSettings(ctx).Execute()

Get general session management settings.

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
    resp, r, err := apiClient.SessionAPI.GetSessionSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.GetSessionSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionSettings`: SessionSettings
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.GetSessionSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionSettingsRequest struct via the builder pattern


### Return type

[**SessionSettings**](SessionSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourcePolicies

> AuthenticationSessionPolicies GetSourcePolicies(ctx).Execute()

Get list of session policies.



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
    resp, r, err := apiClient.SessionAPI.GetSourcePolicies(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.GetSourcePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourcePolicies`: AuthenticationSessionPolicies
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.GetSourcePolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourcePoliciesRequest struct via the builder pattern


### Return type

[**AuthenticationSessionPolicies**](AuthenticationSessionPolicies.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourcePolicy

> AuthenticationSessionPolicy GetSourcePolicy(ctx, id).Execute()

Find session policy by ID.



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
    id := "id_example" // string | ID of the session policy to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.GetSourcePolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.GetSourcePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourcePolicy`: AuthenticationSessionPolicy
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.GetSourcePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the session policy to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourcePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticationSessionPolicy**](AuthenticationSessionPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationPolicy

> ApplicationSessionPolicy UpdateApplicationPolicy(ctx).Body(body).Execute()

Update the application session policy.

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
    body := *openapiclient.NewApplicationSessionPolicy() // ApplicationSessionPolicy | Application session policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.UpdateApplicationPolicy(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.UpdateApplicationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationPolicy`: ApplicationSessionPolicy
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.UpdateApplicationPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationSessionPolicy**](ApplicationSessionPolicy.md) | Application session policy. | 

### Return type

[**ApplicationSessionPolicy**](ApplicationSessionPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGlobalPolicy

> GlobalAuthenticationSessionPolicy UpdateGlobalPolicy(ctx).Body(body).Execute()

Update the global authentication session policy.

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
    body := *openapiclient.NewGlobalAuthenticationSessionPolicy(false) // GlobalAuthenticationSessionPolicy | Global authentication session policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.UpdateGlobalPolicy(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.UpdateGlobalPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGlobalPolicy`: GlobalAuthenticationSessionPolicy
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.UpdateGlobalPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GlobalAuthenticationSessionPolicy**](GlobalAuthenticationSessionPolicy.md) | Global authentication session policy. | 

### Return type

[**GlobalAuthenticationSessionPolicy**](GlobalAuthenticationSessionPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSessionSettings

> SessionSettings UpdateSessionSettings(ctx).Body(body).Execute()

Update general session management settings.

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
    body := *openapiclient.NewSessionSettings() // SessionSettings | Session settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.UpdateSessionSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.UpdateSessionSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSessionSettings`: SessionSettings
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.UpdateSessionSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSessionSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SessionSettings**](SessionSettings.md) | Session settings. | 

### Return type

[**SessionSettings**](SessionSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSourcePolicy

> AuthenticationSessionPolicy UpdateSourcePolicy(ctx, id).Body(body).Execute()

Update a session policy.



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
    id := "id_example" // string | ID of the session policy to update.
    body := *openapiclient.NewAuthenticationSessionPolicy(*openapiclient.NewAuthenticationSource("Type_example", *openapiclient.NewResourceLink("Id_example")), false) // AuthenticationSessionPolicy | Configuration for updated policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.UpdateSourcePolicy(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.UpdateSourcePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSourcePolicy`: AuthenticationSessionPolicy
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.UpdateSourcePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the session policy to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourcePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthenticationSessionPolicy**](AuthenticationSessionPolicy.md) | Configuration for updated policy. | 

### Return type

[**AuthenticationSessionPolicy**](AuthenticationSessionPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

