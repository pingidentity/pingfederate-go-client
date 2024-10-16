# \OauthAuthServerSettingsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCommonScope**](OauthAuthServerSettingsAPI.md#AddCommonScope) | **Post** /oauth/authServerSettings/scopes/commonScopes | Add a new common scope.
[**AddCommonScopeGroup**](OauthAuthServerSettingsAPI.md#AddCommonScopeGroup) | **Post** /oauth/authServerSettings/scopes/commonScopeGroups | Create a new common scope group.
[**AddExclusiveScope**](OauthAuthServerSettingsAPI.md#AddExclusiveScope) | **Post** /oauth/authServerSettings/scopes/exclusiveScopes | Add a new exclusive scope.
[**AddExclusiveScopeGroup**](OauthAuthServerSettingsAPI.md#AddExclusiveScopeGroup) | **Post** /oauth/authServerSettings/scopes/exclusiveScopeGroups | Create a new exclusive scope group.
[**GetAuthorizationServerSettings**](OauthAuthServerSettingsAPI.md#GetAuthorizationServerSettings) | **Get** /oauth/authServerSettings | Get the Authorization Server Settings.
[**GetCommonScope**](OauthAuthServerSettingsAPI.md#GetCommonScope) | **Get** /oauth/authServerSettings/scopes/commonScopes/{name} | Get an existing common scope.
[**GetCommonScopeGroup**](OauthAuthServerSettingsAPI.md#GetCommonScopeGroup) | **Get** /oauth/authServerSettings/scopes/commonScopeGroups/{name} | Get an existing common scope group.
[**GetCommonScopeGroups**](OauthAuthServerSettingsAPI.md#GetCommonScopeGroups) | **Get** /oauth/authServerSettings/scopes/commonScopeGroups | Get the common scope groups.
[**GetCommonScopes**](OauthAuthServerSettingsAPI.md#GetCommonScopes) | **Get** /oauth/authServerSettings/scopes/commonScopes | Get the common scopes.
[**GetExclusiveScope**](OauthAuthServerSettingsAPI.md#GetExclusiveScope) | **Get** /oauth/authServerSettings/scopes/exclusiveScopes/{name} | Get an existing exclusive scope.
[**GetExclusiveScopeGroup**](OauthAuthServerSettingsAPI.md#GetExclusiveScopeGroup) | **Get** /oauth/authServerSettings/scopes/exclusiveScopeGroups/{name} | Get an existing exclusive scope group.
[**GetExclusiveScopeGroups**](OauthAuthServerSettingsAPI.md#GetExclusiveScopeGroups) | **Get** /oauth/authServerSettings/scopes/exclusiveScopeGroups | Get the exclusive scope groups.
[**GetExclusiveScopes**](OauthAuthServerSettingsAPI.md#GetExclusiveScopes) | **Get** /oauth/authServerSettings/scopes/exclusiveScopes | Get the exclusive scopes.
[**RemoveCommonScope**](OauthAuthServerSettingsAPI.md#RemoveCommonScope) | **Delete** /oauth/authServerSettings/scopes/commonScopes/{name} | Remove an existing common scope.
[**RemoveCommonScopeGroup**](OauthAuthServerSettingsAPI.md#RemoveCommonScopeGroup) | **Delete** /oauth/authServerSettings/scopes/commonScopeGroups/{name} | Remove an existing common scope group.
[**RemoveExclusiveScope**](OauthAuthServerSettingsAPI.md#RemoveExclusiveScope) | **Delete** /oauth/authServerSettings/scopes/exclusiveScopes/{name} | Remove an existing exclusive scope.
[**RemoveExclusiveScopeGroup**](OauthAuthServerSettingsAPI.md#RemoveExclusiveScopeGroup) | **Delete** /oauth/authServerSettings/scopes/exclusiveScopeGroups/{name} | Remove an existing exclusive scope group.
[**UpdateAuthorizationServerSettings**](OauthAuthServerSettingsAPI.md#UpdateAuthorizationServerSettings) | **Put** /oauth/authServerSettings | Update the Authorization Server Settings.
[**UpdateCommonScope**](OauthAuthServerSettingsAPI.md#UpdateCommonScope) | **Put** /oauth/authServerSettings/scopes/commonScopes/{name} | Update an existing common scope.
[**UpdateCommonScopeGroup**](OauthAuthServerSettingsAPI.md#UpdateCommonScopeGroup) | **Put** /oauth/authServerSettings/scopes/commonScopeGroups/{name} | Update an existing common scope group.
[**UpdateExclusiveScope**](OauthAuthServerSettingsAPI.md#UpdateExclusiveScope) | **Put** /oauth/authServerSettings/scopes/exclusiveScopes/{name} | Update an existing exclusive scope.
[**UpdateExclusiveScopeGroups**](OauthAuthServerSettingsAPI.md#UpdateExclusiveScopeGroups) | **Put** /oauth/authServerSettings/scopes/exclusiveScopeGroups/{name} | Update an existing exclusive scope group.



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
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.AddCommonScope(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.AddCommonScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCommonScope`: ScopeEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.AddCommonScope`: %v\n", resp)
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

[bearer](../README.md#bearer)

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
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.AddCommonScopeGroup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.AddCommonScopeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCommonScopeGroup`: ScopeGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.AddCommonScopeGroup`: %v\n", resp)
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

[bearer](../README.md#bearer)

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
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.AddExclusiveScope(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.AddExclusiveScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddExclusiveScope`: ScopeEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.AddExclusiveScope`: %v\n", resp)
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

[bearer](../README.md#bearer)

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
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.AddExclusiveScopeGroup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.AddExclusiveScopeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddExclusiveScopeGroup`: ScopeGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.AddExclusiveScopeGroup`: %v\n", resp)
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

[bearer](../README.md#bearer)

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
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.GetAuthorizationServerSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.GetAuthorizationServerSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationServerSettings`: AuthorizationServerSettings
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.GetAuthorizationServerSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationServerSettingsRequest struct via the builder pattern


### Return type

[**AuthorizationServerSettings**](AuthorizationServerSettings.md)

### Authorization

[bearer](../README.md#bearer)

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
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.GetCommonScope(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.GetCommonScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommonScope`: ScopeEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.GetCommonScope`: %v\n", resp)
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

[bearer](../README.md#bearer)

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
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.GetCommonScopeGroup(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.GetCommonScopeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommonScopeGroup`: ScopeGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.GetCommonScopeGroup`: %v\n", resp)
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommonScopeGroups

> ScopeGroupEntries GetCommonScopeGroups(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).SortOrder(sortOrder).Execute()

Get the common scope groups.



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
    numberPerPage := int64(56) // int64 | Number of common scope groups per page. (optional)
    filter := "filter_example" // string | Filter criteria limits the common scope groups that are returned to only those that match it. The filter criteria is compared to the scope group name. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. (optional)
    sortOrder := "sortOrder_example" // string | Sort order of scope group names to retrieve group by. By default, scope groups are sorted in ascending order by scope group name. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.GetCommonScopeGroups(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.GetCommonScopeGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommonScopeGroups`: ScopeGroupEntries
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.GetCommonScopeGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCommonScopeGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve. | 
 **numberPerPage** | **int64** | Number of common scope groups per page. | 
 **filter** | **string** | Filter criteria limits the common scope groups that are returned to only those that match it. The filter criteria is compared to the scope group name. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. | 
 **sortOrder** | **string** | Sort order of scope group names to retrieve group by. By default, scope groups are sorted in ascending order by scope group name. | 

### Return type

[**ScopeGroupEntries**](ScopeGroupEntries.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommonScopes

> ScopeEntries GetCommonScopes(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).SortOrder(sortOrder).IncludeDefaultOpenId(includeDefaultOpenId).Execute()

Get the common scopes.



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
    numberPerPage := int64(56) // int64 | Number of common scopes per page. (optional)
    filter := "filter_example" // string | Filter criteria limits the common scopes that are returned to only those that match it. The filter criteria is compared to the scope name. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. (optional)
    sortOrder := "sortOrder_example" // string | Sort order of scopes names to retrieve scopes by. By default, scopes are sorted in ascending order by scope name. (optional)
    includeDefaultOpenId := true // bool | Include a default 'openid' scope and description if one is not already configured. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.GetCommonScopes(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).SortOrder(sortOrder).IncludeDefaultOpenId(includeDefaultOpenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.GetCommonScopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommonScopes`: ScopeEntries
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.GetCommonScopes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCommonScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve. | 
 **numberPerPage** | **int64** | Number of common scopes per page. | 
 **filter** | **string** | Filter criteria limits the common scopes that are returned to only those that match it. The filter criteria is compared to the scope name. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. | 
 **sortOrder** | **string** | Sort order of scopes names to retrieve scopes by. By default, scopes are sorted in ascending order by scope name. | 
 **includeDefaultOpenId** | **bool** | Include a default &#39;openid&#39; scope and description if one is not already configured. | [default to false]

### Return type

[**ScopeEntries**](ScopeEntries.md)

### Authorization

[bearer](../README.md#bearer)

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
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.GetExclusiveScope(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.GetExclusiveScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExclusiveScope`: ScopeEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.GetExclusiveScope`: %v\n", resp)
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

[bearer](../README.md#bearer)

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
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.GetExclusiveScopeGroup(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.GetExclusiveScopeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExclusiveScopeGroup`: ScopeGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.GetExclusiveScopeGroup`: %v\n", resp)
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExclusiveScopeGroups

> ScopeGroupEntries GetExclusiveScopeGroups(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).SortOrder(sortOrder).Execute()

Get the exclusive scope groups.



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
    numberPerPage := int64(56) // int64 | Number of exclusive scope groups per page. (optional)
    filter := "filter_example" // string | Filter criteria limits the exclusive scope groups that are returned to only those that match it. The filter criteria is compared to the scope group name. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. (optional)
    sortOrder := "sortOrder_example" // string | Sort order of scope group names to retrieve group by. By default, scope groups are sorted in ascending order by scope group name. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.GetExclusiveScopeGroups(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.GetExclusiveScopeGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExclusiveScopeGroups`: ScopeGroupEntries
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.GetExclusiveScopeGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExclusiveScopeGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve. | 
 **numberPerPage** | **int64** | Number of exclusive scope groups per page. | 
 **filter** | **string** | Filter criteria limits the exclusive scope groups that are returned to only those that match it. The filter criteria is compared to the scope group name. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. | 
 **sortOrder** | **string** | Sort order of scope group names to retrieve group by. By default, scope groups are sorted in ascending order by scope group name. | 

### Return type

[**ScopeGroupEntries**](ScopeGroupEntries.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExclusiveScopes

> ScopeEntries GetExclusiveScopes(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).SortOrder(sortOrder).Execute()

Get the exclusive scopes.



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
    numberPerPage := int64(56) // int64 | Number of exclusive scopes per page. (optional)
    filter := "filter_example" // string | Filter criteria limits the exclusive scopes that are returned to only those that match it. The filter criteria is compared to the scope name. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. (optional)
    sortOrder := "sortOrder_example" // string | Sort order of scopes names to retrieve scopes by. By default, scopes are sorted in ascending order by scope name. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.GetExclusiveScopes(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.GetExclusiveScopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExclusiveScopes`: ScopeEntries
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.GetExclusiveScopes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExclusiveScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve. | 
 **numberPerPage** | **int64** | Number of exclusive scopes per page. | 
 **filter** | **string** | Filter criteria limits the exclusive scopes that are returned to only those that match it. The filter criteria is compared to the scope name. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. | 
 **sortOrder** | **string** | Sort order of scopes names to retrieve scopes by. By default, scopes are sorted in ascending order by scope name. | 

### Return type

[**ScopeEntries**](ScopeEntries.md)

### Authorization

[bearer](../README.md#bearer)

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
    r, err := apiClient.OauthAuthServerSettingsAPI.RemoveCommonScope(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.RemoveCommonScope``: %v\n", err)
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

[bearer](../README.md#bearer)

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
    r, err := apiClient.OauthAuthServerSettingsAPI.RemoveCommonScopeGroup(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.RemoveCommonScopeGroup``: %v\n", err)
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

[bearer](../README.md#bearer)

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
    r, err := apiClient.OauthAuthServerSettingsAPI.RemoveExclusiveScope(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.RemoveExclusiveScope``: %v\n", err)
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

[bearer](../README.md#bearer)

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
    r, err := apiClient.OauthAuthServerSettingsAPI.RemoveExclusiveScopeGroup(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.RemoveExclusiveScopeGroup``: %v\n", err)
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

[bearer](../README.md#bearer)

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
    body := *openapiclient.NewAuthorizationServerSettings(int64(123), int64(123), int64(123), int64(123)) // AuthorizationServerSettings | Configuration for updated server settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.UpdateAuthorizationServerSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.UpdateAuthorizationServerSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthorizationServerSettings`: AuthorizationServerSettings
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.UpdateAuthorizationServerSettings`: %v\n", resp)
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

[bearer](../README.md#bearer)

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
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.UpdateCommonScope(context.Background(), name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.UpdateCommonScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCommonScope`: ScopeEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.UpdateCommonScope`: %v\n", resp)
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

[bearer](../README.md#bearer)

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
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.UpdateCommonScopeGroup(context.Background(), name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.UpdateCommonScopeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCommonScopeGroup`: ScopeGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.UpdateCommonScopeGroup`: %v\n", resp)
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

[bearer](../README.md#bearer)

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
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.UpdateExclusiveScope(context.Background(), name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.UpdateExclusiveScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExclusiveScope`: ScopeEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.UpdateExclusiveScope`: %v\n", resp)
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

[bearer](../README.md#bearer)

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
    resp, r, err := apiClient.OauthAuthServerSettingsAPI.UpdateExclusiveScopeGroups(context.Background(), name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthServerSettingsAPI.UpdateExclusiveScopeGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExclusiveScopeGroups`: ScopeGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthServerSettingsAPI.UpdateExclusiveScopeGroups`: %v\n", resp)
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

