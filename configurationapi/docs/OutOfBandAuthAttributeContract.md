# OutOfBandAuthAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | [**[]OutOfBandAuthAttribute**](OutOfBandAuthAttribute.md) | A list of out of band authenticator attributes. | 
**ExtendedAttributes** | Pointer to [**[]OutOfBandAuthAttribute**](OutOfBandAuthAttribute.md) | A list of additional attributes that can be returned by the out of band authenticator plugin instance. The extended attributes are only used if the plugin supports them. | [optional] 

## Methods

### NewOutOfBandAuthAttributeContract

`func NewOutOfBandAuthAttributeContract(coreAttributes []OutOfBandAuthAttribute, ) *OutOfBandAuthAttributeContract`

NewOutOfBandAuthAttributeContract instantiates a new OutOfBandAuthAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutOfBandAuthAttributeContractWithDefaults

`func NewOutOfBandAuthAttributeContractWithDefaults() *OutOfBandAuthAttributeContract`

NewOutOfBandAuthAttributeContractWithDefaults instantiates a new OutOfBandAuthAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *OutOfBandAuthAttributeContract) GetCoreAttributes() []OutOfBandAuthAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *OutOfBandAuthAttributeContract) GetCoreAttributesOk() (*[]OutOfBandAuthAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *OutOfBandAuthAttributeContract) SetCoreAttributes(v []OutOfBandAuthAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.


### GetExtendedAttributes

`func (o *OutOfBandAuthAttributeContract) GetExtendedAttributes() []OutOfBandAuthAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *OutOfBandAuthAttributeContract) GetExtendedAttributesOk() (*[]OutOfBandAuthAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *OutOfBandAuthAttributeContract) SetExtendedAttributes(v []OutOfBandAuthAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *OutOfBandAuthAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


