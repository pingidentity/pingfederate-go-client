# \OauthCimdPoliciesAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCimdPolicy**](OauthCimdPoliciesAPI.md#CreateCimdPolicy) | **Post** /oauth/cimd/policies | Create a new CIMD policy.
[**DeleteCimdPolicy**](OauthCimdPoliciesAPI.md#DeleteCimdPolicy) | **Delete** /oauth/cimd/policies/{id} | Delete a CIMD policy.
[**GetCimdPolicies**](OauthCimdPoliciesAPI.md#GetCimdPolicies) | **Get** /oauth/cimd/policies | Get list of CIMD policies.
[**GetCimdPolicy**](OauthCimdPoliciesAPI.md#GetCimdPolicy) | **Get** /oauth/cimd/policies/{id} | Find a CIMD policy by ID.
[**UpdateCimdPolicy**](OauthCimdPoliciesAPI.md#UpdateCimdPolicy) | **Put** /oauth/cimd/policies/{id} | Update a CIMD policy.



## CreateCimdPolicy

> CimdPolicy CreateCimdPolicy(ctx).Body(body).Execute()

Create a new CIMD policy.



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
    body := *openapiclient.NewCimdPolicy("Name_example", false, []string{"MetadataUrls_example"}) // CimdPolicy | Configuration for the new CIMD policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthCimdPoliciesAPI.CreateCimdPolicy(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthCimdPoliciesAPI.CreateCimdPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCimdPolicy`: CimdPolicy
    fmt.Fprintf(os.Stdout, "Response from `OauthCimdPoliciesAPI.CreateCimdPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCimdPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CimdPolicy**](CimdPolicy.md) | Configuration for the new CIMD policy. | 

### Return type

[**CimdPolicy**](CimdPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCimdPolicy

> DeleteCimdPolicy(ctx, id).Execute()

Delete a CIMD policy.



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
    id := "id_example" // string | ID of the CIMD policy to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthCimdPoliciesAPI.DeleteCimdPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthCimdPoliciesAPI.DeleteCimdPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the CIMD policy to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCimdPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCimdPolicies

> CimdPolicies GetCimdPolicies(ctx).Execute()

Get list of CIMD policies.

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
    resp, r, err := apiClient.OauthCimdPoliciesAPI.GetCimdPolicies(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthCimdPoliciesAPI.GetCimdPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCimdPolicies`: CimdPolicies
    fmt.Fprintf(os.Stdout, "Response from `OauthCimdPoliciesAPI.GetCimdPolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCimdPoliciesRequest struct via the builder pattern


### Return type

[**CimdPolicies**](CimdPolicies.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCimdPolicy

> CimdPolicy GetCimdPolicy(ctx, id).Execute()

Find a CIMD policy by ID.



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
    id := "id_example" // string | ID of the CIMD policy to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthCimdPoliciesAPI.GetCimdPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthCimdPoliciesAPI.GetCimdPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCimdPolicy`: CimdPolicy
    fmt.Fprintf(os.Stdout, "Response from `OauthCimdPoliciesAPI.GetCimdPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the CIMD policy to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCimdPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CimdPolicy**](CimdPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCimdPolicy

> CimdPolicy UpdateCimdPolicy(ctx, id).Body(body).Execute()

Update a CIMD policy.



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
    id := "id_example" // string | ID of the CIMD policy to update.
    body := *openapiclient.NewCimdPolicy("Name_example", false, []string{"MetadataUrls_example"}) // CimdPolicy | Configuration for the updated CIMD policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthCimdPoliciesAPI.UpdateCimdPolicy(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthCimdPoliciesAPI.UpdateCimdPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCimdPolicy`: CimdPolicy
    fmt.Fprintf(os.Stdout, "Response from `OauthCimdPoliciesAPI.UpdateCimdPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the CIMD policy to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCimdPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CimdPolicy**](CimdPolicy.md) | Configuration for the updated CIMD policy. | 

### Return type

[**CimdPolicy**](CimdPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [bearer](../README.md#bearer), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

