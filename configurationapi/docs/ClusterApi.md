# \ClusterAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterAdminNodeStatus**](ClusterAPI.md#GetClusterAdminNodeStatus) | **Get** /cluster/adminNode/status | Get this administrative console&#39;s role and synchronization status.
[**GetClusterSettings**](ClusterAPI.md#GetClusterSettings) | **Get** /cluster/settings | Get the cluster configuration settings.
[**GetClusterStatus**](ClusterAPI.md#GetClusterStatus) | **Get** /cluster/status | Get information on the current status of the cluster.
[**StartReplication**](ClusterAPI.md#StartReplication) | **Post** /cluster/replicate | Replicate configuration updates to all nodes in the cluster.
[**UpdateClusterAdminNodeRole**](ClusterAPI.md#UpdateClusterAdminNodeRole) | **Post** /cluster/adminNode/role/active | Update this administrative console node&#39;s role to active. Possibly responds with warnings related to the update process.
[**UpdateClusterSettings**](ClusterAPI.md#UpdateClusterSettings) | **Put** /cluster/settings | Update the cluster configuration settings.



## GetClusterAdminNodeStatus

> AdminConsoleInfo GetClusterAdminNodeStatus(ctx).Execute()

Get this administrative console's role and synchronization status.

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
    resp, r, err := apiClient.ClusterAPI.GetClusterAdminNodeStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.GetClusterAdminNodeStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterAdminNodeStatus`: AdminConsoleInfo
    fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.GetClusterAdminNodeStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterAdminNodeStatusRequest struct via the builder pattern


### Return type

[**AdminConsoleInfo**](AdminConsoleInfo.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterSettings

> ClusterSettings GetClusterSettings(ctx).Execute()

Get the cluster configuration settings.

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
    resp, r, err := apiClient.ClusterAPI.GetClusterSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.GetClusterSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterSettings`: ClusterSettings
    fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.GetClusterSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterSettingsRequest struct via the builder pattern


### Return type

[**ClusterSettings**](ClusterSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterStatus

> ClusterStatus GetClusterStatus(ctx).Execute()

Get information on the current status of the cluster.

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
    resp, r, err := apiClient.ClusterAPI.GetClusterStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.GetClusterStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterStatus`: ClusterStatus
    fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.GetClusterStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterStatusRequest struct via the builder pattern


### Return type

[**ClusterStatus**](ClusterStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartReplication

> ApiResult StartReplication(ctx).Execute()

Replicate configuration updates to all nodes in the cluster.

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
    resp, r, err := apiClient.ClusterAPI.StartReplication(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.StartReplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartReplication`: ApiResult
    fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.StartReplication`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStartReplicationRequest struct via the builder pattern


### Return type

[**ApiResult**](ApiResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClusterAdminNodeRole

> AdminNodeRoleServiceInfo UpdateClusterAdminNodeRole(ctx).Execute()

Update this administrative console node's role to active. Possibly responds with warnings related to the update process.

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
    resp, r, err := apiClient.ClusterAPI.UpdateClusterAdminNodeRole(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.UpdateClusterAdminNodeRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClusterAdminNodeRole`: AdminNodeRoleServiceInfo
    fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.UpdateClusterAdminNodeRole`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterAdminNodeRoleRequest struct via the builder pattern


### Return type

[**AdminNodeRoleServiceInfo**](AdminNodeRoleServiceInfo.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClusterSettings

> ClusterSettings UpdateClusterSettings(ctx).Body(body).Execute()

Update the cluster configuration settings.

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
    body := *openapiclient.NewClusterSettings() // ClusterSettings | Configuration for cluster settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterAPI.UpdateClusterSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.UpdateClusterSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClusterSettings`: ClusterSettings
    fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.UpdateClusterSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterSettings**](ClusterSettings.md) | Configuration for cluster settings. | 

### Return type

[**ClusterSettings**](ClusterSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

