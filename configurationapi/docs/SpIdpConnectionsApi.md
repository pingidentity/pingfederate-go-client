# \SpIdpConnectionsApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConnectionCert**](SpIdpConnectionsApi.md#AddConnectionCert) | **Post** /sp/idpConnections/{id}/credentials/certs | Add a new IdP connection certificate.
[**CreateConnection**](SpIdpConnectionsApi.md#CreateConnection) | **Post** /sp/idpConnections | Create a new IdP connection.
[**DeleteConnection**](SpIdpConnectionsApi.md#DeleteConnection) | **Delete** /sp/idpConnections/{id} | Delete an IdP connection.
[**GetConnection**](SpIdpConnectionsApi.md#GetConnection) | **Get** /sp/idpConnections/{id} | Find IdP connection by ID.
[**GetConnectionCerts**](SpIdpConnectionsApi.md#GetConnectionCerts) | **Get** /sp/idpConnections/{id}/credentials/certs | Get the IdP connection&#39;s certificates.
[**GetConnections**](SpIdpConnectionsApi.md#GetConnections) | **Get** /sp/idpConnections | Get list of IdP connections.
[**GetIdpConnectionSigningSettings**](SpIdpConnectionsApi.md#GetIdpConnectionSigningSettings) | **Get** /sp/idpConnections/{id}/credentials/signingSettings | Get the IdP connection&#39;s signature settings.
[**GetIdpConnectionsDecryptionKeys**](SpIdpConnectionsApi.md#GetIdpConnectionsDecryptionKeys) | **Get** /sp/idpConnections/{id}/credentials/decryptionKeys | Get the decryption keys of an IdP connection.
[**UpdateConnection**](SpIdpConnectionsApi.md#UpdateConnection) | **Put** /sp/idpConnections/{id} | Update an IdP connection.
[**UpdateConnectionCerts**](SpIdpConnectionsApi.md#UpdateConnectionCerts) | **Put** /sp/idpConnections/{id}/credentials/certs | Update the IdP connection&#39;s certificates.
[**UpdateIdpConnectionSigningSettings**](SpIdpConnectionsApi.md#UpdateIdpConnectionSigningSettings) | **Put** /sp/idpConnections/{id}/credentials/signingSettings | Update the IdP connection&#39;s signature settings.
[**UpdateIdpConnectionsDecryptionKeys**](SpIdpConnectionsApi.md#UpdateIdpConnectionsDecryptionKeys) | **Put** /sp/idpConnections/{id}/credentials/decryptionKeys | Updating the IdP connection&#39;s decryption keys.



## AddConnectionCert

> ConnectionCert AddConnectionCert(ctx, id).Body(body).Execute()

Add a new IdP connection certificate.



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
    id := "id_example" // string | ID of the IdP Connection to update.
    body := *openapiclient.NewConnectionCert(*openapiclient.NewX509File("FileData_example")) // ConnectionCert | Configuration for a verification certificate.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpIdpConnectionsApi.AddConnectionCert(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpIdpConnectionsApi.AddConnectionCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddConnectionCert`: ConnectionCert
    fmt.Fprintf(os.Stdout, "Response from `SpIdpConnectionsApi.AddConnectionCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the IdP Connection to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddConnectionCertRequest struct via the builder pattern


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


## CreateConnection

> IdpConnection CreateConnection(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new IdP connection.



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
    body := *openapiclient.NewIdpConnection("EntityId_example", "Name_example") // IdpConnection | Configuration for new connection.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpIdpConnectionsApi.CreateConnection(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpIdpConnectionsApi.CreateConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnection`: IdpConnection
    fmt.Fprintf(os.Stdout, "Response from `SpIdpConnectionsApi.CreateConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IdpConnection**](IdpConnection.md) | Configuration for new connection. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**IdpConnection**](IdpConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnection

> DeleteConnection(ctx, id).Execute()

Delete an IdP connection.



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
    id := "id_example" // string | ID of IdP Connection to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpIdpConnectionsApi.DeleteConnection(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpIdpConnectionsApi.DeleteConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of IdP Connection to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectionRequest struct via the builder pattern


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


## GetConnection

> IdpConnection GetConnection(ctx, id).Execute()

Find IdP connection by ID.



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
    id := "id_example" // string | ID of IdP Connection to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpIdpConnectionsApi.GetConnection(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpIdpConnectionsApi.GetConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnection`: IdpConnection
    fmt.Fprintf(os.Stdout, "Response from `SpIdpConnectionsApi.GetConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of IdP Connection to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdpConnection**](IdpConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionCerts

> ConnectionCerts GetConnectionCerts(ctx, id).Execute()

Get the IdP connection's certificates.

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
    id := "id_example" // string | ID of the IdP Connection.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpIdpConnectionsApi.GetConnectionCerts(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpIdpConnectionsApi.GetConnectionCerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectionCerts`: ConnectionCerts
    fmt.Fprintf(os.Stdout, "Response from `SpIdpConnectionsApi.GetConnectionCerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the IdP Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionCertsRequest struct via the builder pattern


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


## GetConnections

> IdpConnections GetConnections(ctx).EntityId(entityId).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()

Get list of IdP connections.



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
    filter := "filter_example" // string | Filter criteria limits the IdP connections that are returned to only those that match it. The filter criteria is compared to the IdP connection name and partner entity ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpIdpConnectionsApi.GetConnections(context.Background()).EntityId(entityId).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpIdpConnectionsApi.GetConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnections`: IdpConnections
    fmt.Fprintf(os.Stdout, "Response from `SpIdpConnectionsApi.GetConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | Entity ID of the connection to fetch. (case-sensitive) | 
 **page** | **int64** | Page number to retrieve. | 
 **numberPerPage** | **int64** | Number of connections per page. | 
 **filter** | **string** | Filter criteria limits the IdP connections that are returned to only those that match it. The filter criteria is compared to the IdP connection name and partner entity ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. | 

### Return type

[**IdpConnections**](IdpConnections.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdpConnectionSigningSettings

> SigningSettings GetIdpConnectionSigningSettings(ctx, id).Execute()

Get the IdP connection's signature settings.

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
    id := "id_example" // string | ID of the IdP Connection.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpIdpConnectionsApi.GetIdpConnectionSigningSettings(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpIdpConnectionsApi.GetIdpConnectionSigningSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdpConnectionSigningSettings`: SigningSettings
    fmt.Fprintf(os.Stdout, "Response from `SpIdpConnectionsApi.GetIdpConnectionSigningSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the IdP Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpConnectionSigningSettingsRequest struct via the builder pattern


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


## GetIdpConnectionsDecryptionKeys

> DecryptionKeys GetIdpConnectionsDecryptionKeys(ctx, id).Execute()

Get the decryption keys of an IdP connection.

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
    id := "id_example" // string | ID of the IdP Connection to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpIdpConnectionsApi.GetIdpConnectionsDecryptionKeys(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpIdpConnectionsApi.GetIdpConnectionsDecryptionKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdpConnectionsDecryptionKeys`: DecryptionKeys
    fmt.Fprintf(os.Stdout, "Response from `SpIdpConnectionsApi.GetIdpConnectionsDecryptionKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the IdP Connection to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpConnectionsDecryptionKeysRequest struct via the builder pattern


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


## UpdateConnection

> IdpConnection UpdateConnection(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update an IdP connection.



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
    id := "id_example" // string | ID of IdP Connection to update.
    body := *openapiclient.NewIdpConnection("EntityId_example", "Name_example") // IdpConnection | Configuration for updated connection.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpIdpConnectionsApi.UpdateConnection(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpIdpConnectionsApi.UpdateConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnection`: IdpConnection
    fmt.Fprintf(os.Stdout, "Response from `SpIdpConnectionsApi.UpdateConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of IdP Connection to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IdpConnection**](IdpConnection.md) | Configuration for updated connection. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**IdpConnection**](IdpConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectionCerts

> ConnectionCerts UpdateConnectionCerts(ctx, id).Body(body).Execute()

Update the IdP connection's certificates.

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
    id := "id_example" // string | ID of the IdP Connection to update.
    body := *openapiclient.NewConnectionCerts() // ConnectionCerts | Configuration for a verification certificates.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpIdpConnectionsApi.UpdateConnectionCerts(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpIdpConnectionsApi.UpdateConnectionCerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnectionCerts`: ConnectionCerts
    fmt.Fprintf(os.Stdout, "Response from `SpIdpConnectionsApi.UpdateConnectionCerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the IdP Connection to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectionCertsRequest struct via the builder pattern


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


## UpdateIdpConnectionSigningSettings

> SigningSettings UpdateIdpConnectionSigningSettings(ctx, id).Body(body).Execute()

Update the IdP connection's signature settings.

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
    id := "id_example" // string | ID of the IdP Connection to update.
    body := *openapiclient.NewSigningSettings(*openapiclient.NewResourceLink("Id_example")) // SigningSettings | Signature settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpIdpConnectionsApi.UpdateIdpConnectionSigningSettings(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpIdpConnectionsApi.UpdateIdpConnectionSigningSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdpConnectionSigningSettings`: SigningSettings
    fmt.Fprintf(os.Stdout, "Response from `SpIdpConnectionsApi.UpdateIdpConnectionSigningSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the IdP Connection to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdpConnectionSigningSettingsRequest struct via the builder pattern


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


## UpdateIdpConnectionsDecryptionKeys

> DecryptionKeys UpdateIdpConnectionsDecryptionKeys(ctx, id).Body(body).Execute()

Updating the IdP connection's decryption keys.

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
    id := "id_example" // string | ID of the IdP Connection to update.
    body := *openapiclient.NewDecryptionKeys() // DecryptionKeys | Configuration for decryption keys.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpIdpConnectionsApi.UpdateIdpConnectionsDecryptionKeys(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpIdpConnectionsApi.UpdateIdpConnectionsDecryptionKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdpConnectionsDecryptionKeys`: DecryptionKeys
    fmt.Fprintf(os.Stdout, "Response from `SpIdpConnectionsApi.UpdateIdpConnectionsDecryptionKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the IdP Connection to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdpConnectionsDecryptionKeysRequest struct via the builder pattern


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

