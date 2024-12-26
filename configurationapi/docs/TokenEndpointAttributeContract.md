# TokenEndpointAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]TokenEndpointAttribute**](TokenEndpointAttribute.md) | A list of token endpoint response attributes that are associated with this access token management plugin instance. | [optional] 
**Inherited** | Pointer to **bool** | Whether this attribute contract is inherited from its parent instance. If true, the rest of the properties in this model become read-only. The default value is false. | [optional] 

## Methods

### NewTokenEndpointAttributeContract

`func NewTokenEndpointAttributeContract() *TokenEndpointAttributeContract`

NewTokenEndpointAttributeContract instantiates a new TokenEndpointAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenEndpointAttributeContractWithDefaults

`func NewTokenEndpointAttributeContractWithDefaults() *TokenEndpointAttributeContract`

NewTokenEndpointAttributeContractWithDefaults instantiates a new TokenEndpointAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *TokenEndpointAttributeContract) GetAttributes() []TokenEndpointAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TokenEndpointAttributeContract) GetAttributesOk() (*[]TokenEndpointAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TokenEndpointAttributeContract) SetAttributes(v []TokenEndpointAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TokenEndpointAttributeContract) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetInherited

`func (o *TokenEndpointAttributeContract) GetInherited() bool`

GetInherited returns the Inherited field if non-nil, zero value otherwise.

### GetInheritedOk

`func (o *TokenEndpointAttributeContract) GetInheritedOk() (*bool, bool)`

GetInheritedOk returns a tuple with the Inherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherited

`func (o *TokenEndpointAttributeContract) SetInherited(v bool)`

SetInherited sets Inherited field to given value.

### HasInherited

`func (o *TokenEndpointAttributeContract) HasInherited() bool`

HasInherited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


