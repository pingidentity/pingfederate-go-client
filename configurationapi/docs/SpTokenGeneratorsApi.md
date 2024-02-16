# \SpTokenGeneratorsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTokenGenerator**](SpTokenGeneratorsAPI.md#CreateTokenGenerator) | **Post** /sp/tokenGenerators | Create a new token generator instance.
[**DeleteTokenGenerator**](SpTokenGeneratorsAPI.md#DeleteTokenGenerator) | **Delete** /sp/tokenGenerators/{id} | Delete a token generator instance.
[**GetTokenGenerator**](SpTokenGeneratorsAPI.md#GetTokenGenerator) | **Get** /sp/tokenGenerators/{id} | Find a token generator instance by ID.
[**GetTokenGeneratorDescriptors**](SpTokenGeneratorsAPI.md#GetTokenGeneratorDescriptors) | **Get** /sp/tokenGenerators/descriptors | Get the list of available token generators.
[**GetTokenGeneratorDescriptorsById**](SpTokenGeneratorsAPI.md#GetTokenGeneratorDescriptorsById) | **Get** /sp/tokenGenerators/descriptors/{id} | Get the description of a token generator plugin by ID.
[**GetTokenGenerators**](SpTokenGeneratorsAPI.md#GetTokenGenerators) | **Get** /sp/tokenGenerators | Get the list of token generator instances.
[**UpdateTokenGenerator**](SpTokenGeneratorsAPI.md#UpdateTokenGenerator) | **Put** /sp/tokenGenerators/{id} | Update a token generator instance.



## CreateTokenGenerator

> TokenGenerator CreateTokenGenerator(ctx).Body(body).Execute()

Create a new token generator instance.



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
	body := *openapiclient.NewTokenGenerator("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // TokenGenerator | Configuration for a token generator instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpTokenGeneratorsAPI.CreateTokenGenerator(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpTokenGeneratorsAPI.CreateTokenGenerator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTokenGenerator`: TokenGenerator
	fmt.Fprintf(os.Stdout, "Response from `SpTokenGeneratorsAPI.CreateTokenGenerator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenGeneratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TokenGenerator**](TokenGenerator.md) | Configuration for a token generator instance. | 

### Return type

[**TokenGenerator**](TokenGenerator.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTokenGenerator

> DeleteTokenGenerator(ctx, id).Execute()

Delete a token generator instance.



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
	id := "id_example" // string | ID of the token generator instance to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SpTokenGeneratorsAPI.DeleteTokenGenerator(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpTokenGeneratorsAPI.DeleteTokenGenerator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the token generator instance to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenGeneratorRequest struct via the builder pattern


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


## GetTokenGenerator

> TokenGenerator GetTokenGenerator(ctx, id).Execute()

Find a token generator instance by ID.



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
	id := "id_example" // string | ID of the token generator instance to fetch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpTokenGeneratorsAPI.GetTokenGenerator(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpTokenGeneratorsAPI.GetTokenGenerator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenGenerator`: TokenGenerator
	fmt.Fprintf(os.Stdout, "Response from `SpTokenGeneratorsAPI.GetTokenGenerator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the token generator instance to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenGeneratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenGenerator**](TokenGenerator.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenGeneratorDescriptors

> TokenGeneratorDescriptors GetTokenGeneratorDescriptors(ctx).Execute()

Get the list of available token generators.

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
	resp, r, err := apiClient.SpTokenGeneratorsAPI.GetTokenGeneratorDescriptors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpTokenGeneratorsAPI.GetTokenGeneratorDescriptors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenGeneratorDescriptors`: TokenGeneratorDescriptors
	fmt.Fprintf(os.Stdout, "Response from `SpTokenGeneratorsAPI.GetTokenGeneratorDescriptors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenGeneratorDescriptorsRequest struct via the builder pattern


### Return type

[**TokenGeneratorDescriptors**](TokenGeneratorDescriptors.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenGeneratorDescriptorsById

> TokenGeneratorDescriptor GetTokenGeneratorDescriptorsById(ctx, id).Execute()

Get the description of a token generator plugin by ID.



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
	id := "id_example" // string | ID of a token generator descriptor to fetch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpTokenGeneratorsAPI.GetTokenGeneratorDescriptorsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpTokenGeneratorsAPI.GetTokenGeneratorDescriptorsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenGeneratorDescriptorsById`: TokenGeneratorDescriptor
	fmt.Fprintf(os.Stdout, "Response from `SpTokenGeneratorsAPI.GetTokenGeneratorDescriptorsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a token generator descriptor to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenGeneratorDescriptorsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenGeneratorDescriptor**](TokenGeneratorDescriptor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenGenerators

> TokenGenerators GetTokenGenerators(ctx).Execute()

Get the list of token generator instances.

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
	resp, r, err := apiClient.SpTokenGeneratorsAPI.GetTokenGenerators(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpTokenGeneratorsAPI.GetTokenGenerators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenGenerators`: TokenGenerators
	fmt.Fprintf(os.Stdout, "Response from `SpTokenGeneratorsAPI.GetTokenGenerators`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenGeneratorsRequest struct via the builder pattern


### Return type

[**TokenGenerators**](TokenGenerators.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTokenGenerator

> TokenGenerator UpdateTokenGenerator(ctx, id).Body(body).Execute()

Update a token generator instance.



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
	id := "id_example" // string | ID of the token generator instance.
	body := *openapiclient.NewTokenGenerator("Id_example", "Name_example", *openapiclient.NewResourceLink("Id_example"), *openapiclient.NewPluginConfiguration()) // TokenGenerator | Configuration for the updated token generator instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpTokenGeneratorsAPI.UpdateTokenGenerator(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpTokenGeneratorsAPI.UpdateTokenGenerator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTokenGenerator`: TokenGenerator
	fmt.Fprintf(os.Stdout, "Response from `SpTokenGeneratorsAPI.UpdateTokenGenerator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the token generator instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTokenGeneratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TokenGenerator**](TokenGenerator.md) | Configuration for the updated token generator instance. | 

### Return type

[**TokenGenerator**](TokenGenerator.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

