# \IdpSpConnectionsApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSpConnectionCert**](IdpSpConnectionsApi.md#AddSpConnectionCert) | **Post** /idp/spConnections/{id}/credentials/certs | Add a new SP connection certificate.
[**CreateSpConnection**](IdpSpConnectionsApi.md#CreateSpConnection) | **Post** /idp/spConnections | Create a new SP connection.
[**DeleteSpConnection**](IdpSpConnectionsApi.md#DeleteSpConnection) | **Delete** /idp/spConnections/{id} | Delete an SP connection.
[**GetDecryptionKeys**](IdpSpConnectionsApi.md#GetDecryptionKeys) | **Get** /idp/spConnections/{id}/credentials/decryptionKeys | Get the decryption keys of an SP connection.
[**GetSpConnection**](IdpSpConnectionsApi.md#GetSpConnection) | **Get** /idp/spConnections/{id} | Find SP connection by ID.
[**GetSpConnectionCerts**](IdpSpConnectionsApi.md#GetSpConnectionCerts) | **Get** /idp/spConnections/{id}/credentials/certs | Get the SP connection&#39;s certificates.
[**GetSpConnections**](IdpSpConnectionsApi.md#GetSpConnections) | **Get** /idp/spConnections | Get list of SP connections.
[**GetSpSigningSettings**](IdpSpConnectionsApi.md#GetSpSigningSettings) | **Get** /idp/spConnections/{id}/credentials/signingSettings | Get the SP connection&#39;s signature settings.
[**UpdateDecryptionKeys**](IdpSpConnectionsApi.md#UpdateDecryptionKeys) | **Put** /idp/spConnections/{id}/credentials/decryptionKeys | Updating the SP connection&#39;s decryption keys.
[**UpdateSpConnection**](IdpSpConnectionsApi.md#UpdateSpConnection) | **Put** /idp/spConnections/{id} | Update an SP connection.
[**UpdateSpConnectionCerts**](IdpSpConnectionsApi.md#UpdateSpConnectionCerts) | **Put** /idp/spConnections/{id}/credentials/certs | Update the SP connection&#39;s certificates.
[**UpdateSpSigningSettings**](IdpSpConnectionsApi.md#UpdateSpSigningSettings) | **Put** /idp/spConnections/{id}/credentials/signingSettings | Update the SP connection&#39;s signature settings.



## AddSpConnectionCert

> ConnectionCert AddSpConnectionCert(ctx, id).Body(body).Execute()

Add a new SP connection certificate.



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
    id := "id_example" // string | ID of the SP Connection to update.
    body := *openapiclient.NewConnectionCert(*openapiclient.NewX509File("FileData_example")) // ConnectionCert | Configuration for a verification certificate.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpSpConnectionsApi.AddSpConnectionCert(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpSpConnectionsApi.AddSpConnectionCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSpConnectionCert`: ConnectionCert
    fmt.Fprintf(os.Stdout, "Response from `IdpSpConnectionsApi.AddSpConnectionCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the SP Connection to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSpConnectionCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConnectionCert**](ConnectionCert.md) | Configuration for a verification certificate. | 

### Return type

[**ConnectionCert**](ConnectionCert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSpConnection

> SpConnection CreateSpConnection(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new SP connection.



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
    body := *openapiclient.NewSpConnection("EntityId_example", "Name_example") // SpConnection | Configuration for new connection.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpSpConnectionsApi.CreateSpConnection(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpSpConnectionsApi.CreateSpConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSpConnection`: SpConnection
    fmt.Fprintf(os.Stdout, "Response from `IdpSpConnectionsApi.CreateSpConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SpConnection**](SpConnection.md) | Configuration for new connection. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**SpConnection**](SpConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSpConnection

> DeleteSpConnection(ctx, id).Execute()

Delete an SP connection.



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
    id := "id_example" // string | ID of SP Connection to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdpSpConnectionsApi.DeleteSpConnection(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpSpConnectionsApi.DeleteSpConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of SP Connection to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpConnectionRequest struct via the builder pattern


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


## GetDecryptionKeys

> DecryptionKeys GetDecryptionKeys(ctx, id).Execute()

Get the decryption keys of an SP connection.

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
    id := "id_example" // string | ID of the SP Connection to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpSpConnectionsApi.GetDecryptionKeys(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpSpConnectionsApi.GetDecryptionKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDecryptionKeys`: DecryptionKeys
    fmt.Fprintf(os.Stdout, "Response from `IdpSpConnectionsApi.GetDecryptionKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the SP Connection to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecryptionKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DecryptionKeys**](DecryptionKeys.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpConnection

> SpConnection GetSpConnection(ctx, id).Execute()

Find SP connection by ID.



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
    id := "id_example" // string | ID of the SP Connection to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpSpConnectionsApi.GetSpConnection(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpSpConnectionsApi.GetSpConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpConnection`: SpConnection
    fmt.Fprintf(os.Stdout, "Response from `IdpSpConnectionsApi.GetSpConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the SP Connection to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpConnection**](SpConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpConnectionCerts

> ConnectionCerts GetSpConnectionCerts(ctx, id).Execute()

Get the SP connection's certificates.

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
    id := "id_example" // string | ID of the SP Connection.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpSpConnectionsApi.GetSpConnectionCerts(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpSpConnectionsApi.GetSpConnectionCerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpConnectionCerts`: ConnectionCerts
    fmt.Fprintf(os.Stdout, "Response from `IdpSpConnectionsApi.GetSpConnectionCerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the SP Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpConnectionCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectionCerts**](ConnectionCerts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpConnections

> SpConnections GetSpConnections(ctx).EntityId(entityId).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()

Get list of SP connections.



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
    entityId := "entityId_example" // string | Entity ID of the connection to fetch. (case-sensitive) (optional)
    page := int64(56) // int64 | Page number to retrieve. (optional)
    numberPerPage := int64(56) // int64 | Number of connections per page. (optional)
    filter := "filter_example" // string | Filter criteria limits the SP connections that are returned to only those that match it. The filter criteria is compared to the SP connection name and partner entity ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpSpConnectionsApi.GetSpConnections(context.Background()).EntityId(entityId).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpSpConnectionsApi.GetSpConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpConnections`: SpConnections
    fmt.Fprintf(os.Stdout, "Response from `IdpSpConnectionsApi.GetSpConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSpConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | Entity ID of the connection to fetch. (case-sensitive) | 
 **page** | **int64** | Page number to retrieve. | 
 **numberPerPage** | **int64** | Number of connections per page. | 
 **filter** | **string** | Filter criteria limits the SP connections that are returned to only those that match it. The filter criteria is compared to the SP connection name and partner entity ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. | 

### Return type

[**SpConnections**](SpConnections.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpSigningSettings

> SigningSettings GetSpSigningSettings(ctx, id).Execute()

Get the SP connection's signature settings.

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
    id := "id_example" // string | ID of the SP Connection.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpSpConnectionsApi.GetSpSigningSettings(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpSpConnectionsApi.GetSpSigningSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpSigningSettings`: SigningSettings
    fmt.Fprintf(os.Stdout, "Response from `IdpSpConnectionsApi.GetSpSigningSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the SP Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpSigningSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SigningSettings**](SigningSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDecryptionKeys

> DecryptionKeys UpdateDecryptionKeys(ctx, id).Body(body).Execute()

Updating the SP connection's decryption keys.

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
    id := "id_example" // string | ID of the SP Connection to update.
    body := *openapiclient.NewDecryptionKeys() // DecryptionKeys | Configuration for decryption keys.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpSpConnectionsApi.UpdateDecryptionKeys(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpSpConnectionsApi.UpdateDecryptionKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDecryptionKeys`: DecryptionKeys
    fmt.Fprintf(os.Stdout, "Response from `IdpSpConnectionsApi.UpdateDecryptionKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the SP Connection to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDecryptionKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DecryptionKeys**](DecryptionKeys.md) | Configuration for decryption keys. | 

### Return type

[**DecryptionKeys**](DecryptionKeys.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpConnection

> SpConnection UpdateSpConnection(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update an SP connection.



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
    id := "id_example" // string | ID of SP Connection to update.
    body := *openapiclient.NewSpConnection("EntityId_example", "Name_example") // SpConnection | Configuration for updated connection.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpSpConnectionsApi.UpdateSpConnection(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpSpConnectionsApi.UpdateSpConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSpConnection`: SpConnection
    fmt.Fprintf(os.Stdout, "Response from `IdpSpConnectionsApi.UpdateSpConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of SP Connection to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SpConnection**](SpConnection.md) | Configuration for updated connection. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**SpConnection**](SpConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpConnectionCerts

> ConnectionCerts UpdateSpConnectionCerts(ctx, id).Body(body).Execute()

Update the SP connection's certificates.

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
    id := "id_example" // string | ID of the SP Connection to update.
    body := *openapiclient.NewConnectionCerts() // ConnectionCerts | Configuration for a verification certificates.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpSpConnectionsApi.UpdateSpConnectionCerts(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpSpConnectionsApi.UpdateSpConnectionCerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSpConnectionCerts`: ConnectionCerts
    fmt.Fprintf(os.Stdout, "Response from `IdpSpConnectionsApi.UpdateSpConnectionCerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the SP Connection to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpConnectionCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConnectionCerts**](ConnectionCerts.md) | Configuration for a verification certificates. | 

### Return type

[**ConnectionCerts**](ConnectionCerts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpSigningSettings

> SigningSettings UpdateSpSigningSettings(ctx, id).Body(body).Execute()

Update the SP connection's signature settings.

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
    id := "id_example" // string | ID of the SP Connection to update.
    body := *openapiclient.NewSigningSettings(*openapiclient.NewResourceLink("Id_example")) // SigningSettings | Signature settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpSpConnectionsApi.UpdateSpSigningSettings(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpSpConnectionsApi.UpdateSpSigningSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSpSigningSettings`: SigningSettings
    fmt.Fprintf(os.Stdout, "Response from `IdpSpConnectionsApi.UpdateSpSigningSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the SP Connection to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpSigningSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SigningSettings**](SigningSettings.md) | Signature settings. | 

### Return type

[**SigningSettings**](SigningSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
