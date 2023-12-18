# \SpDefaultUrlsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSpDefaultUrls**](SpDefaultUrlsAPI.md#GetSpDefaultUrls) | **Get** /sp/defaultUrls | Gets the SP Default URLs. These are Values that affect the user&#39;s experience when executing SP-initiated SSO operations.
[**UpdateSpDefaultUrls**](SpDefaultUrlsAPI.md#UpdateSpDefaultUrls) | **Put** /sp/defaultUrls | Update the SP Default URLs. Enter values that affect the user&#39;s experience when executing SP-initiated SSO operations.



## GetSpDefaultUrls

> SpDefaultUrls GetSpDefaultUrls(ctx).Execute()

Gets the SP Default URLs. These are Values that affect the user's experience when executing SP-initiated SSO operations.

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
    resp, r, err := apiClient.SpDefaultUrlsAPI.GetSpDefaultUrls(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpDefaultUrlsAPI.GetSpDefaultUrls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpDefaultUrls`: SpDefaultUrls
    fmt.Fprintf(os.Stdout, "Response from `SpDefaultUrlsAPI.GetSpDefaultUrls`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpDefaultUrlsRequest struct via the builder pattern


### Return type

[**SpDefaultUrls**](SpDefaultUrls.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpDefaultUrls

> SpDefaultUrls UpdateSpDefaultUrls(ctx).Body(body).Execute()

Update the SP Default URLs. Enter values that affect the user's experience when executing SP-initiated SSO operations.

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
    body := *openapiclient.NewSpDefaultUrls() // SpDefaultUrls | Configuration for the IDP Default URL settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpDefaultUrlsAPI.UpdateSpDefaultUrls(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpDefaultUrlsAPI.UpdateSpDefaultUrls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSpDefaultUrls`: SpDefaultUrls
    fmt.Fprintf(os.Stdout, "Response from `SpDefaultUrlsAPI.UpdateSpDefaultUrls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpDefaultUrlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SpDefaultUrls**](SpDefaultUrls.md) | Configuration for the IDP Default URL settings. | 

### Return type

[**SpDefaultUrls**](SpDefaultUrls.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

