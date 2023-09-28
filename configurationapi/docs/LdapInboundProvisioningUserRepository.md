# LdapInboundProvisioningUserRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataStoreRef** | [**ResourceLink**](ResourceLink.md) |  | 
**BaseDn** | Pointer to **string** | The base DN to search from. If not specified, the search will start at the LDAP&#39;s root. | [optional] 
**UniqueUserIdFilter** | **string** | The expression that results in a unique user identifier, when combined with the Base DN. | 
**UniqueGroupIdFilter** | **string** | The expression that results in a unique group identifier, when combined with the Base DN. | 

## Methods

### NewLdapInboundProvisioningUserRepository

`func NewLdapInboundProvisioningUserRepository(dataStoreRef ResourceLink, uniqueUserIdFilter string, uniqueGroupIdFilter string, ) *LdapInboundProvisioningUserRepository`

NewLdapInboundProvisioningUserRepository instantiates a new LdapInboundProvisioningUserRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapInboundProvisioningUserRepositoryWithDefaults

`func NewLdapInboundProvisioningUserRepositoryWithDefaults() *LdapInboundProvisioningUserRepository`

NewLdapInboundProvisioningUserRepositoryWithDefaults instantiates a new LdapInboundProvisioningUserRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataStoreRef

`func (o *LdapInboundProvisioningUserRepository) GetDataStoreRef() ResourceLink`

GetDataStoreRef returns the DataStoreRef field if non-nil, zero value otherwise.

### GetDataStoreRefOk

`func (o *LdapInboundProvisioningUserRepository) GetDataStoreRefOk() (*ResourceLink, bool)`

GetDataStoreRefOk returns a tuple with the DataStoreRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreRef

`func (o *LdapInboundProvisioningUserRepository) SetDataStoreRef(v ResourceLink)`

SetDataStoreRef sets DataStoreRef field to given value.


### GetBaseDn

`func (o *LdapInboundProvisioningUserRepository) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LdapInboundProvisioningUserRepository) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LdapInboundProvisioningUserRepository) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *LdapInboundProvisioningUserRepository) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetUniqueUserIdFilter

`func (o *LdapInboundProvisioningUserRepository) GetUniqueUserIdFilter() string`

GetUniqueUserIdFilter returns the UniqueUserIdFilter field if non-nil, zero value otherwise.

### GetUniqueUserIdFilterOk

`func (o *LdapInboundProvisioningUserRepository) GetUniqueUserIdFilterOk() (*string, bool)`

GetUniqueUserIdFilterOk returns a tuple with the UniqueUserIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueUserIdFilter

`func (o *LdapInboundProvisioningUserRepository) SetUniqueUserIdFilter(v string)`

SetUniqueUserIdFilter sets UniqueUserIdFilter field to given value.


### GetUniqueGroupIdFilter

`func (o *LdapInboundProvisioningUserRepository) GetUniqueGroupIdFilter() string`

GetUniqueGroupIdFilter returns the UniqueGroupIdFilter field if non-nil, zero value otherwise.

### GetUniqueGroupIdFilterOk

`func (o *LdapInboundProvisioningUserRepository) GetUniqueGroupIdFilterOk() (*string, bool)`

GetUniqueGroupIdFilterOk returns a tuple with the UniqueGroupIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueGroupIdFilter

`func (o *LdapInboundProvisioningUserRepository) SetUniqueGroupIdFilter(v string)`

SetUniqueGroupIdFilter sets UniqueGroupIdFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


