# IdpAttributeQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL at your IdP partner&#39;s site where attribute queries are to be sent. | 
**NameMappings** | Pointer to [**[]AttributeQueryNameMapping**](AttributeQueryNameMapping.md) | The attribute name mappings between the SP and the IdP. | [optional] 
**Policy** | Pointer to [**IdpAttributeQueryPolicy**](IdpAttributeQueryPolicy.md) |  | [optional] 

## Methods

### NewIdpAttributeQuery

`func NewIdpAttributeQuery(url string, ) *IdpAttributeQuery`

NewIdpAttributeQuery instantiates a new IdpAttributeQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpAttributeQueryWithDefaults

`func NewIdpAttributeQueryWithDefaults() *IdpAttributeQuery`

NewIdpAttributeQueryWithDefaults instantiates a new IdpAttributeQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *IdpAttributeQuery) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IdpAttributeQuery) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IdpAttributeQuery) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNameMappings

`func (o *IdpAttributeQuery) GetNameMappings() []AttributeQueryNameMapping`

GetNameMappings returns the NameMappings field if non-nil, zero value otherwise.

### GetNameMappingsOk

`func (o *IdpAttributeQuery) GetNameMappingsOk() (*[]AttributeQueryNameMapping, bool)`

GetNameMappingsOk returns a tuple with the NameMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameMappings

`func (o *IdpAttributeQuery) SetNameMappings(v []AttributeQueryNameMapping)`

SetNameMappings sets NameMappings field to given value.

### HasNameMappings

`func (o *IdpAttributeQuery) HasNameMappings() bool`

HasNameMappings returns a boolean if a field has been set.

### GetPolicy

`func (o *IdpAttributeQuery) GetPolicy() IdpAttributeQueryPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *IdpAttributeQuery) GetPolicyOk() (*IdpAttributeQueryPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *IdpAttributeQuery) SetPolicy(v IdpAttributeQueryPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *IdpAttributeQuery) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


