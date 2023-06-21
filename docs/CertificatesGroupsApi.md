# \CertificatesGroupsApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCertificateFromGroup**](CertificatesGroupsApi.md#DeleteCertificateFromGroup) | **Delete** /certificates/groups/{groupName}/{id} | Delete a certificate from a group.
[**GetCertificateFromGroup**](CertificatesGroupsApi.md#GetCertificateFromGroup) | **Get** /certificates/groups/{groupName}/{id} | Retrieve details of a certificate.
[**GetCertificatesForGroup**](CertificatesGroupsApi.md#GetCertificatesForGroup) | **Get** /certificates/groups/{groupName} | Get list of all certificates for a group.
[**ImportFeatureCert**](CertificatesGroupsApi.md#ImportFeatureCert) | **Post** /certificates/groups/{groupName}/import | Import a new certificate to a group.



## DeleteCertificateFromGroup

> DeleteCertificateFromGroup(ctx, groupName, id).Execute()

Delete a certificate from a group.



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
    groupName := "groupName_example" // string | Name of the group to retrieve certificates for.
    id := "id_example" // string | ID of the certificate to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CertificatesGroupsApi.DeleteCertificateFromGroup(context.Background(), groupName, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesGroupsApi.DeleteCertificateFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | Name of the group to retrieve certificates for. | 
**id** | **string** | ID of the certificate to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificateFromGroupRequest struct via the builder pattern


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


## GetCertificateFromGroup

> CertView GetCertificateFromGroup(ctx, groupName, id).Execute()

Retrieve details of a certificate.

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
    groupName := "groupName_example" // string | Name of the group to retrieve certificates for.
    id := "id_example" // string | ID of the certificate to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesGroupsApi.GetCertificateFromGroup(context.Background(), groupName, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesGroupsApi.GetCertificateFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateFromGroup`: CertView
    fmt.Fprintf(os.Stdout, "Response from `CertificatesGroupsApi.GetCertificateFromGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | Name of the group to retrieve certificates for. | 
**id** | **string** | ID of the certificate to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateFromGroupRequest struct via the builder pattern


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


## GetCertificatesForGroup

> CertViews GetCertificatesForGroup(ctx, groupName).Execute()

Get list of all certificates for a group.

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
    groupName := "groupName_example" // string | Name of the group to retrieve certificates for.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesGroupsApi.GetCertificatesForGroup(context.Background(), groupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesGroupsApi.GetCertificatesForGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificatesForGroup`: CertViews
    fmt.Fprintf(os.Stdout, "Response from `CertificatesGroupsApi.GetCertificatesForGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | Name of the group to retrieve certificates for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificatesForGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ImportFeatureCert

> CertView ImportFeatureCert(ctx, groupName).Body(body).Execute()

Import a new certificate to a group.

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
    groupName := "groupName_example" // string | Name of the group to retrieve certificates for.
    body := *openapiclient.NewX509File("FileData_example") // X509File | File data to import.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesGroupsApi.ImportFeatureCert(context.Background(), groupName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesGroupsApi.ImportFeatureCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportFeatureCert`: CertView
    fmt.Fprintf(os.Stdout, "Response from `CertificatesGroupsApi.ImportFeatureCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | Name of the group to retrieve certificates for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportFeatureCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**X509File**](X509File.md) | File data to import. | 

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

