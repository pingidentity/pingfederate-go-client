# \OauthOpenIdConnectAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOIDCPolicy**](OauthOpenIdConnectAPI.md#CreateOIDCPolicy) | **Post** /oauth/openIdConnect/policies | Create a new OpenID Connect Policy.
[**DeleteOIDCPolicy**](OauthOpenIdConnectAPI.md#DeleteOIDCPolicy) | **Delete** /oauth/openIdConnect/policies/{id} | Delete an OpenID Connect Policy.
[**GetOIDCPolicies**](OauthOpenIdConnectAPI.md#GetOIDCPolicies) | **Get** /oauth/openIdConnect/policies | Get list of OpenID Connect Policies.
[**GetOIDCPolicy**](OauthOpenIdConnectAPI.md#GetOIDCPolicy) | **Get** /oauth/openIdConnect/policies/{id} | Find OpenID Connect Policy by ID.
[**GetOIDCSettings**](OauthOpenIdConnectAPI.md#GetOIDCSettings) | **Get** /oauth/openIdConnect/settings | Get the OpenID Connect Settings.
[**UpdateOIDCPolicy**](OauthOpenIdConnectAPI.md#UpdateOIDCPolicy) | **Put** /oauth/openIdConnect/policies/{id} | Update an OpenID Connect Policy.
[**UpdateOIDCSettings**](OauthOpenIdConnectAPI.md#UpdateOIDCSettings) | **Put** /oauth/openIdConnect/settings | Set the OpenID Connect Settings.



## CreateOIDCPolicy

> OpenIdConnectPolicy CreateOIDCPolicy(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new OpenID Connect Policy.



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
    body := *openapiclient.NewOpenIdConnectPolicy("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewOpenIdConnectAttributeContract(), *openapiclient.NewAttributeMapping(map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")})) // OpenIdConnectPolicy | Configuration for new policy.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthOpenIdConnectAPI.CreateOIDCPolicy(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthOpenIdConnectAPI.CreateOIDCPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOIDCPolicy`: OpenIdConnectPolicy
    fmt.Fprintf(os.Stdout, "Response from `OauthOpenIdConnectAPI.CreateOIDCPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOIDCPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OpenIdConnectPolicy**](OpenIdConnectPolicy.md) | Configuration for new policy. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**OpenIdConnectPolicy**](OpenIdConnectPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOIDCPolicy

> DeleteOIDCPolicy(ctx, id).Execute()

Delete an OpenID Connect Policy.



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
    id := "id_example" // string | ID of OpenID Connect Policy to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthOpenIdConnectAPI.DeleteOIDCPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthOpenIdConnectAPI.DeleteOIDCPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of OpenID Connect Policy to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOIDCPolicyRequest struct via the builder pattern


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


## GetOIDCPolicies

> OpenIdConnectPolicies GetOIDCPolicies(ctx).Execute()

Get list of OpenID Connect Policies.

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
    resp, r, err := apiClient.OauthOpenIdConnectAPI.GetOIDCPolicies(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthOpenIdConnectAPI.GetOIDCPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOIDCPolicies`: OpenIdConnectPolicies
    fmt.Fprintf(os.Stdout, "Response from `OauthOpenIdConnectAPI.GetOIDCPolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOIDCPoliciesRequest struct via the builder pattern


### Return type

[**OpenIdConnectPolicies**](OpenIdConnectPolicies.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOIDCPolicy

> OpenIdConnectPolicy GetOIDCPolicy(ctx, id).Execute()

Find OpenID Connect Policy by ID.



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
    id := "id_example" // string | ID of the OpenID Connect Policy to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthOpenIdConnectAPI.GetOIDCPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthOpenIdConnectAPI.GetOIDCPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOIDCPolicy`: OpenIdConnectPolicy
    fmt.Fprintf(os.Stdout, "Response from `OauthOpenIdConnectAPI.GetOIDCPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the OpenID Connect Policy to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOIDCPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpenIdConnectPolicy**](OpenIdConnectPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOIDCSettings

> OpenIdConnectSettings GetOIDCSettings(ctx).Execute()

Get the OpenID Connect Settings.

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
    resp, r, err := apiClient.OauthOpenIdConnectAPI.GetOIDCSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthOpenIdConnectAPI.GetOIDCSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOIDCSettings`: OpenIdConnectSettings
    fmt.Fprintf(os.Stdout, "Response from `OauthOpenIdConnectAPI.GetOIDCSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOIDCSettingsRequest struct via the builder pattern


### Return type

[**OpenIdConnectSettings**](OpenIdConnectSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOIDCPolicy

> OpenIdConnectPolicy UpdateOIDCPolicy(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update an OpenID Connect Policy.



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
    id := "id_example" // string | ID of the OpenID Connect Policy to update.
    body := *openapiclient.NewOpenIdConnectPolicy("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewOpenIdConnectAttributeContract(), *openapiclient.NewAttributeMapping(map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")})) // OpenIdConnectPolicy | Configuration for updated policy.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthOpenIdConnectAPI.UpdateOIDCPolicy(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthOpenIdConnectAPI.UpdateOIDCPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOIDCPolicy`: OpenIdConnectPolicy
    fmt.Fprintf(os.Stdout, "Response from `OauthOpenIdConnectAPI.UpdateOIDCPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the OpenID Connect Policy to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOIDCPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OpenIdConnectPolicy**](OpenIdConnectPolicy.md) | Configuration for updated policy. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**OpenIdConnectPolicy**](OpenIdConnectPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOIDCSettings

> OpenIdConnectSettings UpdateOIDCSettings(ctx).Body(body).Execute()

Set the OpenID Connect Settings.

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
    body := *openapiclient.NewOpenIdConnectSettings() // OpenIdConnectSettings | OpenID Connect Settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthOpenIdConnectAPI.UpdateOIDCSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthOpenIdConnectAPI.UpdateOIDCSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOIDCSettings`: OpenIdConnectSettings
    fmt.Fprintf(os.Stdout, "Response from `OauthOpenIdConnectAPI.UpdateOIDCSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOIDCSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OpenIdConnectSettings**](OpenIdConnectSettings.md) | OpenID Connect Settings. | 

### Return type

[**OpenIdConnectSettings**](OpenIdConnectSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

