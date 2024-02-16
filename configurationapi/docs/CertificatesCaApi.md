# \CertificatesCaAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTrustedCA**](CertificatesCaAPI.md#DeleteTrustedCA) | **Delete** /certificates/ca/{id} | Delete a trusted certificate authority.
[**ExportCaCertificateFile**](CertificatesCaAPI.md#ExportCaCertificateFile) | **Get** /certificates/ca/{id}/file | Download the certificate from a given trusted certificate authority.
[**GetTrustedCAs**](CertificatesCaAPI.md#GetTrustedCAs) | **Get** /certificates/ca | Get list of trusted certificate authorities.
[**GetTrustedCert**](CertificatesCaAPI.md#GetTrustedCert) | **Get** /certificates/ca/{id} | Retrieve details of a trusted certificate authority.
[**ImportTrustedCA**](CertificatesCaAPI.md#ImportTrustedCA) | **Post** /certificates/ca/import | Import a new trusted certificate authority.



## DeleteTrustedCA

> DeleteTrustedCA(ctx, id).Execute()

Delete a trusted certificate authority.



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
	id := "id_example" // string | ID of the trusted certificate authority to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificatesCaAPI.DeleteTrustedCA(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesCaAPI.DeleteTrustedCA``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the trusted certificate authority to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrustedCARequest struct via the builder pattern


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


## ExportCaCertificateFile

> string ExportCaCertificateFile(ctx, id).Execute()

Download the certificate from a given trusted certificate authority.



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
	id := "id_example" // string | ID of the trusted certificate authority.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesCaAPI.ExportCaCertificateFile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesCaAPI.ExportCaCertificateFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportCaCertificateFile`: string
	fmt.Fprintf(os.Stdout, "Response from `CertificatesCaAPI.ExportCaCertificateFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the trusted certificate authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCaCertificateFileRequest struct via the builder pattern


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


## GetTrustedCAs

> CertViews GetTrustedCAs(ctx).Execute()

Get list of trusted certificate authorities.

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
	resp, r, err := apiClient.CertificatesCaAPI.GetTrustedCAs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesCaAPI.GetTrustedCAs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustedCAs`: CertViews
	fmt.Fprintf(os.Stdout, "Response from `CertificatesCaAPI.GetTrustedCAs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustedCAsRequest struct via the builder pattern


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


## GetTrustedCert

> CertView GetTrustedCert(ctx, id).Execute()

Retrieve details of a trusted certificate authority.

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
	id := "id_example" // string | ID of the trusted certificate authority to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesCaAPI.GetTrustedCert(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesCaAPI.GetTrustedCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustedCert`: CertView
	fmt.Fprintf(os.Stdout, "Response from `CertificatesCaAPI.GetTrustedCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the trusted certificate authority to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustedCertRequest struct via the builder pattern


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


## ImportTrustedCA

> CertView ImportTrustedCA(ctx).Body(body).Execute()

Import a new trusted certificate authority.

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
	body := *openapiclient.NewX509File("FileData_example") // X509File | File data to import.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesCaAPI.ImportTrustedCA(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesCaAPI.ImportTrustedCA``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportTrustedCA`: CertView
	fmt.Fprintf(os.Stdout, "Response from `CertificatesCaAPI.ImportTrustedCA`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportTrustedCARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**X509File**](X509File.md) | File data to import. | 

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

