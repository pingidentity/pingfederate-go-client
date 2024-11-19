# LdapDataStoreRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseDn** | Pointer to **string** | The base DN to search from. If not specified, the search will start at the LDAP&#39;s root. | [optional] 
**UniqueUserIdFilter** | **string** | The expression that results in a unique user identifier, when combined with the Base DN. | 

## Methods

### NewLdapDataStoreRepository

`func NewLdapDataStoreRepository(uniqueUserIdFilter string, ) *LdapDataStoreRepository`

NewLdapDataStoreRepository instantiates a new LdapDataStoreRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapDataStoreRepositoryWithDefaults

`func NewLdapDataStoreRepositoryWithDefaults() *LdapDataStoreRepository`

NewLdapDataStoreRepositoryWithDefaults instantiates a new LdapDataStoreRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseDn

`func (o *LdapDataStoreRepository) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LdapDataStoreRepository) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LdapDataStoreRepository) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *LdapDataStoreRepository) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetUniqueUserIdFilter

`func (o *LdapDataStoreRepository) GetUniqueUserIdFilter() string`

GetUniqueUserIdFilter returns the UniqueUserIdFilter field if non-nil, zero value otherwise.

### GetUniqueUserIdFilterOk

`func (o *LdapDataStoreRepository) GetUniqueUserIdFilterOk() (*string, bool)`

GetUniqueUserIdFilterOk returns a tuple with the UniqueUserIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueUserIdFilter

`func (o *LdapDataStoreRepository) SetUniqueUserIdFilter(v string)`

SetUniqueUserIdFilter sets UniqueUserIdFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


