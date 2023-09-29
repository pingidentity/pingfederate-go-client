# SpAdapterUrlMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL that will be compared against the target URL. Use a wildcard (*) to match multiple URLs to the same adapter instance. | [optional] 
**AdapterRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 

## Methods

### NewSpAdapterUrlMapping

`func NewSpAdapterUrlMapping() *SpAdapterUrlMapping`

NewSpAdapterUrlMapping instantiates a new SpAdapterUrlMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpAdapterUrlMappingWithDefaults

`func NewSpAdapterUrlMappingWithDefaults() *SpAdapterUrlMapping`

NewSpAdapterUrlMappingWithDefaults instantiates a new SpAdapterUrlMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *SpAdapterUrlMapping) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SpAdapterUrlMapping) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SpAdapterUrlMapping) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SpAdapterUrlMapping) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAdapterRef

`func (o *SpAdapterUrlMapping) GetAdapterRef() ResourceLink`

GetAdapterRef returns the AdapterRef field if non-nil, zero value otherwise.

### GetAdapterRefOk

`func (o *SpAdapterUrlMapping) GetAdapterRefOk() (*ResourceLink, bool)`

GetAdapterRefOk returns a tuple with the AdapterRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterRef

`func (o *SpAdapterUrlMapping) SetAdapterRef(v ResourceLink)`

SetAdapterRef sets AdapterRef field to given value.

### HasAdapterRef

`func (o *SpAdapterUrlMapping) HasAdapterRef() bool`

HasAdapterRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


