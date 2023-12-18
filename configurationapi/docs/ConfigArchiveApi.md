# \ConfigArchiveAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportConfigArchive**](ConfigArchiveAPI.md#ExportConfigArchive) | **Get** /configArchive/export | Export a configuration archive.
[**ImportConfigArchive**](ConfigArchiveAPI.md#ImportConfigArchive) | **Post** /configArchive/import | Import a configuration archive.



## ExportConfigArchive

> ExportConfigArchive(ctx).Execute()

Export a configuration archive.

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
    r, err := apiClient.ConfigArchiveAPI.ExportConfigArchive(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigArchiveAPI.ExportConfigArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiExportConfigArchiveRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportConfigArchive

> ApiResult ImportConfigArchive(ctx).ForceImport(forceImport).ForceUnsupportedImport(forceUnsupportedImport).ReencryptData(reencryptData).File(file).Execute()

Import a configuration archive.



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
    forceImport := true // bool |  (optional)
    forceUnsupportedImport := true // bool | Force import of unsupported versions. (optional) (default to false)
    reencryptData := true // bool | Reencrypt configuration archive data with the current deployment's encryption key. (optional) (default to false)
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigArchiveAPI.ImportConfigArchive(context.Background()).ForceImport(forceImport).ForceUnsupportedImport(forceUnsupportedImport).ReencryptData(reencryptData).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigArchiveAPI.ImportConfigArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportConfigArchive`: ApiResult
    fmt.Fprintf(os.Stdout, "Response from `ConfigArchiveAPI.ImportConfigArchive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportConfigArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forceImport** | **bool** |  | 
 **forceUnsupportedImport** | **bool** | Force import of unsupported versions. | [default to false]
 **reencryptData** | **bool** | Reencrypt configuration archive data with the current deployment&#39;s encryption key. | [default to false]
 **file** | ***os.File** |  | 

### Return type

[**ApiResult**](ApiResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

