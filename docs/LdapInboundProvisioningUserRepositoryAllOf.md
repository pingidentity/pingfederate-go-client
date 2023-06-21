# LdapInboundProvisioningUserRepositoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataStoreRef** | [**ResourceLink**](ResourceLink.md) |  | 
**BaseDn** | Pointer to **string** | The base DN to search from. If not specified, the search will start at the LDAP&#39;s root. | [optional] 
**UniqueUserIdFilter** | **string** | The expression that results in a unique user identifier, when combined with the Base DN. | 
**UniqueGroupIdFilter** | **string** | The expression that results in a unique group identifier, when combined with the Base DN. | 

## Methods

### NewLdapInboundProvisioningUserRepositoryAllOf

`func NewLdapInboundProvisioningUserRepositoryAllOf(dataStoreRef ResourceLink, uniqueUserIdFilter string, uniqueGroupIdFilter string, ) *LdapInboundProvisioningUserRepositoryAllOf`

NewLdapInboundProvisioningUserRepositoryAllOf instantiates a new LdapInboundProvisioningUserRepositoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapInboundProvisioningUserRepositoryAllOfWithDefaults

`func NewLdapInboundProvisioningUserRepositoryAllOfWithDefaults() *LdapInboundProvisioningUserRepositoryAllOf`

NewLdapInboundProvisioningUserRepositoryAllOfWithDefaults instantiates a new LdapInboundProvisioningUserRepositoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataStoreRef

`func (o *LdapInboundProvisioningUserRepositoryAllOf) GetDataStoreRef() ResourceLink`

GetDataStoreRef returns the DataStoreRef field if non-nil, zero value otherwise.

### GetDataStoreRefOk

`func (o *LdapInboundProvisioningUserRepositoryAllOf) GetDataStoreRefOk() (*ResourceLink, bool)`

GetDataStoreRefOk returns a tuple with the DataStoreRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreRef

`func (o *LdapInboundProvisioningUserRepositoryAllOf) SetDataStoreRef(v ResourceLink)`

SetDataStoreRef sets DataStoreRef field to given value.


### GetBaseDn

`func (o *LdapInboundProvisioningUserRepositoryAllOf) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LdapInboundProvisioningUserRepositoryAllOf) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LdapInboundProvisioningUserRepositoryAllOf) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *LdapInboundProvisioningUserRepositoryAllOf) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetUniqueUserIdFilter

`func (o *LdapInboundProvisioningUserRepositoryAllOf) GetUniqueUserIdFilter() string`

GetUniqueUserIdFilter returns the UniqueUserIdFilter field if non-nil, zero value otherwise.

### GetUniqueUserIdFilterOk

`func (o *LdapInboundProvisioningUserRepositoryAllOf) GetUniqueUserIdFilterOk() (*string, bool)`

GetUniqueUserIdFilterOk returns a tuple with the UniqueUserIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueUserIdFilter

`func (o *LdapInboundProvisioningUserRepositoryAllOf) SetUniqueUserIdFilter(v string)`

SetUniqueUserIdFilter sets UniqueUserIdFilter field to given value.


### GetUniqueGroupIdFilter

`func (o *LdapInboundProvisioningUserRepositoryAllOf) GetUniqueGroupIdFilter() string`

GetUniqueGroupIdFilter returns the UniqueGroupIdFilter field if non-nil, zero value otherwise.

### GetUniqueGroupIdFilterOk

`func (o *LdapInboundProvisioningUserRepositoryAllOf) GetUniqueGroupIdFilterOk() (*string, bool)`

GetUniqueGroupIdFilterOk returns a tuple with the UniqueGroupIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueGroupIdFilter

`func (o *LdapInboundProvisioningUserRepositoryAllOf) SetUniqueGroupIdFilter(v string)`

SetUniqueGroupIdFilter sets UniqueGroupIdFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


