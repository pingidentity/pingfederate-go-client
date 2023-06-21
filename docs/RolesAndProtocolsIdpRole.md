# RolesAndProtocolsIdpRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **interface{}** | Enable Identity Provider Role. | [optional] 
**EnableSaml11** | Pointer to **interface{}** | Enable SAML 1.1. | [optional] 
**EnableSaml10** | Pointer to **interface{}** | Enable SAML 1.0. | [optional] 
**EnableWsFed** | Pointer to **interface{}** | Enable WS Federation. | [optional] 
**EnableWsTrust** | Pointer to **interface{}** | Enable WS Trust. | [optional] 
**Saml20Profile** | Pointer to [**SAML20Profile**](SAML20Profile.md) |  | [optional] 
**EnableOutboundProvisioning** | Pointer to **interface{}** | Enable Outbound Provisioning. | [optional] 

## Methods

### NewRolesAndProtocolsIdpRole

`func NewRolesAndProtocolsIdpRole() *RolesAndProtocolsIdpRole`

NewRolesAndProtocolsIdpRole instantiates a new RolesAndProtocolsIdpRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesAndProtocolsIdpRoleWithDefaults

`func NewRolesAndProtocolsIdpRoleWithDefaults() *RolesAndProtocolsIdpRole`

NewRolesAndProtocolsIdpRoleWithDefaults instantiates a new RolesAndProtocolsIdpRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *RolesAndProtocolsIdpRole) GetEnable() interface{}`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *RolesAndProtocolsIdpRole) GetEnableOk() (*interface{}, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *RolesAndProtocolsIdpRole) SetEnable(v interface{})`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *RolesAndProtocolsIdpRole) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *RolesAndProtocolsIdpRole) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *RolesAndProtocolsIdpRole) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetEnableSaml11

`func (o *RolesAndProtocolsIdpRole) GetEnableSaml11() interface{}`

GetEnableSaml11 returns the EnableSaml11 field if non-nil, zero value otherwise.

### GetEnableSaml11Ok

`func (o *RolesAndProtocolsIdpRole) GetEnableSaml11Ok() (*interface{}, bool)`

GetEnableSaml11Ok returns a tuple with the EnableSaml11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSaml11

`func (o *RolesAndProtocolsIdpRole) SetEnableSaml11(v interface{})`

SetEnableSaml11 sets EnableSaml11 field to given value.

### HasEnableSaml11

`func (o *RolesAndProtocolsIdpRole) HasEnableSaml11() bool`

HasEnableSaml11 returns a boolean if a field has been set.

### SetEnableSaml11Nil

`func (o *RolesAndProtocolsIdpRole) SetEnableSaml11Nil(b bool)`

 SetEnableSaml11Nil sets the value for EnableSaml11 to be an explicit nil

### UnsetEnableSaml11
`func (o *RolesAndProtocolsIdpRole) UnsetEnableSaml11()`

UnsetEnableSaml11 ensures that no value is present for EnableSaml11, not even an explicit nil
### GetEnableSaml10

`func (o *RolesAndProtocolsIdpRole) GetEnableSaml10() interface{}`

GetEnableSaml10 returns the EnableSaml10 field if non-nil, zero value otherwise.

### GetEnableSaml10Ok

`func (o *RolesAndProtocolsIdpRole) GetEnableSaml10Ok() (*interface{}, bool)`

GetEnableSaml10Ok returns a tuple with the EnableSaml10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSaml10

`func (o *RolesAndProtocolsIdpRole) SetEnableSaml10(v interface{})`

SetEnableSaml10 sets EnableSaml10 field to given value.

### HasEnableSaml10

`func (o *RolesAndProtocolsIdpRole) HasEnableSaml10() bool`

HasEnableSaml10 returns a boolean if a field has been set.

### SetEnableSaml10Nil

`func (o *RolesAndProtocolsIdpRole) SetEnableSaml10Nil(b bool)`

 SetEnableSaml10Nil sets the value for EnableSaml10 to be an explicit nil

### UnsetEnableSaml10
`func (o *RolesAndProtocolsIdpRole) UnsetEnableSaml10()`

UnsetEnableSaml10 ensures that no value is present for EnableSaml10, not even an explicit nil
### GetEnableWsFed

`func (o *RolesAndProtocolsIdpRole) GetEnableWsFed() interface{}`

GetEnableWsFed returns the EnableWsFed field if non-nil, zero value otherwise.

### GetEnableWsFedOk

`func (o *RolesAndProtocolsIdpRole) GetEnableWsFedOk() (*interface{}, bool)`

GetEnableWsFedOk returns a tuple with the EnableWsFed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWsFed

`func (o *RolesAndProtocolsIdpRole) SetEnableWsFed(v interface{})`

SetEnableWsFed sets EnableWsFed field to given value.

### HasEnableWsFed

`func (o *RolesAndProtocolsIdpRole) HasEnableWsFed() bool`

HasEnableWsFed returns a boolean if a field has been set.

### SetEnableWsFedNil

`func (o *RolesAndProtocolsIdpRole) SetEnableWsFedNil(b bool)`

 SetEnableWsFedNil sets the value for EnableWsFed to be an explicit nil

### UnsetEnableWsFed
`func (o *RolesAndProtocolsIdpRole) UnsetEnableWsFed()`

UnsetEnableWsFed ensures that no value is present for EnableWsFed, not even an explicit nil
### GetEnableWsTrust

`func (o *RolesAndProtocolsIdpRole) GetEnableWsTrust() interface{}`

GetEnableWsTrust returns the EnableWsTrust field if non-nil, zero value otherwise.

### GetEnableWsTrustOk

`func (o *RolesAndProtocolsIdpRole) GetEnableWsTrustOk() (*interface{}, bool)`

GetEnableWsTrustOk returns a tuple with the EnableWsTrust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWsTrust

`func (o *RolesAndProtocolsIdpRole) SetEnableWsTrust(v interface{})`

SetEnableWsTrust sets EnableWsTrust field to given value.

### HasEnableWsTrust

`func (o *RolesAndProtocolsIdpRole) HasEnableWsTrust() bool`

HasEnableWsTrust returns a boolean if a field has been set.

### SetEnableWsTrustNil

`func (o *RolesAndProtocolsIdpRole) SetEnableWsTrustNil(b bool)`

 SetEnableWsTrustNil sets the value for EnableWsTrust to be an explicit nil

### UnsetEnableWsTrust
`func (o *RolesAndProtocolsIdpRole) UnsetEnableWsTrust()`

UnsetEnableWsTrust ensures that no value is present for EnableWsTrust, not even an explicit nil
### GetSaml20Profile

`func (o *RolesAndProtocolsIdpRole) GetSaml20Profile() SAML20Profile`

GetSaml20Profile returns the Saml20Profile field if non-nil, zero value otherwise.

### GetSaml20ProfileOk

`func (o *RolesAndProtocolsIdpRole) GetSaml20ProfileOk() (*SAML20Profile, bool)`

GetSaml20ProfileOk returns a tuple with the Saml20Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml20Profile

`func (o *RolesAndProtocolsIdpRole) SetSaml20Profile(v SAML20Profile)`

SetSaml20Profile sets Saml20Profile field to given value.

### HasSaml20Profile

`func (o *RolesAndProtocolsIdpRole) HasSaml20Profile() bool`

HasSaml20Profile returns a boolean if a field has been set.

### GetEnableOutboundProvisioning

`func (o *RolesAndProtocolsIdpRole) GetEnableOutboundProvisioning() interface{}`

GetEnableOutboundProvisioning returns the EnableOutboundProvisioning field if non-nil, zero value otherwise.

### GetEnableOutboundProvisioningOk

`func (o *RolesAndProtocolsIdpRole) GetEnableOutboundProvisioningOk() (*interface{}, bool)`

GetEnableOutboundProvisioningOk returns a tuple with the EnableOutboundProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOutboundProvisioning

`func (o *RolesAndProtocolsIdpRole) SetEnableOutboundProvisioning(v interface{})`

SetEnableOutboundProvisioning sets EnableOutboundProvisioning field to given value.

### HasEnableOutboundProvisioning

`func (o *RolesAndProtocolsIdpRole) HasEnableOutboundProvisioning() bool`

HasEnableOutboundProvisioning returns a boolean if a field has been set.

### SetEnableOutboundProvisioningNil

`func (o *RolesAndProtocolsIdpRole) SetEnableOutboundProvisioningNil(b bool)`

 SetEnableOutboundProvisioningNil sets the value for EnableOutboundProvisioning to be an explicit nil

### UnsetEnableOutboundProvisioning
`func (o *RolesAndProtocolsIdpRole) UnsetEnableOutboundProvisioning()`

UnsetEnableOutboundProvisioning ensures that no value is present for EnableOutboundProvisioning, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


