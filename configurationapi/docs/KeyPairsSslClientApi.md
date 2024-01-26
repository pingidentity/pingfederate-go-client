# \KeyPairsSslClientAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSslClientKeyPair**](KeyPairsSslClientAPI.md#CreateSslClientKeyPair) | **Post** /keyPairs/sslClient/generate | Generate a new key pair.
[**DeleteSslClientKeyPair**](KeyPairsSslClientAPI.md#DeleteSslClientKeyPair) | **Delete** /keyPairs/sslClient/{id} | Delete a key pair.
[**ExportSslClientCertificateFile**](KeyPairsSslClientAPI.md#ExportSslClientCertificateFile) | **Get** /keyPairs/sslClient/{id}/certificate | Download the certificate from a given key pair.
[**ExportSslClientCsr**](KeyPairsSslClientAPI.md#ExportSslClientCsr) | **Get** /keyPairs/sslClient/{id}/csr | Generate a new certificate signing request (CSR) for this key pair.
[**ExportSslClientPEMFile**](KeyPairsSslClientAPI.md#ExportSslClientPEMFile) | **Post** /keyPairs/sslClient/{id}/pem | Download the key pair in PEM format.
[**ExportSslClientPKCS12File**](KeyPairsSslClientAPI.md#ExportSslClientPKCS12File) | **Post** /keyPairs/sslClient/{id}/pkcs12 | Download the key pair in PKCS12 format.
[**GetSslClientKeyPair**](KeyPairsSslClientAPI.md#GetSslClientKeyPair) | **Get** /keyPairs/sslClient/{id} | Retrieve details of a key pair.
[**GetSslClientKeyPairs**](KeyPairsSslClientAPI.md#GetSslClientKeyPairs) | **Get** /keyPairs/sslClient | Get list of key pairs.
[**ImportSslClientCsrResponse**](KeyPairsSslClientAPI.md#ImportSslClientCsrResponse) | **Post** /keyPairs/sslClient/{id}/csr | Import a CSR response for this key pair.
[**ImportSslClientKeyPair**](KeyPairsSslClientAPI.md#ImportSslClientKeyPair) | **Post** /keyPairs/sslClient/import | Import a new key pair.



## CreateSslClientKeyPair

> KeyPairView CreateSslClientKeyPair(ctx).Body(body).Execute()

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
    resp, r, err := apiClient.KeyPairsSslClientAPI.CreateSslClientKeyPair(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslClientAPI.CreateSslClientKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSslClientKeyPair`: KeyPairView
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslClientAPI.CreateSslClientKeyPair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSslClientKeyPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewKeyPairSettings**](NewKeyPairSettings.md) | Settings for the new key pair. | 

### Return type

[**KeyPairView**](KeyPairView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSslClientKeyPair

> DeleteSslClientKeyPair(ctx, id).Execute()

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
    r, err := apiClient.KeyPairsSslClientAPI.DeleteSslClientKeyPair(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslClientAPI.DeleteSslClientKeyPair``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSslClientKeyPairRequest struct via the builder pattern


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


## ExportSslClientCertificateFile

> string ExportSslClientCertificateFile(ctx, id).Execute()

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
    resp, r, err := apiClient.KeyPairsSslClientAPI.ExportSslClientCertificateFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslClientAPI.ExportSslClientCertificateFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSslClientCertificateFile`: string
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslClientAPI.ExportSslClientCertificateFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSslClientCertificateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSslClientCsr

> string ExportSslClientCsr(ctx, id).Execute()

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
    resp, r, err := apiClient.KeyPairsSslClientAPI.ExportSslClientCsr(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslClientAPI.ExportSslClientCsr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSslClientCsr`: string
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslClientAPI.ExportSslClientCsr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSslClientCsrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSslClientPEMFile

> string ExportSslClientPEMFile(ctx, id).Body(body).Execute()

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
    resp, r, err := apiClient.KeyPairsSslClientAPI.ExportSslClientPEMFile(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslClientAPI.ExportSslClientPEMFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSslClientPEMFile`: string
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslClientAPI.ExportSslClientPEMFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSslClientPEMFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**KeyPairExportSettings**](KeyPairExportSettings.md) | Parameters for the export request | 

### Return type

**string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSslClientPKCS12File

> string ExportSslClientPKCS12File(ctx, id).Body(body).Execute()

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
    resp, r, err := apiClient.KeyPairsSslClientAPI.ExportSslClientPKCS12File(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslClientAPI.ExportSslClientPKCS12File``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSslClientPKCS12File`: string
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslClientAPI.ExportSslClientPKCS12File`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSslClientPKCS12FileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**KeyPairExportSettings**](KeyPairExportSettings.md) | Parameters for the export request | 

### Return type

**string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSslClientKeyPair

> KeyPairView GetSslClientKeyPair(ctx, id).Execute()

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
    resp, r, err := apiClient.KeyPairsSslClientAPI.GetSslClientKeyPair(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslClientAPI.GetSslClientKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSslClientKeyPair`: KeyPairView
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslClientAPI.GetSslClientKeyPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSslClientKeyPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KeyPairView**](KeyPairView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSslClientKeyPairs

> KeyPairViews GetSslClientKeyPairs(ctx).Execute()

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
    resp, r, err := apiClient.KeyPairsSslClientAPI.GetSslClientKeyPairs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslClientAPI.GetSslClientKeyPairs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSslClientKeyPairs`: KeyPairViews
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslClientAPI.GetSslClientKeyPairs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSslClientKeyPairsRequest struct via the builder pattern


### Return type

[**KeyPairViews**](KeyPairViews.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSslClientCsrResponse

> KeyPairView ImportSslClientCsrResponse(ctx, id).Body(body).Execute()

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
    resp, r, err := apiClient.KeyPairsSslClientAPI.ImportSslClientCsrResponse(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslClientAPI.ImportSslClientCsrResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportSslClientCsrResponse`: KeyPairView
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslClientAPI.ImportSslClientCsrResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportSslClientCsrResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CSRResponse**](CSRResponse.md) | The CSR response. | 

### Return type

[**KeyPairView**](KeyPairView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSslClientKeyPair

> KeyPairView ImportSslClientKeyPair(ctx).Body(body).Execute()

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
    resp, r, err := apiClient.KeyPairsSslClientAPI.ImportSslClientKeyPair(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSslClientAPI.ImportSslClientKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportSslClientKeyPair`: KeyPairView
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSslClientAPI.ImportSslClientKeyPair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportSslClientKeyPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**KeyPairFile**](KeyPairFile.md) | File to import. | 

### Return type

[**KeyPairView**](KeyPairView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

