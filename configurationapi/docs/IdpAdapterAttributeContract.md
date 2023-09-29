# IdpAdapterAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | [**[]IdpAdapterAttribute**](IdpAdapterAttribute.md) | A list of IdP adapter attributes that correspond to the attributes exposed by the IdP adapter type. | 
**ExtendedAttributes** | Pointer to [**[]IdpAdapterAttribute**](IdpAdapterAttribute.md) | A list of additional attributes that can be returned by the IdP adapter. The extended attributes are only used if the adapter supports them. | [optional] 
**UniqueUserKeyAttribute** | Pointer to **string** | The attribute to use for uniquely identify a user&#39;s authentication sessions. | [optional] 
**MaskOgnlValues** | Pointer to **bool** | Whether or not all OGNL expressions used to fulfill an outgoing assertion contract should be masked in the logs. Defaults to false. | [optional] 
**Inherited** | Pointer to **bool** | Whether this attribute contract is inherited from its parent instance. If true, the rest of the properties in this model become read-only. The default value is false. | [optional] 

## Methods

### NewIdpAdapterAttributeContract

`func NewIdpAdapterAttributeContract(coreAttributes []IdpAdapterAttribute, ) *IdpAdapterAttributeContract`

NewIdpAdapterAttributeContract instantiates a new IdpAdapterAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpAdapterAttributeContractWithDefaults

`func NewIdpAdapterAttributeContractWithDefaults() *IdpAdapterAttributeContract`

NewIdpAdapterAttributeContractWithDefaults instantiates a new IdpAdapterAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *IdpAdapterAttributeContract) GetCoreAttributes() []IdpAdapterAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *IdpAdapterAttributeContract) GetCoreAttributesOk() (*[]IdpAdapterAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *IdpAdapterAttributeContract) SetCoreAttributes(v []IdpAdapterAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.


### GetExtendedAttributes

`func (o *IdpAdapterAttributeContract) GetExtendedAttributes() []IdpAdapterAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *IdpAdapterAttributeContract) GetExtendedAttributesOk() (*[]IdpAdapterAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *IdpAdapterAttributeContract) SetExtendedAttributes(v []IdpAdapterAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *IdpAdapterAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.

### GetUniqueUserKeyAttribute

`func (o *IdpAdapterAttributeContract) GetUniqueUserKeyAttribute() string`

GetUniqueUserKeyAttribute returns the UniqueUserKeyAttribute field if non-nil, zero value otherwise.

### GetUniqueUserKeyAttributeOk

`func (o *IdpAdapterAttributeContract) GetUniqueUserKeyAttributeOk() (*string, bool)`

GetUniqueUserKeyAttributeOk returns a tuple with the UniqueUserKeyAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueUserKeyAttribute

`func (o *IdpAdapterAttributeContract) SetUniqueUserKeyAttribute(v string)`

SetUniqueUserKeyAttribute sets UniqueUserKeyAttribute field to given value.

### HasUniqueUserKeyAttribute

`func (o *IdpAdapterAttributeContract) HasUniqueUserKeyAttribute() bool`

HasUniqueUserKeyAttribute returns a boolean if a field has been set.

### GetMaskOgnlValues

`func (o *IdpAdapterAttributeContract) GetMaskOgnlValues() bool`

GetMaskOgnlValues returns the MaskOgnlValues field if non-nil, zero value otherwise.

### GetMaskOgnlValuesOk

`func (o *IdpAdapterAttributeContract) GetMaskOgnlValuesOk() (*bool, bool)`

GetMaskOgnlValuesOk returns a tuple with the MaskOgnlValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskOgnlValues

`func (o *IdpAdapterAttributeContract) SetMaskOgnlValues(v bool)`

SetMaskOgnlValues sets MaskOgnlValues field to given value.

### HasMaskOgnlValues

`func (o *IdpAdapterAttributeContract) HasMaskOgnlValues() bool`

HasMaskOgnlValues returns a boolean if a field has been set.

### GetInherited

`func (o *IdpAdapterAttributeContract) GetInherited() bool`

GetInherited returns the Inherited field if non-nil, zero value otherwise.

### GetInheritedOk

`func (o *IdpAdapterAttributeContract) GetInheritedOk() (*bool, bool)`

GetInheritedOk returns a tuple with the Inherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherited

`func (o *IdpAdapterAttributeContract) SetInherited(v bool)`

SetInherited sets Inherited field to given value.

### HasInherited

`func (o *IdpAdapterAttributeContract) HasInherited() bool`

HasInherited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


