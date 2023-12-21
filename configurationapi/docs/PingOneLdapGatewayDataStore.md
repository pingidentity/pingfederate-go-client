# PingOneLdapGatewayDataStore

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
**LastModified** | Pointer to **time.Time** | The time at which the datastore instance was last changed. This property is read only and is ignored on PUT and POST requests. | [optional] 

## Methods

### NewPingOneLdapGatewayDataStore

`func NewPingOneLdapGatewayDataStore(ldapType string, pingOneConnectionRef ResourceLink, pingOneEnvironmentId string, pingOneLdapGatewayId string, ) *PingOneLdapGatewayDataStore`

NewPingOneLdapGatewayDataStore instantiates a new PingOneLdapGatewayDataStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingOneLdapGatewayDataStoreWithDefaults

`func NewPingOneLdapGatewayDataStoreWithDefaults() *PingOneLdapGatewayDataStore`

NewPingOneLdapGatewayDataStoreWithDefaults instantiates a new PingOneLdapGatewayDataStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PingOneLdapGatewayDataStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PingOneLdapGatewayDataStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PingOneLdapGatewayDataStore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PingOneLdapGatewayDataStore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLdapType

`func (o *PingOneLdapGatewayDataStore) GetLdapType() string`

GetLdapType returns the LdapType field if non-nil, zero value otherwise.

### GetLdapTypeOk

`func (o *PingOneLdapGatewayDataStore) GetLdapTypeOk() (*string, bool)`

GetLdapTypeOk returns a tuple with the LdapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapType

`func (o *PingOneLdapGatewayDataStore) SetLdapType(v string)`

SetLdapType sets LdapType field to given value.


### GetPingOneConnectionRef

`func (o *PingOneLdapGatewayDataStore) GetPingOneConnectionRef() ResourceLink`

GetPingOneConnectionRef returns the PingOneConnectionRef field if non-nil, zero value otherwise.

### GetPingOneConnectionRefOk

`func (o *PingOneLdapGatewayDataStore) GetPingOneConnectionRefOk() (*ResourceLink, bool)`

GetPingOneConnectionRefOk returns a tuple with the PingOneConnectionRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOneConnectionRef

`func (o *PingOneLdapGatewayDataStore) SetPingOneConnectionRef(v ResourceLink)`

SetPingOneConnectionRef sets PingOneConnectionRef field to given value.


### GetPingOneEnvironmentId

`func (o *PingOneLdapGatewayDataStore) GetPingOneEnvironmentId() string`

GetPingOneEnvironmentId returns the PingOneEnvironmentId field if non-nil, zero value otherwise.

### GetPingOneEnvironmentIdOk

`func (o *PingOneLdapGatewayDataStore) GetPingOneEnvironmentIdOk() (*string, bool)`

GetPingOneEnvironmentIdOk returns a tuple with the PingOneEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOneEnvironmentId

`func (o *PingOneLdapGatewayDataStore) SetPingOneEnvironmentId(v string)`

SetPingOneEnvironmentId sets PingOneEnvironmentId field to given value.


### GetPingOneLdapGatewayId

`func (o *PingOneLdapGatewayDataStore) GetPingOneLdapGatewayId() string`

GetPingOneLdapGatewayId returns the PingOneLdapGatewayId field if non-nil, zero value otherwise.

### GetPingOneLdapGatewayIdOk

`func (o *PingOneLdapGatewayDataStore) GetPingOneLdapGatewayIdOk() (*string, bool)`

GetPingOneLdapGatewayIdOk returns a tuple with the PingOneLdapGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOneLdapGatewayId

`func (o *PingOneLdapGatewayDataStore) SetPingOneLdapGatewayId(v string)`

SetPingOneLdapGatewayId sets PingOneLdapGatewayId field to given value.


### GetUseSsl

`func (o *PingOneLdapGatewayDataStore) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *PingOneLdapGatewayDataStore) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *PingOneLdapGatewayDataStore) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *PingOneLdapGatewayDataStore) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### GetBinaryAttributes

`func (o *PingOneLdapGatewayDataStore) GetBinaryAttributes() []string`

GetBinaryAttributes returns the BinaryAttributes field if non-nil, zero value otherwise.

### GetBinaryAttributesOk

`func (o *PingOneLdapGatewayDataStore) GetBinaryAttributesOk() (*[]string, bool)`

GetBinaryAttributesOk returns a tuple with the BinaryAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryAttributes

`func (o *PingOneLdapGatewayDataStore) SetBinaryAttributes(v []string)`

SetBinaryAttributes sets BinaryAttributes field to given value.

### HasBinaryAttributes

`func (o *PingOneLdapGatewayDataStore) HasBinaryAttributes() bool`

HasBinaryAttributes returns a boolean if a field has been set.

### GetLastModified

`func (o *PingOneLdapGatewayDataStore) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *PingOneLdapGatewayDataStore) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *PingOneLdapGatewayDataStore) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *PingOneLdapGatewayDataStore) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


