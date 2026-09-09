# IdentityStoreProvisionerGroupAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | [**[]GroupAttribute**](GroupAttribute.md) | A list of identity store provisioner group attributes that correspond to the group attributes exposed by the identity store provisioner type. | 
**ExtendedAttributes** | Pointer to [**[]GroupAttribute**](GroupAttribute.md) | A list of additional group attributes that can be returned by the identity store provisioner. The extended group attributes are only used if the provisioner supports them. | [optional] 

## Methods

### NewIdentityStoreProvisionerGroupAttributeContract

`func NewIdentityStoreProvisionerGroupAttributeContract(coreAttributes []GroupAttribute, ) *IdentityStoreProvisionerGroupAttributeContract`

NewIdentityStoreProvisionerGroupAttributeContract instantiates a new IdentityStoreProvisionerGroupAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityStoreProvisionerGroupAttributeContractWithDefaults

`func NewIdentityStoreProvisionerGroupAttributeContractWithDefaults() *IdentityStoreProvisionerGroupAttributeContract`

NewIdentityStoreProvisionerGroupAttributeContractWithDefaults instantiates a new IdentityStoreProvisionerGroupAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *IdentityStoreProvisionerGroupAttributeContract) GetCoreAttributes() []GroupAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *IdentityStoreProvisionerGroupAttributeContract) GetCoreAttributesOk() (*[]GroupAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *IdentityStoreProvisionerGroupAttributeContract) SetCoreAttributes(v []GroupAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.


### GetExtendedAttributes

`func (o *IdentityStoreProvisionerGroupAttributeContract) GetExtendedAttributes() []GroupAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *IdentityStoreProvisionerGroupAttributeContract) GetExtendedAttributesOk() (*[]GroupAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *IdentityStoreProvisionerGroupAttributeContract) SetExtendedAttributes(v []GroupAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *IdentityStoreProvisionerGroupAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


