# \OauthAuthServerSettingsApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCommonScope**](OauthAuthServerSettingsApi.md#AddCommonScope) | **Post** /oauth/authServerSettings/scopes/commonScopes | Add a new common scope.
[**AddCommonScopeGroup**](OauthAuthServerSettingsApi.md#AddCommonScopeGroup) | **Post** /oauth/authServerSettings/scopes/commonScopeGroups | Create a new common scope group.
[**AddExclusiveScope**](OauthAuthServerSettingsApi.md#AddExclusiveScope) | **Post** /oauth/authServerSettings/scopes/exclusiveScopes | Add a new exclusive scope.
[**AddExclusiveScopeGroup**](OauthAuthServerSettingsApi.md#AddExclusiveScopeGroup) | **Post** /oauth/authServerSettings/scopes/exclusiveScopeGroups | Create a new exclusive scope group.
[**GetAuthorizationServerSettings**](OauthAuthServerSettingsApi.md#GetAuthorizationServerSettings) | **Get** /oauth/authServerSettings | Get the Authorization Server Settings.
[**GetCommonScope**](OauthAuthServerSettingsApi.md#GetCommonScope) | **Get** /oauth/authServerSettings/scopes/commonScopes/{name} | Get an existing common scope.
[**GetCommonScopeGroup**](OauthAuthServerSettingsApi.md#GetCommonScopeGroup) | **Get** /oauth/authServerSettings/scopes/commonScopeGroups/{name} | Get an existing common scope group.
[**GetExclusiveScope**](OauthAuthServerSettingsApi.md#GetExclusiveScope) | **Get** /oauth/authServerSettings/scopes/exclusiveScopes/{name} | Get an existing exclusive scope.
[**GetExclusiveScopeGroup**](OauthAuthServerSettingsApi.md#GetExclusiveScopeGroup) | **Get** /oauth/authServerSettings/scopes/exclusiveScopeGroups/{name} | Get an existing exclusive scope group.
[**RemoveCommonScope**](OauthAuthServerSettingsApi.md#RemoveCommonScope) | **Delete** /oauth/authServerSettings/scopes/commonScopes/{name} | Remove an existing common scope.
[**RemoveCommonScopeGroup**](OauthAuthServerSettingsApi.md#RemoveCommonScopeGroup) | **Delete** /oauth/authServerSettings/scopes/commonScopeGroups/{name} | Remove an existing common scope group.
[**RemoveExclusiveScope**](OauthAuthServerSettingsApi.md#RemoveExclusiveScope) | **Delete** /oauth/authServerSettings/scopes/exclusiveScopes/{name} | Remove an existing exclusive scope.
[**RemoveExclusiveScopeGroup**](OauthAuthServerSettingsApi.md#RemoveExclusiveScopeGroup) | **Delete** /oauth/authServerSettings/scopes/exclusiveScopeGroups/{name} | Remove an existing exclusive scope group.
[**UpdateAuthorizationServerSettings**](OauthAuthServerSettingsApi.md#UpdateAuthorizationServerSettings) | **Put** /oauth/authServerSettings | Update the Authorization Server Settings.
[**UpdateCommonScope**](OauthAuthServerSettingsApi.md#UpdateCommonScope) | **Put** /oauth/authServerSettings/scopes/commonScopes/{name} | Update an existing common scope.
[**UpdateCommonScopeGroup**](OauthAuthServerSettingsApi.md#UpdateCommonScopeGroup) | **Put** /oauth/authServerSettings/scopes/commonScopeGroups/{name} | Update an existing common scope group.
[**UpdateExclusiveScope**](OauthAuthServerSettingsApi.md#UpdateExclusiveScope) | **Put** /oauth/authServerSettings/scopes/exclusiveScopes/{name} | Update an existing exclusive scope.
[**UpdateExclusiveScopeGroups**](OauthAuthServerSettingsApi.md#UpdateExclusiveScopeGroups) | **Put** /oauth/authServerSettings/scopes/exclusiveScopeGroups/{name} | Update an existing exclusive scope group.



## AddCommonScope

> ScopeEntry AddCommonScope(ctx).Body(body).Execute()

Add a new common scope.

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
    body := *openapiclient.NewScopeEntry("Name_example", "Description_example") // ScopeEntry | The scope definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsApi.AddCommonScope(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.AddCommonScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCommonScope`: ScopeEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsApi.AddCommonScope`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCommonScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ScopeEntry**](ScopeEntry.md) | The scope definition. | 

### Return type

[**ScopeEntry**](ScopeEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddCommonScopeGroup

> ScopeGroupEntry AddCommonScopeGroup(ctx).Body(body).Execute()

Create a new common scope group.

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
    body := *openapiclient.NewScopeGroupEntry("Name_example", "Description_example", []string{"Scopes_example"}) // ScopeGroupEntry | The scope group definition

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsApi.AddCommonScopeGroup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.AddCommonScopeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCommonScopeGroup`: ScopeGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsApi.AddCommonScopeGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCommonScopeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ScopeGroupEntry**](ScopeGroupEntry.md) | The scope group definition | 

### Return type

[**ScopeGroupEntry**](ScopeGroupEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddExclusiveScope

> ScopeEntry AddExclusiveScope(ctx).Body(body).Execute()

Add a new exclusive scope.

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
    body := *openapiclient.NewScopeEntry("Name_example", "Description_example") // ScopeEntry | A new exclusive scope

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsApi.AddExclusiveScope(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.AddExclusiveScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddExclusiveScope`: ScopeEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsApi.AddExclusiveScope`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddExclusiveScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ScopeEntry**](ScopeEntry.md) | A new exclusive scope | 

### Return type

[**ScopeEntry**](ScopeEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddExclusiveScopeGroup

> ScopeGroupEntry AddExclusiveScopeGroup(ctx).Body(body).Execute()

Create a new exclusive scope group.

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
    body := *openapiclient.NewScopeGroupEntry("Name_example", "Description_example", []string{"Scopes_example"}) // ScopeGroupEntry | The scope group definition

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsApi.AddExclusiveScopeGroup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.AddExclusiveScopeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddExclusiveScopeGroup`: ScopeGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsApi.AddExclusiveScopeGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddExclusiveScopeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ScopeGroupEntry**](ScopeGroupEntry.md) | The scope group definition | 

### Return type

[**ScopeGroupEntry**](ScopeGroupEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorizationServerSettings

> AuthorizationServerSettings GetAuthorizationServerSettings(ctx).Execute()

Get the Authorization Server Settings.

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
    resp, r, err := apiClient.OauthAuthServerSettingsApi.GetAuthorizationServerSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.GetAuthorizationServerSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationServerSettings`: AuthorizationServerSettings
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsApi.GetAuthorizationServerSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationServerSettingsRequest struct via the builder pattern


### Return type

[**AuthorizationServerSettings**](AuthorizationServerSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommonScope

> ScopeEntry GetCommonScope(ctx, name).Execute()

Get an existing common scope.

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
    name := "name_example" // string | Name of the common scope.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsApi.GetCommonScope(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.GetCommonScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommonScope`: ScopeEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsApi.GetCommonScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the common scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommonScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScopeEntry**](ScopeEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommonScopeGroup

> ScopeGroupEntry GetCommonScopeGroup(ctx, name).Execute()

Get an existing common scope group.

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
    name := "name_example" // string | Name of the common scope group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsApi.GetCommonScopeGroup(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.GetCommonScopeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommonScopeGroup`: ScopeGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsApi.GetCommonScopeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the common scope group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommonScopeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScopeGroupEntry**](ScopeGroupEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExclusiveScope

> ScopeEntry GetExclusiveScope(ctx, name).Execute()

Get an existing exclusive scope.

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
    name := "name_example" // string | Name of the exclusive scope.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsApi.GetExclusiveScope(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.GetExclusiveScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExclusiveScope`: ScopeEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsApi.GetExclusiveScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exclusive scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExclusiveScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScopeEntry**](ScopeEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExclusiveScopeGroup

> ScopeGroupEntry GetExclusiveScopeGroup(ctx, name).Execute()

Get an existing exclusive scope group.

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
    name := "name_example" // string | Name of the exclusive scope group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsApi.GetExclusiveScopeGroup(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.GetExclusiveScopeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExclusiveScopeGroup`: ScopeGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsApi.GetExclusiveScopeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exclusive scope group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExclusiveScopeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScopeGroupEntry**](ScopeGroupEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCommonScope

> RemoveCommonScope(ctx, name).Execute()

Remove an existing common scope.

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
    name := "name_example" // string | Name of the common scope.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthAuthServerSettingsApi.RemoveCommonScope(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.RemoveCommonScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the common scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCommonScopeRequest struct via the builder pattern


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


## RemoveCommonScopeGroup

> RemoveCommonScopeGroup(ctx, name).Execute()

Remove an existing common scope group.

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
    name := "name_example" // string | Name of the common scope group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthAuthServerSettingsApi.RemoveCommonScopeGroup(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.RemoveCommonScopeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the common scope group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCommonScopeGroupRequest struct via the builder pattern


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


## RemoveExclusiveScope

> RemoveExclusiveScope(ctx, name).Execute()

Remove an existing exclusive scope.

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
    name := "name_example" // string | Name of the exclusive scope.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthAuthServerSettingsApi.RemoveExclusiveScope(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.RemoveExclusiveScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exclusive scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveExclusiveScopeRequest struct via the builder pattern


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


## RemoveExclusiveScopeGroup

> RemoveExclusiveScopeGroup(ctx, name).Execute()

Remove an existing exclusive scope group.

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
    name := "name_example" // string | Name of the exclusive scope group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthAuthServerSettingsApi.RemoveExclusiveScopeGroup(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.RemoveExclusiveScopeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exclusive scope group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveExclusiveScopeGroupRequest struct via the builder pattern


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


## UpdateAuthorizationServerSettings

> AuthorizationServerSettings UpdateAuthorizationServerSettings(ctx).Body(body).Execute()

Update the Authorization Server Settings.

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
    body := *openapiclient.NewAuthorizationServerSettings("DefaultScopeDescription_example", int64(123), int64(123), int64(123), int64(123), "RegisteredAuthorizationPath_example", int64(123), int64(123), false) // AuthorizationServerSettings | Configuration for updated server settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsApi.UpdateAuthorizationServerSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.UpdateAuthorizationServerSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthorizationServerSettings`: AuthorizationServerSettings
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsApi.UpdateAuthorizationServerSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthorizationServerSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthorizationServerSettings**](AuthorizationServerSettings.md) | Configuration for updated server settings. | 

### Return type

[**AuthorizationServerSettings**](AuthorizationServerSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCommonScope

> ScopeEntry UpdateCommonScope(ctx, name).Body(body).Execute()

Update an existing common scope.

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
    name := "name_example" // string | Name of the common scope.
    body := *openapiclient.NewScopeEntry("Name_example", "Description_example") // ScopeEntry | The scope definition

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsApi.UpdateCommonScope(context.Background(), name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.UpdateCommonScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCommonScope`: ScopeEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsApi.UpdateCommonScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the common scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommonScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ScopeEntry**](ScopeEntry.md) | The scope definition | 

### Return type

[**ScopeEntry**](ScopeEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCommonScopeGroup

> ScopeGroupEntry UpdateCommonScopeGroup(ctx, name).Body(body).Execute()

Update an existing common scope group.

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
    name := "name_example" // string | Name of the common scope group.
    body := *openapiclient.NewScopeGroupEntry("Name_example", "Description_example", []string{"Scopes_example"}) // ScopeGroupEntry | The scope group definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsApi.UpdateCommonScopeGroup(context.Background(), name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.UpdateCommonScopeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCommonScopeGroup`: ScopeGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsApi.UpdateCommonScopeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the common scope group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommonScopeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ScopeGroupEntry**](ScopeGroupEntry.md) | The scope group definition. | 

### Return type

[**ScopeGroupEntry**](ScopeGroupEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExclusiveScope

> ScopeEntry UpdateExclusiveScope(ctx, name).Body(body).Execute()

Update an existing exclusive scope.

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
    name := "name_example" // string | Name of the exclusive scope.
    body := *openapiclient.NewScopeEntry("Name_example", "Description_example") // ScopeEntry | The scope definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsApi.UpdateExclusiveScope(context.Background(), name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.UpdateExclusiveScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExclusiveScope`: ScopeEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsApi.UpdateExclusiveScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exclusive scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExclusiveScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ScopeEntry**](ScopeEntry.md) | The scope definition. | 

### Return type

[**ScopeEntry**](ScopeEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExclusiveScopeGroups

> ScopeGroupEntry UpdateExclusiveScopeGroups(ctx, name).Body(body).Execute()

Update an existing exclusive scope group.

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
    name := "name_example" // string | Name of the exclusive scope group.
    body := *openapiclient.NewScopeGroupEntry("Name_example", "Description_example", []string{"Scopes_example"}) // ScopeGroupEntry | The scope group definition

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsApi.UpdateExclusiveScopeGroups(context.Background(), name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsApi.UpdateExclusiveScopeGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExclusiveScopeGroups`: ScopeGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsApi.UpdateExclusiveScopeGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the exclusive scope group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExclusiveScopeGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ScopeGroupEntry**](ScopeGroupEntry.md) | The scope group definition | 

### Return type

[**ScopeGroupEntry**](ScopeGroupEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

