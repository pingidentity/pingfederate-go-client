# \AdministrativeAccountsAPI

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccount**](AdministrativeAccountsAPI.md#AddAccount) | **Post** /administrativeAccounts | Add a new PingFederate native Administrative Account.
[**ChangePassword**](AdministrativeAccountsAPI.md#ChangePassword) | **Post** /administrativeAccounts/changePassword | Change the Password of current PingFederate native Account.
[**DeleteAccount**](AdministrativeAccountsAPI.md#DeleteAccount) | **Delete** /administrativeAccounts/{username} | Delete a PingFederate native Administrative Account information.
[**GetAccount**](AdministrativeAccountsAPI.md#GetAccount) | **Get** /administrativeAccounts/{username} | Get a PingFederate native Administrative Account.
[**GetAccounts**](AdministrativeAccountsAPI.md#GetAccounts) | **Get** /administrativeAccounts | Get all the PingFederate native Administrative Accounts.
[**ResetPassword**](AdministrativeAccountsAPI.md#ResetPassword) | **Post** /administrativeAccounts/{username}/resetPassword | Reset the Password of an existing PingFederate native Administrative Account.
[**UpdateAccount**](AdministrativeAccountsAPI.md#UpdateAccount) | **Put** /administrativeAccounts/{username} | Update the information for a native Administrative Account.



## AddAccount

> AdministrativeAccount AddAccount(ctx).Body(body).Execute()

Add a new PingFederate native Administrative Account.

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
    body := *openapiclient.NewAdministrativeAccount("Username_example") // AdministrativeAccount | Administrative account information.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministrativeAccountsAPI.AddAccount(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrativeAccountsAPI.AddAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAccount`: AdministrativeAccount
    fmt.Fprintf(os.Stdout, "Response from `AdministrativeAccountsAPI.AddAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AdministrativeAccount**](AdministrativeAccount.md) | Administrative account information. | 

### Return type

[**AdministrativeAccount**](AdministrativeAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangePassword

> UserCredentials ChangePassword(ctx).Body(body).Execute()

Change the Password of current PingFederate native Account.

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
    body := *openapiclient.NewUserCredentials("NewPassword_example") // UserCredentials | User Account credential.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministrativeAccountsAPI.ChangePassword(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrativeAccountsAPI.ChangePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangePassword`: UserCredentials
    fmt.Fprintf(os.Stdout, "Response from `AdministrativeAccountsAPI.ChangePassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserCredentials**](UserCredentials.md) | User Account credential. | 

### Return type

[**UserCredentials**](UserCredentials.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccount

> DeleteAccount(ctx, username).Execute()

Delete a PingFederate native Administrative Account information.

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
    username := "username_example" // string | Username of the account to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdministrativeAccountsAPI.DeleteAccount(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrativeAccountsAPI.DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of the account to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


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


## GetAccount

> AdministrativeAccount GetAccount(ctx, username).Execute()

Get a PingFederate native Administrative Account.

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
    username := "username_example" // string | Username of the administrative account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministrativeAccountsAPI.GetAccount(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrativeAccountsAPI.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: AdministrativeAccount
    fmt.Fprintf(os.Stdout, "Response from `AdministrativeAccountsAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of the administrative account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdministrativeAccount**](AdministrativeAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccounts

> AdministrativeAccounts GetAccounts(ctx).Execute()

Get all the PingFederate native Administrative Accounts.

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
    resp, r, err := apiClient.AdministrativeAccountsAPI.GetAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrativeAccountsAPI.GetAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccounts`: AdministrativeAccounts
    fmt.Fprintf(os.Stdout, "Response from `AdministrativeAccountsAPI.GetAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsRequest struct via the builder pattern


### Return type

[**AdministrativeAccounts**](AdministrativeAccounts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPassword

> UserCredentials ResetPassword(ctx, username).Body(body).Execute()

Reset the Password of an existing PingFederate native Administrative Account.

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
    username := "username_example" // string | Username of the administrative account.
    body := *openapiclient.NewUserCredentials("NewPassword_example") // UserCredentials | New password.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministrativeAccountsAPI.ResetPassword(context.Background(), username).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrativeAccountsAPI.ResetPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetPassword`: UserCredentials
    fmt.Fprintf(os.Stdout, "Response from `AdministrativeAccountsAPI.ResetPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of the administrative account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UserCredentials**](UserCredentials.md) | New password. | 

### Return type

[**UserCredentials**](UserCredentials.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> AdministrativeAccount UpdateAccount(ctx, username).Body(body).Execute()

Update the information for a native Administrative Account.

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
    username := "username_example" // string | Username of the account to be updated.
    body := *openapiclient.NewAdministrativeAccount("Username_example") // AdministrativeAccount | Administrative account information.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministrativeAccountsAPI.UpdateAccount(context.Background(), username).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministrativeAccountsAPI.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccount`: AdministrativeAccount
    fmt.Fprintf(os.Stdout, "Response from `AdministrativeAccountsAPI.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of the account to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AdministrativeAccount**](AdministrativeAccount.md) | Administrative account information. | 

### Return type

[**AdministrativeAccount**](AdministrativeAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

