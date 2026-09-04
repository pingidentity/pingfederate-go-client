# SpOAuthTokenExchangeAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | Pointer to [**[]Attribute**](Attribute.md) | A list of read-only attributes that are automatically populated by PingFederate. | [optional] 
**ExtendedAttributes** | Pointer to [**[]Attribute**](Attribute.md) | A list of additional attributes that are added to the outgoing token. | [optional] 

## Methods

### NewSpOAuthTokenExchangeAttributeContract

`func NewSpOAuthTokenExchangeAttributeContract() *SpOAuthTokenExchangeAttributeContract`

NewSpOAuthTokenExchangeAttributeContract instantiates a new SpOAuthTokenExchangeAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpOAuthTokenExchangeAttributeContractWithDefaults

`func NewSpOAuthTokenExchangeAttributeContractWithDefaults() *SpOAuthTokenExchangeAttributeContract`

NewSpOAuthTokenExchangeAttributeContractWithDefaults instantiates a new SpOAuthTokenExchangeAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *SpOAuthTokenExchangeAttributeContract) GetCoreAttributes() []Attribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *SpOAuthTokenExchangeAttributeContract) GetCoreAttributesOk() (*[]Attribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *SpOAuthTokenExchangeAttributeContract) SetCoreAttributes(v []Attribute)`

SetCoreAttributes sets CoreAttributes field to given value.

### HasCoreAttributes

`func (o *SpOAuthTokenExchangeAttributeContract) HasCoreAttributes() bool`

HasCoreAttributes returns a boolean if a field has been set.

### GetExtendedAttributes

`func (o *SpOAuthTokenExchangeAttributeContract) GetExtendedAttributes() []Attribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *SpOAuthTokenExchangeAttributeContract) GetExtendedAttributesOk() (*[]Attribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *SpOAuthTokenExchangeAttributeContract) SetExtendedAttributes(v []Attribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *SpOAuthTokenExchangeAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


