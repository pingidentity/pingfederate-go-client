# SpSsoServiceEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binding** | **string** | The binding of this endpoint, if applicable - usually only required for SAML 2.0 endpoints.  Supported bindings are Artifact and POST. | 
**Url** | **string** | The absolute or relative URL of the endpoint. A relative URL can be specified if a base URL for the connection has been defined. | 
**IsDefault** | Pointer to **bool** | Whether or not this endpoint is the default endpoint. Defaults to false. | [optional] 
**Index** | **int64** | The priority of the endpoint. | 

## Methods

### NewSpSsoServiceEndpoint

`func NewSpSsoServiceEndpoint(binding string, url string, index int64, ) *SpSsoServiceEndpoint`

NewSpSsoServiceEndpoint instantiates a new SpSsoServiceEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpSsoServiceEndpointWithDefaults

`func NewSpSsoServiceEndpointWithDefaults() *SpSsoServiceEndpoint`

NewSpSsoServiceEndpointWithDefaults instantiates a new SpSsoServiceEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinding

`func (o *SpSsoServiceEndpoint) GetBinding() string`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *SpSsoServiceEndpoint) GetBindingOk() (*string, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *SpSsoServiceEndpoint) SetBinding(v string)`

SetBinding sets Binding field to given value.


### GetUrl

`func (o *SpSsoServiceEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SpSsoServiceEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SpSsoServiceEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetIsDefault

`func (o *SpSsoServiceEndpoint) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *SpSsoServiceEndpoint) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *SpSsoServiceEndpoint) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *SpSsoServiceEndpoint) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIndex

`func (o *SpSsoServiceEndpoint) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SpSsoServiceEndpoint) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SpSsoServiceEndpoint) SetIndex(v int64)`

SetIndex sets Index field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


