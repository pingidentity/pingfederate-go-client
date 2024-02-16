# \OauthAuthorizationDetailProcessorsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthorizationDetailProcessor**](OauthAuthorizationDetailProcessorsAPI.md#CreateAuthorizationDetailProcessor) | **Post** /oauth/authorizationDetailProcessors | Create an authorization detail processor plugin instance.
[**DeleteAuthorizationDetailProcessor**](OauthAuthorizationDetailProcessorsAPI.md#DeleteAuthorizationDetailProcessor) | **Delete** /oauth/authorizationDetailProcessors/{id} | Delete an authorization detail processor plugin instance.
[**GetAuthorizationDetailProcessor**](OauthAuthorizationDetailProcessorsAPI.md#GetAuthorizationDetailProcessor) | **Get** /oauth/authorizationDetailProcessors/{id} | Get a specific authorization detail processor plugin instance.
[**GetAuthorizationDetailProcessorPluginDescriptor**](OauthAuthorizationDetailProcessorsAPI.md#GetAuthorizationDetailProcessorPluginDescriptor) | **Get** /oauth/authorizationDetailProcessors/descriptors/{id} | Get an authorization detail processor plugin descriptor.
[**GetAuthorizationDetailProcessorPluginDescriptors**](OauthAuthorizationDetailProcessorsAPI.md#GetAuthorizationDetailProcessorPluginDescriptors) | **Get** /oauth/authorizationDetailProcessors/descriptors | Get a list of available authorization detail processor plugin descriptors.
[**GetAuthorizationDetailProcessors**](OauthAuthorizationDetailProcessorsAPI.md#GetAuthorizationDetailProcessors) | **Get** /oauth/authorizationDetailProcessors | Get a list of authorization detail processor plugin instances.
[**UpdateAuthorizationDetailProcessor**](OauthAuthorizationDetailProcessorsAPI.md#UpdateAuthorizationDetailProcessor) | **Put** /oauth/authorizationDetailProcessors/{id} | Update an authorization detail processor plugin instance.



## CreateAuthorizationDetailProcessor

> AuthorizationDetailProcessor CreateAuthorizationDetailProcessor(ctx).Body(body).Execute()

Create an authorization detail processor plugin instance.

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
	body := *openapiclient.NewAuthorizationDetailProcessor("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // AuthorizationDetailProcessor | Configuration for a authorization detail processor plugin instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAuthorizationDetailProcessorsAPI.CreateAuthorizationDetailProcessor(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailProcessorsAPI.CreateAuthorizationDetailProcessor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAuthorizationDetailProcessor`: AuthorizationDetailProcessor
	fmt.Fprintf(os.Stdout, "Response from `OauthAuthorizationDetailProcessorsAPI.CreateAuthorizationDetailProcessor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthorizationDetailProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthorizationDetailProcessor**](AuthorizationDetailProcessor.md) | Configuration for a authorization detail processor plugin instance. | 

### Return type

[**AuthorizationDetailProcessor**](AuthorizationDetailProcessor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthorizationDetailProcessor

> DeleteAuthorizationDetailProcessor(ctx, id).Execute()

Delete an authorization detail processor plugin instance.

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
	id := "id_example" // string | ID of an authorization detail processor plugin instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OauthAuthorizationDetailProcessorsAPI.DeleteAuthorizationDetailProcessor(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailProcessorsAPI.DeleteAuthorizationDetailProcessor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an authorization detail processor plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthorizationDetailProcessorRequest struct via the builder pattern


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


## GetAuthorizationDetailProcessor

> AuthorizationDetailProcessor GetAuthorizationDetailProcessor(ctx, id).Execute()

Get a specific authorization detail processor plugin instance.

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
	id := "id_example" // string | ID of an authorization detail processor plugin instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAuthorizationDetailProcessorsAPI.GetAuthorizationDetailProcessor(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailProcessorsAPI.GetAuthorizationDetailProcessor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthorizationDetailProcessor`: AuthorizationDetailProcessor
	fmt.Fprintf(os.Stdout, "Response from `OauthAuthorizationDetailProcessorsAPI.GetAuthorizationDetailProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an authorization detail processor plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationDetailProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthorizationDetailProcessor**](AuthorizationDetailProcessor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorizationDetailProcessorPluginDescriptor

> AuthorizationDetailProcessorDescriptor GetAuthorizationDetailProcessorPluginDescriptor(ctx, id).Execute()

Get an authorization detail processor plugin descriptor.

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
	id := "id_example" // string | ID of authorization detail processor plugin descriptor.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAuthorizationDetailProcessorsAPI.GetAuthorizationDetailProcessorPluginDescriptor(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailProcessorsAPI.GetAuthorizationDetailProcessorPluginDescriptor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthorizationDetailProcessorPluginDescriptor`: AuthorizationDetailProcessorDescriptor
	fmt.Fprintf(os.Stdout, "Response from `OauthAuthorizationDetailProcessorsAPI.GetAuthorizationDetailProcessorPluginDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of authorization detail processor plugin descriptor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationDetailProcessorPluginDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthorizationDetailProcessorDescriptor**](AuthorizationDetailProcessorDescriptor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorizationDetailProcessorPluginDescriptors

> AuthorizationDetailProcessorDescriptors GetAuthorizationDetailProcessorPluginDescriptors(ctx).Execute()

Get a list of available authorization detail processor plugin descriptors.

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
	resp, r, err := apiClient.OauthAuthorizationDetailProcessorsAPI.GetAuthorizationDetailProcessorPluginDescriptors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailProcessorsAPI.GetAuthorizationDetailProcessorPluginDescriptors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthorizationDetailProcessorPluginDescriptors`: AuthorizationDetailProcessorDescriptors
	fmt.Fprintf(os.Stdout, "Response from `OauthAuthorizationDetailProcessorsAPI.GetAuthorizationDetailProcessorPluginDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationDetailProcessorPluginDescriptorsRequest struct via the builder pattern


### Return type

[**AuthorizationDetailProcessorDescriptors**](AuthorizationDetailProcessorDescriptors.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorizationDetailProcessors

> AuthorizationDetailProcessors GetAuthorizationDetailProcessors(ctx).Execute()

Get a list of authorization detail processor plugin instances.

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
	resp, r, err := apiClient.OauthAuthorizationDetailProcessorsAPI.GetAuthorizationDetailProcessors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailProcessorsAPI.GetAuthorizationDetailProcessors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthorizationDetailProcessors`: AuthorizationDetailProcessors
	fmt.Fprintf(os.Stdout, "Response from `OauthAuthorizationDetailProcessorsAPI.GetAuthorizationDetailProcessors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationDetailProcessorsRequest struct via the builder pattern


### Return type

[**AuthorizationDetailProcessors**](AuthorizationDetailProcessors.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthorizationDetailProcessor

> AuthorizationDetailProcessor UpdateAuthorizationDetailProcessor(ctx, id).Body(body).Execute()

Update an authorization detail processor plugin instance.

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
	id := "id_example" // string | ID of an authorization detail processor plugin instance.
	body := *openapiclient.NewAuthorizationDetailProcessor("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // AuthorizationDetailProcessor | Configuration for a authorization detail processor plugin instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAuthorizationDetailProcessorsAPI.UpdateAuthorizationDetailProcessor(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailProcessorsAPI.UpdateAuthorizationDetailProcessor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAuthorizationDetailProcessor`: AuthorizationDetailProcessor
	fmt.Fprintf(os.Stdout, "Response from `OauthAuthorizationDetailProcessorsAPI.UpdateAuthorizationDetailProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an authorization detail processor plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthorizationDetailProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthorizationDetailProcessor**](AuthorizationDetailProcessor.md) | Configuration for a authorization detail processor plugin instance. | 

### Return type

[**AuthorizationDetailProcessor**](AuthorizationDetailProcessor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

