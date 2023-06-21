# IdpRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Enable Identity Provider Role. | [optional] 
**EnableSaml11** | Pointer to **bool** | Enable SAML 1.1. | [optional] 
**EnableSaml10** | Pointer to **bool** | Enable SAML 1.0. | [optional] 
**EnableWsFed** | Pointer to **bool** | Enable WS Federation. | [optional] 
**EnableWsTrust** | Pointer to **bool** | Enable WS Trust. | [optional] 
**Saml20Profile** | Pointer to [**SAML20Profile**](SAML20Profile.md) |  | [optional] 
**EnableOutboundProvisioning** | Pointer to **bool** | Enable Outbound Provisioning. | [optional] 

## Methods

### NewIdpRole

`func NewIdpRole() *IdpRole`

NewIdpRole instantiates a new IdpRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpRoleWithDefaults

`func NewIdpRoleWithDefaults() *IdpRole`

NewIdpRoleWithDefaults instantiates a new IdpRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *IdpRole) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IdpRole) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IdpRole) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *IdpRole) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableSaml11

`func (o *IdpRole) GetEnableSaml11() bool`

GetEnableSaml11 returns the EnableSaml11 field if non-nil, zero value otherwise.

### GetEnableSaml11Ok

`func (o *IdpRole) GetEnableSaml11Ok() (*bool, bool)`

GetEnableSaml11Ok returns a tuple with the EnableSaml11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSaml11

`func (o *IdpRole) SetEnableSaml11(v bool)`

SetEnableSaml11 sets EnableSaml11 field to given value.

### HasEnableSaml11

`func (o *IdpRole) HasEnableSaml11() bool`

HasEnableSaml11 returns a boolean if a field has been set.

### GetEnableSaml10

`func (o *IdpRole) GetEnableSaml10() bool`

GetEnableSaml10 returns the EnableSaml10 field if non-nil, zero value otherwise.

### GetEnableSaml10Ok

`func (o *IdpRole) GetEnableSaml10Ok() (*bool, bool)`

GetEnableSaml10Ok returns a tuple with the EnableSaml10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSaml10

`func (o *IdpRole) SetEnableSaml10(v bool)`

SetEnableSaml10 sets EnableSaml10 field to given value.

### HasEnableSaml10

`func (o *IdpRole) HasEnableSaml10() bool`

HasEnableSaml10 returns a boolean if a field has been set.

### GetEnableWsFed

`func (o *IdpRole) GetEnableWsFed() bool`

GetEnableWsFed returns the EnableWsFed field if non-nil, zero value otherwise.

### GetEnableWsFedOk

`func (o *IdpRole) GetEnableWsFedOk() (*bool, bool)`

GetEnableWsFedOk returns a tuple with the EnableWsFed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWsFed

`func (o *IdpRole) SetEnableWsFed(v bool)`

SetEnableWsFed sets EnableWsFed field to given value.

### HasEnableWsFed

`func (o *IdpRole) HasEnableWsFed() bool`

HasEnableWsFed returns a boolean if a field has been set.

### GetEnableWsTrust

`func (o *IdpRole) GetEnableWsTrust() bool`

GetEnableWsTrust returns the EnableWsTrust field if non-nil, zero value otherwise.

### GetEnableWsTrustOk

`func (o *IdpRole) GetEnableWsTrustOk() (*bool, bool)`

GetEnableWsTrustOk returns a tuple with the EnableWsTrust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWsTrust

`func (o *IdpRole) SetEnableWsTrust(v bool)`

SetEnableWsTrust sets EnableWsTrust field to given value.

### HasEnableWsTrust

`func (o *IdpRole) HasEnableWsTrust() bool`

HasEnableWsTrust returns a boolean if a field has been set.

### GetSaml20Profile

`func (o *IdpRole) GetSaml20Profile() SAML20Profile`

GetSaml20Profile returns the Saml20Profile field if non-nil, zero value otherwise.

### GetSaml20ProfileOk

`func (o *IdpRole) GetSaml20ProfileOk() (*SAML20Profile, bool)`

GetSaml20ProfileOk returns a tuple with the Saml20Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml20Profile

`func (o *IdpRole) SetSaml20Profile(v SAML20Profile)`

SetSaml20Profile sets Saml20Profile field to given value.

### HasSaml20Profile

`func (o *IdpRole) HasSaml20Profile() bool`

HasSaml20Profile returns a boolean if a field has been set.

### GetEnableOutboundProvisioning

`func (o *IdpRole) GetEnableOutboundProvisioning() bool`

GetEnableOutboundProvisioning returns the EnableOutboundProvisioning field if non-nil, zero value otherwise.

### GetEnableOutboundProvisioningOk

`func (o *IdpRole) GetEnableOutboundProvisioningOk() (*bool, bool)`

GetEnableOutboundProvisioningOk returns a tuple with the EnableOutboundProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOutboundProvisioning

`func (o *IdpRole) SetEnableOutboundProvisioning(v bool)`

SetEnableOutboundProvisioning sets EnableOutboundProvisioning field to given value.

### HasEnableOutboundProvisioning

`func (o *IdpRole) HasEnableOutboundProvisioning() bool`

HasEnableOutboundProvisioning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


