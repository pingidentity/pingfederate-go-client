# RolesAndProtocolsSpRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **interface{}** | Enable Service Provider Role. | [optional] 
**EnableSaml11** | Pointer to **interface{}** | Enable SAML 1.1. | [optional] 
**EnableSaml10** | Pointer to **interface{}** | Enable SAML 1.0. | [optional] 
**EnableWsFed** | Pointer to **interface{}** | Enable WS Federation. | [optional] 
**EnableWsTrust** | Pointer to **interface{}** | Enable WS Trust. | [optional] 
**Saml20Profile** | Pointer to [**SpSAML20Profile**](SpSAML20Profile.md) |  | [optional] 
**EnableOpenIDConnect** | Pointer to **interface{}** | Enable OpenID Connect. | [optional] 
**EnableInboundProvisioning** | Pointer to **interface{}** | Enable Inbound Provisioning. | [optional] 

## Methods

### NewRolesAndProtocolsSpRole

`func NewRolesAndProtocolsSpRole() *RolesAndProtocolsSpRole`

NewRolesAndProtocolsSpRole instantiates a new RolesAndProtocolsSpRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesAndProtocolsSpRoleWithDefaults

`func NewRolesAndProtocolsSpRoleWithDefaults() *RolesAndProtocolsSpRole`

NewRolesAndProtocolsSpRoleWithDefaults instantiates a new RolesAndProtocolsSpRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *RolesAndProtocolsSpRole) GetEnable() interface{}`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *RolesAndProtocolsSpRole) GetEnableOk() (*interface{}, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *RolesAndProtocolsSpRole) SetEnable(v interface{})`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *RolesAndProtocolsSpRole) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *RolesAndProtocolsSpRole) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *RolesAndProtocolsSpRole) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetEnableSaml11

`func (o *RolesAndProtocolsSpRole) GetEnableSaml11() interface{}`

GetEnableSaml11 returns the EnableSaml11 field if non-nil, zero value otherwise.

### GetEnableSaml11Ok

`func (o *RolesAndProtocolsSpRole) GetEnableSaml11Ok() (*interface{}, bool)`

GetEnableSaml11Ok returns a tuple with the EnableSaml11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSaml11

`func (o *RolesAndProtocolsSpRole) SetEnableSaml11(v interface{})`

SetEnableSaml11 sets EnableSaml11 field to given value.

### HasEnableSaml11

`func (o *RolesAndProtocolsSpRole) HasEnableSaml11() bool`

HasEnableSaml11 returns a boolean if a field has been set.

### SetEnableSaml11Nil

`func (o *RolesAndProtocolsSpRole) SetEnableSaml11Nil(b bool)`

 SetEnableSaml11Nil sets the value for EnableSaml11 to be an explicit nil

### UnsetEnableSaml11
`func (o *RolesAndProtocolsSpRole) UnsetEnableSaml11()`

UnsetEnableSaml11 ensures that no value is present for EnableSaml11, not even an explicit nil
### GetEnableSaml10

`func (o *RolesAndProtocolsSpRole) GetEnableSaml10() interface{}`

GetEnableSaml10 returns the EnableSaml10 field if non-nil, zero value otherwise.

### GetEnableSaml10Ok

`func (o *RolesAndProtocolsSpRole) GetEnableSaml10Ok() (*interface{}, bool)`

GetEnableSaml10Ok returns a tuple with the EnableSaml10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSaml10

`func (o *RolesAndProtocolsSpRole) SetEnableSaml10(v interface{})`

SetEnableSaml10 sets EnableSaml10 field to given value.

### HasEnableSaml10

`func (o *RolesAndProtocolsSpRole) HasEnableSaml10() bool`

HasEnableSaml10 returns a boolean if a field has been set.

### SetEnableSaml10Nil

`func (o *RolesAndProtocolsSpRole) SetEnableSaml10Nil(b bool)`

 SetEnableSaml10Nil sets the value for EnableSaml10 to be an explicit nil

### UnsetEnableSaml10
`func (o *RolesAndProtocolsSpRole) UnsetEnableSaml10()`

UnsetEnableSaml10 ensures that no value is present for EnableSaml10, not even an explicit nil
### GetEnableWsFed

`func (o *RolesAndProtocolsSpRole) GetEnableWsFed() interface{}`

GetEnableWsFed returns the EnableWsFed field if non-nil, zero value otherwise.

### GetEnableWsFedOk

`func (o *RolesAndProtocolsSpRole) GetEnableWsFedOk() (*interface{}, bool)`

GetEnableWsFedOk returns a tuple with the EnableWsFed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWsFed

`func (o *RolesAndProtocolsSpRole) SetEnableWsFed(v interface{})`

SetEnableWsFed sets EnableWsFed field to given value.

### HasEnableWsFed

`func (o *RolesAndProtocolsSpRole) HasEnableWsFed() bool`

HasEnableWsFed returns a boolean if a field has been set.

### SetEnableWsFedNil

`func (o *RolesAndProtocolsSpRole) SetEnableWsFedNil(b bool)`

 SetEnableWsFedNil sets the value for EnableWsFed to be an explicit nil

### UnsetEnableWsFed
`func (o *RolesAndProtocolsSpRole) UnsetEnableWsFed()`

UnsetEnableWsFed ensures that no value is present for EnableWsFed, not even an explicit nil
### GetEnableWsTrust

`func (o *RolesAndProtocolsSpRole) GetEnableWsTrust() interface{}`

GetEnableWsTrust returns the EnableWsTrust field if non-nil, zero value otherwise.

### GetEnableWsTrustOk

`func (o *RolesAndProtocolsSpRole) GetEnableWsTrustOk() (*interface{}, bool)`

GetEnableWsTrustOk returns a tuple with the EnableWsTrust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWsTrust

`func (o *RolesAndProtocolsSpRole) SetEnableWsTrust(v interface{})`

SetEnableWsTrust sets EnableWsTrust field to given value.

### HasEnableWsTrust

`func (o *RolesAndProtocolsSpRole) HasEnableWsTrust() bool`

HasEnableWsTrust returns a boolean if a field has been set.

### SetEnableWsTrustNil

`func (o *RolesAndProtocolsSpRole) SetEnableWsTrustNil(b bool)`

 SetEnableWsTrustNil sets the value for EnableWsTrust to be an explicit nil

### UnsetEnableWsTrust
`func (o *RolesAndProtocolsSpRole) UnsetEnableWsTrust()`

UnsetEnableWsTrust ensures that no value is present for EnableWsTrust, not even an explicit nil
### GetSaml20Profile

`func (o *RolesAndProtocolsSpRole) GetSaml20Profile() SpSAML20Profile`

GetSaml20Profile returns the Saml20Profile field if non-nil, zero value otherwise.

### GetSaml20ProfileOk

`func (o *RolesAndProtocolsSpRole) GetSaml20ProfileOk() (*SpSAML20Profile, bool)`

GetSaml20ProfileOk returns a tuple with the Saml20Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml20Profile

`func (o *RolesAndProtocolsSpRole) SetSaml20Profile(v SpSAML20Profile)`

SetSaml20Profile sets Saml20Profile field to given value.

### HasSaml20Profile

`func (o *RolesAndProtocolsSpRole) HasSaml20Profile() bool`

HasSaml20Profile returns a boolean if a field has been set.

### GetEnableOpenIDConnect

`func (o *RolesAndProtocolsSpRole) GetEnableOpenIDConnect() interface{}`

GetEnableOpenIDConnect returns the EnableOpenIDConnect field if non-nil, zero value otherwise.

### GetEnableOpenIDConnectOk

`func (o *RolesAndProtocolsSpRole) GetEnableOpenIDConnectOk() (*interface{}, bool)`

GetEnableOpenIDConnectOk returns a tuple with the EnableOpenIDConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOpenIDConnect

`func (o *RolesAndProtocolsSpRole) SetEnableOpenIDConnect(v interface{})`

SetEnableOpenIDConnect sets EnableOpenIDConnect field to given value.

### HasEnableOpenIDConnect

`func (o *RolesAndProtocolsSpRole) HasEnableOpenIDConnect() bool`

HasEnableOpenIDConnect returns a boolean if a field has been set.

### SetEnableOpenIDConnectNil

`func (o *RolesAndProtocolsSpRole) SetEnableOpenIDConnectNil(b bool)`

 SetEnableOpenIDConnectNil sets the value for EnableOpenIDConnect to be an explicit nil

### UnsetEnableOpenIDConnect
`func (o *RolesAndProtocolsSpRole) UnsetEnableOpenIDConnect()`

UnsetEnableOpenIDConnect ensures that no value is present for EnableOpenIDConnect, not even an explicit nil
### GetEnableInboundProvisioning

`func (o *RolesAndProtocolsSpRole) GetEnableInboundProvisioning() interface{}`

GetEnableInboundProvisioning returns the EnableInboundProvisioning field if non-nil, zero value otherwise.

### GetEnableInboundProvisioningOk

`func (o *RolesAndProtocolsSpRole) GetEnableInboundProvisioningOk() (*interface{}, bool)`

GetEnableInboundProvisioningOk returns a tuple with the EnableInboundProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInboundProvisioning

`func (o *RolesAndProtocolsSpRole) SetEnableInboundProvisioning(v interface{})`

SetEnableInboundProvisioning sets EnableInboundProvisioning field to given value.

### HasEnableInboundProvisioning

`func (o *RolesAndProtocolsSpRole) HasEnableInboundProvisioning() bool`

HasEnableInboundProvisioning returns a boolean if a field has been set.

### SetEnableInboundProvisioningNil

`func (o *RolesAndProtocolsSpRole) SetEnableInboundProvisioningNil(b bool)`

 SetEnableInboundProvisioningNil sets the value for EnableInboundProvisioning to be an explicit nil

### UnsetEnableInboundProvisioning
`func (o *RolesAndProtocolsSpRole) UnsetEnableInboundProvisioning()`

UnsetEnableInboundProvisioning ensures that no value is present for EnableInboundProvisioning, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


