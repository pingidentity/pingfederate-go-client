# \MetadataUrlsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMetadataUrl**](MetadataUrlsAPI.md#AddMetadataUrl) | **Post** /metadataUrls | Add a new Metadata URL.
[**DeleteMetadataUrl**](MetadataUrlsAPI.md#DeleteMetadataUrl) | **Delete** /metadataUrls/{id} | Delete a Metadata URL by ID.
[**GetMetadataUrl**](MetadataUrlsAPI.md#GetMetadataUrl) | **Get** /metadataUrls/{id} | Get a Metadata URL by ID.
[**GetMetadataUrls**](MetadataUrlsAPI.md#GetMetadataUrls) | **Get** /metadataUrls | Get a list of Metadata URLs
[**UpdateMetadataUrl**](MetadataUrlsAPI.md#UpdateMetadataUrl) | **Put** /metadataUrls/{id} | Update a Metadata URL by ID.



## AddMetadataUrl

> MetadataUrl AddMetadataUrl(ctx).Body(body).Execute()

Add a new Metadata URL.



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
    body := *openapiclient.NewMetadataUrl("Name_example", "Url_example") // MetadataUrl | Configuration for a new Metadata URL.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataUrlsAPI.AddMetadataUrl(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataUrlsAPI.AddMetadataUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMetadataUrl`: MetadataUrl
    fmt.Fprintf(os.Stdout, "Response from `MetadataUrlsAPI.AddMetadataUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddMetadataUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MetadataUrl**](MetadataUrl.md) | Configuration for a new Metadata URL. | 

### Return type

[**MetadataUrl**](MetadataUrl.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMetadataUrl

> DeleteMetadataUrl(ctx, id).Execute()

Delete a Metadata URL by ID.



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
    id := "id_example" // string | ID of Metadata URL to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MetadataUrlsAPI.DeleteMetadataUrl(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataUrlsAPI.DeleteMetadataUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Metadata URL to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMetadataUrlRequest struct via the builder pattern


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


## GetMetadataUrl

> MetadataUrl GetMetadataUrl(ctx, id).Execute()

Get a Metadata URL by ID.



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
    id := "id_example" // string | ID of Metadata URL to fetch

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataUrlsAPI.GetMetadataUrl(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataUrlsAPI.GetMetadataUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetadataUrl`: MetadataUrl
    fmt.Fprintf(os.Stdout, "Response from `MetadataUrlsAPI.GetMetadataUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Metadata URL to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetadataUrl**](MetadataUrl.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadataUrls

> MetadataUrls GetMetadataUrls(ctx).Execute()

Get a list of Metadata URLs

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
    resp, r, err := apiClient.MetadataUrlsAPI.GetMetadataUrls(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataUrlsAPI.GetMetadataUrls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetadataUrls`: MetadataUrls
    fmt.Fprintf(os.Stdout, "Response from `MetadataUrlsAPI.GetMetadataUrls`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataUrlsRequest struct via the builder pattern


### Return type

[**MetadataUrls**](MetadataUrls.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetadataUrl

> MetadataUrl UpdateMetadataUrl(ctx, id).Body(body).Execute()

Update a Metadata URL by ID.



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
    id := "id_example" // string | ID of the Metadata URL to update.
    body := *openapiclient.NewMetadataUrl("Name_example", "Url_example") // MetadataUrl | Configuration for the Metadata URL.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataUrlsAPI.UpdateMetadataUrl(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataUrlsAPI.UpdateMetadataUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMetadataUrl`: MetadataUrl
    fmt.Fprintf(os.Stdout, "Response from `MetadataUrlsAPI.UpdateMetadataUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Metadata URL to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetadataUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MetadataUrl**](MetadataUrl.md) | Configuration for the Metadata URL. | 

### Return type

[**MetadataUrl**](MetadataUrl.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

