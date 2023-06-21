# IdpWsTrustAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | Pointer to [**[]IdpWsTrustAttribute**](IdpWsTrustAttribute.md) | A list of read-only assertion attributes that are automatically populated by PingFederate. | [optional] 
**ExtendedAttributes** | Pointer to [**[]IdpWsTrustAttribute**](IdpWsTrustAttribute.md) | A list of additional attributes that are receive in the incoming assertion. | [optional] 

## Methods

### NewIdpWsTrustAttributeContract

`func NewIdpWsTrustAttributeContract() *IdpWsTrustAttributeContract`

NewIdpWsTrustAttributeContract instantiates a new IdpWsTrustAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpWsTrustAttributeContractWithDefaults

`func NewIdpWsTrustAttributeContractWithDefaults() *IdpWsTrustAttributeContract`

NewIdpWsTrustAttributeContractWithDefaults instantiates a new IdpWsTrustAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *IdpWsTrustAttributeContract) GetCoreAttributes() []IdpWsTrustAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *IdpWsTrustAttributeContract) GetCoreAttributesOk() (*[]IdpWsTrustAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *IdpWsTrustAttributeContract) SetCoreAttributes(v []IdpWsTrustAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.

### HasCoreAttributes

`func (o *IdpWsTrustAttributeContract) HasCoreAttributes() bool`

HasCoreAttributes returns a boolean if a field has been set.

### GetExtendedAttributes

`func (o *IdpWsTrustAttributeContract) GetExtendedAttributes() []IdpWsTrustAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *IdpWsTrustAttributeContract) GetExtendedAttributesOk() (*[]IdpWsTrustAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *IdpWsTrustAttributeContract) SetExtendedAttributes(v []IdpWsTrustAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *IdpWsTrustAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


