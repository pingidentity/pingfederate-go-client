# IdentityStoreProvisionerAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | [**[]Attribute**](Attribute.md) | A list of identity store provisioner attributes that correspond to the attributes exposed by the identity store provisioner type. | 
**ExtendedAttributes** | Pointer to [**[]Attribute**](Attribute.md) | A list of additional attributes that can be returned by the identity store provisioner. The extended attributes are only used if the provisioner supports them. | [optional] 

## Methods

### NewIdentityStoreProvisionerAttributeContract

`func NewIdentityStoreProvisionerAttributeContract(coreAttributes []Attribute, ) *IdentityStoreProvisionerAttributeContract`

NewIdentityStoreProvisionerAttributeContract instantiates a new IdentityStoreProvisionerAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityStoreProvisionerAttributeContractWithDefaults

`func NewIdentityStoreProvisionerAttributeContractWithDefaults() *IdentityStoreProvisionerAttributeContract`

NewIdentityStoreProvisionerAttributeContractWithDefaults instantiates a new IdentityStoreProvisionerAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *IdentityStoreProvisionerAttributeContract) GetCoreAttributes() []Attribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *IdentityStoreProvisionerAttributeContract) GetCoreAttributesOk() (*[]Attribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *IdentityStoreProvisionerAttributeContract) SetCoreAttributes(v []Attribute)`

SetCoreAttributes sets CoreAttributes field to given value.


### GetExtendedAttributes

`func (o *IdentityStoreProvisionerAttributeContract) GetExtendedAttributes() []Attribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *IdentityStoreProvisionerAttributeContract) GetExtendedAttributesOk() (*[]Attribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *IdentityStoreProvisionerAttributeContract) SetExtendedAttributes(v []Attribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *IdentityStoreProvisionerAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


