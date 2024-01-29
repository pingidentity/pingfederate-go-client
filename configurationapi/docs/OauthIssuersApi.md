# \OauthIssuersAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOauthIssuer**](OauthIssuersAPI.md#AddOauthIssuer) | **Post** /oauth/issuers | Create a new virtual issuer.
[**DeleteOauthIssuer**](OauthIssuersAPI.md#DeleteOauthIssuer) | **Delete** /oauth/issuers/{id} | Delete a virtual issuer.
[**GetOauthIssuerById**](OauthIssuersAPI.md#GetOauthIssuerById) | **Get** /oauth/issuers/{id} | Find a virtual issuer by ID.
[**GetOauthIssuers**](OauthIssuersAPI.md#GetOauthIssuers) | **Get** /oauth/issuers | Get the list of virtual issuers.
[**UpdateOauthIssuer**](OauthIssuersAPI.md#UpdateOauthIssuer) | **Put** /oauth/issuers/{id} | Update a virtual issuer.



## AddOauthIssuer

> Issuer AddOauthIssuer(ctx).Body(body).Execute()

Create a new virtual issuer.



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
    body := *openapiclient.NewIssuer("Name_example", "Host_example") // Issuer | Configuration for new virtual issuer.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthIssuersAPI.AddOauthIssuer(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthIssuersAPI.AddOauthIssuer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOauthIssuer`: Issuer
    fmt.Fprintf(os.Stdout, "Response from `OauthIssuersAPI.AddOauthIssuer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddOauthIssuerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Issuer**](Issuer.md) | Configuration for new virtual issuer. | 

### Return type

[**Issuer**](Issuer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOauthIssuer

> DeleteOauthIssuer(ctx, id).Execute()

Delete a virtual issuer.



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
    id := "id_example" // string | ID of the virtual issuer to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthIssuersAPI.DeleteOauthIssuer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthIssuersAPI.DeleteOauthIssuer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the virtual issuer to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOauthIssuerRequest struct via the builder pattern


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


## GetOauthIssuerById

> Issuer GetOauthIssuerById(ctx, id).Execute()

Find a virtual issuer by ID.



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
    id := "id_example" // string | ID of the virtual issuer to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthIssuersAPI.GetOauthIssuerById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthIssuersAPI.GetOauthIssuerById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthIssuerById`: Issuer
    fmt.Fprintf(os.Stdout, "Response from `OauthIssuersAPI.GetOauthIssuerById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the virtual issuer to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthIssuerByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Issuer**](Issuer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthIssuers

> Issuers GetOauthIssuers(ctx).Execute()

Get the list of virtual issuers.

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
    resp, r, err := apiClient.OauthIssuersAPI.GetOauthIssuers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthIssuersAPI.GetOauthIssuers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthIssuers`: Issuers
    fmt.Fprintf(os.Stdout, "Response from `OauthIssuersAPI.GetOauthIssuers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthIssuersRequest struct via the builder pattern


### Return type

[**Issuers**](Issuers.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOauthIssuer

> Issuer UpdateOauthIssuer(ctx, id).Body(body).Execute()

Update a virtual issuer.



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
    id := "id_example" // string | ID of the virtual issuer to update.
    body := *openapiclient.NewIssuer("Name_example", "Host_example") // Issuer | Configuration for updated virtual issuer.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthIssuersAPI.UpdateOauthIssuer(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthIssuersAPI.UpdateOauthIssuer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOauthIssuer`: Issuer
    fmt.Fprintf(os.Stdout, "Response from `OauthIssuersAPI.UpdateOauthIssuer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the virtual issuer to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOauthIssuerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Issuer**](Issuer.md) | Configuration for updated virtual issuer. | 

### Return type

[**Issuer**](Issuer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

