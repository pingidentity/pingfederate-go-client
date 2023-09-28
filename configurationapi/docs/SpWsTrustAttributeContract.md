# SpWsTrustAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | Pointer to [**[]SpWsTrustAttribute**](SpWsTrustAttribute.md) | A list of read-only assertion attributes that are automatically populated by PingFederate. | [optional] 
**ExtendedAttributes** | Pointer to [**[]SpWsTrustAttribute**](SpWsTrustAttribute.md) | A list of additional attributes that are added to the outgoing assertion. | [optional] 

## Methods

### NewSpWsTrustAttributeContract

`func NewSpWsTrustAttributeContract() *SpWsTrustAttributeContract`

NewSpWsTrustAttributeContract instantiates a new SpWsTrustAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpWsTrustAttributeContractWithDefaults

`func NewSpWsTrustAttributeContractWithDefaults() *SpWsTrustAttributeContract`

NewSpWsTrustAttributeContractWithDefaults instantiates a new SpWsTrustAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *SpWsTrustAttributeContract) GetCoreAttributes() []SpWsTrustAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *SpWsTrustAttributeContract) GetCoreAttributesOk() (*[]SpWsTrustAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *SpWsTrustAttributeContract) SetCoreAttributes(v []SpWsTrustAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.

### HasCoreAttributes

`func (o *SpWsTrustAttributeContract) HasCoreAttributes() bool`

HasCoreAttributes returns a boolean if a field has been set.

### GetExtendedAttributes

`func (o *SpWsTrustAttributeContract) GetExtendedAttributes() []SpWsTrustAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *SpWsTrustAttributeContract) GetExtendedAttributesOk() (*[]SpWsTrustAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *SpWsTrustAttributeContract) SetExtendedAttributes(v []SpWsTrustAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *SpWsTrustAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


