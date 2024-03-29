# \CertificatesRevocationAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOcspCertificateById**](CertificatesRevocationAPI.md#DeleteOcspCertificateById) | **Delete** /certificates/revocation/ocspCertificates/{id} | Delete an OCSP responder signature verification certificate by ID.
[**GetOcspCertificateById**](CertificatesRevocationAPI.md#GetOcspCertificateById) | **Get** /certificates/revocation/ocspCertificates/{id} | Get an OCSP responder signature verification certificate by ID.
[**GetOcspCertificates**](CertificatesRevocationAPI.md#GetOcspCertificates) | **Get** /certificates/revocation/ocspCertificates | Get the list of available OCSP responder signature verification certificates.
[**GetRevocationSettings**](CertificatesRevocationAPI.md#GetRevocationSettings) | **Get** /certificates/revocation/settings | Get certificate revocation settings.
[**ImportOcspCertificate**](CertificatesRevocationAPI.md#ImportOcspCertificate) | **Post** /certificates/revocation/ocspCertificates | Import an OCSP responder signature verification certificate.
[**UpdateRevocationSettings**](CertificatesRevocationAPI.md#UpdateRevocationSettings) | **Put** /certificates/revocation/settings | Update certificate revocation settings.



## DeleteOcspCertificateById

> DeleteOcspCertificateById(ctx, id).Execute()

Delete an OCSP responder signature verification certificate by ID.

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
    id := "id_example" // string | Certificate ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CertificatesRevocationAPI.DeleteOcspCertificateById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesRevocationAPI.DeleteOcspCertificateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Certificate ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOcspCertificateByIdRequest struct via the builder pattern


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


## GetOcspCertificateById

> CertView GetOcspCertificateById(ctx, id).Execute()

Get an OCSP responder signature verification certificate by ID.

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
    id := "id_example" // string | Certificate ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesRevocationAPI.GetOcspCertificateById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesRevocationAPI.GetOcspCertificateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOcspCertificateById`: CertView
    fmt.Fprintf(os.Stdout, "Response from `CertificatesRevocationAPI.GetOcspCertificateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Certificate ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOcspCertificateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CertView**](CertView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOcspCertificates

> CertViews GetOcspCertificates(ctx).Execute()

Get the list of available OCSP responder signature verification certificates.

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
    resp, r, err := apiClient.CertificatesRevocationAPI.GetOcspCertificates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesRevocationAPI.GetOcspCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOcspCertificates`: CertViews
    fmt.Fprintf(os.Stdout, "Response from `CertificatesRevocationAPI.GetOcspCertificates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOcspCertificatesRequest struct via the builder pattern


### Return type

[**CertViews**](CertViews.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevocationSettings

> CertificateRevocationSettings GetRevocationSettings(ctx).Execute()

Get certificate revocation settings.

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
    resp, r, err := apiClient.CertificatesRevocationAPI.GetRevocationSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesRevocationAPI.GetRevocationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRevocationSettings`: CertificateRevocationSettings
    fmt.Fprintf(os.Stdout, "Response from `CertificatesRevocationAPI.GetRevocationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevocationSettingsRequest struct via the builder pattern


### Return type

[**CertificateRevocationSettings**](CertificateRevocationSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportOcspCertificate

> CertView ImportOcspCertificate(ctx).Body(body).Execute()

Import an OCSP responder signature verification certificate.

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
    body := *openapiclient.NewX509File("FileData_example") // X509File | File to import.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesRevocationAPI.ImportOcspCertificate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesRevocationAPI.ImportOcspCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportOcspCertificate`: CertView
    fmt.Fprintf(os.Stdout, "Response from `CertificatesRevocationAPI.ImportOcspCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportOcspCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**X509File**](X509File.md) | File to import. | 

### Return type

[**CertView**](CertView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRevocationSettings

> CertificateRevocationSettings UpdateRevocationSettings(ctx).Body(body).Execute()

Update certificate revocation settings.

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
    body := *openapiclient.NewCertificateRevocationSettings() // CertificateRevocationSettings | Certificate revocation settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesRevocationAPI.UpdateRevocationSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesRevocationAPI.UpdateRevocationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRevocationSettings`: CertificateRevocationSettings
    fmt.Fprintf(os.Stdout, "Response from `CertificatesRevocationAPI.UpdateRevocationSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRevocationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CertificateRevocationSettings**](CertificateRevocationSettings.md) | Certificate revocation settings. | 

### Return type

[**CertificateRevocationSettings**](CertificateRevocationSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

