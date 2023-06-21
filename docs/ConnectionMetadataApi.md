# \ConnectionMetadataApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Convert**](ConnectionMetadataApi.md#Convert) | **Post** /connectionMetadata/convert | Convert a partner&#39;s SAML metadata into a JSON representation.
[**Export**](ConnectionMetadataApi.md#Export) | **Post** /connectionMetadata/export | Export a connection&#39;s SAML metadata that can be given to a partner.



## Convert

> ConvertMetadataResponse Convert(ctx).Body(body).Execute()

Convert a partner's SAML metadata into a JSON representation.



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
    body := *openapiclient.NewConvertMetadataRequest("ConnectionType_example", "ExpectedProtocol_example", "SamlMetadata_example") // ConvertMetadataRequest | Convert metadata request.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionMetadataApi.Convert(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionMetadataApi.Convert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Convert`: ConvertMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionMetadataApi.Convert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConvertMetadataRequest**](ConvertMetadataRequest.md) | Convert metadata request. | 

### Return type

[**ConvertMetadataResponse**](ConvertMetadataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Export

> string Export(ctx).Body(body).Execute()

Export a connection's SAML metadata that can be given to a partner.

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
    body := *openapiclient.NewExportMetadataRequest("ConnectionType_example", "ConnectionId_example") // ExportMetadataRequest | Export metadata request.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionMetadataApi.Export(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionMetadataApi.Export``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Export`: string
    fmt.Fprintf(os.Stdout, "Response from `ConnectionMetadataApi.Export`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ExportMetadataRequest**](ExportMetadataRequest.md) | Export metadata request. | 

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

