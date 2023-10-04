# \BulkAPI

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportConfiguration**](BulkAPI.md#ExportConfiguration) | **Get** /bulk/export | Export all API resources to a JSON file.
[**ImportConfiguration**](BulkAPI.md#ImportConfiguration) | **Post** /bulk/import | Import configuration for a PingFederate deployment from a JSON file.



## ExportConfiguration

> BulkConfig ExportConfiguration(ctx).IncludeExternalResources(includeExternalResources).Execute()

Export all API resources to a JSON file.



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
    includeExternalResources := true // bool | Include external resources like OAuth clients stored outside of PingFederate. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkAPI.ExportConfiguration(context.Background()).IncludeExternalResources(includeExternalResources).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkAPI.ExportConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportConfiguration`: BulkConfig
    fmt.Fprintf(os.Stdout, "Response from `BulkAPI.ExportConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeExternalResources** | **bool** | Include external resources like OAuth clients stored outside of PingFederate. | [default to false]

### Return type

[**BulkConfig**](BulkConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportConfiguration

> ImportConfiguration(ctx).Body(body).FailFast(failFast).XBypassExternalValidation(xBypassExternalValidation).Execute()

Import configuration for a PingFederate deployment from a JSON file.



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
    body := *openapiclient.NewBulkConfig(*openapiclient.NewBulkConfigMetadata("PfVersion_example"), []openapiclient.ConfigOperation{*openapiclient.NewConfigOperation("ResourceType_example", "OperationType_example")}) // BulkConfig | Configuration to import.
    failFast := true // bool | When set to true, stops the import as soon as any validation errors are encountered. When false, import will continue to validate configuration after the first failure to identify all validation errors. If any validation errors are present PingFederate will roll back to the state prior to the import attempt. (optional) (default to true)
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BulkAPI.ImportConfiguration(context.Background()).Body(body).FailFast(failFast).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkAPI.ImportConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BulkConfig**](BulkConfig.md) | Configuration to import. | 
 **failFast** | **bool** | When set to true, stops the import as soon as any validation errors are encountered. When false, import will continue to validate configuration after the first failure to identify all validation errors. If any validation errors are present PingFederate will roll back to the state prior to the import attempt. | [default to true]
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

