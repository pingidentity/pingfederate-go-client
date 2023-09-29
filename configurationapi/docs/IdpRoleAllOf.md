# IdpRoleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Enable Identity Provider Role. | [optional] 
**Saml20Profile** | Pointer to [**SAML20Profile**](SAML20Profile.md) |  | [optional] 
**EnableOutboundProvisioning** | Pointer to **bool** | Enable Outbound Provisioning. | [optional] 

## Methods

### NewIdpRoleAllOf

`func NewIdpRoleAllOf() *IdpRoleAllOf`

NewIdpRoleAllOf instantiates a new IdpRoleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpRoleAllOfWithDefaults

`func NewIdpRoleAllOfWithDefaults() *IdpRoleAllOf`

NewIdpRoleAllOfWithDefaults instantiates a new IdpRoleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *IdpRoleAllOf) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IdpRoleAllOf) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IdpRoleAllOf) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *IdpRoleAllOf) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetSaml20Profile

`func (o *IdpRoleAllOf) GetSaml20Profile() SAML20Profile`

GetSaml20Profile returns the Saml20Profile field if non-nil, zero value otherwise.

### GetSaml20ProfileOk

`func (o *IdpRoleAllOf) GetSaml20ProfileOk() (*SAML20Profile, bool)`

GetSaml20ProfileOk returns a tuple with the Saml20Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml20Profile

`func (o *IdpRoleAllOf) SetSaml20Profile(v SAML20Profile)`

SetSaml20Profile sets Saml20Profile field to given value.

### HasSaml20Profile

`func (o *IdpRoleAllOf) HasSaml20Profile() bool`

HasSaml20Profile returns a boolean if a field has been set.

### GetEnableOutboundProvisioning

`func (o *IdpRoleAllOf) GetEnableOutboundProvisioning() bool`

GetEnableOutboundProvisioning returns the EnableOutboundProvisioning field if non-nil, zero value otherwise.

### GetEnableOutboundProvisioningOk

`func (o *IdpRoleAllOf) GetEnableOutboundProvisioningOk() (*bool, bool)`

GetEnableOutboundProvisioningOk returns a tuple with the EnableOutboundProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOutboundProvisioning

`func (o *IdpRoleAllOf) SetEnableOutboundProvisioning(v bool)`

SetEnableOutboundProvisioning sets EnableOutboundProvisioning field to given value.

### HasEnableOutboundProvisioning

`func (o *IdpRoleAllOf) HasEnableOutboundProvisioning() bool`

HasEnableOutboundProvisioning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


