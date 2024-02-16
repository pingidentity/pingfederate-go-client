# \IdpDefaultUrlsAPI

All URIs are relative to *https://localhost:9999/pf-admin-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDefaultUrl**](IdpDefaultUrlsAPI.md#GetDefaultUrl) | **Get** /idp/defaultUrls | Gets the IDP Default URL settings.
[**UpdateDefaultUrlSettings**](IdpDefaultUrlsAPI.md#UpdateDefaultUrlSettings) | **Put** /idp/defaultUrls | Update the IDP Default URL settings.



## GetDefaultUrl

> IdpDefaultUrl GetDefaultUrl(ctx).Execute()

Gets the IDP Default URL settings.

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
	resp, r, err := apiClient.IdpDefaultUrlsAPI.GetDefaultUrl(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdpDefaultUrlsAPI.GetDefaultUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultUrl`: IdpDefaultUrl
	fmt.Fprintf(os.Stdout, "Response from `IdpDefaultUrlsAPI.GetDefaultUrl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultUrlRequest struct via the builder pattern


### Return type

[**IdpDefaultUrl**](IdpDefaultUrl.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDefaultUrlSettings

> IdpDefaultUrl UpdateDefaultUrlSettings(ctx).Body(body).Execute()

Update the IDP Default URL settings.

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
	body := *openapiclient.NewIdpDefaultUrl("IdpErrorMsg_example") // IdpDefaultUrl | Configuration for the IdP Default URL settings.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdpDefaultUrlsAPI.UpdateDefaultUrlSettings(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdpDefaultUrlsAPI.UpdateDefaultUrlSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDefaultUrlSettings`: IdpDefaultUrl
	fmt.Fprintf(os.Stdout, "Response from `IdpDefaultUrlsAPI.UpdateDefaultUrlSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDefaultUrlSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IdpDefaultUrl**](IdpDefaultUrl.md) | Configuration for the IdP Default URL settings. | 

### Return type

[**IdpDefaultUrl**](IdpDefaultUrl.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

