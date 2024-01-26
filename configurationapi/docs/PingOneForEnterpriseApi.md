# \PingOneForEnterpriseAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Disconnect**](PingOneForEnterpriseAPI.md#Disconnect) | **Post** /pingOneForEnterprise/disconnect | Disconnect from PingOne for Enterprise
[**GetKeyPairs**](PingOneForEnterpriseAPI.md#GetKeyPairs) | **Get** /pingOneForEnterprise/keyPairs | Get the PingOne for Enterprise key pair settings
[**GetPingOneForEnterpriseSettings**](PingOneForEnterpriseAPI.md#GetPingOneForEnterpriseSettings) | **Get** /pingOneForEnterprise | Get the PingOne for Enterprise settings
[**RotateKeys**](PingOneForEnterpriseAPI.md#RotateKeys) | **Post** /pingOneForEnterprise/keyPairs/rotate | Rotate the authentication key
[**UpdatePingOneForEnterpriseIdentityRepository**](PingOneForEnterpriseAPI.md#UpdatePingOneForEnterpriseIdentityRepository) | **Post** /pingOneForEnterprise/updateIdentityRepository | Update the PingOne Identity Repository
[**UpdatePingOneSettings**](PingOneForEnterpriseAPI.md#UpdatePingOneSettings) | **Put** /pingOneForEnterprise | Update the PingOne for Enterprise settings.



## Disconnect

> PingOneForEnterpriseSettings Disconnect(ctx).Execute()

Disconnect from PingOne for Enterprise

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
    resp, r, err := apiClient.PingOneForEnterpriseAPI.Disconnect(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneForEnterpriseAPI.Disconnect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Disconnect`: PingOneForEnterpriseSettings
    fmt.Fprintf(os.Stdout, "Response from `PingOneForEnterpriseAPI.Disconnect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectRequest struct via the builder pattern


### Return type

[**PingOneForEnterpriseSettings**](PingOneForEnterpriseSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyPairs

> P14EKeysView GetKeyPairs(ctx).Execute()

Get the PingOne for Enterprise key pair settings

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
    resp, r, err := apiClient.PingOneForEnterpriseAPI.GetKeyPairs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneForEnterpriseAPI.GetKeyPairs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeyPairs`: P14EKeysView
    fmt.Fprintf(os.Stdout, "Response from `PingOneForEnterpriseAPI.GetKeyPairs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyPairsRequest struct via the builder pattern


### Return type

[**P14EKeysView**](P14EKeysView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingOneForEnterpriseSettings

> PingOneForEnterpriseSettings GetPingOneForEnterpriseSettings(ctx).Execute()

Get the PingOne for Enterprise settings

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
    resp, r, err := apiClient.PingOneForEnterpriseAPI.GetPingOneForEnterpriseSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneForEnterpriseAPI.GetPingOneForEnterpriseSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPingOneForEnterpriseSettings`: PingOneForEnterpriseSettings
    fmt.Fprintf(os.Stdout, "Response from `PingOneForEnterpriseAPI.GetPingOneForEnterpriseSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingOneForEnterpriseSettingsRequest struct via the builder pattern


### Return type

[**PingOneForEnterpriseSettings**](PingOneForEnterpriseSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateKeys

> P14EKeysView RotateKeys(ctx).Execute()

Rotate the authentication key

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
    resp, r, err := apiClient.PingOneForEnterpriseAPI.RotateKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneForEnterpriseAPI.RotateKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RotateKeys`: P14EKeysView
    fmt.Fprintf(os.Stdout, "Response from `PingOneForEnterpriseAPI.RotateKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRotateKeysRequest struct via the builder pattern


### Return type

[**P14EKeysView**](P14EKeysView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePingOneForEnterpriseIdentityRepository

> PingOneForEnterpriseSettings UpdatePingOneForEnterpriseIdentityRepository(ctx).Execute()

Update the PingOne Identity Repository



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
    resp, r, err := apiClient.PingOneForEnterpriseAPI.UpdatePingOneForEnterpriseIdentityRepository(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneForEnterpriseAPI.UpdatePingOneForEnterpriseIdentityRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePingOneForEnterpriseIdentityRepository`: PingOneForEnterpriseSettings
    fmt.Fprintf(os.Stdout, "Response from `PingOneForEnterpriseAPI.UpdatePingOneForEnterpriseIdentityRepository`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePingOneForEnterpriseIdentityRepositoryRequest struct via the builder pattern


### Return type

[**PingOneForEnterpriseSettings**](PingOneForEnterpriseSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePingOneSettings

> PingOneForEnterpriseSettings UpdatePingOneSettings(ctx).Body(body).Execute()

Update the PingOne for Enterprise settings.

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
    body := *openapiclient.NewPingOneForEnterpriseSettings() // PingOneForEnterpriseSettings | PingOne for Enterprise connection settings

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingOneForEnterpriseAPI.UpdatePingOneSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneForEnterpriseAPI.UpdatePingOneSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePingOneSettings`: PingOneForEnterpriseSettings
    fmt.Fprintf(os.Stdout, "Response from `PingOneForEnterpriseAPI.UpdatePingOneSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePingOneSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PingOneForEnterpriseSettings**](PingOneForEnterpriseSettings.md) | PingOne for Enterprise connection settings | 

### Return type

[**PingOneForEnterpriseSettings**](PingOneForEnterpriseSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

