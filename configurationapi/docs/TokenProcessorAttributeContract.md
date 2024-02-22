# TokenProcessorAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | [**[]TokenProcessorAttribute**](TokenProcessorAttribute.md) | A list of token processor attributes that correspond to the attributes exposed by the token processor type. | 
**ExtendedAttributes** | Pointer to [**[]TokenProcessorAttribute**](TokenProcessorAttribute.md) | A list of additional attributes that can be returned by the token processor. The extended attributes are only used if the token processor supports them. | [optional] 
**MaskOgnlValues** | Pointer to **bool** | Whether or not all OGNL expressions used to fulfill an outgoing assertion contract should be masked in the logs. Defaults to false. | [optional] 

## Methods

### NewTokenProcessorAttributeContract

`func NewTokenProcessorAttributeContract(coreAttributes []TokenProcessorAttribute, ) *TokenProcessorAttributeContract`

NewTokenProcessorAttributeContract instantiates a new TokenProcessorAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenProcessorAttributeContractWithDefaults

`func NewTokenProcessorAttributeContractWithDefaults() *TokenProcessorAttributeContract`

NewTokenProcessorAttributeContractWithDefaults instantiates a new TokenProcessorAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *TokenProcessorAttributeContract) GetCoreAttributes() []TokenProcessorAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *TokenProcessorAttributeContract) GetCoreAttributesOk() (*[]TokenProcessorAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *TokenProcessorAttributeContract) SetCoreAttributes(v []TokenProcessorAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.


### GetExtendedAttributes

`func (o *TokenProcessorAttributeContract) GetExtendedAttributes() []TokenProcessorAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *TokenProcessorAttributeContract) GetExtendedAttributesOk() (*[]TokenProcessorAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *TokenProcessorAttributeContract) SetExtendedAttributes(v []TokenProcessorAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *TokenProcessorAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.

### GetMaskOgnlValues

`func (o *TokenProcessorAttributeContract) GetMaskOgnlValues() bool`

GetMaskOgnlValues returns the MaskOgnlValues field if non-nil, zero value otherwise.

### GetMaskOgnlValuesOk

`func (o *TokenProcessorAttributeContract) GetMaskOgnlValuesOk() (*bool, bool)`

GetMaskOgnlValuesOk returns a tuple with the MaskOgnlValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskOgnlValues

`func (o *TokenProcessorAttributeContract) SetMaskOgnlValues(v bool)`

SetMaskOgnlValues sets MaskOgnlValues field to given value.

### HasMaskOgnlValues

`func (o *TokenProcessorAttributeContract) HasMaskOgnlValues() bool`

HasMaskOgnlValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


