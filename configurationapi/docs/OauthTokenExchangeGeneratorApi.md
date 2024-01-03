# \OauthTokenExchangeGeneratorAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroup**](OauthTokenExchangeGeneratorAPI.md#CreateGroup) | **Post** /oauth/tokenExchange/generator/groups | Create a new OAuth 2.0 Token Exchange Generator group.
[**DeleteOauthTokenExchangeGroup**](OauthTokenExchangeGeneratorAPI.md#DeleteOauthTokenExchangeGroup) | **Delete** /oauth/tokenExchange/generator/groups/{id} | Delete an OAuth 2.0 Token Exchange Generator group.
[**GetOauthTokenExchangeGroupById**](OauthTokenExchangeGeneratorAPI.md#GetOauthTokenExchangeGroupById) | **Get** /oauth/tokenExchange/generator/groups/{id} | Find an OAuth 2.0 Token Exchange Generator group by ID.
[**GetOauthTokenExchangeGroups**](OauthTokenExchangeGeneratorAPI.md#GetOauthTokenExchangeGroups) | **Get** /oauth/tokenExchange/generator/groups | Get list of OAuth 2.0 Token Exchange Generator groups.
[**GetOauthTokenExchangeSettings**](OauthTokenExchangeGeneratorAPI.md#GetOauthTokenExchangeSettings) | **Get** /oauth/tokenExchange/generator/settings | Get general OAuth 2.0 Token Exchange Generator settings.
[**UpdateOauthTokenExchangeGroup**](OauthTokenExchangeGeneratorAPI.md#UpdateOauthTokenExchangeGroup) | **Put** /oauth/tokenExchange/generator/groups/{id} | Update an OAuth 2.0 Token Exchange Generator group.
[**UpdateOauthTokenExchangeSettings**](OauthTokenExchangeGeneratorAPI.md#UpdateOauthTokenExchangeSettings) | **Put** /oauth/tokenExchange/generator/settings | Update general OAuth 2.0 Token Exchange Generator settings.



## CreateGroup

> TokenExchangeGeneratorGroup CreateGroup(ctx).Body(body).BypassExternalValidation(bypassExternalValidation).Execute()

Create a new OAuth 2.0 Token Exchange Generator group.



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
    body := *openapiclient.NewTokenExchangeGeneratorGroup("Id_example", "Name_example", []openapiclient.TokenExchangeGeneratorMapping{*openapiclient.NewTokenExchangeGeneratorMapping("RequestedTokenType_example", *openapiclient.NewResourceLink("Id_example"))}) // TokenExchangeGeneratorGroup | Configuration for new OAuth 2.0 Token Exchange Generator.
    bypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthTokenExchangeGeneratorAPI.CreateGroup(context.Background()).Body(body).BypassExternalValidation(bypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeGeneratorAPI.CreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroup`: TokenExchangeGeneratorGroup
    fmt.Fprintf(os.Stdout, "Response from `OauthTokenExchangeGeneratorAPI.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TokenExchangeGeneratorGroup**](TokenExchangeGeneratorGroup.md) | Configuration for new OAuth 2.0 Token Exchange Generator. | 
 **bypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | 

### Return type

[**TokenExchangeGeneratorGroup**](TokenExchangeGeneratorGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOauthTokenExchangeGroup

> DeleteOauthTokenExchangeGroup(ctx, id).Execute()

Delete an OAuth 2.0 Token Exchange Generator group.



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
    id := "id_example" // string | ID of OAuth 2.0 Token Exchange Generator group to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthTokenExchangeGeneratorAPI.DeleteOauthTokenExchangeGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeGeneratorAPI.DeleteOauthTokenExchangeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of OAuth 2.0 Token Exchange Generator group to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOauthTokenExchangeGroupRequest struct via the builder pattern


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


## GetOauthTokenExchangeGroupById

> TokenExchangeGeneratorGroup GetOauthTokenExchangeGroupById(ctx, id).Execute()

Find an OAuth 2.0 Token Exchange Generator group by ID.



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
    id := "id_example" // string | ID of the OAuth 2.0 Token Exchange Generator group to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthTokenExchangeGeneratorAPI.GetOauthTokenExchangeGroupById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeGeneratorAPI.GetOauthTokenExchangeGroupById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthTokenExchangeGroupById`: TokenExchangeGeneratorGroup
    fmt.Fprintf(os.Stdout, "Response from `OauthTokenExchangeGeneratorAPI.GetOauthTokenExchangeGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the OAuth 2.0 Token Exchange Generator group to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthTokenExchangeGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenExchangeGeneratorGroup**](TokenExchangeGeneratorGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthTokenExchangeGroups

> TokenExchangeGeneratorGroups GetOauthTokenExchangeGroups(ctx).Execute()

Get list of OAuth 2.0 Token Exchange Generator groups.

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
    resp, r, err := apiClient.OauthTokenExchangeGeneratorAPI.GetOauthTokenExchangeGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeGeneratorAPI.GetOauthTokenExchangeGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthTokenExchangeGroups`: TokenExchangeGeneratorGroups
    fmt.Fprintf(os.Stdout, "Response from `OauthTokenExchangeGeneratorAPI.GetOauthTokenExchangeGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthTokenExchangeGroupsRequest struct via the builder pattern


### Return type

[**TokenExchangeGeneratorGroups**](TokenExchangeGeneratorGroups.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthTokenExchangeSettings

> TokenExchangeGeneratorSettings GetOauthTokenExchangeSettings(ctx).Execute()

Get general OAuth 2.0 Token Exchange Generator settings.

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
    resp, r, err := apiClient.OauthTokenExchangeGeneratorAPI.GetOauthTokenExchangeSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeGeneratorAPI.GetOauthTokenExchangeSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthTokenExchangeSettings`: TokenExchangeGeneratorSettings
    fmt.Fprintf(os.Stdout, "Response from `OauthTokenExchangeGeneratorAPI.GetOauthTokenExchangeSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthTokenExchangeSettingsRequest struct via the builder pattern


### Return type

[**TokenExchangeGeneratorSettings**](TokenExchangeGeneratorSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOauthTokenExchangeGroup

> TokenExchangeGeneratorGroup UpdateOauthTokenExchangeGroup(ctx, id).Body(body).BypassExternalValidation(bypassExternalValidation).Execute()

Update an OAuth 2.0 Token Exchange Generator group.



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
    id := "id_example" // string | ID of the OAuth 2.0 Token Exchange Generator group to update.
    body := *openapiclient.NewTokenExchangeGeneratorGroup("Id_example", "Name_example", []openapiclient.TokenExchangeGeneratorMapping{*openapiclient.NewTokenExchangeGeneratorMapping("RequestedTokenType_example", *openapiclient.NewResourceLink("Id_example"))}) // TokenExchangeGeneratorGroup | Configuration for updated OAuth 2.0 Token Exchange Generator group.
    bypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthTokenExchangeGeneratorAPI.UpdateOauthTokenExchangeGroup(context.Background(), id).Body(body).BypassExternalValidation(bypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeGeneratorAPI.UpdateOauthTokenExchangeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOauthTokenExchangeGroup`: TokenExchangeGeneratorGroup
    fmt.Fprintf(os.Stdout, "Response from `OauthTokenExchangeGeneratorAPI.UpdateOauthTokenExchangeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the OAuth 2.0 Token Exchange Generator group to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOauthTokenExchangeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TokenExchangeGeneratorGroup**](TokenExchangeGeneratorGroup.md) | Configuration for updated OAuth 2.0 Token Exchange Generator group. | 
 **bypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | 

### Return type

[**TokenExchangeGeneratorGroup**](TokenExchangeGeneratorGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOauthTokenExchangeSettings

> TokenExchangeGeneratorSettings UpdateOauthTokenExchangeSettings(ctx).Body(body).BypassExternalValidation(bypassExternalValidation).Execute()

Update general OAuth 2.0 Token Exchange Generator settings.

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
    body := *openapiclient.NewTokenExchangeGeneratorSettings() // TokenExchangeGeneratorSettings | OAuth 2.0 Token Exchange Generator settings.
    bypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthTokenExchangeGeneratorAPI.UpdateOauthTokenExchangeSettings(context.Background()).Body(body).BypassExternalValidation(bypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeGeneratorAPI.UpdateOauthTokenExchangeSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOauthTokenExchangeSettings`: TokenExchangeGeneratorSettings
    fmt.Fprintf(os.Stdout, "Response from `OauthTokenExchangeGeneratorAPI.UpdateOauthTokenExchangeSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOauthTokenExchangeSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TokenExchangeGeneratorSettings**](TokenExchangeGeneratorSettings.md) | OAuth 2.0 Token Exchange Generator settings. | 
 **bypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | 

### Return type

[**TokenExchangeGeneratorSettings**](TokenExchangeGeneratorSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

