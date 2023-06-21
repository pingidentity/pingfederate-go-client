# IdpInboundProvisioningAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | Pointer to [**[]IdpInboundProvisioningAttribute**](IdpInboundProvisioningAttribute.md) | A list of read-only assertion attributes that are automatically populated by PingFederate. | [optional] 
**ExtendedAttributes** | Pointer to [**[]IdpInboundProvisioningAttribute**](IdpInboundProvisioningAttribute.md) | A list of additional attributes that are added to the SCIM response. | [optional] 

## Methods

### NewIdpInboundProvisioningAttributeContract

`func NewIdpInboundProvisioningAttributeContract() *IdpInboundProvisioningAttributeContract`

NewIdpInboundProvisioningAttributeContract instantiates a new IdpInboundProvisioningAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpInboundProvisioningAttributeContractWithDefaults

`func NewIdpInboundProvisioningAttributeContractWithDefaults() *IdpInboundProvisioningAttributeContract`

NewIdpInboundProvisioningAttributeContractWithDefaults instantiates a new IdpInboundProvisioningAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *IdpInboundProvisioningAttributeContract) GetCoreAttributes() []IdpInboundProvisioningAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *IdpInboundProvisioningAttributeContract) GetCoreAttributesOk() (*[]IdpInboundProvisioningAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *IdpInboundProvisioningAttributeContract) SetCoreAttributes(v []IdpInboundProvisioningAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.

### HasCoreAttributes

`func (o *IdpInboundProvisioningAttributeContract) HasCoreAttributes() bool`

HasCoreAttributes returns a boolean if a field has been set.

### GetExtendedAttributes

`func (o *IdpInboundProvisioningAttributeContract) GetExtendedAttributes() []IdpInboundProvisioningAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *IdpInboundProvisioningAttributeContract) GetExtendedAttributesOk() (*[]IdpInboundProvisioningAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *IdpInboundProvisioningAttributeContract) SetExtendedAttributes(v []IdpInboundProvisioningAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *IdpInboundProvisioningAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


