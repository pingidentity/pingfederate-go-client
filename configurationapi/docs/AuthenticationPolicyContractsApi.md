# \AuthenticationPolicyContractsAPI

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthenticationPolicyContract**](AuthenticationPolicyContractsAPI.md#CreateAuthenticationPolicyContract) | **Post** /authenticationPolicyContracts | Create a new Authentication Policy Contract.
[**DeleteAuthenticationPolicyContract**](AuthenticationPolicyContractsAPI.md#DeleteAuthenticationPolicyContract) | **Delete** /authenticationPolicyContracts/{id} | Delete an Authentication Policy Contract.
[**GetAuthenticationPolicyContract**](AuthenticationPolicyContractsAPI.md#GetAuthenticationPolicyContract) | **Get** /authenticationPolicyContracts/{id} | Gets the Authentication Policy Contract by ID.
[**GetAuthenticationPolicyContracts**](AuthenticationPolicyContractsAPI.md#GetAuthenticationPolicyContracts) | **Get** /authenticationPolicyContracts | Gets the Authentication Policy Contracts.
[**UpdateAuthenticationPolicyContract**](AuthenticationPolicyContractsAPI.md#UpdateAuthenticationPolicyContract) | **Put** /authenticationPolicyContracts/{id} | Update an Authentication Policy Contract by ID.



## CreateAuthenticationPolicyContract

> AuthenticationPolicyContract CreateAuthenticationPolicyContract(ctx).Body(body).Execute()

Create a new Authentication Policy Contract.



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
    body := *openapiclient.NewAuthenticationPolicyContract() // AuthenticationPolicyContract | Configuration for a new contract.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationPolicyContractsAPI.CreateAuthenticationPolicyContract(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPolicyContractsAPI.CreateAuthenticationPolicyContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthenticationPolicyContract`: AuthenticationPolicyContract
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationPolicyContractsAPI.CreateAuthenticationPolicyContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthenticationPolicyContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthenticationPolicyContract**](AuthenticationPolicyContract.md) | Configuration for a new contract. | 

### Return type

[**AuthenticationPolicyContract**](AuthenticationPolicyContract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthenticationPolicyContract

> DeleteAuthenticationPolicyContract(ctx, id).Execute()

Delete an Authentication Policy Contract.



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
    id := "id_example" // string | ID of Authentication Policy Contract to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthenticationPolicyContractsAPI.DeleteAuthenticationPolicyContract(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPolicyContractsAPI.DeleteAuthenticationPolicyContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Authentication Policy Contract to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthenticationPolicyContractRequest struct via the builder pattern


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


## GetAuthenticationPolicyContract

> AuthenticationPolicyContract GetAuthenticationPolicyContract(ctx, id).Execute()

Gets the Authentication Policy Contract by ID.



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
    id := "id_example" // string | ID of contract to fetch

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationPolicyContractsAPI.GetAuthenticationPolicyContract(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPolicyContractsAPI.GetAuthenticationPolicyContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticationPolicyContract`: AuthenticationPolicyContract
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationPolicyContractsAPI.GetAuthenticationPolicyContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of contract to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationPolicyContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticationPolicyContract**](AuthenticationPolicyContract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticationPolicyContracts

> AuthenticationPolicyContracts GetAuthenticationPolicyContracts(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()

Gets the Authentication Policy Contracts.

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
    numberPerPage := int64(56) // int64 | Number of contracts per page. (optional)
    filter := "filter_example" // string | Filter criteria limits the authentication policy contracts that are returned to only those that match it. The filter criteria is compared to the authentication policy contract name and ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationPolicyContractsAPI.GetAuthenticationPolicyContracts(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPolicyContractsAPI.GetAuthenticationPolicyContracts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticationPolicyContracts`: AuthenticationPolicyContracts
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationPolicyContractsAPI.GetAuthenticationPolicyContracts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationPolicyContractsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve. | 
 **numberPerPage** | **int64** | Number of contracts per page. | 
 **filter** | **string** | Filter criteria limits the authentication policy contracts that are returned to only those that match it. The filter criteria is compared to the authentication policy contract name and ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. | 

### Return type

[**AuthenticationPolicyContracts**](AuthenticationPolicyContracts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthenticationPolicyContract

> AuthenticationPolicyContract UpdateAuthenticationPolicyContract(ctx, id).Body(body).Execute()

Update an Authentication Policy Contract by ID.



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
    id := "id_example" // string | ID of the Authentication Policy Contract to update.
    body := *openapiclient.NewAuthenticationPolicyContract() // AuthenticationPolicyContract | Configuration for updated Authentication Policy Contract.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationPolicyContractsAPI.UpdateAuthenticationPolicyContract(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationPolicyContractsAPI.UpdateAuthenticationPolicyContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthenticationPolicyContract`: AuthenticationPolicyContract
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationPolicyContractsAPI.UpdateAuthenticationPolicyContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Authentication Policy Contract to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthenticationPolicyContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthenticationPolicyContract**](AuthenticationPolicyContract.md) | Configuration for updated Authentication Policy Contract. | 

### Return type

[**AuthenticationPolicyContract**](AuthenticationPolicyContract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

