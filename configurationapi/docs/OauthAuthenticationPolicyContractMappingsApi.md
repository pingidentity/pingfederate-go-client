# \OauthAuthenticationPolicyContractMappingsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApcMapping**](OauthAuthenticationPolicyContractMappingsAPI.md#CreateApcMapping) | **Post** /oauth/authenticationPolicyContractMappings | Create a new authentication policy contract to persistent grant mapping.
[**DeleteApcMapping**](OauthAuthenticationPolicyContractMappingsAPI.md#DeleteApcMapping) | **Delete** /oauth/authenticationPolicyContractMappings/{id} | Delete an authentication policy contract to persistent grant mapping.
[**GetApcMapping**](OauthAuthenticationPolicyContractMappingsAPI.md#GetApcMapping) | **Get** /oauth/authenticationPolicyContractMappings/{id} | Find the authentication policy contract to persistent grant mapping by ID.
[**GetApcMappings**](OauthAuthenticationPolicyContractMappingsAPI.md#GetApcMappings) | **Get** /oauth/authenticationPolicyContractMappings | Get the list of authentication policy contract to persistent grant mappings.
[**UpdateApcMapping**](OauthAuthenticationPolicyContractMappingsAPI.md#UpdateApcMapping) | **Put** /oauth/authenticationPolicyContractMappings/{id} | Update an authentication policy contract to persistent grant mapping.



## CreateApcMapping

> ApcToPersistentGrantMapping CreateApcMapping(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new authentication policy contract to persistent grant mapping.



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
	body := *openapiclient.NewApcToPersistentGrantMapping("Id_example", *openapiclient.NewResourceLink("Id_example"), map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}) // ApcToPersistentGrantMapping | Configuration for an authentication policy contract to persistent grant mapping.
	xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAuthenticationPolicyContractMappingsAPI.CreateApcMapping(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthenticationPolicyContractMappingsAPI.CreateApcMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApcMapping`: ApcToPersistentGrantMapping
	fmt.Fprintf(os.Stdout, "Response from `OauthAuthenticationPolicyContractMappingsAPI.CreateApcMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApcMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApcToPersistentGrantMapping**](ApcToPersistentGrantMapping.md) | Configuration for an authentication policy contract to persistent grant mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**ApcToPersistentGrantMapping**](ApcToPersistentGrantMapping.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApcMapping

> DeleteApcMapping(ctx, id).Execute()

Delete an authentication policy contract to persistent grant mapping.

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
	id := "id_example" // string | ID of the authentication policy contract to persistent grant mapping.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OauthAuthenticationPolicyContractMappingsAPI.DeleteApcMapping(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthenticationPolicyContractMappingsAPI.DeleteApcMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the authentication policy contract to persistent grant mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApcMappingRequest struct via the builder pattern


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


## GetApcMapping

> ApcToPersistentGrantMapping GetApcMapping(ctx, id).Execute()

Find the authentication policy contract to persistent grant mapping by ID.

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
	id := "id_example" // string | ID of the authentication policy contract to persistent grant mapping.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAuthenticationPolicyContractMappingsAPI.GetApcMapping(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthenticationPolicyContractMappingsAPI.GetApcMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApcMapping`: ApcToPersistentGrantMapping
	fmt.Fprintf(os.Stdout, "Response from `OauthAuthenticationPolicyContractMappingsAPI.GetApcMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the authentication policy contract to persistent grant mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApcMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApcToPersistentGrantMapping**](ApcToPersistentGrantMapping.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApcMappings

> ApcToPersistentGrantMappings GetApcMappings(ctx).Execute()

Get the list of authentication policy contract to persistent grant mappings.

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
	resp, r, err := apiClient.OauthAuthenticationPolicyContractMappingsAPI.GetApcMappings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthenticationPolicyContractMappingsAPI.GetApcMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApcMappings`: ApcToPersistentGrantMappings
	fmt.Fprintf(os.Stdout, "Response from `OauthAuthenticationPolicyContractMappingsAPI.GetApcMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApcMappingsRequest struct via the builder pattern


### Return type

[**ApcToPersistentGrantMappings**](ApcToPersistentGrantMappings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApcMapping

> ApcToPersistentGrantMapping UpdateApcMapping(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update an authentication policy contract to persistent grant mapping.

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
	id := "id_example" // string | ID of the authentication policy contract to persistent grant mapping to update.
	body := *openapiclient.NewApcToPersistentGrantMapping("Id_example", *openapiclient.NewResourceLink("Id_example"), map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}) // ApcToPersistentGrantMapping | Configuration for an authentication policy contract to persistent grant mapping.
	xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAuthenticationPolicyContractMappingsAPI.UpdateApcMapping(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthenticationPolicyContractMappingsAPI.UpdateApcMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApcMapping`: ApcToPersistentGrantMapping
	fmt.Fprintf(os.Stdout, "Response from `OauthAuthenticationPolicyContractMappingsAPI.UpdateApcMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the authentication policy contract to persistent grant mapping to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApcMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApcToPersistentGrantMapping**](ApcToPersistentGrantMapping.md) | Configuration for an authentication policy contract to persistent grant mapping. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**ApcToPersistentGrantMapping**](ApcToPersistentGrantMapping.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

