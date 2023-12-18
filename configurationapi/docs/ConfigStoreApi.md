# \ConfigStoreAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSetting**](ConfigStoreAPI.md#DeleteSetting) | **Delete** /configStore/{bundle}/{id} | Delete a setting.
[**GetConfigStoreSettings**](ConfigStoreAPI.md#GetConfigStoreSettings) | **Get** /configStore/{bundle} | Get all settings from a bundle.
[**GetSetting**](ConfigStoreAPI.md#GetSetting) | **Get** /configStore/{bundle}/{id} | Get a single setting from a bundle.
[**UpdateSetting**](ConfigStoreAPI.md#UpdateSetting) | **Put** /configStore/{bundle}/{id} | Create or update a setting/bundle.



## DeleteSetting

> DeleteSetting(ctx, bundle, id).Execute()

Delete a setting.



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
    bundle := "bundle_example" // string | This field represents a configuration file that contains a bundle of settings.
    id := "id_example" // string | ID of setting to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConfigStoreAPI.DeleteSetting(context.Background(), bundle, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreAPI.DeleteSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundle** | **string** | This field represents a configuration file that contains a bundle of settings. | 
**id** | **string** | ID of setting to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSettingRequest struct via the builder pattern


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


## GetConfigStoreSettings

> ConfigStoreBundle GetConfigStoreSettings(ctx, bundle).Execute()

Get all settings from a bundle.

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
    bundle := "bundle_example" // string | This field represents a configuration file that contains a bundle of settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigStoreAPI.GetConfigStoreSettings(context.Background(), bundle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreAPI.GetConfigStoreSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigStoreSettings`: ConfigStoreBundle
    fmt.Fprintf(os.Stdout, "Response from `ConfigStoreAPI.GetConfigStoreSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundle** | **string** | This field represents a configuration file that contains a bundle of settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigStoreSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigStoreBundle**](ConfigStoreBundle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSetting

> ConfigStoreSetting GetSetting(ctx, bundle, id).Execute()

Get a single setting from a bundle.

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
    bundle := "bundle_example" // string | This field represents a configuration file that contains a bundle of settings.
    id := "id_example" // string | ID of setting to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigStoreAPI.GetSetting(context.Background(), bundle, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreAPI.GetSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSetting`: ConfigStoreSetting
    fmt.Fprintf(os.Stdout, "Response from `ConfigStoreAPI.GetSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundle** | **string** | This field represents a configuration file that contains a bundle of settings. | 
**id** | **string** | ID of setting to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConfigStoreSetting**](ConfigStoreSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSetting

> ConfigStoreSetting UpdateSetting(ctx, bundle, id).Body(body).Execute()

Create or update a setting/bundle.



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
    bundle := "bundle_example" // string | This field represents a configuration file that contains a bundle of settings.
    id := "id_example" // string | ID of setting to create/update.
    body := *openapiclient.NewConfigStoreSetting("Id_example", "Type_example") // ConfigStoreSetting | Configuration setting.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigStoreAPI.UpdateSetting(context.Background(), bundle, id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreAPI.UpdateSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSetting`: ConfigStoreSetting
    fmt.Fprintf(os.Stdout, "Response from `ConfigStoreAPI.UpdateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundle** | **string** | This field represents a configuration file that contains a bundle of settings. | 
**id** | **string** | ID of setting to create/update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ConfigStoreSetting**](ConfigStoreSetting.md) | Configuration setting. | 

### Return type

[**ConfigStoreSetting**](ConfigStoreSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

