# SpRoleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Enable Service Provider Role. | [optional] 
**Saml20Profile** | Pointer to [**SpSAML20Profile**](SpSAML20Profile.md) |  | [optional] 
**EnableOpenIDConnect** | Pointer to **bool** | Enable OpenID Connect. | [optional] 
**EnableInboundProvisioning** | Pointer to **bool** | Enable Inbound Provisioning. | [optional] 

## Methods

### NewSpRoleAllOf

`func NewSpRoleAllOf() *SpRoleAllOf`

NewSpRoleAllOf instantiates a new SpRoleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpRoleAllOfWithDefaults

`func NewSpRoleAllOfWithDefaults() *SpRoleAllOf`

NewSpRoleAllOfWithDefaults instantiates a new SpRoleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *SpRoleAllOf) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *SpRoleAllOf) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *SpRoleAllOf) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *SpRoleAllOf) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetSaml20Profile

`func (o *SpRoleAllOf) GetSaml20Profile() SpSAML20Profile`

GetSaml20Profile returns the Saml20Profile field if non-nil, zero value otherwise.

### GetSaml20ProfileOk

`func (o *SpRoleAllOf) GetSaml20ProfileOk() (*SpSAML20Profile, bool)`

GetSaml20ProfileOk returns a tuple with the Saml20Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml20Profile

`func (o *SpRoleAllOf) SetSaml20Profile(v SpSAML20Profile)`

SetSaml20Profile sets Saml20Profile field to given value.

### HasSaml20Profile

`func (o *SpRoleAllOf) HasSaml20Profile() bool`

HasSaml20Profile returns a boolean if a field has been set.

### GetEnableOpenIDConnect

`func (o *SpRoleAllOf) GetEnableOpenIDConnect() bool`

GetEnableOpenIDConnect returns the EnableOpenIDConnect field if non-nil, zero value otherwise.

### GetEnableOpenIDConnectOk

`func (o *SpRoleAllOf) GetEnableOpenIDConnectOk() (*bool, bool)`

GetEnableOpenIDConnectOk returns a tuple with the EnableOpenIDConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOpenIDConnect

`func (o *SpRoleAllOf) SetEnableOpenIDConnect(v bool)`

SetEnableOpenIDConnect sets EnableOpenIDConnect field to given value.

### HasEnableOpenIDConnect

`func (o *SpRoleAllOf) HasEnableOpenIDConnect() bool`

HasEnableOpenIDConnect returns a boolean if a field has been set.

### GetEnableInboundProvisioning

`func (o *SpRoleAllOf) GetEnableInboundProvisioning() bool`

GetEnableInboundProvisioning returns the EnableInboundProvisioning field if non-nil, zero value otherwise.

### GetEnableInboundProvisioningOk

`func (o *SpRoleAllOf) GetEnableInboundProvisioningOk() (*bool, bool)`

GetEnableInboundProvisioningOk returns a tuple with the EnableInboundProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInboundProvisioning

`func (o *SpRoleAllOf) SetEnableInboundProvisioning(v bool)`

SetEnableInboundProvisioning sets EnableInboundProvisioning field to given value.

### HasEnableInboundProvisioning

`func (o *SpRoleAllOf) HasEnableInboundProvisioning() bool`

HasEnableInboundProvisioning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


