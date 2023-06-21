# \OauthClientsApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOauthClient**](OauthClientsApi.md#CreateOauthClient) | **Post** /oauth/clients | Create a new OAuth client.
[**DeleteOauthClient**](OauthClientsApi.md#DeleteOauthClient) | **Delete** /oauth/clients/{id} | Delete an OAuth client.
[**GetOauthClientById**](OauthClientsApi.md#GetOauthClientById) | **Get** /oauth/clients/{id} | Find the OAuth client by ID.
[**GetOauthClientSecret**](OauthClientsApi.md#GetOauthClientSecret) | **Get** /oauth/clients/{id}/clientAuth/clientSecret | Get the client secret of an existing OAuth client.
[**GetOauthClients**](OauthClientsApi.md#GetOauthClients) | **Get** /oauth/clients | Get the list of OAuth clients.
[**UpdateOauthClient**](OauthClientsApi.md#UpdateOauthClient) | **Put** /oauth/clients/{id} | Updates the OAuth client.
[**UpdateOauthClientSecret**](OauthClientsApi.md#UpdateOauthClientSecret) | **Put** /oauth/clients/{id}/clientAuth/clientSecret | Update the client secret of an existing OAuth client.



## CreateOauthClient

> Client CreateOauthClient(ctx).Body(body).Execute()

Create a new OAuth client.



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
    body := *openapiclient.NewClient("ClientId_example", []string{"GrantTypes_example"}, "Name_example") // Client | Configuration for new client.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthClientsApi.CreateOauthClient(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientsApi.CreateOauthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOauthClient`: Client
    fmt.Fprintf(os.Stdout, "Response from `OauthClientsApi.CreateOauthClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOauthClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Client**](Client.md) | Configuration for new client. | 

### Return type

[**Client**](Client.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOauthClient

> DeleteOauthClient(ctx, id).Execute()

Delete an OAuth client.

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
    id := "id_example" // string | ID of the client.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthClientsApi.DeleteOauthClient(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientsApi.DeleteOauthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOauthClientRequest struct via the builder pattern


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


## GetOauthClientById

> Client GetOauthClientById(ctx, id).Execute()

Find the OAuth client by ID.

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
    id := "id_example" // string | ID of the client.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthClientsApi.GetOauthClientById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientsApi.GetOauthClientById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthClientById`: Client
    fmt.Fprintf(os.Stdout, "Response from `OauthClientsApi.GetOauthClientById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthClientByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Client**](Client.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthClientSecret

> ClientSecret GetOauthClientSecret(ctx, id).Execute()

Get the client secret of an existing OAuth client.

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
    id := "id_example" // string | ID of the client.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthClientsApi.GetOauthClientSecret(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientsApi.GetOauthClientSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthClientSecret`: ClientSecret
    fmt.Fprintf(os.Stdout, "Response from `OauthClientsApi.GetOauthClientSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthClientSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientSecret**](ClientSecret.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthClients

> Clients GetOauthClients(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()

Get the list of OAuth clients.

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
    numberPerPage := int64(56) // int64 | Number of OAuth clients per page. (uncapped if unspecified) (optional)
    filter := "filter_example" // string | Filter criteria limits the OAuth clients that are returned to only those that match it. The filter criteria is compared to the OAuth client name and ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthClientsApi.GetOauthClients(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientsApi.GetOauthClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthClients`: Clients
    fmt.Fprintf(os.Stdout, "Response from `OauthClientsApi.GetOauthClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve. | 
 **numberPerPage** | **int64** | Number of OAuth clients per page. (uncapped if unspecified) | 
 **filter** | **string** | Filter criteria limits the OAuth clients that are returned to only those that match it. The filter criteria is compared to the OAuth client name and ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. | 

### Return type

[**Clients**](Clients.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOauthClient

> Client UpdateOauthClient(ctx, id).Body(body).Execute()

Updates the OAuth client.

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
    id := "id_example" // string | ID of the client to be updated.
    body := *openapiclient.NewClient("ClientId_example", []string{"GrantTypes_example"}, "Name_example") // Client | Configuration for the client.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthClientsApi.UpdateOauthClient(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientsApi.UpdateOauthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOauthClient`: Client
    fmt.Fprintf(os.Stdout, "Response from `OauthClientsApi.UpdateOauthClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the client to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOauthClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Client**](Client.md) | Configuration for the client. | 

### Return type

[**Client**](Client.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOauthClientSecret

> ClientSecret UpdateOauthClientSecret(ctx, id).Body(body).Execute()

Update the client secret of an existing OAuth client.

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
    id := "id_example" // string | ID of the client to be updated.
    body := *openapiclient.NewClientSecret() // ClientSecret | Client Secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthClientsApi.UpdateOauthClientSecret(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientsApi.UpdateOauthClientSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOauthClientSecret`: ClientSecret
    fmt.Fprintf(os.Stdout, "Response from `OauthClientsApi.UpdateOauthClientSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the client to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOauthClientSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ClientSecret**](ClientSecret.md) | Client Secret. | 

### Return type

[**ClientSecret**](ClientSecret.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

