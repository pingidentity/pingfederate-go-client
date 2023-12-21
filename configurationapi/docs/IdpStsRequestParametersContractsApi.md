# \IdpStsRequestParametersContractsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStsRequestParamContract**](IdpStsRequestParametersContractsAPI.md#CreateStsRequestParamContract) | **Post** /idp/stsRequestParametersContracts | Create a new STS Request Parameters Contract.
[**DeleteStsRequestParamContractById**](IdpStsRequestParametersContractsAPI.md#DeleteStsRequestParamContractById) | **Delete** /idp/stsRequestParametersContracts/{id} | Delete a STS Request Parameters Contract.
[**GetStsRequestParamContractById**](IdpStsRequestParametersContractsAPI.md#GetStsRequestParamContractById) | **Get** /idp/stsRequestParametersContracts/{id} | Get a STS Request Parameters Contract.
[**GetStsRequestParamContracts**](IdpStsRequestParametersContractsAPI.md#GetStsRequestParamContracts) | **Get** /idp/stsRequestParametersContracts | Get the list of STS Request Parameters Contracts.
[**UpdateStsRequestParamContractById**](IdpStsRequestParametersContractsAPI.md#UpdateStsRequestParamContractById) | **Put** /idp/stsRequestParametersContracts/{id} | Update a STS Request Parameters Contract.



## CreateStsRequestParamContract

> StsRequestParametersContract CreateStsRequestParamContract(ctx).Body(body).Execute()

Create a new STS Request Parameters Contract.

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
    body := *openapiclient.NewStsRequestParametersContract("Id_example", "Name_example", []string{"Parameters_example"}) // StsRequestParametersContract | Details for the STS Request Parameters Contract.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpStsRequestParametersContractsAPI.CreateStsRequestParamContract(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpStsRequestParametersContractsAPI.CreateStsRequestParamContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStsRequestParamContract`: StsRequestParametersContract
    fmt.Fprintf(os.Stdout, "Response from `IdpStsRequestParametersContractsAPI.CreateStsRequestParamContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStsRequestParamContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**StsRequestParametersContract**](StsRequestParametersContract.md) | Details for the STS Request Parameters Contract. | 

### Return type

[**StsRequestParametersContract**](StsRequestParametersContract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStsRequestParamContractById

> DeleteStsRequestParamContractById(ctx, id).Execute()

Delete a STS Request Parameters Contract.

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
    id := "id_example" // string | ID of STS Request Parameters Contract to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdpStsRequestParametersContractsAPI.DeleteStsRequestParamContractById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpStsRequestParametersContractsAPI.DeleteStsRequestParamContractById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of STS Request Parameters Contract to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStsRequestParamContractByIdRequest struct via the builder pattern


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


## GetStsRequestParamContractById

> StsRequestParametersContract GetStsRequestParamContractById(ctx, id).Execute()

Get a STS Request Parameters Contract.

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
    id := "id_example" // string | ID of STS Request Parameters Contract to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpStsRequestParametersContractsAPI.GetStsRequestParamContractById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpStsRequestParametersContractsAPI.GetStsRequestParamContractById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStsRequestParamContractById`: StsRequestParametersContract
    fmt.Fprintf(os.Stdout, "Response from `IdpStsRequestParametersContractsAPI.GetStsRequestParamContractById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of STS Request Parameters Contract to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStsRequestParamContractByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StsRequestParametersContract**](StsRequestParametersContract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStsRequestParamContracts

> StsRequestParametersContracts GetStsRequestParamContracts(ctx).Execute()

Get the list of STS Request Parameters Contracts.

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
    resp, r, err := apiClient.IdpStsRequestParametersContractsAPI.GetStsRequestParamContracts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpStsRequestParametersContractsAPI.GetStsRequestParamContracts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStsRequestParamContracts`: StsRequestParametersContracts
    fmt.Fprintf(os.Stdout, "Response from `IdpStsRequestParametersContractsAPI.GetStsRequestParamContracts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStsRequestParamContractsRequest struct via the builder pattern


### Return type

[**StsRequestParametersContracts**](StsRequestParametersContracts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStsRequestParamContractById

> StsRequestParametersContract UpdateStsRequestParamContractById(ctx, id).Body(body).Execute()

Update a STS Request Parameters Contract.

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
    id := "id_example" // string | ID of STS Request Parameters Contract to update.
    body := *openapiclient.NewStsRequestParametersContract("Id_example", "Name_example", []string{"Parameters_example"}) // StsRequestParametersContract | Details for updated STS Request Parameters Contract.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpStsRequestParametersContractsAPI.UpdateStsRequestParamContractById(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpStsRequestParametersContractsAPI.UpdateStsRequestParamContractById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStsRequestParamContractById`: StsRequestParametersContract
    fmt.Fprintf(os.Stdout, "Response from `IdpStsRequestParametersContractsAPI.UpdateStsRequestParamContractById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of STS Request Parameters Contract to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStsRequestParamContractByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**StsRequestParametersContract**](StsRequestParametersContract.md) | Details for updated STS Request Parameters Contract. | 

### Return type

[**StsRequestParametersContract**](StsRequestParametersContract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

