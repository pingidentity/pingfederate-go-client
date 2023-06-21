# TokenExchangeProcessorAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | Pointer to [**[]TokenExchangeProcessorAttribute**](TokenExchangeProcessorAttribute.md) | A list of read-only attributes (for example, subject) that are automatically populated by PingFederate. | [optional] 
**ExtendedAttributes** | Pointer to [**[]TokenExchangeProcessorAttribute**](TokenExchangeProcessorAttribute.md) | A list of additional attributes. | [optional] 

## Methods

### NewTokenExchangeProcessorAttributeContract

`func NewTokenExchangeProcessorAttributeContract() *TokenExchangeProcessorAttributeContract`

NewTokenExchangeProcessorAttributeContract instantiates a new TokenExchangeProcessorAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenExchangeProcessorAttributeContractWithDefaults

`func NewTokenExchangeProcessorAttributeContractWithDefaults() *TokenExchangeProcessorAttributeContract`

NewTokenExchangeProcessorAttributeContractWithDefaults instantiates a new TokenExchangeProcessorAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *TokenExchangeProcessorAttributeContract) GetCoreAttributes() []TokenExchangeProcessorAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *TokenExchangeProcessorAttributeContract) GetCoreAttributesOk() (*[]TokenExchangeProcessorAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *TokenExchangeProcessorAttributeContract) SetCoreAttributes(v []TokenExchangeProcessorAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.

### HasCoreAttributes

`func (o *TokenExchangeProcessorAttributeContract) HasCoreAttributes() bool`

HasCoreAttributes returns a boolean if a field has been set.

### GetExtendedAttributes

`func (o *TokenExchangeProcessorAttributeContract) GetExtendedAttributes() []TokenExchangeProcessorAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *TokenExchangeProcessorAttributeContract) GetExtendedAttributesOk() (*[]TokenExchangeProcessorAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *TokenExchangeProcessorAttributeContract) SetExtendedAttributes(v []TokenExchangeProcessorAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *TokenExchangeProcessorAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


