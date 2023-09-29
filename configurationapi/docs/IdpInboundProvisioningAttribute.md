# IdpInboundProvisioningAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this attribute. | 
**Masked** | Pointer to **bool** | Specifies whether this attribute is masked in PingFederate logs. Defaults to false. | [optional] 

## Methods

### NewIdpInboundProvisioningAttribute

`func NewIdpInboundProvisioningAttribute(name string, ) *IdpInboundProvisioningAttribute`

NewIdpInboundProvisioningAttribute instantiates a new IdpInboundProvisioningAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpInboundProvisioningAttributeWithDefaults

`func NewIdpInboundProvisioningAttributeWithDefaults() *IdpInboundProvisioningAttribute`

NewIdpInboundProvisioningAttributeWithDefaults instantiates a new IdpInboundProvisioningAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdpInboundProvisioningAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdpInboundProvisioningAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdpInboundProvisioningAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetMasked

`func (o *IdpInboundProvisioningAttribute) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *IdpInboundProvisioningAttribute) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *IdpInboundProvisioningAttribute) SetMasked(v bool)`

SetMasked sets Masked field to given value.

### HasMasked

`func (o *IdpInboundProvisioningAttribute) HasMasked() bool`

HasMasked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


