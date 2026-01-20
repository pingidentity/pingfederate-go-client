# SpAdapterAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | Pointer to [**[]SpAdapterAttribute**](SpAdapterAttribute.md) | A list of read-only attributes that are automatically populated by the SP adapter descriptor. | [optional] 
**ExtendedAttributes** | Pointer to [**[]SpAdapterAttribute**](SpAdapterAttribute.md) | A list of additional attributes that can be returned by the SP adapter. The extended attributes are only used if the adapter supports them. | [optional] 
**Inherited** | Pointer to **bool** | Whether this attribute contract is inherited from its parent instance. If true, the rest of the properties in this model become read-only. The default value is false. | [optional] 

## Methods

### NewSpAdapterAttributeContract

`func NewSpAdapterAttributeContract() *SpAdapterAttributeContract`

NewSpAdapterAttributeContract instantiates a new SpAdapterAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpAdapterAttributeContractWithDefaults

`func NewSpAdapterAttributeContractWithDefaults() *SpAdapterAttributeContract`

NewSpAdapterAttributeContractWithDefaults instantiates a new SpAdapterAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *SpAdapterAttributeContract) GetCoreAttributes() []SpAdapterAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *SpAdapterAttributeContract) GetCoreAttributesOk() (*[]SpAdapterAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *SpAdapterAttributeContract) SetCoreAttributes(v []SpAdapterAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.

### HasCoreAttributes

`func (o *SpAdapterAttributeContract) HasCoreAttributes() bool`

HasCoreAttributes returns a boolean if a field has been set.

### GetExtendedAttributes

`func (o *SpAdapterAttributeContract) GetExtendedAttributes() []SpAdapterAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *SpAdapterAttributeContract) GetExtendedAttributesOk() (*[]SpAdapterAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *SpAdapterAttributeContract) SetExtendedAttributes(v []SpAdapterAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *SpAdapterAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.

### GetInherited

`func (o *SpAdapterAttributeContract) GetInherited() bool`

GetInherited returns the Inherited field if non-nil, zero value otherwise.

### GetInheritedOk

`func (o *SpAdapterAttributeContract) GetInheritedOk() (*bool, bool)`

GetInheritedOk returns a tuple with the Inherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherited

`func (o *SpAdapterAttributeContract) SetInherited(v bool)`

SetInherited sets Inherited field to given value.

### HasInherited

`func (o *SpAdapterAttributeContract) HasInherited() bool`

HasInherited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


