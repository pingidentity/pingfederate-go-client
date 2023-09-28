# \KeyPairsOauthOpenIdConnectApi

All URIs are relative to *https://localhost/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKeySet**](KeyPairsOauthOpenIdConnectApi.md#CreateKeySet) | **Post** /keyPairs/oauthOpenIdConnect/additionalKeySets | Create a new OAuth/OpenID Connect additional signing key set.
[**DeleteKeySet**](KeyPairsOauthOpenIdConnectApi.md#DeleteKeySet) | **Delete** /keyPairs/oauthOpenIdConnect/additionalKeySets/{id} | Delete an existing OAuth/OpenID Connect additional signing key set.
[**GetKeySet**](KeyPairsOauthOpenIdConnectApi.md#GetKeySet) | **Get** /keyPairs/oauthOpenIdConnect/additionalKeySets/{id} | Retrieve an OAuth/OpenID Connect additional signing key set.
[**GetKeySets**](KeyPairsOauthOpenIdConnectApi.md#GetKeySets) | **Get** /keyPairs/oauthOpenIdConnect/additionalKeySets | Retrieve OAuth/OpenID Connect additional signing key sets.
[**GetOauthOidcKeysSettings**](KeyPairsOauthOpenIdConnectApi.md#GetOauthOidcKeysSettings) | **Get** /keyPairs/oauthOpenIdConnect | Retrieve OAuth/OpenID Connect key settings.
[**UpdateKeySet**](KeyPairsOauthOpenIdConnectApi.md#UpdateKeySet) | **Put** /keyPairs/oauthOpenIdConnect/additionalKeySets/{id} | Update an existing OAuth/OpenID Connect additional signing key set.
[**UpdateOAuthOidcKeysSettings**](KeyPairsOauthOpenIdConnectApi.md#UpdateOAuthOidcKeysSettings) | **Put** /keyPairs/oauthOpenIdConnect | Update OAuth/OpenID Connect key settings.



## CreateKeySet

> AdditionalKeySet CreateKeySet(ctx).Body(body).Execute()

Create a new OAuth/OpenID Connect additional signing key set.



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
    body := *openapiclient.NewAdditionalKeySet("Name_example", *openapiclient.NewSigningKeys(), []openapiclient.ResourceLink{*openapiclient.NewResourceLink("Id_example")}) // AdditionalKeySet | Configuration for a new signing key set.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsOauthOpenIdConnectApi.CreateKeySet(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsOauthOpenIdConnectApi.CreateKeySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKeySet`: AdditionalKeySet
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsOauthOpenIdConnectApi.CreateKeySet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AdditionalKeySet**](AdditionalKeySet.md) | Configuration for a new signing key set. | 

### Return type

[**AdditionalKeySet**](AdditionalKeySet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKeySet

> DeleteKeySet(ctx, id).Execute()

Delete an existing OAuth/OpenID Connect additional signing key set.

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
    id := "id_example" // string | ID of an additional key set to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeyPairsOauthOpenIdConnectApi.DeleteKeySet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsOauthOpenIdConnectApi.DeleteKeySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an additional key set to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeySetRequest struct via the builder pattern


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


## GetKeySet

> AdditionalKeySet GetKeySet(ctx, id).Execute()

Retrieve an OAuth/OpenID Connect additional signing key set.

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
    id := "id_example" // string | ID of an OAuth/OpenID Connect additional signing key set to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsOauthOpenIdConnectApi.GetKeySet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsOauthOpenIdConnectApi.GetKeySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeySet`: AdditionalKeySet
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsOauthOpenIdConnectApi.GetKeySet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an OAuth/OpenID Connect additional signing key set to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdditionalKeySet**](AdditionalKeySet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeySets

> AdditionalKeySets GetKeySets(ctx).Execute()

Retrieve OAuth/OpenID Connect additional signing key sets.

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
    resp, r, err := apiClient.KeyPairsOauthOpenIdConnectApi.GetKeySets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsOauthOpenIdConnectApi.GetKeySets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeySets`: AdditionalKeySets
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsOauthOpenIdConnectApi.GetKeySets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeySetsRequest struct via the builder pattern


### Return type

[**AdditionalKeySets**](AdditionalKeySets.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthOidcKeysSettings

> OAuthOidcKeysSettings GetOauthOidcKeysSettings(ctx).Execute()

Retrieve OAuth/OpenID Connect key settings.

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
    resp, r, err := apiClient.KeyPairsOauthOpenIdConnectApi.GetOauthOidcKeysSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsOauthOpenIdConnectApi.GetOauthOidcKeysSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthOidcKeysSettings`: OAuthOidcKeysSettings
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsOauthOpenIdConnectApi.GetOauthOidcKeysSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthOidcKeysSettingsRequest struct via the builder pattern


### Return type

[**OAuthOidcKeysSettings**](OAuthOidcKeysSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKeySet

> AdditionalKeySet UpdateKeySet(ctx, id).Body(body).Execute()

Update an existing OAuth/OpenID Connect additional signing key set.



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
    id := "id_example" // string | ID of an OAuth/OpenID Connect additional signing key set to update.
    body := *openapiclient.NewAdditionalKeySet("Name_example", *openapiclient.NewSigningKeys(), []openapiclient.ResourceLink{*openapiclient.NewResourceLink("Id_example")}) // AdditionalKeySet | Configuration for updated OAuth/OpenID Connect additional signing key set.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsOauthOpenIdConnectApi.UpdateKeySet(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsOauthOpenIdConnectApi.UpdateKeySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKeySet`: AdditionalKeySet
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsOauthOpenIdConnectApi.UpdateKeySet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an OAuth/OpenID Connect additional signing key set to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AdditionalKeySet**](AdditionalKeySet.md) | Configuration for updated OAuth/OpenID Connect additional signing key set. | 

### Return type

[**AdditionalKeySet**](AdditionalKeySet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOAuthOidcKeysSettings

> OAuthOidcKeysSettings UpdateOAuthOidcKeysSettings(ctx).Body(body).Execute()

Update OAuth/OpenID Connect key settings.

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
    body := *openapiclient.NewOAuthOidcKeysSettings(false) // OAuthOidcKeysSettings | OAuth and OpenID Connect static key settings

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairsOauthOpenIdConnectApi.UpdateOAuthOidcKeysSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsOauthOpenIdConnectApi.UpdateOAuthOidcKeysSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOAuthOidcKeysSettings`: OAuthOidcKeysSettings
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsOauthOpenIdConnectApi.UpdateOAuthOidcKeysSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOAuthOidcKeysSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OAuthOidcKeysSettings**](OAuthOidcKeysSettings.md) | OAuth and OpenID Connect static key settings | 

### Return type

[**OAuthOidcKeysSettings**](OAuthOidcKeysSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

