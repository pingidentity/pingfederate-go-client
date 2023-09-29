# SloServiceEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binding** | **string** | The binding of this endpoint, if applicable - usually only required for SAML 2.0 endpoints. | 
**Url** | **string** | The absolute or relative URL of the endpoint. A relative URL can be specified if a base URL for the connection has been defined. | 
**ResponseUrl** | Pointer to **string** | The absolute or relative URL to which logout responses are sent. A relative URL can be specified if a base URL for the connection has been defined. | [optional] 

## Methods

### NewSloServiceEndpoint

`func NewSloServiceEndpoint(binding string, url string, ) *SloServiceEndpoint`

NewSloServiceEndpoint instantiates a new SloServiceEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloServiceEndpointWithDefaults

`func NewSloServiceEndpointWithDefaults() *SloServiceEndpoint`

NewSloServiceEndpointWithDefaults instantiates a new SloServiceEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinding

`func (o *SloServiceEndpoint) GetBinding() string`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *SloServiceEndpoint) GetBindingOk() (*string, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *SloServiceEndpoint) SetBinding(v string)`

SetBinding sets Binding field to given value.


### GetUrl

`func (o *SloServiceEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SloServiceEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SloServiceEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetResponseUrl

`func (o *SloServiceEndpoint) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *SloServiceEndpoint) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *SloServiceEndpoint) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.

### HasResponseUrl

`func (o *SloServiceEndpoint) HasResponseUrl() bool`

HasResponseUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


