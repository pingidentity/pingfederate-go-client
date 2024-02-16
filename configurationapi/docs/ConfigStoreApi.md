# \ConfigStoreAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfigStoreSetting**](ConfigStoreAPI.md#DeleteConfigStoreSetting) | **Delete** /configStore/{bundle}/{id} | Delete a setting.
[**GetConfigStoreSetting**](ConfigStoreAPI.md#GetConfigStoreSetting) | **Get** /configStore/{bundle}/{id} | Get a single setting from a bundle.
[**GetConfigStoreSettings**](ConfigStoreAPI.md#GetConfigStoreSettings) | **Get** /configStore/{bundle} | Get all settings from a bundle.
[**UpdateConfigStoreSetting**](ConfigStoreAPI.md#UpdateConfigStoreSetting) | **Put** /configStore/{bundle}/{id} | Create or update a setting/bundle.



## DeleteConfigStoreSetting

> DeleteConfigStoreSetting(ctx, bundle, id).Execute()

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
	r, err := apiClient.ConfigStoreAPI.DeleteConfigStoreSetting(context.Background(), bundle, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreAPI.DeleteConfigStoreSetting``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteConfigStoreSettingRequest struct via the builder pattern


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


## GetConfigStoreSetting

> ConfigStoreSetting GetConfigStoreSetting(ctx, bundle, id).Execute()

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
	resp, r, err := apiClient.ConfigStoreAPI.GetConfigStoreSetting(context.Background(), bundle, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreAPI.GetConfigStoreSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigStoreSetting`: ConfigStoreSetting
	fmt.Fprintf(os.Stdout, "Response from `ConfigStoreAPI.GetConfigStoreSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundle** | **string** | This field represents a configuration file that contains a bundle of settings. | 
**id** | **string** | ID of setting to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigStoreSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConfigStoreSetting**](ConfigStoreSetting.md)

### Authorization

[bearer](../README.md#bearer)

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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigStoreSetting

> ConfigStoreSetting UpdateConfigStoreSetting(ctx, bundle, id).Body(body).Execute()

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
	resp, r, err := apiClient.ConfigStoreAPI.UpdateConfigStoreSetting(context.Background(), bundle, id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigStoreAPI.UpdateConfigStoreSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConfigStoreSetting`: ConfigStoreSetting
	fmt.Fprintf(os.Stdout, "Response from `ConfigStoreAPI.UpdateConfigStoreSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundle** | **string** | This field represents a configuration file that contains a bundle of settings. | 
**id** | **string** | ID of setting to create/update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigStoreSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ConfigStoreSetting**](ConfigStoreSetting.md) | Configuration setting. | 

### Return type

[**ConfigStoreSetting**](ConfigStoreSetting.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

