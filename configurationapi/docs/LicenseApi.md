# \LicenseApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLicense**](LicenseApi.md#GetLicense) | **Get** /license | Get a license summary.
[**GetLicenseAgreement**](LicenseApi.md#GetLicenseAgreement) | **Get** /license/agreement | Get license agreement link.
[**UpdateLicense**](LicenseApi.md#UpdateLicense) | **Put** /license | Import a license.
[**UpdateLicenseAgreement**](LicenseApi.md#UpdateLicenseAgreement) | **Put** /license/agreement | Accept license agreement.



## GetLicense

> LicenseView GetLicense(ctx).Execute()

Get a license summary.

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
    resp, r, err := apiClient.LicenseApi.GetLicense(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.GetLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicense`: LicenseView
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.GetLicense`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseRequest struct via the builder pattern


### Return type

[**LicenseView**](LicenseView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseAgreement

> LicenseAgreementInfo GetLicenseAgreement(ctx).Execute()

Get license agreement link.



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
    resp, r, err := apiClient.LicenseApi.GetLicenseAgreement(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.GetLicenseAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseAgreement`: LicenseAgreementInfo
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.GetLicenseAgreement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseAgreementRequest struct via the builder pattern


### Return type

[**LicenseAgreementInfo**](LicenseAgreementInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLicense

> LicenseView UpdateLicense(ctx).Body(body).Execute()

Import a license.

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
    body := *openapiclient.NewLicenseFile("FileData_example") // LicenseFile | Base64 encoded value of a license.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicenseApi.UpdateLicense(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.UpdateLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLicense`: LicenseView
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.UpdateLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LicenseFile**](LicenseFile.md) | Base64 encoded value of a license. | 

### Return type

[**LicenseView**](LicenseView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLicenseAgreement

> LicenseAgreementInfo UpdateLicenseAgreement(ctx).Body(body).Execute()

Accept license agreement.



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
    body := *openapiclient.NewLicenseAgreementInfo() // LicenseAgreementInfo | License Agreement reference.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicenseApi.UpdateLicenseAgreement(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.UpdateLicenseAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLicenseAgreement`: LicenseAgreementInfo
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.UpdateLicenseAgreement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLicenseAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LicenseAgreementInfo**](LicenseAgreementInfo.md) | License Agreement reference. | 

### Return type

[**LicenseAgreementInfo**](LicenseAgreementInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

