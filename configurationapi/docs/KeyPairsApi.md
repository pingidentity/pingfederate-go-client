# \KeyPairsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKeyAlgorithms**](KeyPairsAPI.md#GetKeyAlgorithms) | **Get** /keyPairs/keyAlgorithms | Get list of the key algorithms supported for key pair generation.



## GetKeyAlgorithms

> KeyAlgorithms GetKeyAlgorithms(ctx).Execute()

Get list of the key algorithms supported for key pair generation.

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
    resp, r, err := apiClient.KeyPairsAPI.GetKeyAlgorithms(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsAPI.GetKeyAlgorithms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeyAlgorithms`: KeyAlgorithms
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsAPI.GetKeyAlgorithms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyAlgorithmsRequest struct via the builder pattern


### Return type

[**KeyAlgorithms**](KeyAlgorithms.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

