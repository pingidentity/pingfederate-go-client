# \KeyPairsSslServerAPI

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSslServerKeyPair**](KeyPairsSslServerAPI.md#CreateSslServerKeyPair) | **Post** /keyPairs/sslServer/generate | Generate a new key pair.
[**DeleteSslServerKeyPair**](KeyPairsSslServerAPI.md#DeleteSslServerKeyPair) | **Delete** /keyPairs/sslServer/{id} | Delete a key pair.
[**ExportSslServerCertificateFile**](KeyPairsSslServerAPI.md#ExportSslServerCertificateFile) | **Get** /keyPairs/sslServer/{id}/certificate | Download the certificate from a given key pair.
[**ExportSslServerCsr**](KeyPairsSslServerAPI.md#ExportSslServerCsr) | **Get** /keyPairs/sslServer/{id}/csr | Generate a new certificate signing request (CSR) for this key pair.
[**ExportSslServerPEMFile**](KeyPairsSslServerAPI.md#ExportSslServerPEMFile) | **Post** /keyPairs/sslServer/{id}/pem | Download the key pair in PEM format.
[**ExportSslServerPKCS12File**](KeyPairsSslServerAPI.md#ExportSslServerPKCS12File) | **Post** /keyPairs/sslServer/{id}/pkcs12 | Download the key pair in PKCS12 format.
[**GetSslServerKeyPair**](KeyPairsSslServerAPI.md#GetSslServerKeyPair) | **Get** /keyPairs/sslServer/{id} | Retrieve details of a key pair.
[**GetSslServerKeyPairs**](KeyPairsSslServerAPI.md#GetSslServerKeyPairs) | **Get** /keyPairs/sslServer | Get list of key pairs.
[**GetSslServerSettings**](KeyPairsSslServerAPI.md#GetSslServerSettings) | **Get** /keyPairs/sslServer/settings | Get the SSL Server Certificate Settings.
[**ImportSslServerCsrResponse**](KeyPairsSslServerAPI.md#ImportSslServerCsrResponse) | **Post** /keyPairs/sslServer/{id}/csr | Import a CSR response for this key pair.
[**ImportSslServerKeyPair**](KeyPairsSslServerAPI.md#ImportSslServerKeyPair) | **Post** /keyPairs/sslServer/import | Import a new key pair.
[**UpdateSslServerSettings**](KeyPairsSslServerAPI.md#UpdateSslServerSettings) | **Put** /keyPairs/sslServer/settings | Update the SSL Server Certificate Settings.



## CreateSslServerKeyPair

> KeyPairView CreateSslServerKeyPair(ctx).Body(body).Execute()

Generate a new key pair.

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
    body := *openapiclient.NewNewKeyPairSettings("CommonName_example", "Organization_example", "Country_example", int64(123), "KeyAlgorithm_example") // NewKeyPairSettings | Settings for the new key pair.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsSslServerAPI.CreateSslServerKeyPair(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslServerAPI.CreateSslServerKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSslServerKeyPair`: KeyPairView
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslServerAPI.CreateSslServerKeyPair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSslServerKeyPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewKeyPairSettings**](NewKeyPairSettings.md) | Settings for the new key pair. | 

### Return type

[**KeyPairView**](KeyPairView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSslServerKeyPair

> DeleteSslServerKeyPair(ctx, id).Execute()

Delete a key pair.



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
    id := "id_example" // string | ID of the key pair to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeyPairsSslServerAPI.DeleteSslServerKeyPair(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslServerAPI.DeleteSslServerKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSslServerKeyPairRequest struct via the builder pattern


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


## ExportSslServerCertificateFile

> string ExportSslServerCertificateFile(ctx, id).Execute()

Download the certificate from a given key pair.



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
    id := "id_example" // string | ID of the key pair.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsSslServerAPI.ExportSslServerCertificateFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslServerAPI.ExportSslServerCertificateFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSslServerCertificateFile`: string
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslServerAPI.ExportSslServerCertificateFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSslServerCertificateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSslServerCsr

> string ExportSslServerCsr(ctx, id).Execute()

Generate a new certificate signing request (CSR) for this key pair.



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
    id := "id_example" // string | ID of the key pair.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsSslServerAPI.ExportSslServerCsr(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslServerAPI.ExportSslServerCsr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSslServerCsr`: string
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslServerAPI.ExportSslServerCsr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSslServerCsrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSslServerPEMFile

> string ExportSslServerPEMFile(ctx, id).Body(body).Execute()

Download the key pair in PEM format.



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
    id := "id_example" // string | ID of the key pair.
    body := *openapiclient.NewKeyPairExportSettings("Password_example") // KeyPairExportSettings | Parameters for the export request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsSslServerAPI.ExportSslServerPEMFile(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslServerAPI.ExportSslServerPEMFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSslServerPEMFile`: string
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslServerAPI.ExportSslServerPEMFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSslServerPEMFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**KeyPairExportSettings**](KeyPairExportSettings.md) | Parameters for the export request | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSslServerPKCS12File

> string ExportSslServerPKCS12File(ctx, id).Body(body).Execute()

Download the key pair in PKCS12 format.



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
    id := "id_example" // string | ID of the key pair.
    body := *openapiclient.NewKeyPairExportSettings("Password_example") // KeyPairExportSettings | Parameters for the export request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsSslServerAPI.ExportSslServerPKCS12File(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslServerAPI.ExportSslServerPKCS12File``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSslServerPKCS12File`: string
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslServerAPI.ExportSslServerPKCS12File`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSslServerPKCS12FileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**KeyPairExportSettings**](KeyPairExportSettings.md) | Parameters for the export request | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSslServerKeyPair

> KeyPairView GetSslServerKeyPair(ctx, id).Execute()

Retrieve details of a key pair.

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
    id := "id_example" // string | ID of the key pair to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsSslServerAPI.GetSslServerKeyPair(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslServerAPI.GetSslServerKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSslServerKeyPair`: KeyPairView
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslServerAPI.GetSslServerKeyPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSslServerKeyPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KeyPairView**](KeyPairView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSslServerKeyPairs

> KeyPairViews GetSslServerKeyPairs(ctx).Execute()

Get list of key pairs.

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
    resp, r, err := apiClient.KeyPairsSslServerAPI.GetSslServerKeyPairs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslServerAPI.GetSslServerKeyPairs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSslServerKeyPairs`: KeyPairViews
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslServerAPI.GetSslServerKeyPairs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSslServerKeyPairsRequest struct via the builder pattern


### Return type

[**KeyPairViews**](KeyPairViews.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSslServerSettings

> SslServerSettings GetSslServerSettings(ctx).Execute()

Get the SSL Server Certificate Settings.

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
    resp, r, err := apiClient.KeyPairsSslServerAPI.GetSslServerSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslServerAPI.GetSslServerSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSslServerSettings`: SslServerSettings
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslServerAPI.GetSslServerSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSslServerSettingsRequest struct via the builder pattern


### Return type

[**SslServerSettings**](SslServerSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSslServerCsrResponse

> KeyPairView ImportSslServerCsrResponse(ctx, id).Body(body).Execute()

Import a CSR response for this key pair.

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
    id := "id_example" // string | ID of the key pair.
    body := *openapiclient.NewCSRResponse("FileData_example") // CSRResponse | The CSR response.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsSslServerAPI.ImportSslServerCsrResponse(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslServerAPI.ImportSslServerCsrResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportSslServerCsrResponse`: KeyPairView
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslServerAPI.ImportSslServerCsrResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportSslServerCsrResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CSRResponse**](CSRResponse.md) | The CSR response. | 

### Return type

[**KeyPairView**](KeyPairView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSslServerKeyPair

> KeyPairView ImportSslServerKeyPair(ctx).Body(body).Execute()

Import a new key pair.

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
    body := *openapiclient.NewKeyPairFile("FileData_example", "Password_example") // KeyPairFile | File to import.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsSslServerAPI.ImportSslServerKeyPair(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslServerAPI.ImportSslServerKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportSslServerKeyPair`: KeyPairView
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslServerAPI.ImportSslServerKeyPair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportSslServerKeyPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**KeyPairFile**](KeyPairFile.md) | File to import. | 

### Return type

[**KeyPairView**](KeyPairView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSslServerSettings

> SslServerSettings UpdateSslServerSettings(ctx).Body(body).Execute()

Update the SSL Server Certificate Settings.

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
    body := *openapiclient.NewSslServerSettings(*openapiclient.NewResourceLink("Id_example"), *openapiclient.NewResourceLink("Id_example")) // SslServerSettings | Configuration for activation of SSL server certificates.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsSslServerAPI.UpdateSslServerSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslServerAPI.UpdateSslServerSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSslServerSettings`: SslServerSettings
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslServerAPI.UpdateSslServerSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSslServerSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SslServerSettings**](SslServerSettings.md) | Configuration for activation of SSL server certificates. | 

### Return type

[**SslServerSettings**](SslServerSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

