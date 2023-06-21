# \CertificatesRevocationApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOcspCertificateById**](CertificatesRevocationApi.md#DeleteOcspCertificateById) | **Delete** /certificates/revocation/ocspCertificates/{id} | Delete an OCSP responder signature verification certificate by ID.
[**GetOcspCertificateById**](CertificatesRevocationApi.md#GetOcspCertificateById) | **Get** /certificates/revocation/ocspCertificates/{id} | Get an OCSP responder signature verification certificate by ID.
[**GetOcspCertificates**](CertificatesRevocationApi.md#GetOcspCertificates) | **Get** /certificates/revocation/ocspCertificates | Get the list of available OCSP responder signature verification certificates.
[**GetRevocationSettings**](CertificatesRevocationApi.md#GetRevocationSettings) | **Get** /certificates/revocation/settings | Get certificate revocation settings.
[**ImportOcspCertificate**](CertificatesRevocationApi.md#ImportOcspCertificate) | **Post** /certificates/revocation/ocspCertificates | Import an OCSP responder signature verification certificate.
[**UpdateRevocationSettings**](CertificatesRevocationApi.md#UpdateRevocationSettings) | **Put** /certificates/revocation/settings | Update certificate revocation settings.



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
    r, err := apiClient.CertificatesRevocationApi.DeleteOcspCertificateById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesRevocationApi.DeleteOcspCertificateById``: %v\n", err)
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

No authorization required

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
    resp, r, err := apiClient.CertificatesRevocationApi.GetOcspCertificateById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesRevocationApi.GetOcspCertificateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOcspCertificateById`: CertView
    fmt.Fprintf(os.Stdout, "Response from `CertificatesRevocationApi.GetOcspCertificateById`: %v\n", resp)
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

No authorization required

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
    resp, r, err := apiClient.CertificatesRevocationApi.GetOcspCertificates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesRevocationApi.GetOcspCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOcspCertificates`: CertViews
    fmt.Fprintf(os.Stdout, "Response from `CertificatesRevocationApi.GetOcspCertificates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOcspCertificatesRequest struct via the builder pattern


### Return type

[**CertViews**](CertViews.md)

### Authorization

No authorization required

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
    resp, r, err := apiClient.CertificatesRevocationApi.GetRevocationSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesRevocationApi.GetRevocationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRevocationSettings`: CertificateRevocationSettings
    fmt.Fprintf(os.Stdout, "Response from `CertificatesRevocationApi.GetRevocationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevocationSettingsRequest struct via the builder pattern


### Return type

[**CertificateRevocationSettings**](CertificateRevocationSettings.md)

### Authorization

No authorization required

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
    resp, r, err := apiClient.CertificatesRevocationApi.ImportOcspCertificate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesRevocationApi.ImportOcspCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportOcspCertificate`: CertView
    fmt.Fprintf(os.Stdout, "Response from `CertificatesRevocationApi.ImportOcspCertificate`: %v\n", resp)
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

No authorization required

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
    resp, r, err := apiClient.CertificatesRevocationApi.UpdateRevocationSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesRevocationApi.UpdateRevocationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRevocationSettings`: CertificateRevocationSettings
    fmt.Fprintf(os.Stdout, "Response from `CertificatesRevocationApi.UpdateRevocationSettings`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

