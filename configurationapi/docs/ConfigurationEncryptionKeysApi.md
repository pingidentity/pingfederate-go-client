# \ConfigurationEncryptionKeysApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigurationEncryptionKeys**](ConfigurationEncryptionKeysApi.md#GetConfigurationEncryptionKeys) | **Get** /configurationEncryptionKeys | Get the list of Configuration Encryption Keys.
[**RotateConfigurationEncryptionKey**](ConfigurationEncryptionKeysApi.md#RotateConfigurationEncryptionKey) | **Post** /configurationEncryptionKeys/rotate | Rotate the current Configuration Encryption Key.



## GetConfigurationEncryptionKeys

> ConfigurationEncryptionKeys GetConfigurationEncryptionKeys(ctx).Execute()

Get the list of Configuration Encryption Keys.



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
    resp, r, err := apiClient.ConfigurationEncryptionKeysApi.GetConfigurationEncryptionKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationEncryptionKeysApi.GetConfigurationEncryptionKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigurationEncryptionKeys`: ConfigurationEncryptionKeys
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationEncryptionKeysApi.GetConfigurationEncryptionKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationEncryptionKeysRequest struct via the builder pattern


### Return type

[**ConfigurationEncryptionKeys**](ConfigurationEncryptionKeys.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateConfigurationEncryptionKey

> ConfigurationEncryptionKeys RotateConfigurationEncryptionKey(ctx).Execute()

Rotate the current Configuration Encryption Key.



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
    resp, r, err := apiClient.ConfigurationEncryptionKeysApi.RotateConfigurationEncryptionKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationEncryptionKeysApi.RotateConfigurationEncryptionKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RotateConfigurationEncryptionKey`: ConfigurationEncryptionKeys
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationEncryptionKeysApi.RotateConfigurationEncryptionKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRotateConfigurationEncryptionKeyRequest struct via the builder pattern


### Return type

[**ConfigurationEncryptionKeys**](ConfigurationEncryptionKeys.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

