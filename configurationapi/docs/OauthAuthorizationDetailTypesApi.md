# \OauthAuthorizationDetailTypesAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAuthorizationDetailType**](OauthAuthorizationDetailTypesAPI.md#AddAuthorizationDetailType) | **Post** /oauth/authorizationDetailTypes | Create a new authorization detail type.
[**DeleteAuthorizationDetailType**](OauthAuthorizationDetailTypesAPI.md#DeleteAuthorizationDetailType) | **Delete** /oauth/authorizationDetailTypes/{id} | Delete an authorization detail type.
[**GetAuthorizationDetailTypeById**](OauthAuthorizationDetailTypesAPI.md#GetAuthorizationDetailTypeById) | **Get** /oauth/authorizationDetailTypes/{id} | Get an authorization detail type.
[**GetAuthorizationDetailTypes**](OauthAuthorizationDetailTypesAPI.md#GetAuthorizationDetailTypes) | **Get** /oauth/authorizationDetailTypes | Get the list of authorization detail types.
[**UpdateAuthorizationDetailType**](OauthAuthorizationDetailTypesAPI.md#UpdateAuthorizationDetailType) | **Put** /oauth/authorizationDetailTypes/{id} | Update an authorization detail type.



## AddAuthorizationDetailType

> AuthorizationDetailType AddAuthorizationDetailType(ctx).Body(body).Execute()

Create a new authorization detail type.



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
    body := *openapiclient.NewAuthorizationDetailType("Type_example", *openapiclient.NewResourceLink("Id_example")) // AuthorizationDetailType | Configuration for new authorization detail type.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthorizationDetailTypesAPI.AddAuthorizationDetailType(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailTypesAPI.AddAuthorizationDetailType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAuthorizationDetailType`: AuthorizationDetailType
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthorizationDetailTypesAPI.AddAuthorizationDetailType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAuthorizationDetailTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthorizationDetailType**](AuthorizationDetailType.md) | Configuration for new authorization detail type. | 

### Return type

[**AuthorizationDetailType**](AuthorizationDetailType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthorizationDetailType

> DeleteAuthorizationDetailType(ctx, id).Execute()

Delete an authorization detail type.



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
    id := "id_example" // string | ID of the authorization detail type.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthAuthorizationDetailTypesAPI.DeleteAuthorizationDetailType(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailTypesAPI.DeleteAuthorizationDetailType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the authorization detail type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthorizationDetailTypeRequest struct via the builder pattern


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


## GetAuthorizationDetailTypeById

> AuthorizationDetailType GetAuthorizationDetailTypeById(ctx, id).Execute()

Get an authorization detail type.



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
    id := "id_example" // string | ID of the authorization detail type.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthorizationDetailTypesAPI.GetAuthorizationDetailTypeById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailTypesAPI.GetAuthorizationDetailTypeById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationDetailTypeById`: AuthorizationDetailType
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthorizationDetailTypesAPI.GetAuthorizationDetailTypeById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the authorization detail type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationDetailTypeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthorizationDetailType**](AuthorizationDetailType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorizationDetailTypes

> AuthorizationDetailTypes GetAuthorizationDetailTypes(ctx).Execute()

Get the list of authorization detail types.

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
    resp, r, err := apiClient.OauthAuthorizationDetailTypesAPI.GetAuthorizationDetailTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailTypesAPI.GetAuthorizationDetailTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationDetailTypes`: AuthorizationDetailTypes
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthorizationDetailTypesAPI.GetAuthorizationDetailTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationDetailTypesRequest struct via the builder pattern


### Return type

[**AuthorizationDetailTypes**](AuthorizationDetailTypes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthorizationDetailType

> AuthorizationDetailType UpdateAuthorizationDetailType(ctx, id).Body(body).Execute()

Update an authorization detail type.



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
    id := "id_example" // string | ID of the authorization detail type.
    body := *openapiclient.NewAuthorizationDetailType("Type_example", *openapiclient.NewResourceLink("Id_example")) // AuthorizationDetailType | Configuration for updated authorization detail type.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthAuthorizationDetailTypesAPI.UpdateAuthorizationDetailType(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthAuthorizationDetailTypesAPI.UpdateAuthorizationDetailType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthorizationDetailType`: AuthorizationDetailType
    fmt.Fprintf(os.Stdout, "Response from `OauthAuthorizationDetailTypesAPI.UpdateAuthorizationDetailType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the authorization detail type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthorizationDetailTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AuthorizationDetailType**](AuthorizationDetailType.md) | Configuration for updated authorization detail type. | 

### Return type

[**AuthorizationDetailType**](AuthorizationDetailType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

