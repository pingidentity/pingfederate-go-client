# \OauthAccessTokenManagersAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTokenManager**](OauthAccessTokenManagersAPI.md#CreateTokenManager) | **Post** /oauth/accessTokenManagers | Create a token management plugin instance.
[**DeleteTokenManager**](OauthAccessTokenManagersAPI.md#DeleteTokenManager) | **Delete** /oauth/accessTokenManagers/{id} | Delete a token management plugin instance.
[**GetOauthAccessTokenManagersSettings**](OauthAccessTokenManagersAPI.md#GetOauthAccessTokenManagersSettings) | **Get** /oauth/accessTokenManagers/settings | Get general access token management settings.
[**GetTokenManager**](OauthAccessTokenManagersAPI.md#GetTokenManager) | **Get** /oauth/accessTokenManagers/{id} | Get a specific token management plugin instance.
[**GetTokenManagerDescriptor**](OauthAccessTokenManagersAPI.md#GetTokenManagerDescriptor) | **Get** /oauth/accessTokenManagers/descriptors/{id} | Get the description of a token management plugin descriptor.
[**GetTokenManagerDescriptors**](OauthAccessTokenManagersAPI.md#GetTokenManagerDescriptors) | **Get** /oauth/accessTokenManagers/descriptors | Get the list of available token management plugin descriptors.
[**GetTokenManagers**](OauthAccessTokenManagersAPI.md#GetTokenManagers) | **Get** /oauth/accessTokenManagers | Get a list of all token management plugin instances.
[**UpdateOauthAccessTokenManagersSettings**](OauthAccessTokenManagersAPI.md#UpdateOauthAccessTokenManagersSettings) | **Put** /oauth/accessTokenManagers/settings | Update general access token management settings.
[**UpdateTokenManager**](OauthAccessTokenManagersAPI.md#UpdateTokenManager) | **Put** /oauth/accessTokenManagers/{id} | Update a token management plugin instance.



## CreateTokenManager

> AccessTokenManager CreateTokenManager(ctx).Body(body).Execute()

Create a token management plugin instance.

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
	body := *openapiclient.NewAccessTokenManager("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // AccessTokenManager | Configuration for plugin instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAccessTokenManagersAPI.CreateTokenManager(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAccessTokenManagersAPI.CreateTokenManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTokenManager`: AccessTokenManager
	fmt.Fprintf(os.Stdout, "Response from `OauthAccessTokenManagersAPI.CreateTokenManager`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccessTokenManager**](AccessTokenManager.md) | Configuration for plugin instance. | 

### Return type

[**AccessTokenManager**](AccessTokenManager.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTokenManager

> DeleteTokenManager(ctx, id).Execute()

Delete a token management plugin instance.

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
	id := "id_example" // string | ID of token management plugin instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OauthAccessTokenManagersAPI.DeleteTokenManager(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAccessTokenManagersAPI.DeleteTokenManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of token management plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenManagerRequest struct via the builder pattern


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


## GetOauthAccessTokenManagersSettings

> AccessTokenManagementSettings GetOauthAccessTokenManagersSettings(ctx).Execute()

Get general access token management settings.

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
	resp, r, err := apiClient.OauthAccessTokenManagersAPI.GetOauthAccessTokenManagersSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAccessTokenManagersAPI.GetOauthAccessTokenManagersSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOauthAccessTokenManagersSettings`: AccessTokenManagementSettings
	fmt.Fprintf(os.Stdout, "Response from `OauthAccessTokenManagersAPI.GetOauthAccessTokenManagersSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthAccessTokenManagersSettingsRequest struct via the builder pattern


### Return type

[**AccessTokenManagementSettings**](AccessTokenManagementSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenManager

> AccessTokenManager GetTokenManager(ctx, id).Execute()

Get a specific token management plugin instance.

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
	id := "id_example" // string | ID of token management plugin instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAccessTokenManagersAPI.GetTokenManager(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAccessTokenManagersAPI.GetTokenManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenManager`: AccessTokenManager
	fmt.Fprintf(os.Stdout, "Response from `OauthAccessTokenManagersAPI.GetTokenManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of token management plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessTokenManager**](AccessTokenManager.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenManagerDescriptor

> AccessTokenManagerDescriptor GetTokenManagerDescriptor(ctx, id).Execute()

Get the description of a token management plugin descriptor.

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
	id := "id_example" // string | ID of token management plugin descriptor.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAccessTokenManagersAPI.GetTokenManagerDescriptor(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAccessTokenManagersAPI.GetTokenManagerDescriptor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenManagerDescriptor`: AccessTokenManagerDescriptor
	fmt.Fprintf(os.Stdout, "Response from `OauthAccessTokenManagersAPI.GetTokenManagerDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of token management plugin descriptor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenManagerDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessTokenManagerDescriptor**](AccessTokenManagerDescriptor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenManagerDescriptors

> AccessTokenManagerDescriptors GetTokenManagerDescriptors(ctx).Execute()

Get the list of available token management plugin descriptors.

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
	resp, r, err := apiClient.OauthAccessTokenManagersAPI.GetTokenManagerDescriptors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAccessTokenManagersAPI.GetTokenManagerDescriptors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenManagerDescriptors`: AccessTokenManagerDescriptors
	fmt.Fprintf(os.Stdout, "Response from `OauthAccessTokenManagersAPI.GetTokenManagerDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenManagerDescriptorsRequest struct via the builder pattern


### Return type

[**AccessTokenManagerDescriptors**](AccessTokenManagerDescriptors.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenManagers

> AccessTokenManagers GetTokenManagers(ctx).Execute()

Get a list of all token management plugin instances.

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
	resp, r, err := apiClient.OauthAccessTokenManagersAPI.GetTokenManagers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAccessTokenManagersAPI.GetTokenManagers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenManagers`: AccessTokenManagers
	fmt.Fprintf(os.Stdout, "Response from `OauthAccessTokenManagersAPI.GetTokenManagers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenManagersRequest struct via the builder pattern


### Return type

[**AccessTokenManagers**](AccessTokenManagers.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOauthAccessTokenManagersSettings

> AccessTokenManagementSettings UpdateOauthAccessTokenManagersSettings(ctx).Body(body).Execute()

Update general access token management settings.

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
	body := *openapiclient.NewAccessTokenManagementSettings() // AccessTokenManagementSettings | Access token management settings.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAccessTokenManagersAPI.UpdateOauthAccessTokenManagersSettings(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAccessTokenManagersAPI.UpdateOauthAccessTokenManagersSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOauthAccessTokenManagersSettings`: AccessTokenManagementSettings
	fmt.Fprintf(os.Stdout, "Response from `OauthAccessTokenManagersAPI.UpdateOauthAccessTokenManagersSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOauthAccessTokenManagersSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccessTokenManagementSettings**](AccessTokenManagementSettings.md) | Access token management settings. | 

### Return type

[**AccessTokenManagementSettings**](AccessTokenManagementSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTokenManager

> AccessTokenManager UpdateTokenManager(ctx, id).Body(body).Execute()

Update a token management plugin instance.

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
	id := "id_example" // string | ID of token management plugin instance.
	body := *openapiclient.NewAccessTokenManager("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // AccessTokenManager | Configuration for token management plugin instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAccessTokenManagersAPI.UpdateTokenManager(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAccessTokenManagersAPI.UpdateTokenManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTokenManager`: AccessTokenManager
	fmt.Fprintf(os.Stdout, "Response from `OauthAccessTokenManagersAPI.UpdateTokenManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of token management plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTokenManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AccessTokenManager**](AccessTokenManager.md) | Configuration for token management plugin instance. | 

### Return type

[**AccessTokenManager**](AccessTokenManager.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

