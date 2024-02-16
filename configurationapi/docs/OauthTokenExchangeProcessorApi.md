# \OauthTokenExchangeProcessorAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOauthTokenExchangeProcessorPolicy**](OauthTokenExchangeProcessorAPI.md#CreateOauthTokenExchangeProcessorPolicy) | **Post** /oauth/tokenExchange/processor/policies | Create a new OAuth 2.0 Token Exchange Processor policy.
[**DeleteOauthTokenExchangeProcessorPolicyy**](OauthTokenExchangeProcessorAPI.md#DeleteOauthTokenExchangeProcessorPolicyy) | **Delete** /oauth/tokenExchange/processor/policies/{id} | Delete an OAuth 2.0 Token Exchange Processor policy.
[**GetOauthTokenExchangeProcessorPolicyById**](OauthTokenExchangeProcessorAPI.md#GetOauthTokenExchangeProcessorPolicyById) | **Get** /oauth/tokenExchange/processor/policies/{id} | Find an OAuth 2.0 Token Exchange Processor policy by ID.
[**GetOauthTokenExchangeProcessorPolicyPolicies**](OauthTokenExchangeProcessorAPI.md#GetOauthTokenExchangeProcessorPolicyPolicies) | **Get** /oauth/tokenExchange/processor/policies | Get list of OAuth 2.0 Token Exchange Processor policies.
[**GetOauthTokenExchangeProcessorPolicySettings**](OauthTokenExchangeProcessorAPI.md#GetOauthTokenExchangeProcessorPolicySettings) | **Get** /oauth/tokenExchange/processor/settings | Get general OAuth 2.0 Token Exchange Processor settings.
[**UpdateOauthTokenExchangeProcessorPolicy**](OauthTokenExchangeProcessorAPI.md#UpdateOauthTokenExchangeProcessorPolicy) | **Put** /oauth/tokenExchange/processor/policies/{id} | Update an OAuth 2.0 Token Exchange Processor policy.
[**UpdateOauthTokenExchangeProcessorPolicySettings**](OauthTokenExchangeProcessorAPI.md#UpdateOauthTokenExchangeProcessorPolicySettings) | **Put** /oauth/tokenExchange/processor/settings | Update general OAuth 2.0 Token Exchange Processor settings.



## CreateOauthTokenExchangeProcessorPolicy

> TokenExchangeProcessorPolicy CreateOauthTokenExchangeProcessorPolicy(ctx).Body(body).BypassExternalValidation(bypassExternalValidation).Execute()

Create a new OAuth 2.0 Token Exchange Processor policy.



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
	body := *openapiclient.NewTokenExchangeProcessorPolicy("Id_example", "Name_example", *openapiclient.NewTokenExchangeProcessorAttributeContract(), []openapiclient.TokenExchangeProcessorMapping{*openapiclient.NewTokenExchangeProcessorMapping(map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}, "SubjectTokenType_example", *openapiclient.NewResourceLink("Id_example"))}) // TokenExchangeProcessorPolicy | Configuration for new OAuth 2.0 Token Exchange Processor.
	bypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthTokenExchangeProcessorAPI.CreateOauthTokenExchangeProcessorPolicy(context.Background()).Body(body).BypassExternalValidation(bypassExternalValidation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeProcessorAPI.CreateOauthTokenExchangeProcessorPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOauthTokenExchangeProcessorPolicy`: TokenExchangeProcessorPolicy
	fmt.Fprintf(os.Stdout, "Response from `OauthTokenExchangeProcessorAPI.CreateOauthTokenExchangeProcessorPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOauthTokenExchangeProcessorPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TokenExchangeProcessorPolicy**](TokenExchangeProcessorPolicy.md) | Configuration for new OAuth 2.0 Token Exchange Processor. | 
 **bypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | 

### Return type

[**TokenExchangeProcessorPolicy**](TokenExchangeProcessorPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOauthTokenExchangeProcessorPolicyy

> DeleteOauthTokenExchangeProcessorPolicyy(ctx, id).Execute()

Delete an OAuth 2.0 Token Exchange Processor policy.



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
	id := "id_example" // string | ID of OAuth 2.0 Token Exchange Processor policy to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OauthTokenExchangeProcessorAPI.DeleteOauthTokenExchangeProcessorPolicyy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeProcessorAPI.DeleteOauthTokenExchangeProcessorPolicyy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of OAuth 2.0 Token Exchange Processor policy to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOauthTokenExchangeProcessorPolicyyRequest struct via the builder pattern


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


## GetOauthTokenExchangeProcessorPolicyById

> TokenExchangeProcessorPolicy GetOauthTokenExchangeProcessorPolicyById(ctx, id).Execute()

Find an OAuth 2.0 Token Exchange Processor policy by ID.



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
	id := "id_example" // string | ID of the OAuth 2.0 Token Exchange Processor policy to fetch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthTokenExchangeProcessorAPI.GetOauthTokenExchangeProcessorPolicyById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeProcessorAPI.GetOauthTokenExchangeProcessorPolicyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOauthTokenExchangeProcessorPolicyById`: TokenExchangeProcessorPolicy
	fmt.Fprintf(os.Stdout, "Response from `OauthTokenExchangeProcessorAPI.GetOauthTokenExchangeProcessorPolicyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the OAuth 2.0 Token Exchange Processor policy to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthTokenExchangeProcessorPolicyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenExchangeProcessorPolicy**](TokenExchangeProcessorPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthTokenExchangeProcessorPolicyPolicies

> TokenExchangeProcessorPolicies GetOauthTokenExchangeProcessorPolicyPolicies(ctx).Execute()

Get list of OAuth 2.0 Token Exchange Processor policies.

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
	resp, r, err := apiClient.OauthTokenExchangeProcessorAPI.GetOauthTokenExchangeProcessorPolicyPolicies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeProcessorAPI.GetOauthTokenExchangeProcessorPolicyPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOauthTokenExchangeProcessorPolicyPolicies`: TokenExchangeProcessorPolicies
	fmt.Fprintf(os.Stdout, "Response from `OauthTokenExchangeProcessorAPI.GetOauthTokenExchangeProcessorPolicyPolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthTokenExchangeProcessorPolicyPoliciesRequest struct via the builder pattern


### Return type

[**TokenExchangeProcessorPolicies**](TokenExchangeProcessorPolicies.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthTokenExchangeProcessorPolicySettings

> TokenExchangeProcessorSettings GetOauthTokenExchangeProcessorPolicySettings(ctx).Execute()

Get general OAuth 2.0 Token Exchange Processor settings.

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
	resp, r, err := apiClient.OauthTokenExchangeProcessorAPI.GetOauthTokenExchangeProcessorPolicySettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeProcessorAPI.GetOauthTokenExchangeProcessorPolicySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOauthTokenExchangeProcessorPolicySettings`: TokenExchangeProcessorSettings
	fmt.Fprintf(os.Stdout, "Response from `OauthTokenExchangeProcessorAPI.GetOauthTokenExchangeProcessorPolicySettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthTokenExchangeProcessorPolicySettingsRequest struct via the builder pattern


### Return type

[**TokenExchangeProcessorSettings**](TokenExchangeProcessorSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOauthTokenExchangeProcessorPolicy

> TokenExchangeProcessorPolicy UpdateOauthTokenExchangeProcessorPolicy(ctx, id).Body(body).BypassExternalValidation(bypassExternalValidation).Execute()

Update an OAuth 2.0 Token Exchange Processor policy.



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
	id := "id_example" // string | ID of the OAuth 2.0 Token Exchange Processor policy to update.
	body := *openapiclient.NewTokenExchangeProcessorPolicy("Id_example", "Name_example", *openapiclient.NewTokenExchangeProcessorAttributeContract(), []openapiclient.TokenExchangeProcessorMapping{*openapiclient.NewTokenExchangeProcessorMapping(map[string]AttributeFulfillmentValue{"key": *openapiclient.NewAttributeFulfillmentValue(*openapiclient.NewSourceTypeIdKey("Type_example"), "Value_example")}, "SubjectTokenType_example", *openapiclient.NewResourceLink("Id_example"))}) // TokenExchangeProcessorPolicy | Configuration for updated OAuth 2.0 Token Exchange Processor policy.
	bypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthTokenExchangeProcessorAPI.UpdateOauthTokenExchangeProcessorPolicy(context.Background(), id).Body(body).BypassExternalValidation(bypassExternalValidation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeProcessorAPI.UpdateOauthTokenExchangeProcessorPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOauthTokenExchangeProcessorPolicy`: TokenExchangeProcessorPolicy
	fmt.Fprintf(os.Stdout, "Response from `OauthTokenExchangeProcessorAPI.UpdateOauthTokenExchangeProcessorPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the OAuth 2.0 Token Exchange Processor policy to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOauthTokenExchangeProcessorPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TokenExchangeProcessorPolicy**](TokenExchangeProcessorPolicy.md) | Configuration for updated OAuth 2.0 Token Exchange Processor policy. | 
 **bypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | 

### Return type

[**TokenExchangeProcessorPolicy**](TokenExchangeProcessorPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOauthTokenExchangeProcessorPolicySettings

> TokenExchangeProcessorSettings UpdateOauthTokenExchangeProcessorPolicySettings(ctx).Body(body).BypassExternalValidation(bypassExternalValidation).Execute()

Update general OAuth 2.0 Token Exchange Processor settings.

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
	body := *openapiclient.NewTokenExchangeProcessorSettings() // TokenExchangeProcessorSettings | OAuth 2.0 Token Exchange Processor settings.
	bypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthTokenExchangeProcessorAPI.UpdateOauthTokenExchangeProcessorPolicySettings(context.Background()).Body(body).BypassExternalValidation(bypassExternalValidation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenExchangeProcessorAPI.UpdateOauthTokenExchangeProcessorPolicySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOauthTokenExchangeProcessorPolicySettings`: TokenExchangeProcessorSettings
	fmt.Fprintf(os.Stdout, "Response from `OauthTokenExchangeProcessorAPI.UpdateOauthTokenExchangeProcessorPolicySettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOauthTokenExchangeProcessorPolicySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TokenExchangeProcessorSettings**](TokenExchangeProcessorSettings.md) | OAuth 2.0 Token Exchange Processor settings. | 
 **bypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | 

### Return type

[**TokenExchangeProcessorSettings**](TokenExchangeProcessorSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

