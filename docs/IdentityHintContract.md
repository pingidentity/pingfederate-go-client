# IdentityHintContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | [**[]IdentityHintAttribute**](IdentityHintAttribute.md) | A list of required identity hint contract attributes. | 
**ExtendedAttributes** | Pointer to [**[]IdentityHintAttribute**](IdentityHintAttribute.md) | A list of additional identity hint contract attributes. | [optional] 

## Methods

### NewIdentityHintContract

`func NewIdentityHintContract(coreAttributes []IdentityHintAttribute, ) *IdentityHintContract`

NewIdentityHintContract instantiates a new IdentityHintContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityHintContractWithDefaults

`func NewIdentityHintContractWithDefaults() *IdentityHintContract`

NewIdentityHintContractWithDefaults instantiates a new IdentityHintContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *IdentityHintContract) GetCoreAttributes() []IdentityHintAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *IdentityHintContract) GetCoreAttributesOk() (*[]IdentityHintAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *IdentityHintContract) SetCoreAttributes(v []IdentityHintAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.


### GetExtendedAttributes

`func (o *IdentityHintContract) GetExtendedAttributes() []IdentityHintAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *IdentityHintContract) GetExtendedAttributesOk() (*[]IdentityHintAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *IdentityHintContract) SetExtendedAttributes(v []IdentityHintAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *IdentityHintContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


