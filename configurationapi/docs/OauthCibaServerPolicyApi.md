# \OauthCibaServerPolicyAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCibaServerPolicy**](OauthCibaServerPolicyAPI.md#CreateCibaServerPolicy) | **Post** /oauth/cibaServerPolicy/requestPolicies | Create a new request policy.
[**DeleteCibaServerPolicy**](OauthCibaServerPolicyAPI.md#DeleteCibaServerPolicy) | **Delete** /oauth/cibaServerPolicy/requestPolicies/{id} | Delete a request policy.
[**GetCibaServerPolicies**](OauthCibaServerPolicyAPI.md#GetCibaServerPolicies) | **Get** /oauth/cibaServerPolicy/requestPolicies | Get list of request policies.
[**GetCibaServerPolicyById**](OauthCibaServerPolicyAPI.md#GetCibaServerPolicyById) | **Get** /oauth/cibaServerPolicy/requestPolicies/{id} | Find request policy by ID.
[**GetCibaServerPolicySettings**](OauthCibaServerPolicyAPI.md#GetCibaServerPolicySettings) | **Get** /oauth/cibaServerPolicy/settings | Get general ciba server request policy settings.
[**UpdateCibaServerPolicy**](OauthCibaServerPolicyAPI.md#UpdateCibaServerPolicy) | **Put** /oauth/cibaServerPolicy/requestPolicies/{id} | Update a request policy.
[**UpdateCibaServerPolicySettings**](OauthCibaServerPolicyAPI.md#UpdateCibaServerPolicySettings) | **Put** /oauth/cibaServerPolicy/settings | Update general ciba server request policy settings.



## CreateCibaServerPolicy

> RequestPolicy CreateCibaServerPolicy(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new request policy.



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
    body := *openapiclient.NewRequestPolicy("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewIdentityHintContract([]openapiclient.IdentityHintAttribute{*openapiclient.NewIdentityHintAttribute("Name_example")})) // RequestPolicy | Configuration for new policy.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthCibaServerPolicyAPI.CreateCibaServerPolicy(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthCibaServerPolicyAPI.CreateCibaServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCibaServerPolicy`: RequestPolicy
    fmt.Fprintf(os.Stdout, "Response from `OauthCibaServerPolicyAPI.CreateCibaServerPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCibaServerPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RequestPolicy**](RequestPolicy.md) | Configuration for new policy. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**RequestPolicy**](RequestPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCibaServerPolicy

> DeleteCibaServerPolicy(ctx, id).Execute()

Delete a request policy.



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
    id := "id_example" // string | ID of request policy to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthCibaServerPolicyAPI.DeleteCibaServerPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthCibaServerPolicyAPI.DeleteCibaServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of request policy to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCibaServerPolicyRequest struct via the builder pattern


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


## GetCibaServerPolicies

> RequestPolicies GetCibaServerPolicies(ctx).Execute()

Get list of request policies.

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
    resp, r, err := apiClient.OauthCibaServerPolicyAPI.GetCibaServerPolicies(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthCibaServerPolicyAPI.GetCibaServerPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCibaServerPolicies`: RequestPolicies
    fmt.Fprintf(os.Stdout, "Response from `OauthCibaServerPolicyAPI.GetCibaServerPolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCibaServerPoliciesRequest struct via the builder pattern


### Return type

[**RequestPolicies**](RequestPolicies.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCibaServerPolicyById

> RequestPolicy GetCibaServerPolicyById(ctx, id).Execute()

Find request policy by ID.



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
    id := "id_example" // string | ID of the request policy to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthCibaServerPolicyAPI.GetCibaServerPolicyById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthCibaServerPolicyAPI.GetCibaServerPolicyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCibaServerPolicyById`: RequestPolicy
    fmt.Fprintf(os.Stdout, "Response from `OauthCibaServerPolicyAPI.GetCibaServerPolicyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the request policy to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCibaServerPolicyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestPolicy**](RequestPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCibaServerPolicySettings

> CibaServerPolicySettings GetCibaServerPolicySettings(ctx).Execute()

Get general ciba server request policy settings.

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
    resp, r, err := apiClient.OauthCibaServerPolicyAPI.GetCibaServerPolicySettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthCibaServerPolicyAPI.GetCibaServerPolicySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCibaServerPolicySettings`: CibaServerPolicySettings
    fmt.Fprintf(os.Stdout, "Response from `OauthCibaServerPolicyAPI.GetCibaServerPolicySettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCibaServerPolicySettingsRequest struct via the builder pattern


### Return type

[**CibaServerPolicySettings**](CibaServerPolicySettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCibaServerPolicy

> RequestPolicy UpdateCibaServerPolicy(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update a request policy.



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
    id := "id_example" // string | ID of the request policy to update.
    body := *openapiclient.NewRequestPolicy("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewIdentityHintContract([]openapiclient.IdentityHintAttribute{*openapiclient.NewIdentityHintAttribute("Name_example")})) // RequestPolicy | Configuration for updated policy.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthCibaServerPolicyAPI.UpdateCibaServerPolicy(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthCibaServerPolicyAPI.UpdateCibaServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCibaServerPolicy`: RequestPolicy
    fmt.Fprintf(os.Stdout, "Response from `OauthCibaServerPolicyAPI.UpdateCibaServerPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the request policy to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCibaServerPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RequestPolicy**](RequestPolicy.md) | Configuration for updated policy. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**RequestPolicy**](RequestPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCibaServerPolicySettings

> CibaServerPolicySettings UpdateCibaServerPolicySettings(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update general ciba server request policy settings.

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
    body := *openapiclient.NewCibaServerPolicySettings() // CibaServerPolicySettings | Ciba server request policy settings.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthCibaServerPolicyAPI.UpdateCibaServerPolicySettings(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthCibaServerPolicyAPI.UpdateCibaServerPolicySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCibaServerPolicySettings`: CibaServerPolicySettings
    fmt.Fprintf(os.Stdout, "Response from `OauthCibaServerPolicyAPI.UpdateCibaServerPolicySettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCibaServerPolicySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CibaServerPolicySettings**](CibaServerPolicySettings.md) | Ciba server request policy settings. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**CibaServerPolicySettings**](CibaServerPolicySettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

