# IdpSsoServiceEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binding** | Pointer to **string** | The binding of this endpoint, if applicable - usually only required for SAML 2.0 endpoints. | [optional] 
**Url** | **string** | The absolute or relative URL of the endpoint. A relative URL can be specified if a base URL for the connection has been defined. | 

## Methods

### NewIdpSsoServiceEndpoint

`func NewIdpSsoServiceEndpoint(url string, ) *IdpSsoServiceEndpoint`

NewIdpSsoServiceEndpoint instantiates a new IdpSsoServiceEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpSsoServiceEndpointWithDefaults

`func NewIdpSsoServiceEndpointWithDefaults() *IdpSsoServiceEndpoint`

NewIdpSsoServiceEndpointWithDefaults instantiates a new IdpSsoServiceEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinding

`func (o *IdpSsoServiceEndpoint) GetBinding() string`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *IdpSsoServiceEndpoint) GetBindingOk() (*string, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *IdpSsoServiceEndpoint) SetBinding(v string)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *IdpSsoServiceEndpoint) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

### GetUrl

`func (o *IdpSsoServiceEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IdpSsoServiceEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IdpSsoServiceEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


