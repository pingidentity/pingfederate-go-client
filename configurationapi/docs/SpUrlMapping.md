# SpUrlMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL that will be compared against the target URL. Use a wildcard (*) to match multiple URLs to the same adapter or connection instance. | [optional] 
**Type** | Pointer to **string** | The URL mapping type | [optional] 
**Ref** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 

## Methods

### NewSpUrlMapping

`func NewSpUrlMapping() *SpUrlMapping`

NewSpUrlMapping instantiates a new SpUrlMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpUrlMappingWithDefaults

`func NewSpUrlMappingWithDefaults() *SpUrlMapping`

NewSpUrlMappingWithDefaults instantiates a new SpUrlMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *SpUrlMapping) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SpUrlMapping) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SpUrlMapping) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SpUrlMapping) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetType

`func (o *SpUrlMapping) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpUrlMapping) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpUrlMapping) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SpUrlMapping) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRef

`func (o *SpUrlMapping) GetRef() ResourceLink`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *SpUrlMapping) GetRefOk() (*ResourceLink, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *SpUrlMapping) SetRef(v ResourceLink)`

SetRef sets Ref field to given value.

### HasRef

`func (o *SpUrlMapping) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


