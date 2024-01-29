# \KerberosRealmsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKerberosRealm**](KerberosRealmsAPI.md#CreateKerberosRealm) | **Post** /kerberos/realms | Create a new Kerberos Realm.
[**DeleteKerberosRealm**](KerberosRealmsAPI.md#DeleteKerberosRealm) | **Delete** /kerberos/realms/{id} | Delete a Kerberos Realm.
[**GetKerberosRealm**](KerberosRealmsAPI.md#GetKerberosRealm) | **Get** /kerberos/realms/{id} | Find a Kerberos Realm by ID.
[**GetKerberosRealmSettings**](KerberosRealmsAPI.md#GetKerberosRealmSettings) | **Get** /kerberos/realms/settings | Gets the Kerberos Realms Settings.
[**GetKerberosRealms**](KerberosRealmsAPI.md#GetKerberosRealms) | **Get** /kerberos/realms | Gets the Kerberos Realms.
[**UpdateKerberosRealm**](KerberosRealmsAPI.md#UpdateKerberosRealm) | **Put** /kerberos/realms/{id} | Update a Kerberos Realm by ID.
[**UpdateKerberosRealmSettings**](KerberosRealmsAPI.md#UpdateKerberosRealmSettings) | **Put** /kerberos/realms/settings | Set/Update the Kerberos Realms Settings.



## CreateKerberosRealm

> KerberosRealm CreateKerberosRealm(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new Kerberos Realm.



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
    body := *openapiclient.NewKerberosRealm("KerberosRealmName_example") // KerberosRealm | Configuration for new policy.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Defaults to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KerberosRealmsAPI.CreateKerberosRealm(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosRealmsAPI.CreateKerberosRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKerberosRealm`: KerberosRealm
    fmt.Fprintf(os.Stdout, "Response from `KerberosRealmsAPI.CreateKerberosRealm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKerberosRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**KerberosRealm**](KerberosRealm.md) | Configuration for new policy. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Defaults to false. | [default to false]

### Return type

[**KerberosRealm**](KerberosRealm.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKerberosRealm

> DeleteKerberosRealm(ctx, id).Execute()

Delete a Kerberos Realm.



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
    id := "id_example" // string | ID of Kerberos Realm to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KerberosRealmsAPI.DeleteKerberosRealm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosRealmsAPI.DeleteKerberosRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Kerberos Realm to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKerberosRealmRequest struct via the builder pattern


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


## GetKerberosRealm

> KerberosRealm GetKerberosRealm(ctx, id).Execute()

Find a Kerberos Realm by ID.



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
    id := "id_example" // string | ID of the Kerberos Realm to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KerberosRealmsAPI.GetKerberosRealm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosRealmsAPI.GetKerberosRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKerberosRealm`: KerberosRealm
    fmt.Fprintf(os.Stdout, "Response from `KerberosRealmsAPI.GetKerberosRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Kerberos Realm to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKerberosRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KerberosRealm**](KerberosRealm.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKerberosRealmSettings

> KerberosRealmsSettings GetKerberosRealmSettings(ctx).Execute()

Gets the Kerberos Realms Settings.

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
    resp, r, err := apiClient.KerberosRealmsAPI.GetKerberosRealmSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosRealmsAPI.GetKerberosRealmSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKerberosRealmSettings`: KerberosRealmsSettings
    fmt.Fprintf(os.Stdout, "Response from `KerberosRealmsAPI.GetKerberosRealmSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKerberosRealmSettingsRequest struct via the builder pattern


### Return type

[**KerberosRealmsSettings**](KerberosRealmsSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKerberosRealms

> KerberosRealms GetKerberosRealms(ctx).Execute()

Gets the Kerberos Realms.

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
    resp, r, err := apiClient.KerberosRealmsAPI.GetKerberosRealms(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosRealmsAPI.GetKerberosRealms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKerberosRealms`: KerberosRealms
    fmt.Fprintf(os.Stdout, "Response from `KerberosRealmsAPI.GetKerberosRealms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKerberosRealmsRequest struct via the builder pattern


### Return type

[**KerberosRealms**](KerberosRealms.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKerberosRealm

> KerberosRealm UpdateKerberosRealm(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update a Kerberos Realm by ID.



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
    id := "id_example" // string | ID of the Kerberos Realm to update.
    body := *openapiclient.NewKerberosRealm("KerberosRealmName_example") // KerberosRealm | Configuration for updated Domain/Realm.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Defaults to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KerberosRealmsAPI.UpdateKerberosRealm(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosRealmsAPI.UpdateKerberosRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKerberosRealm`: KerberosRealm
    fmt.Fprintf(os.Stdout, "Response from `KerberosRealmsAPI.UpdateKerberosRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Kerberos Realm to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKerberosRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**KerberosRealm**](KerberosRealm.md) | Configuration for updated Domain/Realm. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Defaults to false. | [default to false]

### Return type

[**KerberosRealm**](KerberosRealm.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKerberosRealmSettings

> KerberosRealmsSettings UpdateKerberosRealmSettings(ctx).Body(body).Execute()

Set/Update the Kerberos Realms Settings.

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
    body := *openapiclient.NewKerberosRealmsSettings("KdcRetries_example", "KdcTimeout_example") // KerberosRealmsSettings | Kerberos Realms Settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KerberosRealmsAPI.UpdateKerberosRealmSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KerberosRealmsAPI.UpdateKerberosRealmSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKerberosRealmSettings`: KerberosRealmsSettings
    fmt.Fprintf(os.Stdout, "Response from `KerberosRealmsAPI.UpdateKerberosRealmSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKerberosRealmSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**KerberosRealmsSettings**](KerberosRealmsSettings.md) | Kerberos Realms Settings. | 

### Return type

[**KerberosRealmsSettings**](KerberosRealmsSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

