# PingOneLdapGatewayDataStoreAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The data store name with a unique value across all data sources. Omitting this attribute will set the value to a combination of the hostname(s) and the principal. | [optional] 
**LdapType** | **string** | A type that allows PingFederate to configure many provisioning settings automatically. The value is validated against the LDAP gateway configuration in PingOne unless the header &#39;X-BypassExternalValidation&#39; is set to true. | 
**PingOneConnectionRef** | [**ResourceLink**](ResourceLink.md) |  | 
**PingOneEnvironmentId** | **string** | The environment ID that the gateway belongs to. | 
**PingOneLdapGatewayId** | **string** | The ID of the PingOne LDAP Gateway this data store uses. | 
**UseSsl** | Pointer to **bool** | Connects to the LDAP data store using secure SSL/TLS encryption (LDAPS). The default value is false. The value is validated against the LDAP gateway configuration in PingOne unless the header &#39;X-BypassExternalValidation&#39; is set to true. | [optional] 
**BinaryAttributes** | Pointer to **[]string** | The list of LDAP attributes to be handled as binary data. | [optional] 

## Methods

### NewPingOneLdapGatewayDataStoreAllOf

`func NewPingOneLdapGatewayDataStoreAllOf(ldapType string, pingOneConnectionRef ResourceLink, pingOneEnvironmentId string, pingOneLdapGatewayId string, ) *PingOneLdapGatewayDataStoreAllOf`

NewPingOneLdapGatewayDataStoreAllOf instantiates a new PingOneLdapGatewayDataStoreAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingOneLdapGatewayDataStoreAllOfWithDefaults

`func NewPingOneLdapGatewayDataStoreAllOfWithDefaults() *PingOneLdapGatewayDataStoreAllOf`

NewPingOneLdapGatewayDataStoreAllOfWithDefaults instantiates a new PingOneLdapGatewayDataStoreAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PingOneLdapGatewayDataStoreAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PingOneLdapGatewayDataStoreAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PingOneLdapGatewayDataStoreAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PingOneLdapGatewayDataStoreAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLdapType

`func (o *PingOneLdapGatewayDataStoreAllOf) GetLdapType() string`

GetLdapType returns the LdapType field if non-nil, zero value otherwise.

### GetLdapTypeOk

`func (o *PingOneLdapGatewayDataStoreAllOf) GetLdapTypeOk() (*string, bool)`

GetLdapTypeOk returns a tuple with the LdapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapType

`func (o *PingOneLdapGatewayDataStoreAllOf) SetLdapType(v string)`

SetLdapType sets LdapType field to given value.


### GetPingOneConnectionRef

`func (o *PingOneLdapGatewayDataStoreAllOf) GetPingOneConnectionRef() ResourceLink`

GetPingOneConnectionRef returns the PingOneConnectionRef field if non-nil, zero value otherwise.

### GetPingOneConnectionRefOk

`func (o *PingOneLdapGatewayDataStoreAllOf) GetPingOneConnectionRefOk() (*ResourceLink, bool)`

GetPingOneConnectionRefOk returns a tuple with the PingOneConnectionRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOneConnectionRef

`func (o *PingOneLdapGatewayDataStoreAllOf) SetPingOneConnectionRef(v ResourceLink)`

SetPingOneConnectionRef sets PingOneConnectionRef field to given value.


### GetPingOneEnvironmentId

`func (o *PingOneLdapGatewayDataStoreAllOf) GetPingOneEnvironmentId() string`

GetPingOneEnvironmentId returns the PingOneEnvironmentId field if non-nil, zero value otherwise.

### GetPingOneEnvironmentIdOk

`func (o *PingOneLdapGatewayDataStoreAllOf) GetPingOneEnvironmentIdOk() (*string, bool)`

GetPingOneEnvironmentIdOk returns a tuple with the PingOneEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOneEnvironmentId

`func (o *PingOneLdapGatewayDataStoreAllOf) SetPingOneEnvironmentId(v string)`

SetPingOneEnvironmentId sets PingOneEnvironmentId field to given value.


### GetPingOneLdapGatewayId

`func (o *PingOneLdapGatewayDataStoreAllOf) GetPingOneLdapGatewayId() string`

GetPingOneLdapGatewayId returns the PingOneLdapGatewayId field if non-nil, zero value otherwise.

### GetPingOneLdapGatewayIdOk

`func (o *PingOneLdapGatewayDataStoreAllOf) GetPingOneLdapGatewayIdOk() (*string, bool)`

GetPingOneLdapGatewayIdOk returns a tuple with the PingOneLdapGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOneLdapGatewayId

`func (o *PingOneLdapGatewayDataStoreAllOf) SetPingOneLdapGatewayId(v string)`

SetPingOneLdapGatewayId sets PingOneLdapGatewayId field to given value.


### GetUseSsl

`func (o *PingOneLdapGatewayDataStoreAllOf) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *PingOneLdapGatewayDataStoreAllOf) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *PingOneLdapGatewayDataStoreAllOf) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *PingOneLdapGatewayDataStoreAllOf) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### GetBinaryAttributes

`func (o *PingOneLdapGatewayDataStoreAllOf) GetBinaryAttributes() []string`

GetBinaryAttributes returns the BinaryAttributes field if non-nil, zero value otherwise.

### GetBinaryAttributesOk

`func (o *PingOneLdapGatewayDataStoreAllOf) GetBinaryAttributesOk() (*[]string, bool)`

GetBinaryAttributesOk returns a tuple with the BinaryAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryAttributes

`func (o *PingOneLdapGatewayDataStoreAllOf) SetBinaryAttributes(v []string)`

SetBinaryAttributes sets BinaryAttributes field to given value.

### HasBinaryAttributes

`func (o *PingOneLdapGatewayDataStoreAllOf) HasBinaryAttributes() bool`

HasBinaryAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


