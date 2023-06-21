# SpRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Enable Service Provider Role. | [optional] 
**EnableSaml11** | Pointer to **bool** | Enable SAML 1.1. | [optional] 
**EnableSaml10** | Pointer to **bool** | Enable SAML 1.0. | [optional] 
**EnableWsFed** | Pointer to **bool** | Enable WS Federation. | [optional] 
**EnableWsTrust** | Pointer to **bool** | Enable WS Trust. | [optional] 
**Saml20Profile** | Pointer to [**SpSAML20Profile**](SpSAML20Profile.md) |  | [optional] 
**EnableOpenIDConnect** | Pointer to **bool** | Enable OpenID Connect. | [optional] 
**EnableInboundProvisioning** | Pointer to **bool** | Enable Inbound Provisioning. | [optional] 

## Methods

### NewSpRole

`func NewSpRole() *SpRole`

NewSpRole instantiates a new SpRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpRoleWithDefaults

`func NewSpRoleWithDefaults() *SpRole`

NewSpRoleWithDefaults instantiates a new SpRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *SpRole) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *SpRole) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *SpRole) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *SpRole) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableSaml11

`func (o *SpRole) GetEnableSaml11() bool`

GetEnableSaml11 returns the EnableSaml11 field if non-nil, zero value otherwise.

### GetEnableSaml11Ok

`func (o *SpRole) GetEnableSaml11Ok() (*bool, bool)`

GetEnableSaml11Ok returns a tuple with the EnableSaml11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSaml11

`func (o *SpRole) SetEnableSaml11(v bool)`

SetEnableSaml11 sets EnableSaml11 field to given value.

### HasEnableSaml11

`func (o *SpRole) HasEnableSaml11() bool`

HasEnableSaml11 returns a boolean if a field has been set.

### GetEnableSaml10

`func (o *SpRole) GetEnableSaml10() bool`

GetEnableSaml10 returns the EnableSaml10 field if non-nil, zero value otherwise.

### GetEnableSaml10Ok

`func (o *SpRole) GetEnableSaml10Ok() (*bool, bool)`

GetEnableSaml10Ok returns a tuple with the EnableSaml10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSaml10

`func (o *SpRole) SetEnableSaml10(v bool)`

SetEnableSaml10 sets EnableSaml10 field to given value.

### HasEnableSaml10

`func (o *SpRole) HasEnableSaml10() bool`

HasEnableSaml10 returns a boolean if a field has been set.

### GetEnableWsFed

`func (o *SpRole) GetEnableWsFed() bool`

GetEnableWsFed returns the EnableWsFed field if non-nil, zero value otherwise.

### GetEnableWsFedOk

`func (o *SpRole) GetEnableWsFedOk() (*bool, bool)`

GetEnableWsFedOk returns a tuple with the EnableWsFed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWsFed

`func (o *SpRole) SetEnableWsFed(v bool)`

SetEnableWsFed sets EnableWsFed field to given value.

### HasEnableWsFed

`func (o *SpRole) HasEnableWsFed() bool`

HasEnableWsFed returns a boolean if a field has been set.

### GetEnableWsTrust

`func (o *SpRole) GetEnableWsTrust() bool`

GetEnableWsTrust returns the EnableWsTrust field if non-nil, zero value otherwise.

### GetEnableWsTrustOk

`func (o *SpRole) GetEnableWsTrustOk() (*bool, bool)`

GetEnableWsTrustOk returns a tuple with the EnableWsTrust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWsTrust

`func (o *SpRole) SetEnableWsTrust(v bool)`

SetEnableWsTrust sets EnableWsTrust field to given value.

### HasEnableWsTrust

`func (o *SpRole) HasEnableWsTrust() bool`

HasEnableWsTrust returns a boolean if a field has been set.

### GetSaml20Profile

`func (o *SpRole) GetSaml20Profile() SpSAML20Profile`

GetSaml20Profile returns the Saml20Profile field if non-nil, zero value otherwise.

### GetSaml20ProfileOk

`func (o *SpRole) GetSaml20ProfileOk() (*SpSAML20Profile, bool)`

GetSaml20ProfileOk returns a tuple with the Saml20Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml20Profile

`func (o *SpRole) SetSaml20Profile(v SpSAML20Profile)`

SetSaml20Profile sets Saml20Profile field to given value.

### HasSaml20Profile

`func (o *SpRole) HasSaml20Profile() bool`

HasSaml20Profile returns a boolean if a field has been set.

### GetEnableOpenIDConnect

`func (o *SpRole) GetEnableOpenIDConnect() bool`

GetEnableOpenIDConnect returns the EnableOpenIDConnect field if non-nil, zero value otherwise.

### GetEnableOpenIDConnectOk

`func (o *SpRole) GetEnableOpenIDConnectOk() (*bool, bool)`

GetEnableOpenIDConnectOk returns a tuple with the EnableOpenIDConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOpenIDConnect

`func (o *SpRole) SetEnableOpenIDConnect(v bool)`

SetEnableOpenIDConnect sets EnableOpenIDConnect field to given value.

### HasEnableOpenIDConnect

`func (o *SpRole) HasEnableOpenIDConnect() bool`

HasEnableOpenIDConnect returns a boolean if a field has been set.

### GetEnableInboundProvisioning

`func (o *SpRole) GetEnableInboundProvisioning() bool`

GetEnableInboundProvisioning returns the EnableInboundProvisioning field if non-nil, zero value otherwise.

### GetEnableInboundProvisioningOk

`func (o *SpRole) GetEnableInboundProvisioningOk() (*bool, bool)`

GetEnableInboundProvisioningOk returns a tuple with the EnableInboundProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInboundProvisioning

`func (o *SpRole) SetEnableInboundProvisioning(v bool)`

SetEnableInboundProvisioning sets EnableInboundProvisioning field to given value.

### HasEnableInboundProvisioning

`func (o *SpRole) HasEnableInboundProvisioning() bool`

HasEnableInboundProvisioning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


