# TokenEndpointAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this attribute. | 
**MultiValued** | Pointer to **bool** | Indicates whether attribute value is always returned as an array. | [optional] 
**MappedScopes** | Pointer to **[]string** | List of scopes that will trigger this attribute to be included in the token endpoint response. | [optional] 

## Methods

### NewTokenEndpointAttribute

`func NewTokenEndpointAttribute(name string, ) *TokenEndpointAttribute`

NewTokenEndpointAttribute instantiates a new TokenEndpointAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenEndpointAttributeWithDefaults

`func NewTokenEndpointAttributeWithDefaults() *TokenEndpointAttribute`

NewTokenEndpointAttributeWithDefaults instantiates a new TokenEndpointAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TokenEndpointAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenEndpointAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenEndpointAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetMultiValued

`func (o *TokenEndpointAttribute) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *TokenEndpointAttribute) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *TokenEndpointAttribute) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.

### HasMultiValued

`func (o *TokenEndpointAttribute) HasMultiValued() bool`

HasMultiValued returns a boolean if a field has been set.

### GetMappedScopes

`func (o *TokenEndpointAttribute) GetMappedScopes() []string`

GetMappedScopes returns the MappedScopes field if non-nil, zero value otherwise.

### GetMappedScopesOk

`func (o *TokenEndpointAttribute) GetMappedScopesOk() (*[]string, bool)`

GetMappedScopesOk returns a tuple with the MappedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedScopes

`func (o *TokenEndpointAttribute) SetMappedScopes(v []string)`

SetMappedScopes sets MappedScopes field to given value.

### HasMappedScopes

`func (o *TokenEndpointAttribute) HasMappedScopes() bool`

HasMappedScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


