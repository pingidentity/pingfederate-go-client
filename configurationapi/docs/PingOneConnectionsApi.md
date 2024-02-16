# \PingOneConnectionsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePingOneConnection**](PingOneConnectionsAPI.md#CreatePingOneConnection) | **Post** /pingOneConnections | Create a new PingOne connection.
[**DeletePingOneConnection**](PingOneConnectionsAPI.md#DeletePingOneConnection) | **Delete** /pingOneConnections/{id} | Delete a PingOne connection.
[**GetCredentialStatus**](PingOneConnectionsAPI.md#GetCredentialStatus) | **Get** /pingOneConnections/{id}/credentialStatus | Get the status of the credential associated with the PingOne connection
[**GetPingOneConnection**](PingOneConnectionsAPI.md#GetPingOneConnection) | **Get** /pingOneConnections/{id} | Get a PingOne connection with the specified ID.
[**GetPingOneConnectionAssociations**](PingOneConnectionsAPI.md#GetPingOneConnectionAssociations) | **Get** /pingOneConnections/{id}/serviceAssociations | Get information about components using this connection to access PingOne services.
[**GetPingOneConnectionEnvironments**](PingOneConnectionsAPI.md#GetPingOneConnectionEnvironments) | **Get** /pingOneConnections/{id}/environments | Get the list of environments that the PingOne connection has access to.
[**GetPingOneConnectionUsages**](PingOneConnectionsAPI.md#GetPingOneConnectionUsages) | **Get** /pingOneConnections/{id}/usage | Get the list of resources that reference this PingOne connection.
[**GetPingOneConnections**](PingOneConnectionsAPI.md#GetPingOneConnections) | **Get** /pingOneConnections | Get the list of all PingOne connections.
[**UpdatePingOneConnection**](PingOneConnectionsAPI.md#UpdatePingOneConnection) | **Put** /pingOneConnections/{id} | Update a PingOne connection.



## CreatePingOneConnection

> PingOneConnection CreatePingOneConnection(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new PingOne connection.

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
	body := *openapiclient.NewPingOneConnection("Name_example") // PingOneConnection | Configuration for the new PingOne connection.
	xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PingOneConnectionsAPI.CreatePingOneConnection(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PingOneConnectionsAPI.CreatePingOneConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePingOneConnection`: PingOneConnection
	fmt.Fprintf(os.Stdout, "Response from `PingOneConnectionsAPI.CreatePingOneConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePingOneConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PingOneConnection**](PingOneConnection.md) | Configuration for the new PingOne connection. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**PingOneConnection**](PingOneConnection.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePingOneConnection

> DeletePingOneConnection(ctx, id).Execute()

Delete a PingOne connection.

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
	id := "id_example" // string | ID of the PingOne connection to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PingOneConnectionsAPI.DeletePingOneConnection(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PingOneConnectionsAPI.DeletePingOneConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the PingOne connection to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePingOneConnectionRequest struct via the builder pattern


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


## GetCredentialStatus

> PingOneCredentialStatus GetCredentialStatus(ctx, id).Execute()

Get the status of the credential associated with the PingOne connection

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
	id := "id_example" // string | ID of the PingOne connection.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PingOneConnectionsAPI.GetCredentialStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PingOneConnectionsAPI.GetCredentialStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredentialStatus`: PingOneCredentialStatus
	fmt.Fprintf(os.Stdout, "Response from `PingOneConnectionsAPI.GetCredentialStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the PingOne connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PingOneCredentialStatus**](PingOneCredentialStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingOneConnection

> PingOneConnection GetPingOneConnection(ctx, id).Execute()

Get a PingOne connection with the specified ID.

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
	id := "id_example" // string | ID of the connection to fetch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PingOneConnectionsAPI.GetPingOneConnection(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PingOneConnectionsAPI.GetPingOneConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPingOneConnection`: PingOneConnection
	fmt.Fprintf(os.Stdout, "Response from `PingOneConnectionsAPI.GetPingOneConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the connection to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingOneConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PingOneConnection**](PingOneConnection.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingOneConnectionAssociations

> ServiceAssociations GetPingOneConnectionAssociations(ctx, id).Execute()

Get information about components using this connection to access PingOne services.

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
	id := "id_example" // string | ID of the PingOne connection.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PingOneConnectionsAPI.GetPingOneConnectionAssociations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PingOneConnectionsAPI.GetPingOneConnectionAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPingOneConnectionAssociations`: ServiceAssociations
	fmt.Fprintf(os.Stdout, "Response from `PingOneConnectionsAPI.GetPingOneConnectionAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the PingOne connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingOneConnectionAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAssociations**](ServiceAssociations.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingOneConnectionEnvironments

> PingOneEnvironments GetPingOneConnectionEnvironments(ctx, id).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()

Get the list of environments that the PingOne connection has access to.

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
	id := "id_example" // string | ID of the PingOne connection.
	page := int64(56) // int64 | Page number to retrieve. (optional)
	numberPerPage := int64(56) // int64 | Number of environments per page. (optional)
	filter := "filter_example" // string | Filter criteria limits the environments that are returned to only those that match it. The filter criteria is compared to the environment name and ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PingOneConnectionsAPI.GetPingOneConnectionEnvironments(context.Background(), id).Page(page).NumberPerPage(numberPerPage).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PingOneConnectionsAPI.GetPingOneConnectionEnvironments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPingOneConnectionEnvironments`: PingOneEnvironments
	fmt.Fprintf(os.Stdout, "Response from `PingOneConnectionsAPI.GetPingOneConnectionEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the PingOne connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingOneConnectionEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | Page number to retrieve. | 
 **numberPerPage** | **int64** | Number of environments per page. | 
 **filter** | **string** | Filter criteria limits the environments that are returned to only those that match it. The filter criteria is compared to the environment name and ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. | 

### Return type

[**PingOneEnvironments**](PingOneEnvironments.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingOneConnectionUsages

> ResourceUsages GetPingOneConnectionUsages(ctx, id).Execute()

Get the list of resources that reference this PingOne connection.

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
	id := "id_example" // string | ID of the PingOne connection.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PingOneConnectionsAPI.GetPingOneConnectionUsages(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PingOneConnectionsAPI.GetPingOneConnectionUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPingOneConnectionUsages`: ResourceUsages
	fmt.Fprintf(os.Stdout, "Response from `PingOneConnectionsAPI.GetPingOneConnectionUsages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the PingOne connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingOneConnectionUsagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourceUsages**](ResourceUsages.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingOneConnections

> PingOneConnections GetPingOneConnections(ctx).Execute()

Get the list of all PingOne connections.

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
	resp, r, err := apiClient.PingOneConnectionsAPI.GetPingOneConnections(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PingOneConnectionsAPI.GetPingOneConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPingOneConnections`: PingOneConnections
	fmt.Fprintf(os.Stdout, "Response from `PingOneConnectionsAPI.GetPingOneConnections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingOneConnectionsRequest struct via the builder pattern


### Return type

[**PingOneConnections**](PingOneConnections.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePingOneConnection

> PingOneConnection UpdatePingOneConnection(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update a PingOne connection.

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
	id := "id_example" // string | ID of the PingOne connection to update.
	body := *openapiclient.NewPingOneConnection("Name_example") // PingOneConnection | Configuration for the updated connection.
	xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PingOneConnectionsAPI.UpdatePingOneConnection(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PingOneConnectionsAPI.UpdatePingOneConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePingOneConnection`: PingOneConnection
	fmt.Fprintf(os.Stdout, "Response from `PingOneConnectionsAPI.UpdatePingOneConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the PingOne connection to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePingOneConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PingOneConnection**](PingOneConnection.md) | Configuration for the updated connection. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**PingOneConnection**](PingOneConnection.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

