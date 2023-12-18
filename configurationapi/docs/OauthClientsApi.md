# \OauthClientsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOauthClient**](OauthClientsAPI.md#CreateOauthClient) | **Post** /oauth/clients | Create a new OAuth client.
[**DeleteOauthClient**](OauthClientsAPI.md#DeleteOauthClient) | **Delete** /oauth/clients/{id} | Delete an OAuth client.
[**GetClientSecret**](OauthClientsAPI.md#GetClientSecret) | **Get** /oauth/clients/{id}/clientAuth/clientSecret | Get the client secret of an existing OAuth client.
[**GetClients**](OauthClientsAPI.md#GetClients) | **Get** /oauth/clients | Get the list of OAuth clients.
[**GetOauthClientById**](OauthClientsAPI.md#GetOauthClientById) | **Get** /oauth/clients/{id} | Find the OAuth client by ID.
[**UpdateClientSecret**](OauthClientsAPI.md#UpdateClientSecret) | **Put** /oauth/clients/{id}/clientAuth/clientSecret | Update the client secret of an existing OAuth client.
[**UpdateOauthClient**](OauthClientsAPI.md#UpdateOauthClient) | **Put** /oauth/clients/{id} | Updates the OAuth client.



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
    resp, r, err := apiClient.OauthClientsAPI.CreateOauthClient(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientsAPI.CreateOauthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOauthClient`: Client
    fmt.Fprintf(os.Stdout, "Response from `OauthClientsAPI.CreateOauthClient`: %v\n", resp)
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
    r, err := apiClient.OauthClientsAPI.DeleteOauthClient(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientsAPI.DeleteOauthClient``: %v\n", err)
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


## GetClientSecret

> ClientSecret GetClientSecret(ctx, id).Execute()

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
    resp, r, err := apiClient.OauthClientsAPI.GetClientSecret(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientsAPI.GetClientSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientSecret`: ClientSecret
    fmt.Fprintf(os.Stdout, "Response from `OauthClientsAPI.GetClientSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientSecretRequest struct via the builder pattern


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


## GetClients

> Clients GetClients(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()

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
    resp, r, err := apiClient.OauthClientsAPI.GetClients(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientsAPI.GetClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClients`: Clients
    fmt.Fprintf(os.Stdout, "Response from `OauthClientsAPI.GetClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsRequest struct via the builder pattern


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
    resp, r, err := apiClient.OauthClientsAPI.GetOauthClientById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientsAPI.GetOauthClientById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthClientById`: Client
    fmt.Fprintf(os.Stdout, "Response from `OauthClientsAPI.GetOauthClientById`: %v\n", resp)
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


## UpdateClientSecret

> ClientSecret UpdateClientSecret(ctx, id).Body(body).Execute()

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
    resp, r, err := apiClient.OauthClientsAPI.UpdateClientSecret(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientsAPI.UpdateClientSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClientSecret`: ClientSecret
    fmt.Fprintf(os.Stdout, "Response from `OauthClientsAPI.UpdateClientSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the client to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientSecretRequest struct via the builder pattern


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
    resp, r, err := apiClient.OauthClientsAPI.UpdateOauthClient(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthClientsAPI.UpdateOauthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOauthClient`: Client
    fmt.Fprintf(os.Stdout, "Response from `OauthClientsAPI.UpdateOauthClient`: %v\n", resp)
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

