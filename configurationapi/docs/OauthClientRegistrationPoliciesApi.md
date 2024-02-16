# \OauthClientRegistrationPoliciesAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDynamicClientRegistrationPolicy**](OauthClientRegistrationPoliciesAPI.md#CreateDynamicClientRegistrationPolicy) | **Post** /oauth/clientRegistrationPolicies | Create a client registration policy plugin instance.
[**DeleteDynamicClientRegistrationPolicy**](OauthClientRegistrationPoliciesAPI.md#DeleteDynamicClientRegistrationPolicy) | **Delete** /oauth/clientRegistrationPolicies/{id} | Delete a client registration policy plugin instance.
[**GetDynamicClientRegistrationDescriptor**](OauthClientRegistrationPoliciesAPI.md#GetDynamicClientRegistrationDescriptor) | **Get** /oauth/clientRegistrationPolicies/descriptors/{id} | Get the description of a client registration policy plugin descriptor.
[**GetDynamicClientRegistrationDescriptors**](OauthClientRegistrationPoliciesAPI.md#GetDynamicClientRegistrationDescriptors) | **Get** /oauth/clientRegistrationPolicies/descriptors | Get the list of available client registration policy plugin descriptors.
[**GetDynamicClientRegistrationPolicies**](OauthClientRegistrationPoliciesAPI.md#GetDynamicClientRegistrationPolicies) | **Get** /oauth/clientRegistrationPolicies | Get a list of client registration policy plugin instances.
[**GetDynamicClientRegistrationPolicy**](OauthClientRegistrationPoliciesAPI.md#GetDynamicClientRegistrationPolicy) | **Get** /oauth/clientRegistrationPolicies/{id} | Get a specific client registration policy plugin instance.
[**UpdateDynamicClientRegistrationPolicy**](OauthClientRegistrationPoliciesAPI.md#UpdateDynamicClientRegistrationPolicy) | **Put** /oauth/clientRegistrationPolicies/{id} | Update a client registration policy plugin instance.



## CreateDynamicClientRegistrationPolicy

> ClientRegistrationPolicy CreateDynamicClientRegistrationPolicy(ctx).Body(body).Execute()

Create a client registration policy plugin instance.

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
	body := *openapiclient.NewClientRegistrationPolicy("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // ClientRegistrationPolicy | Configuration for a client registration policy plugin instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthClientRegistrationPoliciesAPI.CreateDynamicClientRegistrationPolicy(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthClientRegistrationPoliciesAPI.CreateDynamicClientRegistrationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDynamicClientRegistrationPolicy`: ClientRegistrationPolicy
	fmt.Fprintf(os.Stdout, "Response from `OauthClientRegistrationPoliciesAPI.CreateDynamicClientRegistrationPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDynamicClientRegistrationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClientRegistrationPolicy**](ClientRegistrationPolicy.md) | Configuration for a client registration policy plugin instance. | 

### Return type

[**ClientRegistrationPolicy**](ClientRegistrationPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDynamicClientRegistrationPolicy

> DeleteDynamicClientRegistrationPolicy(ctx, id).Execute()

Delete a client registration policy plugin instance.

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
	id := "id_example" // string | ID of a client registration policy plugin instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OauthClientRegistrationPoliciesAPI.DeleteDynamicClientRegistrationPolicy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthClientRegistrationPoliciesAPI.DeleteDynamicClientRegistrationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a client registration policy plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDynamicClientRegistrationPolicyRequest struct via the builder pattern


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


## GetDynamicClientRegistrationDescriptor

> ClientRegistrationPolicyDescriptor GetDynamicClientRegistrationDescriptor(ctx, id).Execute()

Get the description of a client registration policy plugin descriptor.

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
	id := "id_example" // string | ID of client registration policy plugin descriptor.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthClientRegistrationPoliciesAPI.GetDynamicClientRegistrationDescriptor(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthClientRegistrationPoliciesAPI.GetDynamicClientRegistrationDescriptor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDynamicClientRegistrationDescriptor`: ClientRegistrationPolicyDescriptor
	fmt.Fprintf(os.Stdout, "Response from `OauthClientRegistrationPoliciesAPI.GetDynamicClientRegistrationDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of client registration policy plugin descriptor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDynamicClientRegistrationDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientRegistrationPolicyDescriptor**](ClientRegistrationPolicyDescriptor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDynamicClientRegistrationDescriptors

> ClientRegistrationPolicyDescriptors GetDynamicClientRegistrationDescriptors(ctx).Execute()

Get the list of available client registration policy plugin descriptors.

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
	resp, r, err := apiClient.OauthClientRegistrationPoliciesAPI.GetDynamicClientRegistrationDescriptors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthClientRegistrationPoliciesAPI.GetDynamicClientRegistrationDescriptors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDynamicClientRegistrationDescriptors`: ClientRegistrationPolicyDescriptors
	fmt.Fprintf(os.Stdout, "Response from `OauthClientRegistrationPoliciesAPI.GetDynamicClientRegistrationDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDynamicClientRegistrationDescriptorsRequest struct via the builder pattern


### Return type

[**ClientRegistrationPolicyDescriptors**](ClientRegistrationPolicyDescriptors.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDynamicClientRegistrationPolicies

> ClientRegistrationPolicies GetDynamicClientRegistrationPolicies(ctx).Execute()

Get a list of client registration policy plugin instances.

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
	resp, r, err := apiClient.OauthClientRegistrationPoliciesAPI.GetDynamicClientRegistrationPolicies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthClientRegistrationPoliciesAPI.GetDynamicClientRegistrationPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDynamicClientRegistrationPolicies`: ClientRegistrationPolicies
	fmt.Fprintf(os.Stdout, "Response from `OauthClientRegistrationPoliciesAPI.GetDynamicClientRegistrationPolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDynamicClientRegistrationPoliciesRequest struct via the builder pattern


### Return type

[**ClientRegistrationPolicies**](ClientRegistrationPolicies.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDynamicClientRegistrationPolicy

> ClientRegistrationPolicy GetDynamicClientRegistrationPolicy(ctx, id).Execute()

Get a specific client registration policy plugin instance.

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
	id := "id_example" // string | ID of client registration policy plugin instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthClientRegistrationPoliciesAPI.GetDynamicClientRegistrationPolicy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthClientRegistrationPoliciesAPI.GetDynamicClientRegistrationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDynamicClientRegistrationPolicy`: ClientRegistrationPolicy
	fmt.Fprintf(os.Stdout, "Response from `OauthClientRegistrationPoliciesAPI.GetDynamicClientRegistrationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of client registration policy plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDynamicClientRegistrationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientRegistrationPolicy**](ClientRegistrationPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDynamicClientRegistrationPolicy

> ClientRegistrationPolicy UpdateDynamicClientRegistrationPolicy(ctx, id).Body(body).Execute()

Update a client registration policy plugin instance.

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
	id := "id_example" // string | ID of a client registration policy plugin instance.
	body := *openapiclient.NewClientRegistrationPolicy("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // ClientRegistrationPolicy | Configuration for a client registration policy plugin instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthClientRegistrationPoliciesAPI.UpdateDynamicClientRegistrationPolicy(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthClientRegistrationPoliciesAPI.UpdateDynamicClientRegistrationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDynamicClientRegistrationPolicy`: ClientRegistrationPolicy
	fmt.Fprintf(os.Stdout, "Response from `OauthClientRegistrationPoliciesAPI.UpdateDynamicClientRegistrationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a client registration policy plugin instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDynamicClientRegistrationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ClientRegistrationPolicy**](ClientRegistrationPolicy.md) | Configuration for a client registration policy plugin instance. | 

### Return type

[**ClientRegistrationPolicy**](ClientRegistrationPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

