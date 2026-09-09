# TokenGeneratorAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | [**[]TokenGeneratorAttribute**](TokenGeneratorAttribute.md) | A list of token generator attributes that correspond to the attributes exposed by the token generator type. | 
**ExtendedAttributes** | Pointer to [**[]TokenGeneratorAttribute**](TokenGeneratorAttribute.md) | A list of additional attributes that can be returned by the token processor. The extended attributes are only used if the token generator supports them. | [optional] 

## Methods

### NewTokenGeneratorAttributeContract

`func NewTokenGeneratorAttributeContract(coreAttributes []TokenGeneratorAttribute, ) *TokenGeneratorAttributeContract`

NewTokenGeneratorAttributeContract instantiates a new TokenGeneratorAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenGeneratorAttributeContractWithDefaults

`func NewTokenGeneratorAttributeContractWithDefaults() *TokenGeneratorAttributeContract`

NewTokenGeneratorAttributeContractWithDefaults instantiates a new TokenGeneratorAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *TokenGeneratorAttributeContract) GetCoreAttributes() []TokenGeneratorAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *TokenGeneratorAttributeContract) GetCoreAttributesOk() (*[]TokenGeneratorAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *TokenGeneratorAttributeContract) SetCoreAttributes(v []TokenGeneratorAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.


### GetExtendedAttributes

`func (o *TokenGeneratorAttributeContract) GetExtendedAttributes() []TokenGeneratorAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *TokenGeneratorAttributeContract) GetExtendedAttributesOk() (*[]TokenGeneratorAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *TokenGeneratorAttributeContract) SetExtendedAttributes(v []TokenGeneratorAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *TokenGeneratorAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


