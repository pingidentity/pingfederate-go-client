# \LocalIdentityIdentityProfilesAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentityProfile**](LocalIdentityIdentityProfilesAPI.md#CreateIdentityProfile) | **Post** /localIdentity/identityProfiles | Create a new local identity profile.
[**DeleteIdentityProfile**](LocalIdentityIdentityProfilesAPI.md#DeleteIdentityProfile) | **Delete** /localIdentity/identityProfiles/{id} | Delete the local identity profile by ID.
[**GetIdentityProfile**](LocalIdentityIdentityProfilesAPI.md#GetIdentityProfile) | **Get** /localIdentity/identityProfiles/{id} | Get the local identity profile by ID.
[**GetIdentityProfiles**](LocalIdentityIdentityProfilesAPI.md#GetIdentityProfiles) | **Get** /localIdentity/identityProfiles | Get the list of configured local identity profiles.
[**UpdateIdentityProfile**](LocalIdentityIdentityProfilesAPI.md#UpdateIdentityProfile) | **Put** /localIdentity/identityProfiles/{id} | Update the local identity profile by ID.



## CreateIdentityProfile

> LocalIdentityProfile CreateIdentityProfile(ctx).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Create a new local identity profile.



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
    body := *openapiclient.NewLocalIdentityProfile("Name_example", *openapiclient.NewResourceLink("Id_example")) // LocalIdentityProfile | Configuration for a new profile.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalIdentityIdentityProfilesAPI.CreateIdentityProfile(context.Background()).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalIdentityIdentityProfilesAPI.CreateIdentityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdentityProfile`: LocalIdentityProfile
    fmt.Fprintf(os.Stdout, "Response from `LocalIdentityIdentityProfilesAPI.CreateIdentityProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LocalIdentityProfile**](LocalIdentityProfile.md) | Configuration for a new profile. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**LocalIdentityProfile**](LocalIdentityProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityProfile

> DeleteIdentityProfile(ctx, id).Execute()

Delete the local identity profile by ID.



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
    id := "id_example" // string | ID of the profile to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LocalIdentityIdentityProfilesAPI.DeleteIdentityProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalIdentityIdentityProfilesAPI.DeleteIdentityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the profile to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityProfileRequest struct via the builder pattern


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


## GetIdentityProfile

> LocalIdentityProfile GetIdentityProfile(ctx, id).Execute()

Get the local identity profile by ID.



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
    id := "id_example" // string | ID of profile to fetch

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalIdentityIdentityProfilesAPI.GetIdentityProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalIdentityIdentityProfilesAPI.GetIdentityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProfile`: LocalIdentityProfile
    fmt.Fprintf(os.Stdout, "Response from `LocalIdentityIdentityProfilesAPI.GetIdentityProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of profile to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LocalIdentityProfile**](LocalIdentityProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityProfiles

> LocalIdentityProfiles GetIdentityProfiles(ctx).Page(page).NumberPerPage(numberPerPage).Filter(filter).ApcId(apcId).Execute()

Get the list of configured local identity profiles.

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
    page := int64(56) // int64 | Page number to retrieve. (optional)
    numberPerPage := int64(56) // int64 | Number of local identity profiles per page. (optional)
    filter := "filter_example" // string | Filter criteria limits the local identity profiles that are returned to only those that match it. The filter criteria is compared to the local identity profile name and ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. (optional)
    apcId := "apcId_example" // string | Filter the local identity profiles by matching policy contract ID. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalIdentityIdentityProfilesAPI.GetIdentityProfiles(context.Background()).Page(page).NumberPerPage(numberPerPage).Filter(filter).ApcId(apcId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalIdentityIdentityProfilesAPI.GetIdentityProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProfiles`: LocalIdentityProfiles
    fmt.Fprintf(os.Stdout, "Response from `LocalIdentityIdentityProfilesAPI.GetIdentityProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number to retrieve. | 
 **numberPerPage** | **int64** | Number of local identity profiles per page. | 
 **filter** | **string** | Filter criteria limits the local identity profiles that are returned to only those that match it. The filter criteria is compared to the local identity profile name and ID fields. The comparison is a case-insensitive partial match. No additional pattern based matching is supported. | 
 **apcId** | **string** | Filter the local identity profiles by matching policy contract ID. | 

### Return type

[**LocalIdentityProfiles**](LocalIdentityProfiles.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentityProfile

> LocalIdentityProfile UpdateIdentityProfile(ctx, id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()

Update the local identity profile by ID.



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
    id := "id_example" // string | ID of the profile to update
    body := *openapiclient.NewLocalIdentityProfile("Name_example", *openapiclient.NewResourceLink("Id_example")) // LocalIdentityProfile | Configuration for updated local identity profile.
    xBypassExternalValidation := true // bool | External validation will be bypassed when set to true. Default to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalIdentityIdentityProfilesAPI.UpdateIdentityProfile(context.Background(), id).Body(body).XBypassExternalValidation(xBypassExternalValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalIdentityIdentityProfilesAPI.UpdateIdentityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdentityProfile`: LocalIdentityProfile
    fmt.Fprintf(os.Stdout, "Response from `LocalIdentityIdentityProfilesAPI.UpdateIdentityProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the profile to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentityProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LocalIdentityProfile**](LocalIdentityProfile.md) | Configuration for updated local identity profile. | 
 **xBypassExternalValidation** | **bool** | External validation will be bypassed when set to true. Default to false. | [default to false]

### Return type

[**LocalIdentityProfile**](LocalIdentityProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

