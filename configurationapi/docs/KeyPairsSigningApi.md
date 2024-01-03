# \KeyPairsSigningAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSigningKeyPair**](KeyPairsSigningAPI.md#CreateSigningKeyPair) | **Post** /keyPairs/signing/generate | Generate a new key pair.
[**DeleteKeyPairRotationSettings**](KeyPairsSigningAPI.md#DeleteKeyPairRotationSettings) | **Delete** /keyPairs/signing/{id}/rotationSettings | Delete rotation settings for a signing key pair.
[**DeleteSigningKeyPair**](KeyPairsSigningAPI.md#DeleteSigningKeyPair) | **Delete** /keyPairs/signing/{id} | Delete a key pair.
[**ExportCertificateFile**](KeyPairsSigningAPI.md#ExportCertificateFile) | **Get** /keyPairs/signing/{id}/certificate | Download the certificate from a given key pair.
[**ExportCsr**](KeyPairsSigningAPI.md#ExportCsr) | **Get** /keyPairs/signing/{id}/csr | Generate a new certificate signing request (CSR) for this key pair.
[**ExportPEMFile**](KeyPairsSigningAPI.md#ExportPEMFile) | **Post** /keyPairs/signing/{id}/pem | Download the key pair in PEM format.
[**ExportPKCS12File**](KeyPairsSigningAPI.md#ExportPKCS12File) | **Post** /keyPairs/signing/{id}/pkcs12 | Download the key pair in PKCS12 format.
[**GetRotationSettings**](KeyPairsSigningAPI.md#GetRotationSettings) | **Get** /keyPairs/signing/{id}/rotationSettings | Retrieve details of rotation settings for a key pair.
[**GetSigningKeyPair**](KeyPairsSigningAPI.md#GetSigningKeyPair) | **Get** /keyPairs/signing/{id} | Retrieve details of a key pair.
[**GetSigningKeyPairs**](KeyPairsSigningAPI.md#GetSigningKeyPairs) | **Get** /keyPairs/signing | Get list of key pairs.
[**ImportCsrResponse**](KeyPairsSigningAPI.md#ImportCsrResponse) | **Post** /keyPairs/signing/{id}/csr | Import a CSR response for this key pair.
[**ImportSigningKeyPair**](KeyPairsSigningAPI.md#ImportSigningKeyPair) | **Post** /keyPairs/signing/import | Import a new key pair.
[**UpdateRotationSettings**](KeyPairsSigningAPI.md#UpdateRotationSettings) | **Put** /keyPairs/signing/{id}/rotationSettings | Add rotation settings to a key pair



## CreateSigningKeyPair

> KeyPairView CreateSigningKeyPair(ctx).Body(body).Execute()

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
    resp, r, err := apiClient.KeyPairsSigningAPI.CreateSigningKeyPair(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSigningAPI.CreateSigningKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSigningKeyPair`: KeyPairView
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSigningAPI.CreateSigningKeyPair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSigningKeyPairRequest struct via the builder pattern


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


## DeleteKeyPairRotationSettings

> DeleteKeyPairRotationSettings(ctx, id).Execute()

Delete rotation settings for a signing key pair.



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
    id := "id_example" // string | ID of the key pair's rotation settings to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeyPairsSigningAPI.DeleteKeyPairRotationSettings(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSigningAPI.DeleteKeyPairRotationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair&#39;s rotation settings to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyPairRotationSettingsRequest struct via the builder pattern


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


## DeleteSigningKeyPair

> DeleteSigningKeyPair(ctx, id).Execute()

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
    r, err := apiClient.KeyPairsSigningAPI.DeleteSigningKeyPair(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSigningAPI.DeleteSigningKeyPair``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSigningKeyPairRequest struct via the builder pattern


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


## ExportCertificateFile

> string ExportCertificateFile(ctx, id).Execute()

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
    resp, r, err := apiClient.KeyPairsSigningAPI.ExportCertificateFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSigningAPI.ExportCertificateFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportCertificateFile`: string
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSigningAPI.ExportCertificateFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCertificateFileRequest struct via the builder pattern


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


## ExportCsr

> string ExportCsr(ctx, id).Execute()

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
    resp, r, err := apiClient.KeyPairsSigningAPI.ExportCsr(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSigningAPI.ExportCsr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportCsr`: string
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSigningAPI.ExportCsr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCsrRequest struct via the builder pattern


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


## ExportPEMFile

> string ExportPEMFile(ctx, id).Body(body).Execute()

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
    resp, r, err := apiClient.KeyPairsSigningAPI.ExportPEMFile(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSigningAPI.ExportPEMFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportPEMFile`: string
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSigningAPI.ExportPEMFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPEMFileRequest struct via the builder pattern


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


## ExportPKCS12File

> string ExportPKCS12File(ctx, id).Body(body).Execute()

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
    resp, r, err := apiClient.KeyPairsSigningAPI.ExportPKCS12File(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSigningAPI.ExportPKCS12File``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportPKCS12File`: string
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSigningAPI.ExportPKCS12File`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPKCS12FileRequest struct via the builder pattern


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


## GetRotationSettings

> KeyPairRotationSettings GetRotationSettings(ctx, id).Execute()

Retrieve details of rotation settings for a key pair.

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
    id := "id_example" // string | ID of the key pair to retrieve its rotation settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsSigningAPI.GetRotationSettings(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSigningAPI.GetRotationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRotationSettings`: KeyPairRotationSettings
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSigningAPI.GetRotationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair to retrieve its rotation settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRotationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KeyPairRotationSettings**](KeyPairRotationSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSigningKeyPair

> KeyPairView GetSigningKeyPair(ctx, id).Execute()

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
    resp, r, err := apiClient.KeyPairsSigningAPI.GetSigningKeyPair(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSigningAPI.GetSigningKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSigningKeyPair`: KeyPairView
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSigningAPI.GetSigningKeyPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSigningKeyPairRequest struct via the builder pattern


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


## GetSigningKeyPairs

> KeyPairViews GetSigningKeyPairs(ctx).Execute()

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
    resp, r, err := apiClient.KeyPairsSigningAPI.GetSigningKeyPairs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSigningAPI.GetSigningKeyPairs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSigningKeyPairs`: KeyPairViews
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSigningAPI.GetSigningKeyPairs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSigningKeyPairsRequest struct via the builder pattern


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


## ImportCsrResponse

> KeyPairView ImportCsrResponse(ctx, id).Body(body).Execute()

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
    resp, r, err := apiClient.KeyPairsSigningAPI.ImportCsrResponse(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSigningAPI.ImportCsrResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportCsrResponse`: KeyPairView
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSigningAPI.ImportCsrResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportCsrResponseRequest struct via the builder pattern


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


## ImportSigningKeyPair

> KeyPairView ImportSigningKeyPair(ctx).Body(body).Execute()

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
    resp, r, err := apiClient.KeyPairsSigningAPI.ImportSigningKeyPair(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSigningAPI.ImportSigningKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportSigningKeyPair`: KeyPairView
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSigningAPI.ImportSigningKeyPair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportSigningKeyPairRequest struct via the builder pattern


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


## UpdateRotationSettings

> KeyPairRotationSettings UpdateRotationSettings(ctx, id).Body(body).Execute()

Add rotation settings to a key pair

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
    body := *openapiclient.NewKeyPairRotationSettings(int64(123), int64(123)) // KeyPairRotationSettings | The key rotation settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsSigningAPI.UpdateRotationSettings(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsSigningAPI.UpdateRotationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRotationSettings`: KeyPairRotationSettings
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsSigningAPI.UpdateRotationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the key pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRotationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**KeyPairRotationSettings**](KeyPairRotationSettings.md) | The key rotation settings. | 

### Return type

[**KeyPairRotationSettings**](KeyPairRotationSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

